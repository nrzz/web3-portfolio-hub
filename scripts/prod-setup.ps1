# Web3 Portfolio Dashboard - Production Setup Script (PowerShell)
# This script sets up the production environment with PostgreSQL and Redis

Write-Host "🚀 Setting up Web3 Portfolio Dashboard Production Environment..." -ForegroundColor Green

# Check if Docker is running
try {
    docker info | Out-Null
    Write-Host "✅ Docker is running" -ForegroundColor Green
} catch {
    Write-Host "❌ Docker is not running. Please start Docker and try again." -ForegroundColor Red
    exit 1
}

# Check if Docker Compose is available
try {
    docker-compose --version | Out-Null
    Write-Host "✅ Docker Compose is available" -ForegroundColor Green
} catch {
    Write-Host "❌ Docker Compose is not installed. Please install Docker Compose and try again." -ForegroundColor Red
    exit 1
}

# Check if environment file exists
if (-not (Test-Path "backend\env.production")) {
    Write-Host "❌ Production environment file not found. Please create backend\env.production" -ForegroundColor Red
    exit 1
}

Write-Host "✅ Docker and Docker Compose are available" -ForegroundColor Green

# Create necessary directories
Write-Host "📁 Creating necessary directories..." -ForegroundColor Yellow
New-Item -ItemType Directory -Force -Path "logs" | Out-Null
New-Item -ItemType Directory -Force -Path "ssl" | Out-Null
New-Item -ItemType Directory -Force -Path "data\postgres" | Out-Null
New-Item -ItemType Directory -Force -Path "data\redis" | Out-Null

# Copy environment file
Write-Host "📝 Setting up production environment..." -ForegroundColor Yellow
Copy-Item "backend\env.production" "backend\.env"

# Stop any existing containers
Write-Host "🛑 Stopping any existing containers..." -ForegroundColor Yellow
docker-compose -f docker-compose.prod.yml down --remove-orphans

# Build and start the services
Write-Host "🔨 Building and starting production services..." -ForegroundColor Yellow
docker-compose -f docker-compose.prod.yml up --build -d

# Wait for services to be healthy
Write-Host "⏳ Waiting for services to be ready..." -ForegroundColor Yellow
Start-Sleep -Seconds 15

# Check service health
Write-Host "🔍 Checking service health..." -ForegroundColor Yellow

# Check PostgreSQL
try {
    docker-compose -f docker-compose.prod.yml exec -T postgres pg_isready -U portfolio_user -d web3_portfolio | Out-Null
    Write-Host "✅ PostgreSQL is ready" -ForegroundColor Green
} catch {
    Write-Host "❌ PostgreSQL is not ready. Check logs with: docker-compose -f docker-compose.prod.yml logs postgres" -ForegroundColor Red
}

# Check Redis
try {
    docker-compose -f docker-compose.prod.yml exec -T redis redis-cli ping | Out-Null
    Write-Host "✅ Redis is ready" -ForegroundColor Green
} catch {
    Write-Host "❌ Redis is not ready. Check logs with: docker-compose -f docker-compose.prod.yml logs redis" -ForegroundColor Red
}

# Check Backend
try {
    Invoke-WebRequest -Uri "http://localhost:8080/health" -UseBasicParsing | Out-Null
    Write-Host "✅ Backend API is ready" -ForegroundColor Green
} catch {
    Write-Host "⏳ Backend API is starting up..." -ForegroundColor Yellow
    Write-Host "   You can check the logs with: docker-compose -f docker-compose.prod.yml logs backend" -ForegroundColor Gray
}

# Check Frontend
try {
    Invoke-WebRequest -Uri "http://localhost:3000" -UseBasicParsing | Out-Null
    Write-Host "✅ Frontend is ready" -ForegroundColor Green
} catch {
    Write-Host "⏳ Frontend is starting up..." -ForegroundColor Yellow
    Write-Host "   You can check the logs with: docker-compose -f docker-compose.prod.yml logs frontend" -ForegroundColor Gray
}

# Check Nginx
try {
    Invoke-WebRequest -Uri "http://localhost:80" -UseBasicParsing | Out-Null
    Write-Host "✅ Nginx is ready" -ForegroundColor Green
} catch {
    Write-Host "⏳ Nginx is starting up..." -ForegroundColor Yellow
    Write-Host "   You can check the logs with: docker-compose -f docker-compose.prod.yml logs nginx" -ForegroundColor Gray
}

Write-Host ""
Write-Host "🎉 Production environment setup complete!" -ForegroundColor Green
Write-Host ""
Write-Host "📊 Services:" -ForegroundColor Cyan
Write-Host "   • PostgreSQL: localhost:5432" -ForegroundColor White
Write-Host "   • Redis: localhost:6379" -ForegroundColor White
Write-Host "   • Backend API: http://localhost:8080" -ForegroundColor White
Write-Host "   • Frontend: http://localhost:3000" -ForegroundColor White
Write-Host "   • Nginx (Reverse Proxy): http://localhost:80" -ForegroundColor White
Write-Host ""
Write-Host "🔧 Useful commands:" -ForegroundColor Cyan
Write-Host "   • View logs: docker-compose -f docker-compose.prod.yml logs -f" -ForegroundColor White
Write-Host "   • Stop services: docker-compose -f docker-compose.prod.yml down" -ForegroundColor White
Write-Host "   • Restart services: docker-compose -f docker-compose.prod.yml restart" -ForegroundColor White
Write-Host "   • Access PostgreSQL: docker-compose -f docker-compose.prod.yml exec postgres psql -U portfolio_user -d web3_portfolio" -ForegroundColor White
Write-Host "   • Access Redis: docker-compose -f docker-compose.prod.yml exec redis redis-cli" -ForegroundColor White
Write-Host ""
Write-Host "📝 Database credentials:" -ForegroundColor Cyan
Write-Host "   • Database: web3_portfolio" -ForegroundColor White
Write-Host "   • Username: portfolio_user" -ForegroundColor White
Write-Host "   • Password: secure_password_123" -ForegroundColor White
Write-Host ""
Write-Host "🌐 Open your browser and navigate to:" -ForegroundColor Cyan
Write-Host "   • Frontend: http://localhost:3000" -ForegroundColor White
Write-Host "   • API Documentation: http://localhost:8080/docs" -ForegroundColor White
Write-Host "   • Health Check: http://localhost:8080/health" -ForegroundColor White
Write-Host ""
Write-Host "🔒 Security Notes:" -ForegroundColor Cyan
Write-Host "   • Change default passwords in production" -ForegroundColor White
Write-Host "   • Set up SSL certificates for HTTPS" -ForegroundColor White
Write-Host "   • Configure firewall rules" -ForegroundColor White
Write-Host "   • Set up monitoring and logging" -ForegroundColor White
Write-Host "   • Regular database backups" -ForegroundColor White 