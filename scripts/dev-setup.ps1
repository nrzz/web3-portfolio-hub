# Web3 Portfolio Dashboard - Development Setup Script
# This script helps set up the development environment

Write-Host "🚀 Setting up Web3 Portfolio Dashboard Development Environment" -ForegroundColor Green

# Check if we're in the right directory
if (-not (Test-Path "backend\main.go")) {
    Write-Host "❌ Error: Please run this script from the project root directory" -ForegroundColor Red
    exit 1
}

# Step 1: Check if backend is already running and kill it
Write-Host "📋 Checking for existing backend processes..." -ForegroundColor Yellow
$backendProcess = Get-Process -Name "go" -ErrorAction SilentlyContinue | Where-Object { $_.ProcessName -eq "go" }
if ($backendProcess) {
    Write-Host "🔄 Stopping existing backend process..." -ForegroundColor Yellow
    Stop-Process -Name "go" -Force -ErrorAction SilentlyContinue
    Start-Sleep -Seconds 2
}

# Step 2: Check if frontend is already running and kill it
Write-Host "📋 Checking for existing frontend processes..." -ForegroundColor Yellow
$nodeProcesses = Get-Process -Name "node" -ErrorAction SilentlyContinue
if ($nodeProcesses) {
    Write-Host "🔄 Stopping existing Node.js processes..." -ForegroundColor Yellow
    Stop-Process -Name "node" -Force -ErrorAction SilentlyContinue
    Start-Sleep -Seconds 2
}

# Step 3: Install backend dependencies
Write-Host "📦 Installing Go dependencies..." -ForegroundColor Yellow
Set-Location backend
go mod tidy
if ($LASTEXITCODE -ne 0) {
    Write-Host "❌ Failed to install Go dependencies" -ForegroundColor Red
    exit 1
}

# Step 4: Install frontend dependencies
Write-Host "📦 Installing frontend dependencies..." -ForegroundColor Yellow
Set-Location ..\frontend
npm install
if ($LASTEXITCODE -ne 0) {
    Write-Host "❌ Failed to install frontend dependencies" -ForegroundColor Red
    exit 1
}

# Step 5: Test database connection
Write-Host "🗄️ Testing database connection..." -ForegroundColor Yellow
Set-Location ..\backend
go run main.go --test-db
if ($LASTEXITCODE -ne 0) {
    Write-Host "❌ Database connection failed. Please check your PostgreSQL installation and connection string in backend\env.dev" -ForegroundColor Red
    Write-Host "💡 Make sure PostgreSQL is installed and running" -ForegroundColor Cyan
    exit 1
}

# Step 6: Run database migrations
Write-Host "🗄️ Running database migrations..." -ForegroundColor Yellow
go run main.go --migrate
if ($LASTEXITCODE -ne 0) {
    Write-Host "❌ Database migration failed" -ForegroundColor Red
    exit 1
}

# Step 7: Start backend in background
Write-Host "🚀 Starting backend server..." -ForegroundColor Green
Start-Process -FilePath "go" -ArgumentList "run", "main.go" -WorkingDirectory "D:\Projects\web3-portfolio-dashboard\backend" -WindowStyle Hidden

# Step 8: Wait for backend to start
Write-Host "⏳ Waiting for backend to start..." -ForegroundColor Yellow
Start-Sleep -Seconds 5

# Step 9: Test backend health
Write-Host "🏥 Testing backend health..." -ForegroundColor Yellow
try {
    $response = Invoke-RestMethod -Uri "http://localhost:8080/health" -Method GET -TimeoutSec 10
    Write-Host "✅ Backend is running and healthy!" -ForegroundColor Green
} catch {
    Write-Host "❌ Backend health check failed. Backend might not be running properly." -ForegroundColor Red
    Write-Host "💡 You can manually start the backend with: cd backend; go run main.go" -ForegroundColor Cyan
}

# Step 10: Start frontend
Write-Host "🚀 Starting frontend server..." -ForegroundColor Green
Set-Location ..\frontend
Start-Process -FilePath "npm" -ArgumentList "run", "dev" -WorkingDirectory "D:\Projects\web3-portfolio-dashboard\frontend" -WindowStyle Hidden

# Step 11: Wait for frontend to start
Write-Host "⏳ Waiting for frontend to start..." -ForegroundColor Yellow
Start-Sleep -Seconds 10

# Step 12: Open browser
Write-Host "🌐 Opening application in browser..." -ForegroundColor Green
Start-Process "http://localhost:3000"

Write-Host "✅ Development environment setup complete!" -ForegroundColor Green
Write-Host "📋 Backend: http://localhost:8080" -ForegroundColor Cyan
Write-Host "📋 Frontend: http://localhost:3000" -ForegroundColor Cyan
Write-Host "📋 Health Check: http://localhost:8080/health" -ForegroundColor Cyan
Write-Host "📋 API Docs: http://localhost:8080/api/v1" -ForegroundColor Cyan

Write-Host "`n💡 To stop the servers, use Ctrl+C or close the terminal windows" -ForegroundColor Yellow
Write-Host "💡 To restart, run this script again" -ForegroundColor Yellow

# Keep the script running to show the status
Write-Host "`n🔄 Development servers are running. Press Ctrl+C to stop..." -ForegroundColor Green
try {
    while ($true) {
        Start-Sleep -Seconds 30
        Write-Host "✅ Servers are still running..." -ForegroundColor Green
    }
} catch {
    Write-Host "`n🛑 Stopping development servers..." -ForegroundColor Yellow
    Stop-Process -Name "go" -Force -ErrorAction SilentlyContinue
    Stop-Process -Name "node" -Force -ErrorAction SilentlyContinue
    Write-Host "✅ Development servers stopped" -ForegroundColor Green
} 