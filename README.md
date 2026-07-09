# Web3 Portfolio Dashboard

[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
[![Go](https://img.shields.io/badge/Go-1.23+-00ADD8)](https://go.dev/)
[![React](https://img.shields.io/badge/React-18-61DAFB)](https://react.dev/)

A **full-stack Web3 portfolio tracker** built with Go and React. Track crypto portfolios across Ethereum, Polygon, and BSC with portfolio management, analytics, and price alerts. A community forum UI exists but backend endpoints are **stubs** (not yet wired to persistence).

> **Status:** Development-ready with Docker Compose. Not production-hardened — see [Limitations](#limitations) below.

## Quick Start

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
docker-compose -f docker-compose.dev.yml up --build -d
# Wait 30–60 seconds, then open http://localhost:3000
```

## Application Access

| Service | URL | Description |
|---------|-----|-------------|
| **Frontend** | http://localhost:3000 | Main application UI |
| **Backend API** | http://localhost:8080 | REST API (`/api/v1/*`) |
| **Health** | http://localhost:8080/health or `/api/health` | Service health check |
| **pgAdmin** | http://localhost:5050 | Database management (dev compose only) |

There is **no Swagger UI** — API routes are documented below and in handler code.

## Test Credentials

- **Email:** `test@web3portfolio.dev`
- **Password:** `password`
- **Subscription:** Premium tier (seeded test data)

## Pre-Loaded Test Data (Docker dev)

- 3 test users with different subscription tiers
- Sample wallets across Ethereum, Polygon, and BSC
- Token holdings and transactions (seeded via `init-scripts/`)

## Architecture

```
┌─────────────────┐    ┌─────────────────┐    ┌─────────────────┐
│   Frontend      │    │   Backend       │    │   PostgreSQL    │
│   (React + TS)  │◄──►│   (Go + Gin)    │◄──►│   Port: 5432    │
│   Port: 3000    │    │   Port: 8080    │    └─────────────────┘
└─────────────────┘    └─────────────────┘
```

Redis is included in Docker Compose for future caching but is **not used** by the Go backend today.

## Technology Stack

### Frontend
- React 18, TypeScript, Vite, Tailwind CSS

### Backend
- Go 1.23, Gin, PostgreSQL, JWT authentication

### Infrastructure
- Docker Compose, Nginx (production frontend proxy)

## Features

| Feature | Status |
|---------|--------|
| Auth (register/login/JWT) | Implemented |
| Portfolio & wallet management | Implemented |
| Web3 balance/price lookups | Implemented (RPC-dependent) |
| Analytics & alerts | Implemented (Pro/Premium tier gating) |
| Community forum | **Stub** — routes return placeholder responses |
| Redis caching | **Not wired** — env var exists, unused |
| Email notifications | **Not implemented** |

## Development

### Prerequisites
- Docker Desktop (recommended), or Node.js 18+ and Go 1.23+ for local runs

### Local Backend
```bash
cd backend
go mod download
go run main.go
```

### Local Frontend
```bash
cd frontend
npm ci
npm run dev
```

### Docker Development
```bash
docker-compose -f docker-compose.dev.yml up --build
docker-compose -f docker-compose.dev.yml logs -f
docker-compose -f docker-compose.dev.yml down
```

## API Overview

Base path: `/api/v1` (auth required except public routes).

### Public
- `GET /health` — health check (alias: `GET /api/health`)
- `GET /api/version` — version info

### Authentication
```bash
POST /api/v1/auth/register   # { "email", "password" }
POST /api/v1/auth/login      # { "email", "password" }
```

### Portfolios (authenticated)
```bash
GET    /api/v1/portfolios
POST   /api/v1/portfolios
GET    /api/v1/portfolios/:id
GET    /api/v1/portfolios/:id/balances
POST   /api/v1/portfolios/:id/addresses
```

### User profile (authenticated)
```bash
GET /api/v1/user/profile
GET /api/v1/user/subscription
```

### Forum (stub)
```bash
GET  /api/v1/forum/questions      # returns empty list
POST /api/v1/forum/questions      # returns stub message
```

## Testing

```bash
# Backend
cd backend && go test ./...

# Frontend
cd frontend && npm run test -- --run

# Health check
curl http://localhost:8080/health
curl http://localhost:8080/api/health
```

## Configuration

### Required environment variables
```bash
DATABASE_URL=postgres://user:pass@host:5432/db
JWT_SECRET=your-secret-key
CORS_ALLOWED_ORIGINS=http://localhost:3000,http://localhost:3001
ETHEREUM_RPC_URL=https://...
```

Copy `backend/env.example` to `backend/.env` and adjust values.

### CORS
Set `CORS_ALLOWED_ORIGINS` to a comma-separated list of allowed frontend origins. The API reflects the request `Origin` when it matches — credentials are supported without using `*`.

## Limitations

- Forum backend is stub-only; frontend pages exist but data is not persisted.
- Redis is declared in compose but unused in application code.
- Analytics dashboard uses some mock/demo data on the frontend.
- Production deployment requires you to set strong secrets, RPC URLs, and CORS origins.

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make changes with tests where applicable
4. Submit a pull request

## License

MIT — see [LICENSE](LICENSE).
