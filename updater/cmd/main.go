package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/114windd/DeFiOraclePipeline.git/updater/pkg/ethclient"
	"github.com/114windd/DeFiOraclePipeline.git/updater/pkg/updater"
)

func main() {
	// Command line flags
	var (
		natsURL    = flag.String("nats-url", "nats://localhost:4222", "NATS server URL")
		subject    = flag.String("subject", "prices.ethusd", "NATS subject to subscribe to")
		ethRPCURL  = flag.String("eth-rpc", "https://eth-sepolia.g.alchemy.com/v2/6ChCkEoo-jvGgoa85eb9G", "Ethereum RPC URL")
		privateKey = flag.String("private-key", "", "Private key for transaction signing")
		gasLimit   = flag.Uint64("gas-limit", 200000, "Gas limit for transactions")
		threshold  = flag.Float64("threshold", 0.005, "Price change threshold (0.005 = 0.5%)")
	)
	flag.Parse()

	// Validate required parameters
	if *privateKey == "" {
		log.Fatal("Private key is required. Use -private-key flag or set ETH_PRIVATE_KEY environment variable")
	}

	// Get private key from environment if not provided via flag
	if *privateKey == "" {
		*privateKey = os.Getenv("ETH_PRIVATE_KEY")
		if *privateKey == "" {
			log.Fatal("Private key must be provided via -private-key flag or ETH_PRIVATE_KEY environment variable")
		}
	}

	// Initialize Ethereum client
	ethClient, err := ethclient.NewEthClient(*ethRPCURL, *privateKey, *gasLimit)
	if err != nil {
		log.Fatalf("Failed to initialize Ethereum client: %v", err)
	}
	defer ethClient.Close()

	// Log account information
	log.Printf("Using Ethereum address: %s", ethClient.GetAddress().Hex())

	// Check balance
	balance, err := ethClient.GetBalance()
	if err != nil {
		log.Printf("Warning: Failed to get balance: %v", err)
	} else {
		log.Printf("Account balance: %s ETH", balance.String())
	}

	// Initialize updater
	updater, err := updater.NewUpdater(*natsURL, *subject, ethClient, *threshold)
	if err != nil {
		log.Fatalf("Failed to initialize updater: %v", err)
	}

	// Set up signal handling for graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// Start updater in a goroutine
	go func() {
		if err := updater.Start(); err != nil {
			log.Fatalf("Updater failed: %v", err)
		}
	}()

	log.Printf("Updater worker started. Listening for price updates on subject: %s", *subject)
	log.Printf("Price change threshold: %.2f%%", *threshold*100)

	// Wait for shutdown signal
	<-quit
	log.Println("Shutting down updater worker...")

	// Stop the updater
	updater.Stop()
	log.Println("Updater worker stopped")
}

