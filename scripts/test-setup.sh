#!/bin/bash

# Web3 Portfolio Dashboard - Complete Test Setup Script
# This script sets up the entire application with Docker for end-to-end testing

echo "🚀 Setting up Web3 Portfolio Dashboard for Testing..."

# Check if Docker is running
echo "📋 Checking Docker status..."
if ! docker info > /dev/null 2>&1; then
    echo "❌ Docker is not running. Please start Docker first."
    exit 1
fi
echo "✅ Docker is running"

# Stop any existing containers
echo "🛑 Stopping existing containers..."
docker-compose -f docker-compose.dev.yml down --volumes --remove-orphans

# Build and start all services
echo "🔨 Building and starting services..."
docker-compose -f docker-compose.dev.yml up --build -d

# Wait for services to be ready
echo "⏳ Waiting for services to be ready..."
sleep 30

# Check service health
echo "🏥 Checking service health..."

# Check PostgreSQL
if docker exec web3-portfolio-postgres-dev pg_isready -U dev_user -d web3_portfolio_dev > /dev/null 2>&1; then
    echo "✅ PostgreSQL is ready"
else
    echo "❌ PostgreSQL is not ready"
fi

# Check Redis
if docker exec web3-portfolio-redis-dev redis-cli ping | grep -q "PONG"; then
    echo "✅ Redis is ready"
else
    echo "❌ Redis is not ready"
fi

# Check Backend API
if curl -f http://localhost:8080/health > /dev/null 2>&1; then
    echo "✅ Backend API is ready"
else
    echo "❌ Backend API is not ready yet"
fi

# Check Frontend
if curl -f http://localhost:3000 > /dev/null 2>&1; then
    echo "✅ Frontend is ready"
else
    echo "❌ Frontend is not ready yet"
fi

echo ""
echo "🎉 Setup Complete! Here's how to test the application:"
echo ""
echo "📱 Frontend Application:"
echo "   URL: http://localhost:3000"
echo "   Features: Dashboard, Portfolio, Analytics, Alerts, Settings"
echo ""
echo "🔧 Backend API:"
echo "   URL: http://localhost:8080"
echo "   Health Check: http://localhost:8080/health"
echo "   API Docs: http://localhost:8080/swagger"
echo ""
echo "🗄️  Database Management:"
echo "   pgAdmin: http://localhost:5050"
echo "   Email: admin@web3portfolio.dev"
echo "   Password: admin123"
echo ""
echo "🧪 Test Credentials:"
echo "   Email: test@web3portfolio.dev"
echo "   Password: password"
echo ""
echo "📊 Test Data Available:"
echo "   - 3 test users with different subscription tiers"
echo "   - 4 wallets across Ethereum, Polygon, and BSC"
echo "   - 5 tokens (ETH, USDC, MATIC, BNB, LINK)"
echo "   - 6 holdings with real market data"
echo "   - 4 sample transactions"
echo "   - 3 price alerts"
echo ""
echo "🔍 Testing Steps:"
echo "   1. Open http://localhost:3000 in your browser"
echo "   2. Login with test@web3portfolio.dev / password"
echo "   3. Explore the dashboard and portfolio views"
echo "   4. Test wallet connection and data retrieval"
echo "   5. Create and manage price alerts"
echo "   6. View analytics and transaction history"
echo ""
echo "🛠️  Useful Commands:"
echo "   View logs: docker-compose -f docker-compose.dev.yml logs -f"
echo "   Stop services: docker-compose -f docker-compose.dev.yml down"
echo "   Restart services: docker-compose -f docker-compose.dev.yml restart"
echo ""
echo "🎯 Ready to test your Web3 Portfolio Dashboard!" 