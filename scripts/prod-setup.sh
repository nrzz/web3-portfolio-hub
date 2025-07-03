#!/bin/bash

# Web3 Portfolio Dashboard - Production Setup Script
# This script sets up the production environment with PostgreSQL and Redis

set -e

echo "🚀 Setting up Web3 Portfolio Dashboard Production Environment..."

# Check if Docker is running
if ! docker info > /dev/null 2>&1; then
    echo "❌ Docker is not running. Please start Docker and try again."
    exit 1
fi

# Check if Docker Compose is available
if ! command -v docker-compose &> /dev/null; then
    echo "❌ Docker Compose is not installed. Please install Docker Compose and try again."
    exit 1
fi

# Check if environment file exists
if [ ! -f "backend/env.production" ]; then
    echo "❌ Production environment file not found. Please create backend/env.production"
    exit 1
fi

echo "✅ Docker and Docker Compose are available"

# Create necessary directories
echo "📁 Creating necessary directories..."
mkdir -p logs
mkdir -p ssl
mkdir -p data/postgres
mkdir -p data/redis

# Copy environment file
echo "📝 Setting up production environment..."
cp backend/env.production backend/.env

# Stop any existing containers
echo "🛑 Stopping any existing containers..."
docker-compose -f docker-compose.prod.yml down --remove-orphans

# Build and start the services
echo "🔨 Building and starting production services..."
docker-compose -f docker-compose.prod.yml up --build -d

# Wait for services to be healthy
echo "⏳ Waiting for services to be ready..."
sleep 15

# Check service health
echo "🔍 Checking service health..."

# Check PostgreSQL
if docker-compose -f docker-compose.prod.yml exec -T postgres pg_isready -U portfolio_user -d web3_portfolio > /dev/null 2>&1; then
    echo "✅ PostgreSQL is ready"
else
    echo "❌ PostgreSQL is not ready. Check logs with: docker-compose -f docker-compose.prod.yml logs postgres"
fi

# Check Redis
if docker-compose -f docker-compose.prod.yml exec -T redis redis-cli ping > /dev/null 2>&1; then
    echo "✅ Redis is ready"
else
    echo "❌ Redis is not ready. Check logs with: docker-compose -f docker-compose.prod.yml logs redis"
fi

# Check Backend
if curl -f http://localhost:8080/health > /dev/null 2>&1; then
    echo "✅ Backend API is ready"
else
    echo "⏳ Backend API is starting up..."
    echo "   You can check the logs with: docker-compose -f docker-compose.prod.yml logs backend"
fi

# Check Frontend
if curl -f http://localhost:3000 > /dev/null 2>&1; then
    echo "✅ Frontend is ready"
else
    echo "⏳ Frontend is starting up..."
    echo "   You can check the logs with: docker-compose -f docker-compose.prod.yml logs frontend"
fi

# Check Nginx
if curl -f http://localhost:80 > /dev/null 2>&1; then
    echo "✅ Nginx is ready"
else
    echo "⏳ Nginx is starting up..."
    echo "   You can check the logs with: docker-compose -f docker-compose.prod.yml logs nginx"
fi

echo ""
echo "🎉 Production environment setup complete!"
echo ""
echo "📊 Services:"
echo "   • PostgreSQL: localhost:5432"
echo "   • Redis: localhost:6379"
echo "   • Backend API: http://localhost:8080"
echo "   • Frontend: http://localhost:3000"
echo "   • Nginx (Reverse Proxy): http://localhost:80"
echo ""
echo "🔧 Useful commands:"
echo "   • View logs: docker-compose -f docker-compose.prod.yml logs -f"
echo "   • Stop services: docker-compose -f docker-compose.prod.yml down"
echo "   • Restart services: docker-compose -f docker-compose.prod.yml restart"
echo "   • Access PostgreSQL: docker-compose -f docker-compose.prod.yml exec postgres psql -U portfolio_user -d web3_portfolio"
echo "   • Access Redis: docker-compose -f docker-compose.prod.yml exec redis redis-cli"
echo ""
echo "📝 Database credentials:"
echo "   • Database: web3_portfolio"
echo "   • Username: portfolio_user"
echo "   • Password: secure_password_123"
echo ""
echo "🌐 Open your browser and navigate to:"
echo "   • Frontend: http://localhost:3000"
echo "   • API Documentation: http://localhost:8080/docs"
echo "   • Health Check: http://localhost:8080/health"
echo ""
echo "🔒 Security Notes:"
echo "   • Change default passwords in production"
echo "   • Set up SSL certificates for HTTPS"
echo "   • Configure firewall rules"
echo "   • Set up monitoring and logging"
echo "   • Regular database backups" 