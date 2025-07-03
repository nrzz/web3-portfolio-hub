# Web3 Portfolio Dashboard API End-to-End Testing Script
# Run this script to test all major endpoints

Write-Host "🚀 Starting Web3 Portfolio Dashboard API End-to-End Testing" -ForegroundColor Green
Write-Host "==================================================" -ForegroundColor Green

$baseUrl = "http://localhost:8080"
$token = ""

# Test 1: Health Check
Write-Host "`n1️⃣ Testing Health Endpoint..." -ForegroundColor Yellow
try {
    $response = Invoke-RestMethod -Uri "$baseUrl/api/health" -Method GET
    Write-Host "✅ Health Check: PASSED" -ForegroundColor Green
    Write-Host "   Status: $($response.status)" -ForegroundColor Cyan
    Write-Host "   Database: $($response.services.database)" -ForegroundColor Cyan
    Write-Host "   Web3 Networks: $($response.services.web3 | ConvertTo-Json -Compress)" -ForegroundColor Cyan
} catch {
    Write-Host "❌ Health Check: FAILED - $($_.Exception.Message)" -ForegroundColor Red
}

# Test 2: Version Endpoint
Write-Host "`n2️⃣ Testing Version Endpoint..." -ForegroundColor Yellow
try {
    $response = Invoke-RestMethod -Uri "$baseUrl/api/version" -Method GET
    Write-Host "✅ Version Check: PASSED" -ForegroundColor Green
    Write-Host "   Version: $($response.version)" -ForegroundColor Cyan
    Write-Host "   Environment: $($response.environment)" -ForegroundColor Cyan
} catch {
    Write-Host "❌ Version Check: FAILED - $($_.Exception.Message)" -ForegroundColor Red
}

# Test 3: User Registration
Write-Host "`n3️⃣ Testing User Registration..." -ForegroundColor Yellow
try {
    $registerData = @{
        email = "test@example.com"
        password = "testpassword123"
        discord_id = "test_discord_123"
    } | ConvertTo-Json

    $response = Invoke-RestMethod -Uri "$baseUrl/api/v1/auth/register" -Method POST -Body $registerData -ContentType "application/json"
    $token = $response.token
    Write-Host "✅ User Registration: PASSED" -ForegroundColor Green
    Write-Host "   User ID: $($response.user.id)" -ForegroundColor Cyan
    Write-Host "   Email: $($response.user.email)" -ForegroundColor Cyan
} catch {
    Write-Host "❌ User Registration: FAILED - $($_.Exception.Message)" -ForegroundColor Red
}

# Test 4: User Login
Write-Host "`n4️⃣ Testing User Login..." -ForegroundColor Yellow
try {
    $loginData = @{
        email = "test@example.com"
        password = "testpassword123"
    } | ConvertTo-Json

    $response = Invoke-RestMethod -Uri "$baseUrl/api/v1/auth/login" -Method POST -Body $loginData -ContentType "application/json"
    $token = $response.token
    Write-Host "✅ User Login: PASSED" -ForegroundColor Green
    Write-Host "   Token received: $($token.Substring(0, 20))..." -ForegroundColor Cyan
} catch {
    Write-Host "❌ User Login: FAILED - $($_.Exception.Message)" -ForegroundColor Red
}

# Test 5: Get User Profile (Protected)
Write-Host "`n5️⃣ Testing Get User Profile..." -ForegroundColor Yellow
try {
    $headers = @{
        "Authorization" = "Bearer $token"
    }
    $response = Invoke-RestMethod -Uri "$baseUrl/api/v1/user/profile" -Method GET -Headers $headers
    Write-Host "✅ Get User Profile: PASSED" -ForegroundColor Green
    Write-Host "   Email: $($response.user.email)" -ForegroundColor Cyan
    Write-Host "   Subscription: $($response.user.subscription_tier)" -ForegroundColor Cyan
} catch {
    Write-Host "❌ Get User Profile: FAILED - $($_.Exception.Message)" -ForegroundColor Red
}

