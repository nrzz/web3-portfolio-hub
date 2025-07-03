# Web3 Portfolio Dashboard - Development Setup Guide

This guide will help you set up the development environment with local PostgreSQL database.

## Prerequisites

- Go 1.21+ installed
- Node.js 18+ installed
- npm or yarn installed
- PostgreSQL (local installation)

## Step 1: PostgreSQL Database Setup

### 1.1 Install PostgreSQL
1. Download PostgreSQL from: https://www.postgresql.org/download/windows/
2. Run the installer as Administrator
3. Use password: `postgres123` (remember this!)
4. Keep default port: `5432`

### 1.2 Create Database and User
1. Run the setup script:
   ```powershell
   .\scripts\setup-postgres.ps1
   ```
2. Follow the prompts (use password: `postgres123`)

### 1.3 Environment Configuration
The environment is already configured for local PostgreSQL in `backend/env.dev`:
   ```env
   DATABASE_URL=postgres://dev_user:dev_password123@localhost:5432/web3_portfolio_dev?sslmode=disable
   ```

## Step 2: Install Dependencies

### 2.1 Backend Dependencies
```powershell
cd backend
go mod tidy
```

### 2.2 Frontend Dependencies
```powershell
cd frontend
npm install
```

## Step 3: Test Database Connection

```powershell
cd backend
go run main.go --test-db
```

You should see: `✅ Database connection test successful!`

## Step 4: Run Database Migrations

```powershell
cd backend
go run main.go --migrate
```

You should see: `✅ Database migrations completed successfully!`

## Step 5: Start Development Servers

### Option A: Using the Setup Script (Recommended)
```powershell
# From project root
.\scripts\dev-setup.ps1
```

### Option B: Manual Start

#### Start Backend
```powershell
cd backend
go run main.go
```

#### Start Frontend (in a new terminal)
```powershell
cd frontend
npm run dev
```

## Step 6: Verify Setup

1. **Backend Health Check**: http://localhost:8080/health
2. **Frontend**: http://localhost:3000
3. **API Documentation**: http://localhost:8080/api/v1

## Troubleshooting

### Database Connection Issues

**Error**: `failed to connect to PostgreSQL database`
- Check if PostgreSQL is installed and running
- Verify the connection string in `backend/env.dev`
- Make sure the database and user exist

**Error**: `connection timeout`
- PostgreSQL service is not running
- Start it with: `Start-Service postgresql*`

### Port Conflicts

**Error**: `Only one usage of each socket address is normally permitted`
- Kill existing processes:
  ```powershell
  Stop-Process -Name "go" -Force -ErrorAction SilentlyContinue
  Stop-Process -Name "node" -Force -ErrorAction SilentlyContinue
  ```

### PowerShell Issues

**Error**: `The token '&&' is not a valid statement separator`
- Use `;` instead of `&&` in PowerShell:
  ```powershell
  cd backend; go run main.go
  ```

## Database Schema

The application creates the following tables:
- `users` - User accounts and authentication
- `portfolios` - User portfolios
- `addresses` - Blockchain addresses in portfolios
- `transactions` - Blockchain transactions
- `alerts` - User alerts and notifications
- `balances` - Token balances

## Development Workflow

1. **Start development**: Run `.\scripts\dev-setup.ps1`
2. **Make changes**: Edit files in `frontend/src` or `backend/internal`
3. **Test changes**: Frontend auto-reloads, backend needs restart
4. **Stop development**: Press `Ctrl+C` in the terminal

## Useful Commands

```powershell
# Test database connection
cd backend; go run main.go --test-db

# Run migrations
cd backend; go run main.go --migrate

# Start backend only
cd backend; go run main.go

# Start frontend only
cd frontend; npm run dev

# Install dependencies
cd backend; go mod tidy
cd frontend; npm install
```

## Environment Variables

Key environment variables in `backend/env.dev`:
- `DATABASE_URL` - Local PostgreSQL connection string
- `JWT_SECRET` - Secret for JWT token signing
- `PORT` - Backend server port (default: 8080)
- `CORS_ALLOWED_ORIGINS` - Allowed frontend origins

## Next Steps

1. **Authentication**: Test user registration/login
2. **Portfolio Management**: Add blockchain addresses
3. **Web3 Integration**: Connect to blockchain networks
4. **Alerts**: Set up price and transaction alerts

## Support

If you encounter issues:
1. Check the troubleshooting section above
2. Verify PostgreSQL is installed and running
3. Ensure all dependencies are installed
4. Check that ports 8080 and 3000 are available 