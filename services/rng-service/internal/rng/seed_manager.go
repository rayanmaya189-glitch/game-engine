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

	// Generate server seed
	serverSeed, err := sm.generateRandomBytes(sm.serverSeedLen)
	if err != nil {
		return nil, fmt.Errorf("failed to generate server seed: %w", err)
	}

	// Generate client seed
	clientSeed, err := sm.generateRandomBytes(sm.clientSeedLen)
	if err != nil {
		return nil, fmt.Errorf("failed to generate client seed: %w", err)
	}

	// Create seed pair
	pair := &SeedPair{
		ServerSeed: hex.EncodeToString(serverSeed),
		ClientSeed: hex.EncodeToString(clientSeed),
		Nonce:      0,
		CreatedAt:  time.Now().UTC(),
		GameID:     gameID,
		Hash:       sm.computeHash(serverSeed, clientSeed),
	}

	// Store the seed pair
	sm.activeSeeds[pair.Hash] = pair

	return pair, nil
}

// SetClientSeed sets the client seed for a game (provided by client)
func (sm *SeedManager) SetClientSeed(gameID, serverSeedHash, clientSeed string) (*SeedPair, error) {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	// Validate client seed format
	if len(clientSeed) != sm.clientSeedLen*2 { // hex encoded
		return nil, errors.New("invalid client seed length")
	}

	// Find existing server seed
	pair, exists := sm.activeSeeds[serverSeedHash]
	if !exists {
		return nil, errors.New("server seed not found")
	}

	// Update client seed
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

	// Combine seeds for result
	combined := sm.combineSeeds(pair.ServerSeed, pair.ClientSeed, pair.Nonce)

	// Generate result
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

	// Increment nonce
	pair.Nonce++

	return result, nil
}

// VerifyResult verifies that a result was generated fairly
func (sm *SeedManager) VerifyResult(gameID string, result, max int) (bool, error) {
	sm.mu.RLock()
	defer sm.mu.RUnlock()

	pair, err := sm.getSeedPairRLocked(gameID)
	if err != nil {
		return false, err
	}

	// Compute expected result
	combined := sm.combineSeeds(pair.ServerSeed, pair.ClientSeed, pair.Nonce)

	rng, err := NewCryptoRNG(false, false, 32)
	if err != nil {
		return false, err
	}

	if err := rng.Seed(combined); err != nil {
		return false, err
	}

	expected, err := rng.Intn(max)
	if err != nil {
		return false, err
	}

	return expected == result, nil
}

// GenerateCertificate generates a certificate for the current RNG state
func (sm *SeedManager) GenerateCertificate(gameID string) (*SeedCertificate, error) {
	sm.mu.RLock()
	defer sm.mu.RUnlock()

	pair, err := sm.getSeedPairRLocked(gameID)
	if err != nil {
		return nil, err
	}

	// In production, this would use a proper signing key
	signature := sm.computeSignature(pair)

	cert := &SeedCertificate{
		SeedHash:   pair.Hash,
		GameID:     gameID,
		Signature:  signature,
		Timestamp:  time.Now().UTC(),
		ValidUntil: time.Now().UTC().Add(24 * time.Hour),
		RNGState:   sm.computeRNGState(pair),
	}

	return cert, nil
}

// VerifyCertificate verifies a certificate
func (sm *SeedManager) VerifyCertificate(cert *SeedCertificate) bool {
	sm.mu.RLock()
	defer sm.mu.RUnlock()

	// Check if certificate is expired
	if time.Now().UTC().After(cert.ValidUntil) {
		return false
	}

	// Find matching seed pair
	pair, exists := sm.activeSeeds[cert.SeedHash]
	if !exists {
		return false
	}

	// Verify signature
	expectedSig := sm.computeSignature(pair)
	return cert.Signature == expectedSig
}

// MarkVerified marks a seed pair as verified
func (sm *SeedManager) MarkVerified(gameID string) error {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	pair, err := sm.getSeedPairLocked(gameID)
	if err != nil {
		return err
	}

	pair.IsVerified = true
	return nil
}

// Cleanup removes old/verified seed pairs
func (sm *SeedManager) Cleanup(maxAge time.Duration) int {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	now := time.Now().UTC()
	removed := 0

	for hash, pair := range sm.activeSeeds {
		age := now.Sub(pair.CreatedAt)
		if age > maxAge || pair.IsVerified {
			delete(sm.activeSeeds, hash)
			removed++
		}
	}

	return removed
}

// Helper functions

func (sm *SeedManager) generateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func (sm *SeedManager) computeHash(serverSeed, clientSeed []byte) string {
	h := sha256.Sum256(append(serverSeed, clientSeed...))
	return hex.EncodeToString(h[:])
}

func (sm *SeedManager) combineSeeds(serverSeed, clientSeed string, nonce uint64) []byte {
	data := fmt.Sprintf("%s:%s:%d", serverSeed, clientSeed, nonce)
	h := sha256.Sum256([]byte(data))
	return h[:]
}

func (sm *SeedManager) computeSignature(pair *SeedPair) string {
	data := fmt.Sprintf("%s:%s:%d:%s", pair.ServerSeed, pair.ClientSeed, pair.Nonce, pair.GameID)
	h := sha256.Sum256([]byte(data))
	return hex.EncodeToString(h[:])
}

func (sm *SeedManager) computeRNGState(pair *SeedPair) string {
	data := fmt.Sprintf("%s:%s:%d", pair.ServerSeed, pair.ClientSeed, pair.Nonce)
	h := sha256.Sum256([]byte(data))
	return hex.EncodeToString(h[:])
}

func (sm *SeedManager) getSeedPairLocked(gameID string) (*SeedPair, error) {
	for _, pair := range sm.activeSeeds {
		if pair.GameID == gameID && !pair.IsVerified {
			return pair, nil
		}
	}
	return nil, errors.New("seed pair not found")
}

func (sm *SeedManager) getSeedPairRLocked(gameID string) (*SeedPair, error) {
	for _, pair := range sm.activeSeeds {
		if pair.GameID == gameID && !pair.IsVerified {
			return pair, nil
		}
	}
	return nil, errors.New("seed pair not found")
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
