# Web3 Portfolio Dashboard - Deployment Guide

This guide covers how to deploy the Web3 Portfolio Dashboard for both development and production environments using Docker.

## 🚀 Quick Start

### Prerequisites

- Docker Desktop installed and running
- Docker Compose installed
- Git (for cloning the repository)

### Development Deployment

1. **Clone the repository:**
   ```bash
   git clone <repository-url>
   cd web3-portfolio-dashboard
   ```

2. **Start development environment:**
   ```bash
   # Linux/Mac
   chmod +x scripts/deploy-dev.sh
   ./scripts/deploy-dev.sh

   # Windows PowerShell
   .\scripts\deploy-dev.ps1
   ```

3. **Access the application:**
   - Frontend: http://localhost:3000
   - Backend API: http://localhost:8080
   - pgAdmin: http://localhost:5050 (admin@web3portfolio.dev / admin123)

### Production Deployment

1. **Set up environment variables:**
   ```bash
   cp env.production.example .env.production
   # Edit .env.production with your actual values
   ```

2. **Deploy to production:**
   ```bash
   # Linux/Mac
   chmod +x scripts/deploy-prod.sh
   ./scripts/deploy-prod.sh

   # Windows PowerShell
   .\scripts\deploy-prod.ps1
   ```

3. **Access the application:**
   - Frontend: http://localhost:3000
   - Backend API: http://localhost:8080
   - Nginx: http://localhost:80

## 📋 Environment Configuration

### Development Environment

The development environment uses `docker-compose.dev.yml` with:
- Hot reloading for both frontend and backend
- Development databases with sample data
- Debug logging enabled
- pgAdmin for database management

### Production Environment

The production environment uses `docker-compose.prod.yml` with:
- Optimized builds
- Production databases
- Nginx reverse proxy
- Health checks
- Security hardening

### Required Environment Variables

Create a `.env.production` file with the following variables:

```bash
# Database Configuration
POSTGRES_PASSWORD=your-super-secure-postgres-password-here
REDIS_PASSWORD=your-super-secure-redis-password-here

# JWT Configuration
JWT_SECRET=your-super-secret-jwt-key-minimum-32-characters-long

# Web3 RPC URLs
ETHEREUM_RPC_URL=https://mainnet.infura.io/v3/YOUR_INFURA_PROJECT_ID
POLYGON_RPC_URL=https://polygon-rpc.com
BSC_RPC_URL=https://bsc-dataseed.binance.org
ARBITRUM_RPC_URL=https://arb1.arbitrum.io/rpc

# API Keys
ETHERSCAN_API_KEY=your-etherscan-api-key
POLYGONSCAN_API_KEY=your-polygonscan-api-key
BSCSCAN_API_KEY=your-bscscan-api-key
COINGECKO_API_KEY=your-coingecko-api-key

# CORS Configuration
CORS_ALLOWED_ORIGINS=https://yourdomain.com,https://www.yourdomain.com

# Monitoring
SENTRY_DSN=your-sentry-dsn
GOOGLE_ANALYTICS_ID=your-google-analytics-id

# Security
CSP_NONCE=your-csp-nonce-value
```

## 🐳 Docker Services

### Development Services

- **postgres**: PostgreSQL database with development data
- **redis**: Redis cache for sessions and caching
- **backend**: Go API with hot reloading
- **frontend**: React app with hot reloading
- **pgadmin**: Database management interface

### Production Services

- **postgres**: PostgreSQL database
- **redis**: Redis cache with authentication
- **backend**: Optimized Go API
- **frontend**: Built React app served by Nginx
- **nginx**: Reverse proxy with SSL support

## 🔧 Manual Commands

### Development

```bash
# Start development environment
docker-compose -f docker-compose.dev.yml up -d

# View logs
docker-compose -f docker-compose.dev.yml logs -f

# Stop services
docker-compose -f docker-compose.dev.yml down

# Rebuild and start
docker-compose -f docker-compose.dev.yml up --build -d

# Clean up volumes
docker-compose -f docker-compose.dev.yml down -v
```

### Production

