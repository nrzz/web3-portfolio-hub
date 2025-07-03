# Web3 Portfolio Dashboard - Complete Testing Guide

## 🚀 Quick Start Testing

You can now test the complete Web3 Portfolio Dashboard from subscription login to wallet data retrieval using Docker, without manually setting up any databases!

### Prerequisites
- Docker Desktop installed and running
- Windows PowerShell or Linux/Mac terminal

### One-Command Setup (Windows)
```powershell
.\scripts\test-setup.ps1
```

### One-Command Setup (Linux/Mac)
```bash
./scripts/test-setup.sh
```

### Manual Setup
```bash
# Start all services with test data
docker-compose -f docker-compose.dev.yml up --build -d

# Wait for services to be ready (30-60 seconds)
# Then access the application
```

## 🎯 What You Can Test

### 1. **Complete User Journey**
- ✅ User registration and login
- ✅ Subscription management (Basic, Pro, Premium tiers)
- ✅ Profile management and settings

### 2. **Wallet Integration**
- ✅ Connect multiple wallets (Ethereum, Polygon, BSC)
- ✅ Real-time balance tracking
- ✅ Transaction history
- ✅ Token holdings and values

### 3. **Portfolio Management**
- ✅ Multi-wallet portfolio overview
- ✅ Token allocation and diversification
- ✅ Performance tracking and analytics
- ✅ Historical data visualization

### 4. **Price Alerts & Notifications**
- ✅ Create custom price alerts
- ✅ Portfolio value alerts
- ✅ Alert history and management
- ✅ Real-time notifications

### 5. **Analytics & Reporting**
- ✅ Portfolio performance metrics
- ✅ Token price charts
- ✅ Transaction analysis
- ✅ Risk assessment

## 📱 Application URLs

| Service | URL | Description |
|---------|-----|-------------|
| **Frontend** | http://localhost:3000 | Main application UI |
| **Backend API** | http://localhost:8080 | REST API endpoints |
| **API Health** | http://localhost:8080/health | Service health check |
| **API Docs** | http://localhost:8080/swagger | Interactive API documentation |
| **Database** | http://localhost:5050 | pgAdmin database management |

## 🧪 Test Credentials

### Primary Test User
- **Email:** `test@web3portfolio.dev`
- **Password:** `password`
- **Subscription:** Premium tier
- **Wallets:** 2 wallets (Ethereum + Polygon)

### Additional Test Users
- **Alice:** `alice@web3portfolio.dev` / `password` (Pro tier)
- **Bob:** `bob@web3portfolio.dev` / `password` (Basic tier)

### Database Access
- **pgAdmin:** `admin@web3portfolio.dev` / `admin123`

## 📊 Pre-Loaded Test Data

### Users & Subscriptions
- 3 test users with different subscription tiers
- Active subscriptions with full feature access

### Wallets & Networks
- **Ethereum Mainnet:** `0x742d35Cc6634C0532925a3b8D4C9db96C4b4d8b6`
- **Polygon:** `0x742d35Cc6634C0532925a3b8D4C9db96C4b4d8b7`
- **BSC:** `0x742d35Cc6634C0532925a3b8D4C9db96C4b4d8b9`

### Tokens & Holdings
- **ETH:** 2.5 tokens ($6,250 value)
- **USDC:** 5,000 tokens ($5,000 value)
- **MATIC:** 1,000 tokens ($850 value)
- **BNB:** 10 tokens ($3,200 value)
- **LINK:** 100 tokens ($1,550 value)

### Transactions
- 4 sample transactions (buy/sell)
- Realistic gas fees and timestamps
- Multiple networks and tokens

### Alerts
- ETH price alert (>$3,000)
- Portfolio value alert (<$10,000)
- MATIC price alert (<$0.80)

## 🔍 Step-by-Step Testing Guide

### 1. **Initial Setup & Login**
1. Open http://localhost:3000
2. Click "Login" or "Sign Up"
3. Use test credentials: `test@web3portfolio.dev` / `password`
4. Verify successful login and dashboard access

