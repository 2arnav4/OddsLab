-- Enable UUID extension for generating UUIDs
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- 1. Users Table
CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    wallet_address VARCHAR(255) UNIQUE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL
);

-- 2. Markets Table
CREATE TABLE IF NOT EXISTS markets (
    id VARCHAR(255) PRIMARY KEY, -- Polymarket ID
    title VARCHAR(512) NOT NULL,
    description TEXT,
    category VARCHAR(255),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    closes_at TIMESTAMP WITH TIME ZONE NOT NULL,
    resolved_outcome VARCHAR(50), -- e.g., 'YES', 'NO', or NULL if unresolved
    polymarket_url VARCHAR(1024)
);

-- 3. Market Prices Table (for tracking price history)
CREATE TABLE IF NOT EXISTS market_prices (
    id BIGSERIAL PRIMARY KEY,
    market_id VARCHAR(255) REFERENCES markets(id) ON DELETE CASCADE NOT NULL,
    yes_price DECIMAL(5, 4) NOT NULL CHECK (yes_price >= 0.0 AND yes_price <= 1.0),
    no_price DECIMAL(5, 4) NOT NULL CHECK (no_price >= 0.0 AND no_price <= 1.0),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL
);

-- 4. Transactions Table (Immutable Ledger)
CREATE TABLE IF NOT EXISTS transactions (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID REFERENCES users(id) ON DELETE RESTRICT NOT NULL,
    market_id VARCHAR(255) REFERENCES markets(id) ON DELETE RESTRICT NOT NULL,
    outcome VARCHAR(10) NOT NULL CHECK (outcome IN ('YES', 'NO')),
    amount_bet DECIMAL(15, 4) NOT NULL CHECK (amount_bet > 0.0),
    status VARCHAR(20) NOT NULL CHECK (status IN ('pending', 'won', 'lost')),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    settled_at TIMESTAMP WITH TIME ZONE,
    pnl DECIMAL(15, 4)
);

-- 5. Portfolio State Table (Denormalized Cache)
CREATE TABLE IF NOT EXISTS portfolio_state (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID UNIQUE REFERENCES users(id) ON DELETE CASCADE NOT NULL,
    balance DECIMAL(15, 4) NOT NULL DEFAULT 1000.0000 CHECK (balance >= 0.0),
    total_bet DECIMAL(15, 4) NOT NULL DEFAULT 0.0000 CHECK (total_bet >= 0.0),
    total_won DECIMAL(15, 4) NOT NULL DEFAULT 0.0000 CHECK (total_won >= 0.0),
    total_lost DECIMAL(15, 4) NOT NULL DEFAULT 0.0000 CHECK (total_lost >= 0.0),
    win_count INTEGER NOT NULL DEFAULT 0 CHECK (win_count >= 0),
    loss_count INTEGER NOT NULL DEFAULT 0 CHECK (loss_count >= 0),
    best_bet_roi DECIMAL(10, 4) NOT NULL DEFAULT 0.0000,
    worst_bet_roi DECIMAL(10, 4) NOT NULL DEFAULT 0.0000,
    cache_updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL
);

-- Indexes for performance optimization
CREATE INDEX IF NOT EXISTS idx_users_email ON users(email);
CREATE INDEX IF NOT EXISTS idx_markets_closes_at ON markets(closes_at);
CREATE INDEX IF NOT EXISTS idx_market_prices_market_id_updated_at ON market_prices(market_id, updated_at DESC);
CREATE INDEX IF NOT EXISTS idx_transactions_user_id ON transactions(user_id);
CREATE INDEX IF NOT EXISTS idx_transactions_market_id ON transactions(market_id);
CREATE INDEX IF NOT EXISTS idx_transactions_status ON transactions(status);
CREATE INDEX IF NOT EXISTS idx_portfolio_state_user_id ON portfolio_state(user_id);
