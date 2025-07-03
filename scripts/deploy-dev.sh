#!/bin/bash

# Development Deployment Script
set -e

echo "🚀 Starting Web3 Portfolio Dashboard Development Deployment..."

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

# Stop any existing containers
echo "🛑 Stopping existing containers..."
docker-compose -f docker-compose.dev.yml down --remove-orphans

# Remove old volumes if requested
if [ "$1" = "--clean" ]; then
    echo "🧹 Cleaning up old volumes..."
    docker-compose -f docker-compose.dev.yml down -v
fi

# Build and start services
echo "🔨 Building and starting services..."
docker-compose -f docker-compose.dev.yml up --build -d

# Wait for services to be ready
echo "⏳ Waiting for services to be ready..."
sleep 30

# Check service health
echo "🏥 Checking service health..."
docker-compose -f docker-compose.dev.yml ps

# Show logs
echo "📋 Recent logs:"
docker-compose -f docker-compose.dev.yml logs --tail=20

echo "✅ Development deployment completed!"
echo ""
echo "🌐 Services available at:"
echo "   Frontend: http://localhost:3000"
echo "   Backend API: http://localhost:8080"
echo "   pgAdmin: http://localhost:5050 (admin@web3portfolio.dev / admin123)"
echo ""
echo "📊 To view logs: docker-compose -f docker-compose.dev.yml logs -f"
echo "🛑 To stop: docker-compose -f docker-compose.dev.yml down" 