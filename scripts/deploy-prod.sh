#!/bin/bash

# Production Deployment Script
set -e

echo "🚀 Starting Web3 Portfolio Dashboard Production Deployment..."

# Check if Docker is running
if ! docker info > /dev/null 2>&1; then
    echo "❌ Docker is not running. Please start Docker and try again."
    exit 1
fi

# Check if docker-compose is available
if ! command -v docker-compose &> /dev/null; then
    echo "❌ docker-compose is not installed. Please install it and try again."
    exit 1
fi

# Check if .env.production exists
if [ ! -f ".env.production" ]; then
    echo "❌ .env.production file not found!"
    echo "📝 Please copy env.production.example to .env.production and update the values."
    exit 1
fi

# Load environment variables
echo "📋 Loading environment variables..."
set -a
source .env.production
set +a

# Validate required environment variables
required_vars=("POSTGRES_PASSWORD" "REDIS_PASSWORD" "JWT_SECRET")
for var in "${required_vars[@]}"; do
    if [ -z "${!var}" ]; then
        echo "❌ Required environment variable $var is not set!"
        exit 1
    fi
done

# Stop any existing containers
echo "🛑 Stopping existing containers..."
docker-compose -f docker-compose.prod.yml down --remove-orphans

# Remove old volumes if requested
if [ "$1" = "--clean" ]; then
    echo "🧹 Cleaning up old volumes..."
    docker-compose -f docker-compose.prod.yml down -v
fi

# Build and start services
echo "🔨 Building and starting services..."
docker-compose -f docker-compose.prod.yml up --build -d

# Wait for services to be ready
echo "⏳ Waiting for services to be ready..."
sleep 60

# Check service health
echo "🏥 Checking service health..."
docker-compose -f docker-compose.prod.yml ps

# Show logs
echo "📋 Recent logs:"
docker-compose -f docker-compose.prod.yml logs --tail=20

echo "✅ Production deployment completed!"
echo ""
echo "🌐 Services available at:"
echo "   Frontend: http://localhost:3000"
echo "   Backend API: http://localhost:8080"
echo "   Nginx: http://localhost:80"
echo ""
echo "📊 To view logs: docker-compose -f docker-compose.prod.yml logs -f"
echo "🛑 To stop: docker-compose -f docker-compose.prod.yml down"
echo ""
echo "🔒 Security Notes:"
echo "   - Change default passwords in production"
echo "   - Configure SSL certificates for HTTPS"
echo "   - Set up proper firewall rules"
echo "   - Monitor logs for security issues" 