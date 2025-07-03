# Test Web3 Networks Endpoint
Write-Host "Testing Web3 Networks Endpoint" -ForegroundColor Green

$baseUrl = "http://localhost:8080"

# First get a fresh token
$loginData = @{
    email = "test@example.com"
    password = "testpassword123"
} | ConvertTo-Json

try {
    $response = Invoke-RestMethod -Uri "$baseUrl/api/v1/auth/login" -Method POST -Body $loginData -ContentType "application/json"
    $token = $response.token
    Write-Host "Got token: $($token.Substring(0, 20))..." -ForegroundColor Cyan
    
    # Test Web3 Networks
    $headers = @{
        "Authorization" = "Bearer $token"
    }
    $response = Invoke-RestMethod -Uri "$baseUrl/api/v1/web3/networks" -Method GET -Headers $headers
    Write-Host "PASSED - Networks: $($response.networks | ConvertTo-Json)" -ForegroundColor Green
} catch {
    Write-Host "FAILED - $($_.Exception.Message)" -ForegroundColor Red
} 