```bash
# Start production environment
docker-compose -f docker-compose.prod.yml up -d

# View logs
docker-compose -f docker-compose.prod.yml logs -f

# Stop services
docker-compose -f docker-compose.prod.yml down

# Rebuild and start
docker-compose -f docker-compose.prod.yml up --build -d

# Clean up volumes
docker-compose -f docker-compose.prod.yml down -v
```

## 🔒 Security Considerations

### Production Security Checklist

- [ ] Change default passwords in `.env.production`
- [ ] Configure SSL certificates for HTTPS
- [ ] Set up proper firewall rules
- [ ] Enable rate limiting
- [ ] Configure CORS properly
- [ ] Set up monitoring and alerting
- [ ] Regular security updates
- [ ] Database backups
- [ ] Log monitoring

### SSL Configuration

1. **Generate SSL certificates:**
   ```bash
   mkdir ssl
   # Add your SSL certificates to the ssl/ directory
   ```

2. **Update nginx.conf:**
   Uncomment and configure the HTTPS server block in `nginx.conf`

3. **Update environment variables:**
   Set `CORS_ALLOWED_ORIGINS` to use HTTPS URLs

## 📊 Monitoring and Logs

### Health Checks

All services include health checks:
- Backend: `http://localhost:8080/health`
- Frontend: `http://localhost:3000/`
- Database: PostgreSQL connection check
- Redis: Ping command

### Log Management

```bash
# View all logs
docker-compose -f docker-compose.prod.yml logs

# View specific service logs
docker-compose -f docker-compose.prod.yml logs backend

# Follow logs in real-time
docker-compose -f docker-compose.prod.yml logs -f

# View logs with timestamps
docker-compose -f docker-compose.prod.yml logs -t
```

## 🚨 Troubleshooting

### Common Issues

1. **Port conflicts:**
   ```bash
   # Check what's using the ports
   netstat -tulpn | grep :3000
   netstat -tulpn | grep :8080
   ```

2. **Database connection issues:**
   ```bash
   # Check database logs
   docker-compose logs postgres
   
   # Connect to database
   docker exec -it web3-portfolio-postgres psql -U portfolio_user -d web3_portfolio
   ```

3. **Build failures:**
   ```bash
   # Clean build
   docker-compose down
   docker system prune -f
   docker-compose up --build
   ```

4. **Permission issues:**
   ```bash
   # Fix file permissions
   chmod +x scripts/*.sh
   ```

### Performance Optimization

1. **Database optimization:**
   - Add database indexes
   - Configure connection pooling
   - Regular maintenance

2. **Caching:**
   - Redis is configured for caching
   - Implement application-level caching
   - Use CDN for static assets

3. **Monitoring:**
   - Set up Prometheus/Grafana
   - Configure alerting
   - Monitor resource usage

## 🔄 Updates and Maintenance

### Updating the Application

1. **Pull latest changes:**
   ```bash
   git pull origin main
   ```

2. **Rebuild and restart:**
   ```bash
   # Development
   docker-compose -f docker-compose.dev.yml up --build -d

   # Production
   docker-compose -f docker-compose.prod.yml up --build -d
   ```

### Database Migrations

1. **Backup current data:**
   ```bash
   docker exec web3-portfolio-postgres pg_dump -U portfolio_user web3_portfolio > backup.sql
   ```

2. **Apply migrations:**
   ```bash
   # Run migration scripts
   docker exec -it web3-portfolio-postgres psql -U portfolio_user -d web3_portfolio -f /path/to/migration.sql
   ```

### Backup Strategy

1. **Database backups:**
   ```bash
   # Create backup script
   docker exec web3-portfolio-postgres pg_dump -U portfolio_user web3_portfolio > backup_$(date +%Y%m%d_%H%M%S).sql
   ```

2. **Volume backups:**
   ```bash
   # Backup volumes
   docker run --rm -v web3-portfolio_postgres_data:/data -v $(pwd):/backup alpine tar czf /backup/postgres_backup.tar.gz -C /data .
   ```

## 📞 Support

For issues and questions:
1. Check the troubleshooting section
2. Review logs for error messages
3. Check GitHub issues
4. Contact the development team

---

**Note:** Always test deployments in a staging environment before applying to production. 