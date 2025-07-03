# Update PostgreSQL User Password Script
# This script updates the PostgreSQL user password to match the environment configuration

Write-Host "🔐 Updating PostgreSQL user password..." -ForegroundColor Green

# PostgreSQL connection details
$PG_USER = "dev_user"
$NEW_PASSWORD = "Welcome@18"
$DB_NAME = "web3_portfolio_dev"

Write-Host "📋 Updating password for user: $PG_USER" -ForegroundColor Yellow
Write-Host "📋 New password: $NEW_PASSWORD" -ForegroundColor Yellow
Write-Host "📋 Database: $DB_NAME" -ForegroundColor Yellow

# Method 1: Using psql with postgres user (if you have access)
Write-Host "🔄 Attempting to update password using psql..." -ForegroundColor Yellow

try {
    # Try to connect as postgres user and update the password
    $sqlCommand = "ALTER USER $PG_USER WITH PASSWORD '$NEW_PASSWORD';"
    
    # Execute the SQL command using psql
    $result = psql -U postgres -d postgres -c $sqlCommand 2>&1
    
    if ($LASTEXITCODE -eq 0) {
        Write-Host "✅ Password updated successfully!" -ForegroundColor Green
    } else {
        Write-Host "❌ Failed to update password using psql" -ForegroundColor Red
        Write-Host "Error: $result" -ForegroundColor Red
        
        # Method 2: Manual instructions
        Write-Host "`n💡 Manual steps to update password:" -ForegroundColor Cyan
        Write-Host "1. Open pgAdmin or psql" -ForegroundColor White
        Write-Host "2. Connect as postgres user" -ForegroundColor White
        Write-Host "3. Run: ALTER USER dev_user WITH PASSWORD 'Welcome@18';" -ForegroundColor White
        Write-Host "4. Or run: \password dev_user" -ForegroundColor White
    }
} catch {
    Write-Host "❌ Error executing psql command" -ForegroundColor Red
    Write-Host "Error: $($_.Exception.Message)" -ForegroundColor Red
}

Write-Host "`n💡 Alternative methods:" -ForegroundColor Cyan
Write-Host "1. Use pgAdmin: Right-click user → Properties → Definition → Password" -ForegroundColor White
Write-Host "2. Use psql: \password dev_user" -ForegroundColor White
Write-Host "3. Use SQL: ALTER USER dev_user WITH PASSWORD 'Welcome@18';" -ForegroundColor White

Write-Host "`n✅ Script completed!" -ForegroundColor Green 