# Test 6: Create Portfolio
Write-Host "`n6️⃣ Testing Create Portfolio..." -ForegroundColor Yellow
try {
    $portfolioData = @{
        name = "Test Portfolio"
    } | ConvertTo-Json

    $headers = @{
        "Authorization" = "Bearer $token"
    }
    $response = Invoke-RestMethod -Uri "$baseUrl/api/v1/portfolios" -Method POST -Body $portfolioData -ContentType "application/json" -Headers $headers
    $portfolioId = $response.portfolio.id
    Write-Host "✅ Create Portfolio: PASSED" -ForegroundColor Green
    Write-Host "   Portfolio ID: $portfolioId" -ForegroundColor Cyan
    Write-Host "   Name: $($response.portfolio.name)" -ForegroundColor Cyan
} catch {
    Write-Host "❌ Create Portfolio: FAILED - $($_.Exception.Message)" -ForegroundColor Red
}

# Test 7: Get Portfolios
Write-Host "`n7️⃣ Testing Get Portfolios..." -ForegroundColor Yellow
try {
    $headers = @{
        "Authorization" = "Bearer $token"
    }
    $response = Invoke-RestMethod -Uri "$baseUrl/api/v1/portfolios" -Method GET -Headers $headers
    Write-Host "✅ Get Portfolios: PASSED" -ForegroundColor Green
    Write-Host "   Portfolio Count: $($response.portfolios.Count)" -ForegroundColor Cyan
} catch {
    Write-Host "❌ Get Portfolios: FAILED - $($_.Exception.Message)" -ForegroundColor Red
}

# Test 8: Add Address to Portfolio
Write-Host "`n8️⃣ Testing Add Address to Portfolio..." -ForegroundColor Yellow
try {
    $addressData = @{
        address = "0x742d35Cc6634C0532925a3b8D4C9db96C4b4d8b6"
        network = "ethereum"
        label = "Test Wallet"
    } | ConvertTo-Json

    $headers = @{
        "Authorization" = "Bearer $token"
    }
    $response = Invoke-RestMethod -Uri "$baseUrl/api/v1/portfolios/$portfolioId/addresses" -Method POST -Body $addressData -ContentType "application/json" -Headers $headers
    Write-Host "✅ Add Address: PASSED" -ForegroundColor Green
    Write-Host "   Address ID: $($response.address.id)" -ForegroundColor Cyan
    Write-Host "   Network: $($response.address.network)" -ForegroundColor Cyan
} catch {
    Write-Host "❌ Add Address: FAILED - $($_.Exception.Message)" -ForegroundColor Red
}

# Test 9: Get Portfolio Addresses
Write-Host "`n9️⃣ Testing Get Portfolio Addresses..." -ForegroundColor Yellow
try {
    $headers = @{
        "Authorization" = "Bearer $token"
    }
    $response = Invoke-RestMethod -Uri "$baseUrl/api/v1/portfolios/$portfolioId/addresses" -Method GET -Headers $headers
    Write-Host "✅ Get Portfolio Addresses: PASSED" -ForegroundColor Green
    Write-Host "   Address Count: $($response.addresses.Count)" -ForegroundColor Cyan
} catch {
    Write-Host "❌ Get Portfolio Addresses: FAILED - $($_.Exception.Message)" -ForegroundColor Red
}

# Test 10: Get Portfolio Balances
Write-Host "`n🔟 Testing Get Portfolio Balances..." -ForegroundColor Yellow
try {
    $headers = @{
        "Authorization" = "Bearer $token"
    }
    $response = Invoke-RestMethod -Uri "$baseUrl/api/v1/portfolios/$portfolioId/balances" -Method GET -Headers $headers
    Write-Host "✅ Get Portfolio Balances: PASSED" -ForegroundColor Green
    Write-Host "   Balance Count: $($response.balances.Count)" -ForegroundColor Cyan
} catch {
    Write-Host "❌ Get Portfolio Balances: FAILED - $($_.Exception.Message)" -ForegroundColor Red
}

