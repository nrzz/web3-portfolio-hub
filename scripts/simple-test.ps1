# Simple API Test Script
Write-Host "Testing Web3 Portfolio Dashboard API" -ForegroundColor Green

$baseUrl = "http://localhost:8080"

# Test 1: Health Check
Write-Host "`n1. Testing Health Endpoint..." -ForegroundColor Yellow
try {
    $response = Invoke-RestMethod -Uri "$baseUrl/api/health" -Method GET
    Write-Host "PASSED - Status: $($response.status)" -ForegroundColor Green
} catch {
    Write-Host "FAILED - $($_.Exception.Message)" -ForegroundColor Red
}

# Test 2: Version Check
Write-Host "`n2. Testing Version Endpoint..." -ForegroundColor Yellow
try {
    $response = Invoke-RestMethod -Uri "$baseUrl/api/version" -Method GET
    Write-Host "PASSED - Version: $($response.version)" -ForegroundColor Green
} catch {
    Write-Host "FAILED - $($_.Exception.Message)" -ForegroundColor Red
}

# Test 3: User Registration
Write-Host "`n3. Testing User Registration..." -ForegroundColor Yellow
try {
    $registerData = @{
        email = "test@example.com"
        password = "testpassword123"
        discord_id = "test_discord_123"
    } | ConvertTo-Json

    $response = Invoke-RestMethod -Uri "$baseUrl/api/v1/auth/register" -Method POST -Body $registerData -ContentType "application/json"
    $token = $response.token
    Write-Host "PASSED - User ID: $($response.user.id)" -ForegroundColor Green
} catch {
    Write-Host "FAILED - $($_.Exception.Message)" -ForegroundColor Red
}

# Test 4: User Login
Write-Host "`n4. Testing User Login..." -ForegroundColor Yellow
try {
    $loginData = @{
        email = "test@example.com"
        password = "testpassword123"
    } | ConvertTo-Json

    $response = Invoke-RestMethod -Uri "$baseUrl/api/v1/auth/login" -Method POST -Body $loginData -ContentType "application/json"
    $token = $response.token
    Write-Host "PASSED - Token received" -ForegroundColor Green
} catch {
    Write-Host "FAILED - $($_.Exception.Message)" -ForegroundColor Red
}

# Test 5: Create Portfolio
Write-Host "`n5. Testing Create Portfolio..." -ForegroundColor Yellow
try {
    $portfolioData = @{
        name = "Test Portfolio"
    } | ConvertTo-Json

    $headers = @{
        "Authorization" = "Bearer $token"
    }
    $response = Invoke-RestMethod -Uri "$baseUrl/api/v1/portfolios" -Method POST -Body $portfolioData -ContentType "application/json" -Headers $headers
    $portfolioId = $response.portfolio.id
    Write-Host "PASSED - Portfolio ID: $portfolioId" -ForegroundColor Green
} catch {
    Write-Host "FAILED - $($_.Exception.Message)" -ForegroundColor Red
}

# Test 6: Get Portfolios
Write-Host "`n6. Testing Get Portfolios..." -ForegroundColor Yellow
try {
    $headers = @{
        "Authorization" = "Bearer $token"
    }
    $response = Invoke-RestMethod -Uri "$baseUrl/api/v1/portfolios" -Method GET -Headers $headers
    Write-Host "PASSED - Found $($response.portfolios.Count) portfolios" -ForegroundColor Green
} catch {
    Write-Host "FAILED - $($_.Exception.Message)" -ForegroundColor Red
}

Write-Host "`nTesting Complete!" -ForegroundColor Green 