#!/bin/bash

# Simple deployment script for DeFi Oracle Pipeline
# Make sure Anvil is running on localhost:8545

echo "🚀 Deploying DeFi Oracle Pipeline contracts to Anvil..."

# Deploy contracts
forge script script/Deploy.s.sol --rpc-url http://localhost:8545 --broadcast --private-key $ANVIL_PRIVATE_KEY

echo "✅ Deployment complete!"
echo ""
echo "📋 Contract Addresses:"
echo "Check the deployment output above for contract addresses"
echo ""
echo "🔗 Anvil RPC URL: http://localhost:8545"
echo "🔗 Chain ID: 31337"

