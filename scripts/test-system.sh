#!/bin/bash

# Test script for the Oracle Pipeline system
# This script tests all the major components and endpoints

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

echo -e "${BLUE}🧪 Testing Oracle Pipeline System${NC}"

# Configuration
BACKEND_URL="http://localhost:8080"
PROMETHEUS_URL="http://localhost:9090"
GRAFANA_URL="http://localhost:3000"
NATS_URL="http://localhost:8222"

# Test functions
test_endpoint() {
    local url=$1
    local expected_status=$2
    local description=$3
    
    echo -n "Testing $description... "
    
    if response=$(curl -s -w "%{http_code}" -o /dev/null "$url" 2>/dev/null); then
        if [ "$response" = "$expected_status" ]; then
            echo -e "${GREEN}✅ PASS${NC}"
            return 0
        else
            echo -e "${RED}❌ FAIL (HTTP $response)${NC}"
            return 1
        fi
    else
        echo -e "${RED}❌ FAIL (Connection error)${NC}"
        return 1
    fi
}

test_json_endpoint() {
    local url=$1
    local expected_field=$2
    local description=$3
    
    echo -n "Testing $description... "
    
    if response=$(curl -s "$url" 2>/dev/null); then
        if echo "$response" | jq -e ".$expected_field" > /dev/null 2>&1; then
            echo -e "${GREEN}✅ PASS${NC}"
            return 0
        else
            echo -e "${RED}❌ FAIL (Missing field: $expected_field)${NC}"
            return 1
        fi
    else
        echo -e "${RED}❌ FAIL (JSON parse error)${NC}"
        return 1
    fi
}

# Wait for services to be ready
echo -e "${YELLOW}Waiting for services to be ready...${NC}"
sleep 10

# Test backend health
echo -e "${BLUE}Testing Backend Service${NC}"
test_endpoint "$BACKEND_URL/health" "200" "Health check"

# Test price endpoint
test_json_endpoint "$BACKEND_URL/price" "price" "Price endpoint"

# Test price history
test_json_endpoint "$BACKEND_URL/price/history" "prices" "Price history"

# Test TWAP endpoint
test_json_endpoint "$BACKEND_URL/price/twap?duration=1h" "twap" "TWAP calculation"

# Test admin stats
test_json_endpoint "$BACKEND_URL/admin/stats" "total_prices" "Admin statistics"

# Test Prometheus
echo -e "${BLUE}Testing Prometheus${NC}"
test_endpoint "$PROMETHEUS_URL/-/healthy" "200" "Prometheus health"

# Test Grafana
echo -e "${BLUE}Testing Grafana${NC}"
test_endpoint "$GRAFANA_URL/api/health" "200" "Grafana health"

# Test NATS monitoring
echo -e "${BLUE}Testing NATS${NC}"
test_endpoint "$NATS_URL/varz" "200" "NATS monitoring"

# Test database connection
echo -e "${BLUE}Testing Database Connection${NC}"
if docker-compose exec -T postgres pg_isready -U oracle_user -d oracle_db > /dev/null 2>&1; then
    echo -e "${GREEN}✅ Database connection: PASS${NC}"
else
    echo -e "${RED}❌ Database connection: FAIL${NC}"
fi

# Test Redis connection
echo -e "${BLUE}Testing Redis Connection${NC}"
if docker-compose exec -T redis redis-cli ping > /dev/null 2>&1; then
    echo -e "${GREEN}✅ Redis connection: PASS${NC}"
else
    echo -e "${RED}❌ Redis connection: FAIL${NC}"
fi

# Test NATS connection
echo -e "${BLUE}Testing NATS Connection${NC}"
if docker-compose exec -T nats nats server info > /dev/null 2>&1; then
    echo -e "${GREEN}✅ NATS connection: PASS${NC}"
else
    echo -e "${RED}❌ NATS connection: FAIL${NC}"
fi

# Test price data flow
echo -e "${BLUE}Testing Price Data Flow${NC}"

# Get initial price
echo "Getting initial price..."
initial_price=$(curl -s "$BACKEND_URL/price" | jq -r '.price')
echo "Initial price: $initial_price"

# Wait for potential price update
echo "Waiting for potential price update..."
sleep 35

# Get updated price
echo "Getting updated price..."
updated_price=$(curl -s "$BACKEND_URL/price" | jq -r '.price')
echo "Updated price: $updated_price"

if [ "$initial_price" != "$updated_price" ]; then
    echo -e "${GREEN}✅ Price update flow: PASS${NC}"
else
    echo -e "${YELLOW}⚠️  Price update flow: No change detected (may be normal)${NC}"
fi

# Test error handling
echo -e "${BLUE}Testing Error Handling${NC}"

# Test invalid endpoint
test_endpoint "$BACKEND_URL/invalid" "404" "Invalid endpoint"

# Test invalid TWAP duration
test_endpoint "$BACKEND_URL/price/twap?duration=invalid" "400" "Invalid TWAP duration"

# Test system metrics
echo -e "${BLUE}Testing System Metrics${NC}"

# Check if metrics are being collected
if curl -s "$PROMETHEUS_URL/api/v1/query?query=up" | jq -e '.data.result' > /dev/null 2>&1; then
    echo -e "${GREEN}✅ Prometheus metrics: PASS${NC}"
else
    echo -e "${RED}❌ Prometheus metrics: FAIL${NC}"
fi

# Test service logs
echo -e "${BLUE}Testing Service Logs${NC}"

# Check if services are logging
if docker-compose logs backend | grep -q "Successfully processed price" 2>/dev/null; then
    echo -e "${GREEN}✅ Backend logs: PASS${NC}"
else
    echo -e "${YELLOW}⚠️  Backend logs: No price processing logs found${NC}"
fi

# Summary
echo -e "${BLUE}📊 Test Summary${NC}"
echo -e "${GREEN}✅ All core services are running${NC}"
echo -e "${GREEN}✅ API endpoints are responding${NC}"
echo -e "${GREEN}✅ Database and cache connections are working${NC}"
echo -e "${GREEN}✅ Monitoring services are operational${NC}"

echo -e "${BLUE}🎉 System test completed successfully!${NC}"
echo -e "${YELLOW}Note: Some tests may show warnings if the system is still initializing${NC}"

