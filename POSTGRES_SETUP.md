# PostgreSQL Setup Guide for Web3 Portfolio Dashboard

This guide will help you set up a local PostgreSQL database for development.

## Prerequisites

- Windows 10/11
- Administrator privileges
- Internet connection for download

## Step 1: Install PostgreSQL

### 1.1 Download PostgreSQL
1. Go to: https://www.postgresql.org/download/windows/
2. Click "Download the installer"
3. Choose the latest version (currently 16.x)
4. Download the Windows x86-64 installer

### 1.2 Install PostgreSQL
1. **Run the installer as Administrator**
2. **Installation Directory:** `C:\Program Files\PostgreSQL\16`
3. **Data Directory:** `C:\Program Files\PostgreSQL\16\data`
4. **Password:** `postgres123` (remember this!)
5. **Port:** `5432` (default)
6. **Locale:** `Default locale`
7. **Stack Builder:** Uncheck (not needed for development)

### 1.3 Add to PATH
After installation, add PostgreSQL to your system PATH:
1. Open System Properties → Advanced → Environment Variables
2. Edit the `Path` variable
3. Add: `C:\Program Files\PostgreSQL\16\bin`
4. Click OK and restart your terminal

## Step 2: Create Database and User

### Option A: Using the Setup Script (Recommended)

1. **Run the setup script:**
   ```powershell
   # From project root
   .\scripts\setup-postgres.ps1
   ```

2. **Follow the prompts:**
   - Enter your PostgreSQL superuser password (default: `postgres123`)
   - The script will create everything automatically

### Option B: Manual Setup

1. **Open pgAdmin** (installed with PostgreSQL)
2. **Connect to the server:**
   - Host: `localhost`
   - Port: `5432`
   - Username: `postgres`
   - Password: `postgres123`

3. **Create Database:**
   ```sql
   CREATE DATABASE web3_portfolio_dev;
   ```

4. **Create User:**
   ```sql
   CREATE USER dev_user WITH PASSWORD 'dev_password123';
   ```

5. **Grant Privileges:**
   ```sql
   GRANT ALL PRIVILEGES ON DATABASE web3_portfolio_dev TO dev_user;
   ```

6. **Connect to the new database and grant schema privileges:**
   ```sql
   \c web3_portfolio_dev
   GRANT ALL ON SCHEMA public TO dev_user;
   GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO dev_user;
   GRANT ALL PRIVILEGES ON ALL SEQUENCES IN SCHEMA public TO dev_user;
   ```

## Step 3: Update Environment Configuration

1. **Open `backend/env.dev`**
2. **Update the DATABASE_URL:**
   ```env
   DATABASE_URL=postgres://dev_user:dev_password123@localhost:5432/web3_portfolio_dev?sslmode=disable
   ```

## Step 4: Test Database Connection

```powershell
cd backend
go run main.go --test-db
```

You should see: `✅ Database connection test successful!`

## Step 5: Run Database Migrations

```powershell
cd backend
go run main.go --migrate
```

You should see: `✅ Database migrations completed successfully!`

## Step 6: Start Development Servers

### Start Backend
```powershell
cd backend
go run main.go
```

### Start Frontend (in a new terminal)
```powershell
cd frontend
npm run dev
```

## Troubleshooting

### PostgreSQL Not Found
**Error:** `psql : The term 'psql' is not recognized`
**Solution:**
1. Add PostgreSQL to PATH: `C:\Program Files\PostgreSQL\16\bin`
2. Restart your terminal
3. Or restart your computer

### Connection Refused
**Error:** `connectex: No connection could be made because the target machine actively refused it`
**Solution:**
1. Check if PostgreSQL service is running:
   ```powershell
   Get-Service postgresql*
   ```
2. Start the service if stopped:
   ```powershell
   Start-Service postgresql*
   ```

### Authentication Failed
**Error:** `FATAL: password authentication failed`
**Solution:**
1. Check your password in the connection string
2. Reset PostgreSQL password if needed:
   ```powershell
   # Connect as postgres user
   psql -U postgres -h localhost
   # Change password
   ALTER USER postgres PASSWORD 'new_password';
   ```

### Database Already Exists
**Error:** `database "web3_portfolio_dev" already exists`
**Solution:**
- This is OK! The database already exists
- You can continue with the next steps

### User Already Exists
**Error:** `role "dev_user" already exists`
**Solution:**
- This is OK! The user already exists
- You can continue with the next steps

## Database Connection Details

- **Host:** localhost
- **Port:** 5432
- **Database:** web3_portfolio_dev
- **Username:** dev_user
- **Password:** dev_password123
- **Connection String:** `postgres://dev_user:dev_password123@localhost:5432/web3_portfolio_dev?sslmode=disable`

## Useful Commands

```powershell
# Test database connection
cd backend; go run main.go --test-db

# Run migrations
cd backend; go run main.go --migrate

# Connect to database with psql
psql -U dev_user -h localhost -d web3_portfolio_dev

# List all databases
psql -U postgres -h localhost -c "\l"

# List all users
psql -U postgres -h localhost -c "\du"
```

## Next Steps

1. **Test the full application:**
   - Backend: http://localhost:8080/health
   - Frontend: http://localhost:3000

2. **Create test data:**
   - Register a new user
   - Add portfolio addresses
   - Test Web3 integration

3. **Development workflow:**
   - Make changes to code
   - Test with the local database
   - Deploy to production with your preferred cloud database

## Support

If you encounter issues:
1. Check the troubleshooting section above
2. Verify PostgreSQL is running: `Get-Service postgresql*`
3. Test connection manually: `psql -U dev_user -h localhost -d web3_portfolio_dev`
4. Check logs: Event Viewer → Windows Logs → Application 