### 2. **Dashboard Overview**
1. Check portfolio total value ($16,850)
2. Verify token allocation chart
3. Review recent transactions
4. Check performance metrics

### 3. **Wallet Management**
1. Navigate to Portfolio page
2. View connected wallets (2 wallets)
3. Check individual wallet balances
4. Test wallet connection flow

### 4. **Portfolio Analytics**
1. Go to Analytics page
2. Review performance charts
3. Check token allocation
4. View historical data

### 5. **Price Alerts**
1. Navigate to Alerts page
2. View existing alerts (3 alerts)
3. Create new price alert
4. Test alert management

### 6. **Settings & Profile**
1. Access Settings page
2. Update profile information
3. Manage subscription settings
4. Configure preferences

### 7. **API Testing**
1. Visit http://localhost:8080/swagger
2. Test authentication endpoints
3. Verify portfolio data retrieval
4. Check wallet integration

## 🛠️ Development Commands

### View Logs
```bash
# All services
docker-compose -f docker-compose.dev.yml logs -f

# Specific service
docker-compose -f docker-compose.dev.yml logs -f backend
docker-compose -f docker-compose.dev.yml logs -f frontend
```

### Restart Services
```bash
# Restart all services
docker-compose -f docker-compose.dev.yml restart

# Restart specific service
docker-compose -f docker-compose.dev.yml restart backend
```

### Stop Services
```bash
# Stop and remove containers
docker-compose -f docker-compose.dev.yml down

# Stop and remove containers + volumes
docker-compose -f docker-compose.dev.yml down --volumes
```

### Database Access
```bash
# Connect to PostgreSQL
docker exec -it web3-portfolio-postgres-dev psql -U dev_user -d web3_portfolio_dev

# View database logs
docker logs web3-portfolio-postgres-dev
```

## 🔧 Troubleshooting

### Services Not Starting
1. Check Docker Desktop is running
2. Verify ports 3000, 8080, 5432, 6379, 5050 are available
3. Run `docker-compose -f docker-compose.dev.yml logs` for errors

### Database Connection Issues
1. Wait 30-60 seconds for PostgreSQL to initialize
2. Check database logs: `docker logs web3-portfolio-postgres-dev`
3. Verify test data was loaded correctly

### Frontend Not Loading
1. Check if backend API is responding: http://localhost:8080/health
2. Verify frontend logs: `docker logs web3-portfolio-frontend-dev`
3. Clear browser cache and try again

### API Errors
1. Check backend logs: `docker logs web3-portfolio-backend-dev`
2. Verify Redis connection: `docker exec web3-portfolio-redis-dev redis-cli ping`
3. Check database connectivity

## 📈 Performance Testing

### Load Testing
```bash
# Test API endpoints with multiple requests
curl -X GET http://localhost:8080/api/v1/portfolio \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -w "@curl-format.txt"
```

### Database Performance
```bash
# Check database performance
docker exec web3-portfolio-postgres-dev psql -U dev_user -d web3_portfolio_dev -c "SELECT * FROM pg_stat_activity;"
```

## 🎯 Success Criteria

Your testing is successful when you can:

✅ **Login** with test credentials and access dashboard  
✅ **View** portfolio with $16,850 total value  
✅ **See** 2 connected wallets with real balances  
✅ **Create** and manage price alerts  
✅ **Access** analytics with charts and data  
✅ **Update** profile and subscription settings  
✅ **Connect** to real Web3 APIs for live data  
✅ **Retrieve** wallet data from multiple networks  

## 🚀 Next Steps

After successful testing:

1. **Customize** the application for your needs
2. **Add** your own API keys for production
3. **Deploy** to your preferred cloud platform
4. **Scale** the infrastructure as needed
5. **Monitor** performance and user metrics

## 📞 Support

If you encounter issues:

1. Check the troubleshooting section above
2. Review service logs for error messages
3. Verify all prerequisites are met
4. Ensure Docker has sufficient resources allocated

---

**🎉 You're now ready to test the complete Web3 Portfolio Dashboard!** 