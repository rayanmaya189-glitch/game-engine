package rng

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"sync"
	"time"
)

// SeedManager manages seeds for provably fair gaming
type SeedManager struct {
	mu            sync.RWMutex
	activeSeeds   map[string]*SeedPair
	serverSeedLen int
	clientSeedLen int
	nonceEnabled  bool
}

// SeedPair represents a server/client seed pair for provably fair games
type SeedPair struct {
	ServerSeed string    `json:"server_seed"`
	ClientSeed string    `json:"client_seed"`
	Nonce      uint64    `json:"nonce"`
	CreatedAt  time.Time `json:"created_at"`
	GameID     string    `json:"game_id"`
	IsVerified bool      `json:"is_verified"`
	Hash       string    `json:"hash"`
}

// SeedCertificate is a certificate for a seed pair
type SeedCertificate struct {
	SeedHash   string    `json:"seed_hash"`
	GameID     string    `json:"game_id"`
	Signature  string    `json:"signature"`
	Timestamp  time.Time `json:"timestamp"`
	ValidUntil time.Time `json:"valid_until"`
	RNGState   string    `json:"rng_state"`
}

// NewSeedManager creates a new SeedManager
func NewSeedManager(serverSeedLen, clientSeedLen int, nonceEnabled bool) *SeedManager {
	return &SeedManager{
		activeSeeds:   make(map[string]*SeedPair),
		serverSeedLen: serverSeedLen,
		clientSeedLen: clientSeedLen,
		nonceEnabled:  nonceEnabled,
	}
}

// GenerateSeedPair generates a new server/client seed pair for a game
func (sm *SeedManager) GenerateSeedPair(gameID string) (*SeedPair, error) {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	serverSeed, err := sm.generateRandomBytes(sm.serverSeedLen)
	if err != nil {
		return nil, fmt.Errorf("failed to generate server seed: %w", err)
	}

	clientSeed, err := sm.generateRandomBytes(sm.clientSeedLen)
	if err != nil {
		return nil, fmt.Errorf("failed to generate client seed: %w", err)
	}

	pair := &SeedPair{
		ServerSeed: hex.EncodeToString(serverSeed),
		ClientSeed: hex.EncodeToString(clientSeed),
		Nonce:      0,
		CreatedAt:  time.Now().UTC(),
		GameID:     gameID,
		Hash:       sm.computeHash(serverSeed, clientSeed),
	}

	sm.activeSeeds[pair.Hash] = pair

	return pair, nil
}

// SetClientSeed sets the client seed for a game (provided by client)
func (sm *SeedManager) SetClientSeed(gameID, serverSeedHash, clientSeed string) (*SeedPair, error) {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	if len(clientSeed) != sm.clientSeedLen*2 {
		return nil, errors.New("invalid client seed length")
	}

	pair, exists := sm.activeSeeds[serverSeedHash]
	if !exists {
		return nil, errors.New("server seed not found")
	}

	pair.ClientSeed = clientSeed
	pair.Hash = sm.computeHash([]byte(pair.ServerSeed), []byte(clientSeed))

	return pair, nil
}

// GetSeed gets the seed pair for a game
func (sm *SeedManager) GetSeed(gameID string) (*SeedPair, error) {
	sm.mu.RLock()
	defer sm.mu.RUnlock()

	for _, pair := range sm.activeSeeds {
		if pair.GameID == gameID && !pair.IsVerified {
			return pair, nil
		}
	}

	return nil, errors.New("seed pair not found")
}

// GetSeedByHash gets the seed pair by hash
func (sm *SeedManager) GetSeedByHash(hash string) (*SeedPair, error) {
	sm.mu.RLock()
	defer sm.mu.RUnlock()

	pair, exists := sm.activeSeeds[hash]
	if !exists {
		return nil, errors.New("seed pair not found")
	}

	return pair, nil
}

// IncrementNonce increments the nonce for a game
func (sm *SeedManager) IncrementNonce(gameID string) (uint64, error) {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	pair, err := sm.getSeedPairLocked(gameID)
	if err != nil {
		return 0, err
	}

	pair.Nonce++
	return pair.Nonce, nil
}

// GetNonce gets the current nonce for a game
func (sm *SeedManager) GetNonce(gameID string) (uint64, error) {
	sm.mu.RLock()
	defer sm.mu.RUnlock()

	pair, err := sm.getSeedPairRLocked(gameID)
	if err != nil {
		return 0, err
	}

	return pair.Nonce, nil
}

// ComputeResult computes a result from the seed pair
func (sm *SeedManager) ComputeResult(gameID string, max int) (int, error) {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	pair, err := sm.getSeedPairLocked(gameID)
	if err != nil {
		return 0, err
	}

	combined := sm.combineSeeds(pair.ServerSeed, pair.ClientSeed, pair.Nonce)

	rng, err := NewCryptoRNG(false, false, 32)
	if err != nil {
		return 0, err
	}

	if err := rng.Seed(combined); err != nil {
		return 0, err
	}

	result, err := rng.Intn(max)
	if err != nil {
		return 0, err
	}

	pair.Nonce++

	return result, nil
}

// generateRandomBytes generates random bytes
func (sm *SeedManager) generateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// computeHash computes SHA256 hash
func (sm *SeedManager) computeHash(serverSeed, clientSeed []byte) string {
	h := sha256.Sum256(append(serverSeed, clientSeed...))
	return hex.EncodeToString(h[:])
}

// MarshalJSON implements custom JSON marshaling for SeedPair
func (sp *SeedPair) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		ClientSeed string    `json:"client_seed"`
		Nonce      uint64    `json:"nonce"`
		CreatedAt  time.Time `json:"created_at"`
		GameID     string    `json:"game_id"`
		IsVerified bool      `json:"is_verified"`
	}{
		ClientSeed: sp.ClientSeed,
		Nonce:      sp.Nonce,
		CreatedAt:  sp.CreatedAt,
		GameID:     sp.GameID,
		IsVerified: sp.IsVerified,
	})
}
