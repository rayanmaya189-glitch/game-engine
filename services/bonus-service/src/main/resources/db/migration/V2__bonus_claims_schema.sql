-- Bonus Claims table for tracking user bonus claims
CREATE TABLE IF NOT EXISTS bonus_claims (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL,
    bonus_id UUID NOT NULL,
    bonus_amount DECIMAL(19, 2) NOT NULL,
    wagering_requirement DECIMAL(19, 2) NOT NULL,
    wagering_contributed DECIMAL(19, 2) DEFAULT 0,
    winnings_amount DECIMAL(19, 2),
    status VARCHAR(20) NOT NULL DEFAULT 'ACTIVE',
    claimed_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    completed_at TIMESTAMP WITH TIME ZONE,
    expires_at TIMESTAMP WITH TIME ZONE,
    cancelled_at TIMESTAMP WITH TIME ZONE,
    cancellation_reason VARCHAR(500),
    
    CONSTRAINT fk_bonus_claim_bonus FOREIGN KEY (bonus_id) REFERENCES bonuses(id),
    CONSTRAINT chk_bonus_claim_status CHECK (status IN ('ACTIVE', 'COMPLETED', 'CANCELLED', 'EXPIRED'))
);

-- Index for faster queries
CREATE INDEX idx_bonus_claims_user_id ON bonus_claims(user_id);
CREATE INDEX idx_bonus_claims_bonus_id ON bonus_claims(bonus_id);
CREATE INDEX idx_bonus_claims_status ON bonus_claims(status);
CREATE INDEX idx_bonus_claims_user_bonus_status ON bonus_claims(user_id, bonus_id, status);
CREATE INDEX idx_bonus_claims_expires_at ON bonus_claims(expires_at) WHERE status = 'ACTIVE';

-- Add unique constraint to prevent duplicate bonus claims per user
CREATE UNIQUE INDEX idx_bonus_claims_unique_user_bonus_active 
ON bonus_claims(user_id, bonus_id) 
WHERE status = 'ACTIVE';
