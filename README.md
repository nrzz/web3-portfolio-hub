# 🚀 Web3 Portfolio Dashboard & Community Forum

A comprehensive **full-stack Web3 portfolio tracker** with an integrated **community forum** for cryptocurrency enthusiasts. Track your crypto portfolios across multiple blockchain networks while engaging with a Stack Overflow-style community.

## 🚀 Quick Start Testing

**Test the complete application from subscription login to wallet data retrieval with one command!**

### Windows (PowerShell)
```powershell
.\scripts\test-setup.ps1
```

### Linux/Mac (Bash)
```bash
./scripts/test-setup.sh
```

### Manual Setup
```bash
# Start all services with pre-loaded test data
docker-compose -f docker-compose.dev.yml up --build -d

# Wait 30-60 seconds for services to initialize
# Then access: http://localhost:3000
```

## 🎯 What You Get

✅ **Complete User Journey** - Registration, login, subscription management  
✅ **Multi-Wallet Support** - Ethereum, Polygon, BSC integration  
✅ **Real-Time Data** - Live prices, balances, transactions  
✅ **Portfolio Analytics** - Performance tracking, charts, insights  
✅ **Price Alerts** - Custom notifications and monitoring  
✅ **Production Ready** - Secure, scalable, containerized  

## 📱 Application Access

| Service | URL | Description |
|---------|-----|-------------|
| **Frontend** | http://localhost:3000 | Main application UI |
| **Backend API** | http://localhost:8080 | REST API endpoints |
| **API Health** | http://localhost:8080/health | Service health check |
| **API Docs** | http://localhost:8080/swagger | Interactive API documentation |
| **Database** | http://localhost:5050 | pgAdmin database management |

## 🧪 Test Credentials

- **Email:** `test@web3portfolio.dev`
- **Password:** `password`
- **Subscription:** Premium tier with full features

## 📊 Pre-Loaded Test Data

- **3 test users** with different subscription tiers
- **4 wallets** across Ethereum, Polygon, and BSC
- **5 tokens** (ETH, USDC, MATIC, BNB, LINK) with real market data
- **6 holdings** totaling $16,850 portfolio value
- **4 sample transactions** with realistic data
- **3 price alerts** for testing notifications

## 🏗️ Architecture

```
┌─────────────────┐    ┌─────────────────┐    ┌─────────────────┐
│   Frontend      │    │   Backend       │    │   Database      │
│   (React + TS)  │◄──►│   (Go + Gin)    │◄──►│   (PostgreSQL)  │
│   Port: 3000    │    │   Port: 8080    │    │   Port: 5432    │
└─────────────────┘    └─────────────────┘    └─────────────────┘
                                │
                                ▼
                       ┌─────────────────┐
                       │   Redis Cache   │
                       │   Port: 6379    │
                       └─────────────────┘
```

## 🛠️ Technology Stack

### Frontend
- **React 18** with TypeScript
- **Vite** for fast development
- **Tailwind CSS** for styling
- **Framer Motion** for animations
- **Lucide React** for icons
- **Headless UI** for components

### Backend
- **Go 1.21** with Gin framework
- **PostgreSQL** for primary database
- **Redis** for caching and sessions
- **JWT** for authentication
- **Swagger** for API documentation

### Infrastructure
- **Docker** for containerization
- **Docker Compose** for orchestration
- **Nginx** for production serving
- **pgAdmin** for database management

## 📋 Features

### 🔐 Authentication & Subscriptions
- User registration and login
- JWT-based authentication
- Subscription tier management (Basic, Pro, Premium)
- Role-based access control

### 💼 Portfolio Management
- Multi-wallet support (Ethereum, Polygon, BSC)
- Real-time balance tracking
- Token holdings and values
- Portfolio performance metrics

### 📊 Analytics & Reporting
- Portfolio performance charts
- Token allocation visualization
- Historical data analysis
- Risk assessment metrics

### 🔔 Price Alerts
- Custom price alerts for any token
- Portfolio value alerts
- Email notifications
- Alert history and management

### 🔗 Web3 Integration
- Real-time blockchain data
- Multi-network support
- Transaction history
- Gas fee tracking

### 💬 Community Forum
- **Q&A Platform**: Stack Overflow-style question/answer system
- **Voting System**: Upvote/downvote questions and answers
- **Reputation System**: Earn points for community contributions
- **Tagging**: Organize content with relevant tags
- **Moderation**: Role-based access control

## 🚀 Development

### Prerequisites
- Docker Desktop
- Node.js 18+ (for local development)
- Go 1.21+ (for local development)

