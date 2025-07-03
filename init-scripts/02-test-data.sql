-- Test Data for Web3 Portfolio Dashboard
-- This script populates the database with sample data for testing

-- Insert test users
INSERT INTO users (id, email, password_hash, first_name, last_name, subscription_tier, subscription_status, created_at, updated_at) VALUES
('550e8400-e29b-41d4-a716-446655440001', 'test@web3portfolio.dev', '$2a$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi', 'John', 'Doe', 'premium', 'active', NOW(), NOW()),
('550e8400-e29b-41d4-a716-446655440002', 'alice@web3portfolio.dev', '$2a$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi', 'Alice', 'Smith', 'pro', 'active', NOW(), NOW()),
('550e8400-e29b-41d4-a716-446655440003', 'bob@web3portfolio.dev', '$2a$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi', 'Bob', 'Johnson', 'basic', 'active', NOW(), NOW())
ON CONFLICT (id) DO NOTHING;

-- Insert test wallets
INSERT INTO wallets (id, user_id, name, address, network, wallet_type, is_active, created_at, updated_at) VALUES
('660e8400-e29b-41d4-a716-446655440001', '550e8400-e29b-41d4-a716-446655440001', 'Main Ethereum Wallet', '0x742d35Cc6634C0532925a3b8D4C9db96C4b4d8b6', 'ethereum', 'metamask', true, NOW(), NOW()),
('660e8400-e29b-41d4-a716-446655440002', '550e8400-e29b-41d4-a716-446655440001', 'Polygon Wallet', '0x742d35Cc6634C0532925a3b8D4C9db96C4b4d8b7', 'polygon', 'metamask', true, NOW(), NOW()),
('660e8400-e29b-41d4-a716-446655440003', '550e8400-e29b-41d4-a716-446655440002', 'Alice ETH Wallet', '0x742d35Cc6634C0532925a3b8D4C9db96C4b4d8b8', 'ethereum', 'metamask', true, NOW(), NOW()),
('660e8400-e29b-41d4-a716-446655440004', '550e8400-e29b-41d4-a716-446655440003', 'Bob BSC Wallet', '0x742d35Cc6634C0532925a3b8D4C9db96C4b4d8b9', 'bsc', 'metamask', true, NOW(), NOW())
ON CONFLICT (id) DO NOTHING;

-- Insert test tokens
INSERT INTO tokens (id, symbol, name, network, contract_address, decimals, logo_url, price_usd, market_cap, volume_24h, created_at, updated_at) VALUES
('770e8400-e29b-41d4-a716-446655440001', 'ETH', 'Ethereum', 'ethereum', '0x0000000000000000000000000000000000000000', 18, 'https://assets.coingecko.com/coins/images/279/large/ethereum.png', 2500.00, 300000000000, 15000000000, NOW(), NOW()),
('770e8400-e29b-41d4-a716-446655440002', 'USDC', 'USD Coin', 'ethereum', '0xA0b86a33E6441b8C4C8C8C8C8C8C8C8C8C8C8C8', 6, 'https://assets.coingecko.com/coins/images/6319/large/USD_Coin_icon.png', 1.00, 25000000000, 5000000000, NOW(), NOW()),
('770e8400-e29b-41d4-a716-446655440003', 'MATIC', 'Polygon', 'polygon', '0x0000000000000000000000000000000000001010', 18, 'https://assets.coingecko.com/coins/images/4713/large/matic-token-icon.png', 0.85, 8000000000, 300000000, NOW(), NOW()),
('770e8400-e29b-41d4-a716-446655440004', 'BNB', 'BNB', 'bsc', '0x0000000000000000000000000000000000000000', 18, 'https://assets.coingecko.com/coins/images/825/large/bnb-icon2_2x.png', 320.00, 50000000000, 2000000000, NOW(), NOW()),
('770e8400-e29b-41d4-a716-446655440005', 'LINK', 'Chainlink', 'ethereum', '0x514910771AF9Ca656af840dff83E8264EcF986CA', 18, 'https://assets.coingecko.com/coins/images/877/large/chainlink.png', 15.50, 9000000000, 800000000, NOW(), NOW())
ON CONFLICT (id) DO NOTHING;

