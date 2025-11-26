# Test Add Address Functionality
Write-Host "Testing Add Address Functionality..." -ForegroundColor Green

# Step 1: Register a user
Write-Host "`n1. Registering user..." -ForegroundColor Yellow
$registerBody = @{
    name = "Test User"
    email = "test@example.com"
    password = "password123"
} | ConvertTo-Json

$registerResponse = Invoke-RestMethod -Uri "http://localhost:8080/api/v1/auth/register" -Method POST -Body $registerBody -ContentType "application/json"
Write-Host "Register Response: $($registerResponse | ConvertTo-Json -Depth 3)"

# Step 2: Login to get token
Write-Host "`n2. Logging in..." -ForegroundColor Yellow
$loginBody = @{
    email = "test@example.com"
    password = "password123"
} | ConvertTo-Json

$loginResponse = Invoke-RestMethod -Uri "http://localhost:8080/api/v1/auth/login" -Method POST -Body $loginBody -ContentType "application/json"
$token = $loginResponse.token
Write-Host "Login successful, token: $($token.Substring(0, 20))..."

# Step 3: Create a portfolio
Write-Host "`n3. Creating portfolio..." -ForegroundColor Yellow
$portfolioBody = @{
    name = "Test Portfolio"
} | ConvertTo-Json

$headers = @{
    "Authorization" = "Bearer $token"
    "Content-Type" = "application/json"
}

$portfolioResponse = Invoke-RestMethod -Uri "http://localhost:8080/api/v1/portfolios" -Method POST -Body $portfolioBody -Headers $headers
$portfolioId = $portfolioResponse.portfolio.id
Write-Host "Portfolio created with ID: $portfolioId"

# Step 4: Get portfolios to verify
Write-Host "`n4. Getting portfolios..." -ForegroundColor Yellow
$portfoliosResponse = Invoke-RestMethod -Uri "http://localhost:8080/api/v1/portfolios" -Method GET -Headers $headers
Write-Host "Portfolios: $($portfoliosResponse | ConvertTo-Json -Depth 3)"

# Step 5: Add address
Write-Host "`n5. Adding address..." -ForegroundColor Yellow
$addressBody = @{
    address = "0x742d35Cc6634C0532925a3b8D4C9db96C4b4d8b6"
    network = "ethereum"
    label = "Test Wallet"
} | ConvertTo-Json

try {
    $addressResponse = Invoke-RestMethod -Uri "http://localhost:8080/api/v1/portfolios/$portfolioId/addresses" -Method POST -Body $addressBody -Headers $headers
    Write-Host "Address added successfully: $($addressResponse | ConvertTo-Json -Depth 3)" -ForegroundColor Green
} catch {
    Write-Host "Error adding address: $($_.Exception.Message)" -ForegroundColor Red
    Write-Host "Response: $($_.Exception.Response)" -ForegroundColor Red
}

# Step 6: Get addresses to verify
Write-Host "`n6. Getting addresses..." -ForegroundColor Yellow
try {
    $addressesResponse = Invoke-RestMethod -Uri "http://localhost:8080/api/v1/portfolios/$portfolioId/addresses" -Method GET -Headers $headers
    Write-Host "Addresses: $($addressesResponse | ConvertTo-Json -Depth 3)" -ForegroundColor Green
} catch {
    Write-Host "Error getting addresses: $($_.Exception.Message)" -ForegroundColor Red
}

Write-Host "`nTest completed!" -ForegroundColor Green 