### Local Development

#### Backend
```bash
cd backend
go mod download
go run main.go
```

#### Frontend
```bash
cd frontend
npm install
npm run dev
```

### Docker Development
```bash
# Start development environment
docker-compose -f docker-compose.dev.yml up --build

# View logs
docker-compose -f docker-compose.dev.yml logs -f

# Stop services
docker-compose -f docker-compose.dev.yml down
```

## 🏭 Production Deployment

### Using Docker Compose
```bash
# Start production environment
docker-compose -f docker-compose.prod.yml up --build -d

# Scale services
docker-compose -f docker-compose.prod.yml up --scale backend=3 -d
```

### Environment Variables
Copy and configure the environment files:
```bash
cp backend/env.example backend/.env
cp frontend/env.example frontend/.env
```

### Required Environment Variables
```bash
# Backend
DATABASE_URL=postgres://user:pass@host:5432/db
REDIS_URL=redis://host:6379
JWT_SECRET=your-secret-key
ETHEREUM_RPC_URL=https://eth-mainnet.alchemyapi.io/v2/YOUR_KEY
POLYGON_RPC_URL=https://polygon-rpc.com
BSC_RPC_URL=https://bsc-dataseed1.binance.org
COINGECKO_API_URL=https://api.coingecko.com/api/v3

# Frontend
VITE_API_URL=http://localhost:8080
VITE_APP_NAME=Web3 Portfolio Dashboard
```

## 📚 API Documentation

### Authentication
```bash
# Login
POST /api/v1/auth/login
{
  "email": "user@example.com",
  "password": "password"
}

# Register
POST /api/v1/auth/register
{
  "email": "user@example.com",
  "password": "password",
  "first_name": "John",
  "last_name": "Doe"
}
```

### Portfolio
```bash
# Get portfolio overview
GET /api/v1/portfolio
Authorization: Bearer <jwt_token>

# Get wallet holdings
GET /api/v1/portfolio/wallets
Authorization: Bearer <jwt_token>

# Add wallet
POST /api/v1/portfolio/wallets
Authorization: Bearer <jwt_token>
{
  "name": "My Wallet",
  "address": "0x...",
  "network": "ethereum"
}
```

### Alerts
```bash
# Create price alert
POST /api/v1/alerts
Authorization: Bearer <jwt_token>
{
  "name": "ETH Alert",
  "condition_type": "price_above",
  "token_symbol": "ETH",
  "threshold_value": 3000.00
}
```

### Community Forum
- `GET /api/v1/forum/questions` - List questions
- `POST /api/v1/forum/questions` - Create question
- `GET /api/v1/forum/questions/:id` - Get question details
- `POST /api/v1/forum/answers` - Post answer
- `POST /api/v1/forum/votes` - Vote on content

## 🧪 Testing

### End-to-End Testing
```bash
# Run the complete test setup
./scripts/test-setup.sh  # Linux/Mac
.\scripts\test-setup.ps1 # Windows
```

### API Testing
```bash
# Test health endpoint
curl http://localhost:8080/health

# Test authentication
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email":"test@web3portfolio.dev","password":"password"}'
```

### Database Testing
```bash
# Connect to database
docker exec -it web3-portfolio-postgres-dev psql -U dev_user -d web3_portfolio_dev

# View test data
SELECT * FROM users;
SELECT * FROM wallets;
SELECT * FROM holdings;
```

## 📊 Monitoring & Logs

### View Logs
```bash
# All services
docker-compose -f docker-compose.dev.yml logs -f

# Specific service
docker-compose -f docker-compose.dev.yml logs -f backend
```

### Health Checks
- Backend: http://localhost:8080/health
- Database: `docker exec web3-portfolio-postgres-dev pg_isready`
- Redis: `docker exec web3-portfolio-redis-dev redis-cli ping`

## 🔧 Configuration

### Database Schema
The application automatically creates all necessary tables and indexes on startup.

### Caching
Redis is used for:
- Session storage
- API response caching
- Rate limiting
- Real-time data caching

### Security
- JWT authentication
- Password hashing with bcrypt
- CORS configuration
- Rate limiting
- Input validation

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests if applicable
5. Submit a pull request

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🆘 Support

For support and questions:
1. Check the [TESTING.md](TESTING.md) guide
2. Review the troubleshooting section
3. Check service logs for errors
4. Open an issue on GitHub

---

**🎉 Ready to test? Run `.\scripts\test-setup.ps1` (Windows) or `./scripts/test-setup.sh` (Linux/Mac) to get started!** 

**Built with ❤️ for the Web3 community** 