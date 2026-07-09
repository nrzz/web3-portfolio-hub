# PostgreSQL Setup Script for Web3 Portfolio Dashboard
# This script creates the database and user for local development

Write-Host "🗄️ Setting up PostgreSQL Database for Web3 Portfolio Dashboard" -ForegroundColor Green

# Check if PostgreSQL is installed
try {
    $psqlVersion = psql --version 2>$null
    if ($LASTEXITCODE -eq 0) {
        Write-Host "✅ PostgreSQL is installed" -ForegroundColor Green
        Write-Host "Version: $psqlVersion" -ForegroundColor Cyan
    } else {
        throw "PostgreSQL not found"
    }
} catch {
    Write-Host "❌ PostgreSQL is not installed or not in PATH" -ForegroundColor Red
    Write-Host "Please install PostgreSQL from: https://www.postgresql.org/download/windows/" -ForegroundColor Yellow
    Write-Host "After installation, restart your terminal and run this script again" -ForegroundColor Yellow
    exit 1
}

# Database configuration
$DB_NAME = "web3_portfolio_dev"
$DB_USER = "dev_user"

# Use env var or prompt — never hardcode passwords in source
$DB_PASSWORD = $env:WEB3_DB_PASSWORD
if (-not $DB_PASSWORD) {
    Write-Host "`n🔐 Enter a password for database user '$DB_USER' (or set WEB3_DB_PASSWORD):" -ForegroundColor Yellow
    $securePassword = Read-Host -AsSecureString
    $DB_PASSWORD = [Runtime.InteropServices.Marshal]::PtrToStringAuto([Runtime.InteropServices.Marshal]::SecureStringToBSTR($securePassword))
}

if (-not $DB_PASSWORD) {
    Write-Host "❌ Database password is required" -ForegroundColor Red
    exit 1
}

Write-Host "📋 Database Configuration:" -ForegroundColor Cyan
Write-Host "  Database Name: $DB_NAME" -ForegroundColor White
Write-Host "  Username: $DB_USER" -ForegroundColor White
Write-Host "  Password: (hidden)" -ForegroundColor White

# Prompt for PostgreSQL superuser password
Write-Host "`n🔐 Enter your PostgreSQL superuser password (default: postgres123):" -ForegroundColor Yellow
$PG_PASSWORD = Read-Host -AsSecureString
$PG_PASSWORD_PLAIN = [Runtime.InteropServices.Marshal]::PtrToStringAuto([Runtime.InteropServices.Marshal]::SecureStringToBSTR($PG_PASSWORD))

if (-not $PG_PASSWORD_PLAIN) {
    $PG_PASSWORD_PLAIN = "postgres123"
}

# Set environment variable for psql
$env:PGPASSWORD = $PG_PASSWORD_PLAIN

Write-Host "`n🔧 Creating database and user..." -ForegroundColor Yellow

# Create database
Write-Host "  Creating database '$DB_NAME'..." -ForegroundColor White
try {
    psql -U postgres -h localhost -c "CREATE DATABASE $DB_NAME;" 2>$null
    if ($LASTEXITCODE -eq 0) {
        Write-Host "  ✅ Database created successfully" -ForegroundColor Green
    } else {
        Write-Host "  ⚠️ Database might already exist (this is OK)" -ForegroundColor Yellow
    }
} catch {
    Write-Host "  ❌ Failed to create database" -ForegroundColor Red
    Write-Host "  Error: $_" -ForegroundColor Red
}

# Create user
Write-Host "  Creating user '$DB_USER'..." -ForegroundColor White
try {
    psql -U postgres -h localhost -c "CREATE USER $DB_USER WITH PASSWORD '$DB_PASSWORD';" 2>$null
    if ($LASTEXITCODE -eq 0) {
        Write-Host "  ✅ User created successfully" -ForegroundColor Green
    } else {
        Write-Host "  ⚠️ User might already exist (this is OK)" -ForegroundColor Yellow
    }
} catch {
    Write-Host "  ❌ Failed to create user" -ForegroundColor Red
    Write-Host "  Error: $_" -ForegroundColor Red
}

# Grant privileges
Write-Host "  Granting privileges..." -ForegroundColor White
try {
    psql -U postgres -h localhost -c "GRANT ALL PRIVILEGES ON DATABASE $DB_NAME TO $DB_USER;" 2>$null
    psql -U postgres -h localhost -d $DB_NAME -c "GRANT ALL ON SCHEMA public TO $DB_USER;" 2>$null
    psql -U postgres -h localhost -d $DB_NAME -c "GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO $DB_USER;" 2>$null
    psql -U postgres -h localhost -d $DB_NAME -c "GRANT ALL PRIVILEGES ON ALL SEQUENCES IN SCHEMA public TO $DB_USER;" 2>$null
    Write-Host "  ✅ Privileges granted successfully" -ForegroundColor Green
} catch {
    Write-Host "  ❌ Failed to grant privileges" -ForegroundColor Red
    Write-Host "  Error: $_" -ForegroundColor Red
}

# Test connection
Write-Host "`n🧪 Testing database connection..." -ForegroundColor Yellow
$env:PGPASSWORD = $DB_PASSWORD
try {
    $testResult = psql -U $DB_USER -h localhost -d $DB_NAME -c "SELECT version();" 2>$null
    if ($LASTEXITCODE -eq 0) {
        Write-Host "  ✅ Database connection successful!" -ForegroundColor Green
    } else {
        Write-Host "  ❌ Database connection failed" -ForegroundColor Red
    }
} catch {
    Write-Host "  ❌ Database connection failed" -ForegroundColor Red
    Write-Host "  Error: $_" -ForegroundColor Red
}

# Update environment file
Write-Host "`n📝 Updating environment configuration..." -ForegroundColor Yellow
$envFile = "backend\env.dev"
if (Test-Path $envFile) {
    $content = Get-Content $envFile -Raw
    $newContent = $content -replace 'DATABASE_URL=.*', "DATABASE_URL=postgres://$DB_USER`:$DB_PASSWORD@localhost:5432/$DB_NAME?sslmode=disable"
    Set-Content $envFile $newContent
    Write-Host "  ✅ Environment file updated" -ForegroundColor Green
} else {
    Write-Host "  ❌ Environment file not found: $envFile" -ForegroundColor Red
}

Write-Host "`n🎉 PostgreSQL setup complete!" -ForegroundColor Green
Write-Host "`n📋 Connection Details:" -ForegroundColor Cyan
Write-Host "  Host: localhost" -ForegroundColor White
Write-Host "  Port: 5432" -ForegroundColor White
Write-Host "  Database: $DB_NAME" -ForegroundColor White
Write-Host "  Username: $DB_USER" -ForegroundColor White
Write-Host "  Password: (hidden)" -ForegroundColor White
Write-Host "  Connection String: postgres://$DB_USER`:****@localhost:5432/$DB_NAME?sslmode=disable" -ForegroundColor White

Write-Host "`n🚀 Next steps:" -ForegroundColor Cyan
Write-Host "  1. Run database migrations: cd backend; go run main.go --migrate" -ForegroundColor White
Write-Host "  2. Start the backend: cd backend; go run main.go" -ForegroundColor White
Write-Host "  3. Start the frontend: cd frontend; npm run dev" -ForegroundColor White

# Clean up environment variable
Remove-Item Env:PGPASSWORD -ErrorAction SilentlyContinue 