# Web3 Portfolio Dashboard - Complete Test Setup Script
# This script sets up the entire application with Docker for end-to-end testing

Write-Host "🚀 Setting up Web3 Portfolio Dashboard for Testing..." -ForegroundColor Green

# Check if Docker is running
Write-Host "📋 Checking Docker status..." -ForegroundColor Yellow
try {
    docker version | Out-Null
    Write-Host "✅ Docker is running" -ForegroundColor Green
} catch {
    Write-Host "❌ Docker is not running. Please start Docker Desktop first." -ForegroundColor Red
    exit 1
}

# Stop any existing containers
Write-Host "🛑 Stopping existing containers..." -ForegroundColor Yellow
docker-compose -f docker-compose.dev.yml down --volumes --remove-orphans

# Build and start all services
Write-Host "🔨 Building and starting services..." -ForegroundColor Yellow
docker-compose -f docker-compose.dev.yml up --build -d

# Wait for services to be ready
Write-Host "⏳ Waiting for services to be ready..." -ForegroundColor Yellow
Start-Sleep -Seconds 30

# Check service health
Write-Host "🏥 Checking service health..." -ForegroundColor Yellow

# Check PostgreSQL
try {
    $pgStatus = docker exec web3-portfolio-postgres-dev pg_isready -U dev_user -d web3_portfolio_dev
    if ($pgStatus -like "*accepting connections*") {
        Write-Host "✅ PostgreSQL is ready" -ForegroundColor Green
    } else {
        Write-Host "❌ PostgreSQL is not ready" -ForegroundColor Red
    }
} catch {
    Write-Host "❌ Could not check PostgreSQL status" -ForegroundColor Red
}

# Check Redis
try {
    $redisStatus = docker exec web3-portfolio-redis-dev redis-cli ping
    if ($redisStatus -eq "PONG") {
        Write-Host "✅ Redis is ready" -ForegroundColor Green
    } else {
        Write-Host "❌ Redis is not ready" -ForegroundColor Red
    }
} catch {
    Write-Host "❌ Could not check Redis status" -ForegroundColor Red
}

# Check Backend API
try {
    $backendStatus = Invoke-WebRequest -Uri "http://localhost:8080/health" -UseBasicParsing -TimeoutSec 10
    if ($backendStatus.StatusCode -eq 200) {
        Write-Host "✅ Backend API is ready" -ForegroundColor Green
    } else {
        Write-Host "❌ Backend API returned status: $($backendStatus.StatusCode)" -ForegroundColor Red
    }
} catch {
    Write-Host "❌ Backend API is not ready yet" -ForegroundColor Red
}

# Check Frontend
try {
    $frontendStatus = Invoke-WebRequest -Uri "http://localhost:3000" -UseBasicParsing -TimeoutSec 10
    if ($frontendStatus.StatusCode -eq 200) {
        Write-Host "✅ Frontend is ready" -ForegroundColor Green
    } else {
        Write-Host "❌ Frontend returned status: $($frontendStatus.StatusCode)" -ForegroundColor Red
    }
} catch {
    Write-Host "❌ Frontend is not ready yet" -ForegroundColor Red
}

Write-Host ""
Write-Host "🎉 Setup Complete! Here's how to test the application:" -ForegroundColor Green
Write-Host ""
Write-Host "📱 Frontend Application:" -ForegroundColor Cyan
Write-Host "   URL: http://localhost:3000" -ForegroundColor White
Write-Host "   Features: Dashboard, Portfolio, Analytics, Alerts, Settings" -ForegroundColor White
Write-Host ""
Write-Host "🔧 Backend API:" -ForegroundColor Cyan
Write-Host "   URL: http://localhost:8080" -ForegroundColor White
Write-Host "   Health Check: http://localhost:8080/health" -ForegroundColor White
Write-Host "   API Docs: http://localhost:8080/swagger" -ForegroundColor White
Write-Host ""
Write-Host "🗄️  Database Management:" -ForegroundColor Cyan
Write-Host "   pgAdmin: http://localhost:5050" -ForegroundColor White
Write-Host "   Email: admin@web3portfolio.dev" -ForegroundColor White
Write-Host "   Password: admin123" -ForegroundColor White
Write-Host ""
Write-Host "🧪 Test Credentials:" -ForegroundColor Cyan
Write-Host "   Email: test@web3portfolio.dev" -ForegroundColor White
Write-Host "   Password: password" -ForegroundColor White
Write-Host ""
Write-Host "📊 Test Data Available:" -ForegroundColor Cyan
Write-Host "   - 3 test users with different subscription tiers" -ForegroundColor White
Write-Host "   - 4 wallets across Ethereum, Polygon, and BSC" -ForegroundColor White
Write-Host "   - 5 tokens: ETH, USDC, MATIC, BNB, LINK" -ForegroundColor White
Write-Host "   - 6 holdings with real market data" -ForegroundColor White
Write-Host "   - 4 sample transactions" -ForegroundColor White
Write-Host "   - 3 price alerts" -ForegroundColor White
Write-Host ""
Write-Host "🔍 Testing Steps:" -ForegroundColor Cyan
Write-Host "   1. Open http://localhost:3000 in your browser" -ForegroundColor White
Write-Host "   2. Login with test@web3portfolio.dev / password" -ForegroundColor White
Write-Host "   3. Explore the dashboard and portfolio views" -ForegroundColor White
Write-Host "   4. Test wallet connection and data retrieval" -ForegroundColor White
Write-Host "   5. Create and manage price alerts" -ForegroundColor White
Write-Host "   6. View analytics and transaction history" -ForegroundColor White
Write-Host ""
Write-Host "🛠️  Useful Commands:" -ForegroundColor Cyan
Write-Host "   View logs: docker-compose -f docker-compose.dev.yml logs -f" -ForegroundColor White
Write-Host "   Stop services: docker-compose -f docker-compose.dev.yml down" -ForegroundColor White
Write-Host "   Restart services: docker-compose -f docker-compose.dev.yml restart" -ForegroundColor White
Write-Host ""
Write-Host "🎯 Ready to test your Web3 Portfolio Dashboard!" -ForegroundColor Green 