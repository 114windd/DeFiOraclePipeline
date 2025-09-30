#!/bin/bash

# Quick Start Script for Oracle Pipeline
# Simplified version for quick testing

set -e

# Colors
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m'

echo -e "${BLUE}🚀 Quick Start - Oracle Pipeline${NC}"

# Set private key
export ETH_PRIVATE_KEY=""

echo -e "${YELLOW}Setting up environment...${NC}"

# Create monitoring directories
mkdir -p monitoring/grafana/dashboards
mkdir -p monitoring/grafana/datasources
mkdir -p monitoring/nginx
mkdir -p scripts

# Ensure directories are writable
chmod 755 monitoring/grafana/dashboards
chmod 755 monitoring/grafana/datasources
chmod 755 monitoring/nginx
chmod 755 scripts

# Create minimal Prometheus config
cat > monitoring/prometheus.yml << 'EOF'
global:
  scrape_interval: 15s

scrape_configs:
  - job_name: 'prometheus'
    static_configs:
      - targets: ['localhost:9090']
  - job_name: 'backend'
    static_configs:
      - targets: ['backend:8080']
    metrics_path: '/metrics'
EOF

# Create minimal Grafana datasource
cat > monitoring/grafana/datasources/prometheus.yml << 'EOF'
apiVersion: 1
datasources:
  - name: Prometheus
    type: prometheus
    access: proxy
    url: http://prometheus:9090
    isDefault: true
EOF

# Create minimal dashboard config
cat > monitoring/grafana/dashboards/dashboard.yml << 'EOF'
apiVersion: 1
providers:
  - name: 'Oracle Pipeline'
    orgId: 1
    type: file
    options:
      path: /etc/grafana/provisioning/dashboards
EOF

# Create Nginx config
cat > monitoring/nginx/nginx.conf << 'EOF'
events { worker_connections 1024; }
http {
    upstream backend { server backend:8080; }
    upstream prometheus { server prometheus:9090; }
    upstream grafana { server grafana:3000; }
    
    server {
        listen 80;
        location /api/ { proxy_pass http://backend/; }
        location /metrics/ { proxy_pass http://prometheus/; }
        location /grafana/ { proxy_pass http://grafana/; }
        location / { return 200 'Oracle Pipeline Running!'; add_header Content-Type text/plain; }
    }
}
EOF

# Note: Database initialization is handled by GORM AutoMigrate
# No manual SQL files needed - GORM will create tables automatically

echo -e "${GREEN}Configuration files created${NC}"

# Stop existing containers
echo -e "${YELLOW}Stopping existing containers...${NC}"
docker-compose down -v 2>/dev/null || true

# Start services
echo -e "${YELLOW}Starting services...${NC}"
docker-compose up -d

# Wait for services
echo -e "${YELLOW}Waiting for services to start...${NC}"
sleep 30

# Check health
echo -e "${YELLOW}Checking service health...${NC}"

if curl -f http://localhost:8080/health > /dev/null 2>&1; then
    echo -e "${GREEN}✅ Backend: Healthy${NC}"
else
    echo -e "${RED}❌ Backend: Not responding${NC}"
fi

if curl -f http://localhost:9090/-/healthy > /dev/null 2>&1; then
    echo -e "${GREEN}✅ Prometheus: Healthy${NC}"
else
    echo -e "${RED}❌ Prometheus: Not responding${NC}"
fi

if curl -f http://localhost:3000/api/health > /dev/null 2>&1; then
    echo -e "${GREEN}✅ Grafana: Healthy${NC}"
else
    echo -e "${RED}❌ Grafana: Not responding${NC}"
fi

# Test API
echo -e "${YELLOW}Testing API endpoints...${NC}"
if curl -s http://localhost:8080/price > /dev/null 2>&1; then
    echo -e "${GREEN}✅ Price API: Working${NC}"
    PRICE=$(curl -s http://localhost:8080/price)
    echo -e "${BLUE}Current price: $PRICE${NC}"
else
    echo -e "${RED}❌ Price API: Not working${NC}"
fi

echo ""
echo -e "${GREEN}🎉 Oracle Pipeline is running!${NC}"
echo ""
echo -e "${BLUE}📊 Services:${NC}"
echo -e "  Backend:    http://localhost:8080"
echo -e "  Prometheus: http://localhost:9090"
echo -e "  Grafana:    http://localhost:3000 (admin/admin)"
echo -e "  NATS:       http://localhost:8222"
echo ""
echo -e "${BLUE}📋 API Endpoints:${NC}"
echo -e "  GET /health        - Health check"
echo -e "  GET /price         - Latest ETH/USD price"
echo -e "  GET /price/history - Price history"
echo -e "  GET /metrics       - Prometheus metrics"
echo ""
echo -e "${YELLOW}Use 'docker-compose logs -f' to view logs${NC}"
echo -e "${YELLOW}Use 'docker-compose down' to stop services${NC}"
echo ""

# Show logs
echo -e "${BLUE}📋 Service Logs (Press Ctrl+C to exit):${NC}"
docker-compose logs -f
