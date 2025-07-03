-- Initialize Web3 Portfolio Database
-- This script runs when the PostgreSQL container starts for the first time

-- Create extensions if they don't exist
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE EXTENSION IF NOT EXISTS "pgcrypto";

-- Create a function to update the updated_at timestamp
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ language 'plpgsql';

-- Grant necessary permissions
GRANT ALL PRIVILEGES ON DATABASE web3_portfolio_dev TO dev_user;
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO dev_user;
GRANT ALL PRIVILEGES ON ALL SEQUENCES IN SCHEMA public TO dev_user;

-- Create schema for better organization (optional)
CREATE SCHEMA IF NOT EXISTS web3_portfolio;

-- Set search path
SET search_path TO web3_portfolio, public;

-- Log successful initialization
DO $$
BEGIN
    RAISE NOTICE 'Web3 Portfolio Database initialized successfully';
END $$; 