# ETH/USD Oracle Pipeline

A complete decentralized-style oracle pipeline that fetches off-chain ETH/USD market data, processes it through a Go backend, and publishes it to an on-chain Solidity contract. This enables other DeFi smart contracts to query reliable ETH/USD prices.

## 🏗️ Architecture

The system consists of several interconnected components:

- **Go Backend**: Fetches ETH/USD prices from external APIs, normalizes them, caches in Redis, and publishes to NATS
- **Oracle Updater Worker**: Consumes price updates from NATS and submits them to the blockchain
- **Solidity Contracts**: Oracle contract for price storage and ConsumerDemo for stop-loss/price alerts
- **Infrastructure**: PostgreSQL, Redis, NATS.io, Prometheus, Grafana

## 🚀 Quick Start

### Prerequisites

- Docker and Docker Compose
- Go 1.21+
- Foundry (for contract development)
- Private key for Ethereum transactions

### 1. Clone and Setup

```bash
git clone <repository-url>
cd defioraclepipeline
```

### 2. Set Environment Variables

```bash
export ETH_PRIVATE_KEY="0x..."  # Your private key for blockchain transactions
export ETHERSCAN_API_KEY="..."  # Optional: for contract verification
```

### 3. Run the System

```bash
# Start all services
./deployments/run-system.sh

# Or manually with Docker Compose
docker-compose up -d
```

### 4. Access Services

- **Backend API**: http://localhost:8080
- **Prometheus**: http://localhost:9090
- **Grafana**: http://localhost:3000 (admin/admin)
- **NATS Monitoring**: http://localhost:8222

## 📊 API Endpoints

### Health & Status
- `GET /health` - Health check
- `GET /admin/stats` - System statistics

### Price Data
- `GET /price` - Latest ETH/USD price
- `GET /price/history?limit=100` - Price history
- `GET /price/twap?duration=1h` - Time-weighted average price

### Monitoring
- `GET /metrics` - Prometheus metrics

## 🔧 Development

### Backend Development

```bash
cd backend
go mod tidy
go run cmd/main.go
```

### Contract Development

```bash
cd contracts
forge build
forge test
forge coverage
```

### Deploy Contracts

```bash
# Deploy to Sepolia testnet
./deployments/deploy-contracts.sh
```

## 🏛️ Smart Contracts

### Oracle Contract

The main oracle contract that stores ETH/USD price data:

```solidity
// Update price (only updater can call)
function updatePrice(uint256 newPrice) external;

// Get latest price
function getLatestPrice() external view returns (uint256, uint256, uint256);

// Access control
function setUpdater(address newUpdater) external;
function pause() external;
function unpause() external;
```

### ConsumerDemo Contract

Demonstrates usage with stop-loss and price alert functionality:

```solidity
// Create a position with stop-loss and alert
function createPosition(uint256 stopLossPrice, uint256 alertPrice) external payable;

// Check if position should be liquidated
function isLiquidatable(address user) external view returns (bool);

// Get collateral value in USD
function getCollateralValue(address user) external view returns (uint256);
```

## 🔍 Monitoring

### Prometheus Metrics

The system exposes comprehensive metrics:

- `price_fetch_duration_seconds` - API fetch latency
- `price_fetch_errors_total` - Fetch error count
- `price_updates_total` - Successful price updates
- `cache_hits_total` - Cache hit rate
- `database_operations_total` - DB operation count

### Grafana Dashboards

Pre-configured dashboards for:
- Price data visualization
- System performance metrics
- Error rates and alerts
- Infrastructure health

## 🐳 Docker Services

| Service | Port | Description |
|---------|------|-------------|
| backend | 8080 | Go API service |
| postgres | 5432 | Database |
| redis | 6379 | Cache |
| nats | 4222 | Message queue |
| prometheus | 9090 | Metrics collection |
| grafana | 3000 | Monitoring dashboard |
| nginx | 80 | Reverse proxy |

## 🔐 Security Features

- **Access Control**: Only authorized updater can update prices
- **Input Validation**: Price bounds and staleness checks
- **Pause Mechanism**: Emergency stop functionality
- **Private Key Management**: Secure transaction signing
- **Rate Limiting**: API request throttling

## 🧪 Testing

### Unit Tests

```bash
# Backend tests
cd backend
go test ./...

# Contract tests
cd contracts
forge test
```

### Integration Tests

```bash
# Run full system test
docker-compose up -d
./scripts/test-system.sh
```

## 📈 Performance

- **Price Fetch Interval**: 30 seconds (configurable)
- **Price Change Threshold**: 0.5% (configurable)
- **Cache TTL**: 1 hour
- **Database Retention**: 30 days
- **Gas Limit**: 200,000 (configurable)

## 🚨 Alerts

The system includes alerts for:
- Stale price data (>1 hour old)
- Failed price fetches
- Database connection issues
- High error rates
- Low cache hit rates

## 🔧 Configuration

### Environment Variables

| Variable | Default | Description |
|----------|---------|-------------|
| `SERVER_PORT` | 8080 | API server port |
| `DATABASE_URL` | postgres://... | Database connection |
| `REDIS_URL` | redis://... | Redis connection |
| `NATS_URL` | nats://... | NATS connection |
| `COINGECKO_URL` | https://api.coingecko.com/... | Price API URL |
| `FETCH_INTERVAL` | 30s | Price fetch interval |
| `PRICE_CHANGE_THRESHOLD` | 0.005 | Price change threshold |
| `ETH_PRIVATE_KEY` | - | Private key for transactions |

## 🛠️ Troubleshooting

### Common Issues

1. **Services not starting**: Check Docker logs with `docker-compose logs`
2. **Price not updating**: Verify NATS connection and updater service
3. **Database errors**: Check PostgreSQL connection and migrations
4. **Contract deployment fails**: Verify private key and RPC URL

### Logs

```bash
# View all logs
docker-compose logs -f

# View specific service logs
docker-compose logs -f backend
docker-compose logs -f updater
```

## 📚 Documentation

- [API Documentation](docs/api.md)
- [Contract Documentation](docs/contracts.md)
- [Deployment Guide](docs/deployment.md)
- [Monitoring Guide](docs/monitoring.md)

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests
5. Submit a pull request

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

- [OpenZeppelin](https://openzeppelin.com/) for secure contract libraries
- [Foundry](https://book.getfoundry.sh/) for smart contract development
- [CoinGecko](https://coingecko.com/) for price data
- [Alchemy](https://alchemy.com/) for Ethereum RPC

## 📞 Support

For support and questions:
- Create an issue on GitHub
- Check the documentation
- Review the troubleshooting guide

---

**⚠️ Disclaimer**: This is an educational project. Do not use in production without proper security audits and testing.
