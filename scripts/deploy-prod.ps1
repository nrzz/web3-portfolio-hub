# Production Deployment Script for PowerShell
param(
    [switch]$Clean
)

Write-Host "🚀 Starting Web3 Portfolio Dashboard Production Deployment..." -ForegroundColor Green

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

# Check if .env.production exists
if (-not (Test-Path ".env.production")) {
    Write-Host "❌ .env.production file not found!" -ForegroundColor Red
    Write-Host "📝 Please copy env.production.example to .env.production and update the values." -ForegroundColor Yellow
    exit 1
}

# Load environment variables
Write-Host "📋 Loading environment variables..." -ForegroundColor Yellow
Get-Content ".env.production" | ForEach-Object {
    if ($_ -match "^([^#][^=]+)=(.*)$") {
        [Environment]::SetEnvironmentVariable($matches[1], $matches[2], "Process")
    }
}

# Validate required environment variables
$requiredVars = @("POSTGRES_PASSWORD", "REDIS_PASSWORD", "JWT_SECRET")
foreach ($var in $requiredVars) {
    if ([string]::IsNullOrEmpty([Environment]::GetEnvironmentVariable($var, "Process"))) {
        Write-Host "❌ Required environment variable $var is not set!" -ForegroundColor Red
        exit 1
    }
}

# Stop any existing containers
Write-Host "🛑 Stopping existing containers..." -ForegroundColor Yellow
docker-compose -f docker-compose.prod.yml down --remove-orphans

# Remove old volumes if requested
if ($Clean) {
    Write-Host "🧹 Cleaning up old volumes..." -ForegroundColor Yellow
    docker-compose -f docker-compose.prod.yml down -v
}

# Build and start services
Write-Host "🔨 Building and starting services..." -ForegroundColor Yellow
docker-compose -f docker-compose.prod.yml up --build -d

# Wait for services to be ready
Write-Host "⏳ Waiting for services to be ready..." -ForegroundColor Yellow
Start-Sleep -Seconds 60

# Check service health
Write-Host "🏥 Checking service health..." -ForegroundColor Yellow
docker-compose -f docker-compose.prod.yml ps

# Show logs
Write-Host "📋 Recent logs:" -ForegroundColor Yellow
docker-compose -f docker-compose.prod.yml logs --tail=20

Write-Host "✅ Production deployment completed!" -ForegroundColor Green
Write-Host ""
Write-Host "🌐 Services available at:" -ForegroundColor Cyan
Write-Host "   Frontend: http://localhost:3000" -ForegroundColor White
Write-Host "   Backend API: http://localhost:8080" -ForegroundColor White
Write-Host "   Nginx: http://localhost:80" -ForegroundColor White
Write-Host ""
Write-Host "📊 To view logs: docker-compose -f docker-compose.prod.yml logs -f" -ForegroundColor Gray
Write-Host "🛑 To stop: docker-compose -f docker-compose.prod.yml down" -ForegroundColor Gray
Write-Host ""
Write-Host "🔒 Security Notes:" -ForegroundColor Yellow
Write-Host "   - Change default passwords in production" -ForegroundColor White
Write-Host "   - Configure SSL certificates for HTTPS" -ForegroundColor White
Write-Host "   - Set up proper firewall rules" -ForegroundColor White
Write-Host "   - Monitor logs for security issues" -ForegroundColor White 