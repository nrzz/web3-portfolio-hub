#!/bin/bash

# Web3 Portfolio Dashboard - Development Setup Script
# This script sets up the development environment with PostgreSQL and Redis

set -e

echo "🚀 Setting up Web3 Portfolio Dashboard Development Environment..."

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

echo "✅ Docker and Docker Compose are available"

# Create necessary directories
echo "📁 Creating necessary directories..."
mkdir -p init-scripts
mkdir -p logs

# Copy environment file if it doesn't exist
if [ ! -f "backend/.env" ]; then
    echo "📝 Creating backend environment file..."
    cp backend/env.dev backend/.env
    echo "✅ Backend environment file created"
else
    echo "✅ Backend environment file already exists"
fi

# Stop any existing containers
echo "🛑 Stopping any existing containers..."
docker-compose down --remove-orphans

# Build and start the services
echo "🔨 Building and starting services..."
docker-compose up --build -d

# Wait for services to be healthy
echo "⏳ Waiting for services to be ready..."
sleep 10

# Check service health
echo "🔍 Checking service health..."

# Check PostgreSQL
if docker-compose exec -T postgres pg_isready -U dev_user -d web3_portfolio_dev > /dev/null 2>&1; then
    echo "✅ PostgreSQL is ready"
else
    echo "❌ PostgreSQL is not ready. Check logs with: docker-compose logs postgres"
fi

# Check Redis
if docker-compose exec -T redis redis-cli ping > /dev/null 2>&1; then
    echo "✅ Redis is ready"
else
    echo "❌ Redis is not ready. Check logs with: docker-compose logs redis"
fi

# Check Backend
if curl -f http://localhost:8080/health > /dev/null 2>&1; then
    echo "✅ Backend API is ready"
else
    echo "⏳ Backend API is starting up..."
    echo "   You can check the logs with: docker-compose logs backend"
fi

echo ""
echo "🎉 Development environment setup complete!"
echo ""
echo "📊 Services:"
echo "   • PostgreSQL: localhost:5432"
echo "   • Redis: localhost:6379"
echo "   • Backend API: http://localhost:8080"
echo "   • Frontend: http://localhost:3000"
echo "   • pgAdmin: http://localhost:5050"
echo ""
echo "🔧 Useful commands:"
echo "   • View logs: docker-compose logs -f"
echo "   • Stop services: docker-compose down"
echo "   • Restart services: docker-compose restart"
echo "   • Access PostgreSQL: docker-compose exec postgres psql -U dev_user -d web3_portfolio_dev"
echo "   • Access Redis: docker-compose exec redis redis-cli"
echo ""
echo "📝 Database credentials:"
echo "   • Database: web3_portfolio_dev"
echo "   • Username: dev_user"
echo "   • Password: dev_password"
echo ""
echo "🔐 pgAdmin credentials:"
echo "   • Email: admin@web3portfolio.dev"
echo "   • Password: admin123"
echo ""
echo "🌐 Open your browser and navigate to:"
echo "   • Frontend: http://localhost:3000"
echo "   • API Documentation: http://localhost:8080/docs"
echo "   • Database Admin: http://localhost:5050" 