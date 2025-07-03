# PostgreSQL Installation Guide for Windows

## Step 1: Download PostgreSQL

1. **Open your browser** and go to: https://www.postgresql.org/download/windows/

2. **Click "Download the installer"**

3. **Choose PostgreSQL 16.x** (latest version)

4. **Download the Windows x86-64 installer**

## Step 2: Install PostgreSQL

1. **Right-click the downloaded file** and select "Run as administrator"

2. **Click "Next"** through the welcome screen

3. **Choose installation directory:**
   - Keep default: `C:\Program Files\PostgreSQL\16`
   - Click "Next"

4. **Choose data directory:**
   - Keep default: `C:\Program Files\PostgreSQL\16\data`
   - Click "Next"

5. **Set password:**
   - **Enter:** `postgres123`
   - **Confirm:** `postgres123`
   - **Remember this password!**
   - Click "Next"

6. **Set port:**
   - Keep default: `5432`
   - Click "Next"

7. **Choose locale:**
   - Keep default: `Default locale`
   - Click "Next"

8. **Pre Installation Summary:**
   - Review settings
   - Click "Next"

9. **Installation:**
   - Wait for installation to complete
   - **Uncheck "Stack Builder"** (not needed)
   - Click "Finish"

## Step 3: Add PostgreSQL to PATH

1. **Open System Properties:**
   - Press `Windows + R`
   - Type `sysdm.cpl`
   - Press Enter

2. **Go to Advanced tab:**
   - Click "Environment Variables"

3. **Edit System Variables:**
   - Find "Path" in System Variables
   - Select it and click "Edit"

4. **Add PostgreSQL path:**
   - Click "New"
   - Add: `C:\Program Files\PostgreSQL\16\bin`
   - Click "OK" on all dialogs

5. **Restart your terminal/PowerShell**

## Step 4: Verify Installation

1. **Open a new PowerShell window**

2. **Test PostgreSQL:**
   ```powershell
   psql --version
   ```

3. **You should see something like:**
   ```
   psql (PostgreSQL) 16.x
   ```

## Step 5: Start PostgreSQL Service

1. **Check if service is running:**
   ```powershell
   Get-Service postgresql*
   ```

2. **If not running, start it:**
   ```powershell
   Start-Service postgresql*
   ```

## Step 6: Test Connection

1. **Test connection to PostgreSQL:**
   ```powershell
   psql -U postgres -h localhost
   ```

2. **Enter password when prompted:** `postgres123`

3. **You should see:**
   ```
   postgres=#
   ```

4. **Exit PostgreSQL:**
   ```sql
   \q
   ```

## Step 7: Run Database Setup

After PostgreSQL is installed and running:

1. **Go to your project directory:**
   ```powershell
   cd D:\Projects\web3-portfolio-dashboard
   ```

2. **Run the setup script:**
   ```powershell
   .\scripts\setup-postgres.ps1
   ```

3. **Follow the prompts** (use password: `postgres123`)

## Troubleshooting

### "psql is not recognized"
- PostgreSQL is not in PATH
- Restart your terminal after adding to PATH
- Or restart your computer

### "Connection refused"
- PostgreSQL service is not running
- Start it with: `Start-Service postgresql*`

### "Password authentication failed"
- Check your password
- Default is: `postgres123`

### "Port 5432 is already in use"
- Another PostgreSQL instance is running
- Stop it or use a different port

## Next Steps

After successful installation:

1. **Run database setup:**
   ```powershell
   .\scripts\setup-postgres.ps1
   ```

2. **Test connection:**
   ```powershell
   cd backend
   go run main.go --test-db
   ```

3. **Run migrations:**
   ```powershell
   go run main.go --migrate
   ```

4. **Start development:**
   ```powershell
   go run main.go
   ```

## Support

If you encounter issues:
1. Make sure you ran the installer as Administrator
2. Check that the service is running: `Get-Service postgresql*`
3. Verify PATH is set correctly
4. Restart your terminal after installation 