# Test 11: Web3 Networks
Write-Host "`n1️⃣1️⃣ Testing Web3 Networks..." -ForegroundColor Yellow
try {
    $headers = @{
        "Authorization" = "Bearer $token"
    }
    $response = Invoke-RestMethod -Uri "$baseUrl/api/v1/web3/networks" -Method GET -Headers $headers
    Write-Host "✅ Web3 Networks: PASSED" -ForegroundColor Green
    Write-Host "   Networks: $($response.networks | ConvertTo-Json -Compress)" -ForegroundColor Cyan
} catch {
    Write-Host "❌ Web3 Networks: FAILED - $($_.Exception.Message)" -ForegroundColor Red
}

# Test 12: Ethereum Network Status
Write-Host "`n1️⃣2️⃣ Testing Ethereum Network Status..." -ForegroundColor Yellow
try {
    $headers = @{
        "Authorization" = "Bearer $token"
    }
    $response = Invoke-RestMethod -Uri "$baseUrl/api/v1/web3/networks/ethereum/status" -Method GET -Headers $headers
    Write-Host "✅ Ethereum Network Status: PASSED" -ForegroundColor Green
    Write-Host "   Status: $($response.status)" -ForegroundColor Cyan
} catch {
    Write-Host "❌ Ethereum Network Status: FAILED - $($_.Exception.Message)" -ForegroundColor Red
}

# Test 13: Get Token Price
Write-Host "`n1️⃣3️⃣ Testing Get Token Price..." -ForegroundColor Yellow
try {
    $headers = @{
        "Authorization" = "Bearer $token"
    }
    $response = Invoke-RestMethod -Uri "$baseUrl/api/v1/web3/tokens/ETH/price" -Method GET -Headers $headers
    Write-Host "✅ Get Token Price: PASSED" -ForegroundColor Green
    Write-Host "   Symbol: $($response.symbol)" -ForegroundColor Cyan
    Write-Host "   Price: $($response.price)" -ForegroundColor Cyan
} catch {
    Write-Host "❌ Get Token Price: FAILED - $($_.Exception.Message)" -ForegroundColor Red
}

# Test 14: Create Alert
Write-Host "`n1️⃣4️⃣ Testing Create Alert..." -ForegroundColor Yellow
try {
    $alertData = @{
        type = "price"
        name = "ETH Price Alert"
        conditions = '{"token": "ETH", "operator": "above", "value": "2000"}'
    } | ConvertTo-Json

    $headers = @{
        "Authorization" = "Bearer $token"
    }
    $response = Invoke-RestMethod -Uri "$baseUrl/api/v1/alerts" -Method POST -Body $alertData -ContentType "application/json" -Headers $headers
    Write-Host "✅ Create Alert: PASSED" -ForegroundColor Green
    Write-Host "   Alert ID: $($response.alert.id)" -ForegroundColor Cyan
    Write-Host "   Name: $($response.alert.name)" -ForegroundColor Cyan
} catch {
    Write-Host "❌ Create Alert: FAILED - $($_.Exception.Message)" -ForegroundColor Red
}

# Test 15: Get Alerts
Write-Host "`n1️⃣5️⃣ Testing Get Alerts..." -ForegroundColor Yellow
try {
    $headers = @{
        "Authorization" = "Bearer $token"
    }
    $response = Invoke-RestMethod -Uri "$baseUrl/api/v1/alerts" -Method GET -Headers $headers
    Write-Host "✅ Get Alerts: PASSED" -ForegroundColor Green
    Write-Host "   Alert Count: $($response.alerts.Count)" -ForegroundColor Cyan
} catch {
    Write-Host "❌ Get Alerts: FAILED - $($_.Exception.Message)" -ForegroundColor Red
}

Write-Host "`n🎉 End-to-End Testing Complete!" -ForegroundColor Green
Write-Host "==================================================" -ForegroundColor Green
Write-Host "All major API endpoints have been tested." -ForegroundColor Cyan
Write-Host "Check the results above for any failures." -ForegroundColor Cyan 