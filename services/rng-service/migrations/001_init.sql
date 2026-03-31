-- RNG Service Database Schema

-- RNG Seeds table: stores seed values for random number generation
CREATE TABLE IF NOT EXISTS rng_seeds (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    seed_id VARCHAR(255) UNIQUE NOT NULL,
    server_seed VARCHAR(512) NOT NULL,
    server_seed_hash VARCHAR(512) NOT NULL,
    client_seed VARCHAR(512),
    nonce INTEGER DEFAULT 0,
    combined_hash VARCHAR(512),
    status VARCHAR(50) DEFAULT 'ACTIVE', -- ACTIVE, USED, EXPIRED, REVEALED
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    revealed_at TIMESTAMP WITH TIME ZONE,
    expires_at TIMESTAMP WITH TIME ZONE
);

CREATE INDEX idx_rng_seeds_seed_id ON rng_seeds(seed_id);
CREATE INDEX idx_rng_seeds_server_seed_hash ON rng_seeds(server_seed_hash);
CREATE INDEX idx_rng_seeds_status ON rng_seeds(status);
CREATE INDEX idx_rng_seeds_created_at ON rng_seeds(created_at);

-- RNG Results table: stores generated random values
CREATE TABLE IF NOT EXISTS rng_results (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    result_id VARCHAR(255) UNIQUE NOT NULL,
    seed_id VARCHAR(255) NOT NULL REFERENCES rng_seeds(seed_id),
    service_name VARCHAR(100) NOT NULL,
    request_type VARCHAR(100) NOT NULL,
    random_value VARCHAR(512) NOT NULL,
    random_integer INTEGER,
    random_float DECIMAL(20, 18),
    range_min INTEGER,
    range_max INTEGER,
    metadata JSONB DEFAULT '{}',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

CREATE INDEX idx_rng_results_result_id ON rng_results(result_id);
CREATE INDEX idx_rng_results_seed_id ON rng_results(seed_id);
CREATE INDEX idx_rng_results_service_name ON rng_results(service_name);
CREATE INDEX idx_rng_results_created_at ON rng_results(created_at);

-- Provably Fair Proofs table: stores verification data for provably fair gaming
CREATE TABLE IF NOT EXISTS provably_fair_proofs (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    proof_id VARCHAR(255) UNIQUE NOT NULL,
    seed_id VARCHAR(255) NOT NULL REFERENCES rng_seeds(seed_id),
    result_id VARCHAR(255) NOT NULL REFERENCES rng_results(result_id),
    player_id VARCHAR(255) NOT NULL,
    game_type VARCHAR(100) NOT NULL,
    server_seed VARCHAR(512) NOT NULL,
    server_seed_hash VARCHAR(512) NOT NULL,
    client_seed VARCHAR(512) NOT NULL,
    nonce INTEGER NOT NULL,
    combined_hash VARCHAR(512) NOT NULL,
    verified BOOLEAN DEFAULT FALSE,
    verification_url VARCHAR(512),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    verified_at TIMESTAMP WITH TIME ZONE
);

CREATE INDEX idx_provably_fair_proofs_proof_id ON provably_fair_proofs(proof_id);
CREATE INDEX idx_provably_fair_proofs_seed_id ON provably_fair_proofs(seed_id);
CREATE INDEX idx_provably_fair_proofs_result_id ON provably_fair_proofs(result_id);
CREATE INDEX idx_provably_fair_proofs_player_id ON provably_fair_proofs(player_id);
CREATE INDEX idx_provably_fair_proofs_game_type ON provably_fair_proofs(game_type);
CREATE INDEX idx_provably_fair_proofs_created_at ON provably_fair_proofs(created_at);

-- Trigger to update updated_at on rng_seeds changes
CREATE TRIGGER update_rng_seeds_updated_at BEFORE UPDATE ON rng_seeds
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
