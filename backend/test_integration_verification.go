package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"time"

	"github.com/114windd/DeFiOraclePipeline.git/backend/pkg/blockchain"
	"github.com/114windd/DeFiOraclePipeline.git/backend/pkg/fetcher"
	"github.com/114windd/DeFiOraclePipeline.git/backend/pkg/normalizer"
)

func main() {
	fmt.Println("🔍 DeFi Oracle Pipeline - Integration Verification Test")
	fmt.Println("============================================================")
	fmt.Println("📋 Contract: 0x5FbDB2315678afecb367f032d93F642f64180aa3")
	fmt.Println("🔗 RPC: http://localhost:8545")
	fmt.Println("🌐 API: CoinGecko ETH/USD")
	fmt.Println("")

	// Initialize components
	fetcher := fetcher.NewFetcher("https://api.coingecko.com/api/v3/simple/price?ids=ethereum&vs_currencies=usd", 10*time.Second)
	normalizer := normalizer.NewDefaultNormalizer()

	// Load configuration from environment variables
	blockchainConfig := &blockchain.Config{
		RPCURL:       "http://localhost:8545",
		ContractAddr: "0x5FbDB2315678afecb367f032d93F642f64180aa3",
		PrivateKey:   "", // Will be loaded from environment
		GasLimit:     100000,
	}

	// Load private key from environment
	if privateKey := os.Getenv("ANVIL_PRIVATE_KEY"); privateKey != "" {
		blockchainConfig.PrivateKey = privateKey
	} else {
		log.Fatal("❌ ANVIL_PRIVATE_KEY environment variable is required")
	}

	blockchainClient, err := blockchain.NewRealClient(blockchainConfig)
	if err != nil {
		log.Fatalf("❌ Failed to initialize blockchain client: %v", err)
	}
	defer blockchainClient.Close()

	// Test 1: Get initial blockchain state
	fmt.Println("📊 PHASE 1: Initial Blockchain State")
	fmt.Println("----------------------------------------")
	initialPrice, initialTime, initialRound, err := blockchainClient.GetLatestPrice()
	if err != nil {
		log.Fatalf("❌ Failed to read initial price: %v", err)
	}
	fmt.Printf("💰 Initial contract price: $%.2f\n", initialPrice)
	fmt.Printf("🕐 Initial timestamp: %s\n", initialTime.Format("15:04:05"))
	fmt.Printf("🔢 Initial round ID: %d\n", initialRound)
	fmt.Println("")

	// Test 2: Fetch price from CoinGecko API
	fmt.Println("🌐 PHASE 2: CoinGecko API Fetch")
	fmt.Println("----------------------------------------")
	apiPrice, err := fetcher.FetchETHUSDPrice()
	if err != nil {
		log.Fatalf("❌ Failed to fetch from CoinGecko: %v", err)
	}
	fmt.Printf("📈 CoinGecko ETH price: $%.2f\n", apiPrice)

	// Test 3: Validate and normalize API price
	fmt.Println("\n✅ PHASE 3: Price Validation & Normalization")
	fmt.Println("----------------------------------------")
	if err := normalizer.ValidatePrice(apiPrice); err != nil {
		log.Fatalf("❌ Price validation failed: %v", err)
	}
	normalizedPrice := normalizer.NormalizePrice(apiPrice)
	fmt.Printf("🎯 Normalized price: $%.2f\n", normalizedPrice)
	fmt.Printf("🔢 Contract units: %.0f\n", normalizedPrice*100000000)
	fmt.Println("")

	// Test 4: Update blockchain with new price
	fmt.Println("⛓️  PHASE 4: Blockchain Update")
	fmt.Println("----------------------------------------")
	fmt.Printf("🔄 Updating Oracle contract with: $%.2f\n", normalizedPrice)

	updateStart := time.Now()
	err = blockchainClient.UpdateOraclePrice(normalizedPrice)
	updateDuration := time.Since(updateStart)

	if err != nil {
		log.Fatalf("❌ Failed to update blockchain: %v", err)
	}
	fmt.Printf("✅ Update completed in: %v\n", updateDuration)
	fmt.Println("")

	// Test 5: Verify the update
	fmt.Println("🔍 PHASE 5: Verification")
	fmt.Println("----------------------------------------")

	// Wait a moment for the transaction to be fully processed
	time.Sleep(2 * time.Second)

	finalPrice, finalTime, finalRound, err := blockchainClient.GetLatestPrice()
	if err != nil {
		log.Fatalf("❌ Failed to verify update: %v", err)
	}

	// Test 6: Compare API vs Blockchain values
	fmt.Println("📊 PHASE 6: API vs Blockchain Comparison")
	fmt.Println("----------------------------------------")

	// Calculate difference
	priceDiff := math.Abs(apiPrice - finalPrice)
	priceDiffPercent := (priceDiff / apiPrice) * 100

	fmt.Printf("🌐 CoinGecko API price: $%.2f\n", apiPrice)
	fmt.Printf("⛓️  Blockchain price:    $%.2f\n", finalPrice)
	fmt.Printf("📏 Absolute difference:  $%.2f\n", priceDiff)
	fmt.Printf("📊 Percentage difference: %.4f%%\n", priceDiffPercent)

	// Determine if the update was successful
	success := priceDiffPercent < 0.01 // Less than 0.01% difference

	fmt.Println("\n🎯 VERIFICATION RESULTS")
	fmt.Println("============================================================")

	if success {
		fmt.Println("✅ SUCCESS: API and Blockchain prices match!")
		fmt.Printf("✅ Price accuracy: %.4f%% (within 0.01%% tolerance)\n", priceDiffPercent)
	} else {
		fmt.Println("❌ FAILURE: API and Blockchain prices don't match!")
		fmt.Printf("❌ Price difference exceeds tolerance: %.4f%%\n", priceDiffPercent)
	}

	// Additional verification details
	fmt.Println("\n📋 DETAILED VERIFICATION")
	fmt.Println("----------------------------------------")
	fmt.Printf("🕐 Update timestamp: %s\n", finalTime.Format("15:04:05"))
	fmt.Printf("🔢 Round ID change: %d → %d\n", initialRound, finalRound)
	fmt.Printf("⏱️  Update duration: %v\n", updateDuration)

	// Check if price is stale
	isStale, err := blockchainClient.IsPriceStale()
	if err != nil {
		log.Printf("⚠️ Failed to check staleness: %v", err)
	} else {
		fmt.Printf("🕰️  Price staleness: %v\n", isStale)
	}

	fmt.Println("\n🎉 INTEGRATION TEST COMPLETED")
	fmt.Println("============================================================")

	if success {
		fmt.Println("🚀 Your Go backend is successfully updating the smart contract!")
		fmt.Println("✅ API prices are being correctly transmitted to the blockchain")
		fmt.Println("✅ Real-time price synchronization is working")
	} else {
		fmt.Println("⚠️  There may be an issue with price synchronization")
		fmt.Println("🔧 Check the price normalization logic or blockchain client")
	}
}
