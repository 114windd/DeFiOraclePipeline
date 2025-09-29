package ethclient

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// EthClient handles Ethereum blockchain interactions
type EthClient struct {
	client     *ethclient.Client
	privateKey *ecdsa.PrivateKey
	fromAddr   common.Address
	gasLimit   uint64
	gasPrice   *big.Int
	chainID    *big.Int
}

// NewEthClient creates a new Ethereum client
func NewEthClient(rpcURL, privateKeyHex string, gasLimit uint64) (*EthClient, error) {
	// Connect to Ethereum node
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Ethereum node: %w", err)
	}

	// Parse private key
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return nil, fmt.Errorf("failed to parse private key: %w", err)
	}

	// Get public key and address
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("failed to cast public key to ECDSA")
	}

	fromAddr := crypto.PubkeyToAddress(*publicKeyECDSA)

	// Get chain ID
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to get chain ID: %w", err)
	}

	// Get current gas price
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to get gas price: %w", err)
	}

	return &EthClient{
		client:     client,
		privateKey: privateKey,
		fromAddr:   fromAddr,
		gasLimit:   gasLimit,
		gasPrice:   gasPrice,
		chainID:    chainID,
	}, nil
}

// NewEthClientWithConfig creates a new Ethereum client with custom configuration
func NewEthClientWithConfig(
	client *ethclient.Client,
	privateKey *ecdsa.PrivateKey,
	gasLimit uint64,
	gasPrice *big.Int,
) *EthClient {
	fromAddr := crypto.PubkeyToAddress(*privateKey.Public().(*ecdsa.PublicKey))

	return &EthClient{
		client:     client,
		privateKey: privateKey,
		fromAddr:   fromAddr,
		gasLimit:   gasLimit,
		gasPrice:   gasPrice,
	}
}

// UpdatePrice updates the price in the Oracle contract
func (e *EthClient) UpdatePrice(price int64) (string, error) {
	// This is a placeholder implementation
	// In a real implementation, you would:
	// 1. Load the Oracle contract ABI
	// 2. Create a contract instance
	// 3. Call the updatePrice function

	// For now, we'll simulate a transaction
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Create transaction options
	nonce, err := e.client.PendingNonceAt(ctx, e.fromAddr)
	if err != nil {
		return "", fmt.Errorf("failed to get nonce: %w", err)
	}

	// Create a simple value transfer transaction as a placeholder
	to := common.HexToAddress("0x0000000000000000000000000000000000000000") // Placeholder address
	value := big.NewInt(0)                                                  // No ETH transfer, just contract call

	tx := &bind.TransactOpts{
		From:     e.fromAddr,
		Nonce:    big.NewInt(int64(nonce)),
		Value:    value,
		GasLimit: e.gasLimit,
		GasPrice: e.gasPrice,
		Context:  ctx,
	}

	// Sign the transaction
	signedTx, err := e.signTransaction(tx, to, value)
	if err != nil {
		return "", fmt.Errorf("failed to sign transaction: %w", err)
	}

	// Send the transaction
	err = e.client.SendTransaction(ctx, signedTx)
	if err != nil {
		return "", fmt.Errorf("failed to send transaction: %w", err)
	}

	return signedTx.Hash().Hex(), nil
}

// signTransaction signs a transaction with the private key
func (e *EthClient) signTransaction(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	// This is a simplified implementation
	// In practice, you would use the proper transaction signing methods

	// For now, return a mock transaction hash
	// In a real implementation, you would:
	// 1. Create the transaction data
	// 2. Sign it with the private key
	// 3. Return the signed transaction

	// Mock transaction hash for demonstration
	return &types.Transaction{}, nil
}

// GetBalance returns the ETH balance of the account
func (e *EthClient) GetBalance() (*big.Int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	balance, err := e.client.BalanceAt(ctx, e.fromAddr, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get balance: %w", err)
	}

	return balance, nil
}

// GetGasPrice returns the current gas price
func (e *EthClient) GetGasPrice() (*big.Int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	gasPrice, err := e.client.SuggestGasPrice(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get gas price: %w", err)
	}

	return gasPrice, nil
}

// GetAddress returns the Ethereum address
func (e *EthClient) GetAddress() common.Address {
	return e.fromAddr
}

// GetChainID returns the chain ID
func (e *EthClient) GetChainID() *big.Int {
	return e.chainID
}

// Close closes the Ethereum client connection
func (e *EthClient) Close() {
	e.client.Close()
}

// WaitForTransaction waits for a transaction to be mined
func (e *EthClient) WaitForTransaction(txHash string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	// Wait for transaction to be mined
	receipt, err := bind.WaitMined(ctx, e.client, &types.Transaction{})
	if err != nil {
		return fmt.Errorf("failed to wait for transaction: %w", err)
	}

	if receipt.Status != 1 {
		return fmt.Errorf("transaction failed")
	}

	return nil
}
