# Development Deployment Script for PowerShell
param(
    [switch]$Clean
)

Write-Host "🚀 Starting Web3 Portfolio Dashboard Development Deployment..." -ForegroundColor Green

# Check if Docker is running
try {
    docker info | Out-Null
} catch {
    Write-Host "❌ Docker is not running. Please start Docker and try again." -ForegroundColor Red
    exit 1
}

# Check if docker-compose is available
try {
    docker-compose --version | Out-Null
} catch {
    Write-Host "❌ docker-compose is not installed. Please install it and try again." -ForegroundColor Red
    exit 1
}

# Stop any existing containers
Write-Host "🛑 Stopping existing containers..." -ForegroundColor Yellow
docker-compose -f docker-compose.dev.yml down --remove-orphans

# Remove old volumes if requested
if ($Clean) {
    Write-Host "🧹 Cleaning up old volumes..." -ForegroundColor Yellow
    docker-compose -f docker-compose.dev.yml down -v
}

# Build and start services
Write-Host "🔨 Building and starting services..." -ForegroundColor Yellow
docker-compose -f docker-compose.dev.yml up --build -d

# Wait for services to be ready
Write-Host "⏳ Waiting for services to be ready..." -ForegroundColor Yellow
Start-Sleep -Seconds 30

# Check service health
Write-Host "🏥 Checking service health..." -ForegroundColor Yellow
docker-compose -f docker-compose.dev.yml ps

# Show logs
Write-Host "📋 Recent logs:" -ForegroundColor Yellow
docker-compose -f docker-compose.dev.yml logs --tail=20

Write-Host "✅ Development deployment completed!" -ForegroundColor Green
Write-Host ""
Write-Host "🌐 Services available at:" -ForegroundColor Cyan
Write-Host "   Frontend: http://localhost:3000" -ForegroundColor White
Write-Host "   Backend API: http://localhost:8080" -ForegroundColor White
Write-Host "   pgAdmin: http://localhost:5050 (admin@web3portfolio.dev / admin123)" -ForegroundColor White
Write-Host ""
Write-Host "📊 To view logs: docker-compose -f docker-compose.dev.yml logs -f" -ForegroundColor Gray
Write-Host "🛑 To stop: docker-compose -f docker-compose.dev.yml down" -ForegroundColor Gray 