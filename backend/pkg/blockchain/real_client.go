package blockchain

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// RealClient handles real blockchain interactions using go-ethereum
type RealClient struct {
	client       *ethclient.Client
	oracle       *OracleContract
	privateKey   *ecdsa.PrivateKey
	publicKey    *ecdsa.PublicKey
	fromAddress  common.Address
	chainID      *big.Int
	contractAddr common.Address
}

// Config holds blockchain client configuration
type Config struct {
	RPCURL       string
	ContractAddr string
	PrivateKey   string
	GasLimit     uint64
}

// NewRealClient creates a new real blockchain client
func NewRealClient(config *Config) (*RealClient, error) {
	// Connect to Ethereum client
	client, err := ethclient.Dial(config.RPCURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Ethereum client: %v", err)
	}

	// Parse private key
	privateKey, err := crypto.HexToECDSA(config.PrivateKey[2:]) // Remove 0x prefix
	if err != nil {
		return nil, fmt.Errorf("failed to parse private key: %v", err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("failed to cast public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	// Get chain ID
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to get chain ID: %v", err)
	}

	// Parse contract address
	contractAddr := common.HexToAddress(config.ContractAddr)

	// Create contract instance
	oracle, err := NewOracleContract(contractAddr, client)
	if err != nil {
		return nil, fmt.Errorf("failed to create contract instance: %v", err)
	}

	return &RealClient{
		client:       client,
		oracle:       oracle,
		privateKey:   privateKey,
		publicKey:    publicKeyECDSA,
		fromAddress:  fromAddress,
		chainID:      chainID,
		contractAddr: contractAddr,
	}, nil
}

// Close closes the blockchain client connection
func (c *RealClient) Close() {
	c.client.Close()
}

// UpdateOraclePrice sends a real transaction to update the Oracle price
func (c *RealClient) UpdateOraclePrice(priceUSD float64) error {
	// Convert price to contract units (8 decimals)
	priceWei := new(big.Int)
	priceWei.SetInt64(int64(priceUSD * 100000000)) // Multiply by 10^8

	// Create transaction options
	nonce, err := c.client.PendingNonceAt(context.Background(), c.fromAddress)
	if err != nil {
		return fmt.Errorf("failed to get nonce: %v", err)
	}

	gasPrice, err := c.client.SuggestGasPrice(context.Background())
	if err != nil {
		return fmt.Errorf("failed to suggest gas price: %v", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(c.privateKey, c.chainID)
	if err != nil {
		return fmt.Errorf("failed to create transactor: %v", err)
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = 100000 // Gas limit
	auth.GasPrice = gasPrice

	// Call the updatePrice function
	tx, err := c.oracle.UpdatePrice(auth, priceWei)
	if err != nil {
		return fmt.Errorf("failed to send transaction: %v", err)
	}

	fmt.Printf("✅ Transaction sent: %s\n", tx.Hash().Hex())
	fmt.Printf("💰 Updating Oracle price to: $%.2f (contract units: %s)\n", priceUSD, priceWei.String())

	// Wait for transaction to be mined
	receipt, err := bind.WaitMined(context.Background(), c.client, tx)
	if err != nil {
		return fmt.Errorf("failed to wait for transaction: %v", err)
	}

	if receipt.Status == 0 {
		return fmt.Errorf("transaction failed")
	}

	fmt.Printf("✅ Transaction confirmed in block: %d\n", receipt.BlockNumber.Uint64())
	return nil
}

// GetLatestPrice fetches the latest price from the Oracle contract
func (c *RealClient) GetLatestPrice() (float64, time.Time, uint64, error) {
	// Create call options
	callOpts := &bind.CallOpts{
		Context: context.Background(),
	}

	// Call latestPrice function
	result, err := c.oracle.LatestPrice(callOpts)
	if err != nil {
		return 0, time.Time{}, 0, fmt.Errorf("failed to call latestPrice: %v", err)
	}

	// Convert price back to USD
	priceUSD := float64(result.Price.Int64()) / 100000000

	// Convert timestamp
	timestamp := time.Unix(result.Timestamp.Int64(), 0)

	// Convert round ID
	roundID := result.RoundId.Uint64()

	return priceUSD, timestamp, roundID, nil
}

// IsPriceStale checks if the Oracle price is stale
func (c *RealClient) IsPriceStale() (bool, error) {
	callOpts := &bind.CallOpts{
		Context: context.Background(),
	}

	isStale, err := c.oracle.IsStale(callOpts)
	if err != nil {
		return false, fmt.Errorf("failed to call isStale: %v", err)
	}

	return isStale, nil
}