-- Insert test holdings
INSERT INTO holdings (id, wallet_id, token_id, balance, value_usd, created_at, updated_at) VALUES
('880e8400-e29b-41d4-a716-446655440001', '660e8400-e29b-41d4-a716-446655440001', '770e8400-e29b-41d4-a716-446655440001', 2.5, 6250.00, NOW(), NOW()),
('880e8400-e29b-41d4-a716-446655440002', '660e8400-e29b-41d4-a716-446655440001', '770e8400-e29b-41d4-a716-446655440002', 5000.00, 5000.00, NOW(), NOW()),
('880e8400-e29b-41d4-a716-446655440003', '660e8400-e29b-41d4-a716-446655440002', '770e8400-e29b-41d4-a716-446655440003', 1000.00, 850.00, NOW(), NOW()),
('880e8400-e29b-41d4-a716-446655440004', '660e8400-e29b-41d4-a716-446655440003', '770e8400-e29b-41d4-a716-446655440001', 1.0, 2500.00, NOW(), NOW()),
('880e8400-e29b-41d4-a716-446655440005', '660e8400-e29b-41d4-a716-446655440004', '770e8400-e29b-41d4-a716-446655440004', 10.0, 3200.00, NOW(), NOW()),
('880e8400-e29b-41d4-a716-446655440006', '660e8400-e29b-41d4-a716-446655440001', '770e8400-e29b-41d4-a716-446655440005', 100.0, 1550.00, NOW(), NOW())
ON CONFLICT (id) DO NOTHING;

-- Insert test transactions
INSERT INTO transactions (id, wallet_id, token_id, tx_hash, tx_type, amount, value_usd, gas_fee, block_number, timestamp, created_at, updated_at) VALUES
('990e8400-e29b-41d4-a716-446655440001', '660e8400-e29b-41d4-a716-446655440001', '770e8400-e29b-41d4-a716-446655440001', '0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef', 'buy', 1.0, 2500.00, 50.00, 18000000, EXTRACT(EPOCH FROM NOW() - INTERVAL '7 days'), NOW(), NOW()),
('990e8400-e29b-41d4-a716-446655440002', '660e8400-e29b-41d4-a716-446655440001', '770e8400-e29b-41d4-a716-446655440002', '0x2345678901bcdef1234567890abcdef1234567890abcdef1234567890abcdef', 'buy', 2000.00, 2000.00, 30.00, 18000001, EXTRACT(EPOCH FROM NOW() - INTERVAL '5 days'), NOW(), NOW()),
('990e8400-e29b-41d4-a716-446655440003', '660e8400-e29b-41d4-a716-446655440002', '770e8400-e29b-41d4-a716-446655440003', '0x3456789012cdef1234567890abcdef1234567890abcdef1234567890abcdef', 'buy', 500.00, 425.00, 5.00, 45000000, EXTRACT(EPOCH FROM NOW() - INTERVAL '3 days'), NOW(), NOW()),
('990e8400-e29b-41d4-a716-446655440004', '660e8400-e29b-41d4-a716-446655440001', '770e8400-e29b-41d4-a716-446655440001', '0x4567890123def1234567890abcdef1234567890abcdef1234567890abcdef', 'sell', 0.5, 1250.00, 45.00, 18000002, EXTRACT(EPOCH FROM NOW() - INTERVAL '1 day'), NOW(), NOW())
ON CONFLICT (id) DO NOTHING;

-- Insert test alerts
INSERT INTO alerts (id, user_id, name, condition_type, token_symbol, threshold_value, is_active, created_at, updated_at) VALUES
('aa0e8400-e29b-41d4-a716-446655440001', '550e8400-e29b-41d4-a716-446655440001', 'ETH Price Alert', 'price_above', 'ETH', 3000.00, true, NOW(), NOW()),
('aa0e8400-e29b-41d4-a716-446655440002', '550e8400-e29b-41d4-a716-446655440001', 'Portfolio Value Alert', 'value_below', 'PORTFOLIO', 10000.00, true, NOW(), NOW()),
('aa0e8400-e29b-41d4-a716-446655440003', '550e8400-e29b-41d4-a716-446655440002', 'MATIC Price Alert', 'price_below', 'MATIC', 0.80, true, NOW(), NOW())
ON CONFLICT (id) DO NOTHING;

-- Insert test alert history
INSERT INTO alert_history (id, alert_id, triggered_value, triggered_at, created_at) VALUES
('bb0e8400-e29b-41d4-a716-446655440001', 'aa0e8400-e29b-41d4-a716-446655440001', 3050.00, NOW() - INTERVAL '2 hours', NOW()),
('bb0e8400-e29b-41d4-a716-446655440002', 'aa0e8400-e29b-41d4-a716-446655440003', 0.75, NOW() - INTERVAL '1 day', NOW())
ON CONFLICT (id) DO NOTHING;

-- Log successful test data insertion
DO $$
BEGIN
    RAISE NOTICE 'Test data inserted successfully for Web3 Portfolio Dashboard';
    RAISE NOTICE 'Test users: 3 users created';
    RAISE NOTICE 'Test wallets: 4 wallets created';
    RAISE NOTICE 'Test tokens: 5 tokens created';
    RAISE NOTICE 'Test holdings: 6 holdings created';
    RAISE NOTICE 'Test transactions: 4 transactions created';
    RAISE NOTICE 'Test alerts: 3 alerts created';
END $$; 