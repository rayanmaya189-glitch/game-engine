package rng

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"time"
)

// VerifyResult verifies that a result was generated fairly
func (sm *SeedManager) VerifyResult(gameID string, result, max int) (bool, error) {
	sm.mu.RLock()
	defer sm.mu.RUnlock()

	pair, err := sm.getSeedPairRLocked(gameID)
	if err != nil {
		return false, err
	}

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

	if time.Now().UTC().After(cert.ValidUntil) {
		return false
	}

	pair, exists := sm.activeSeeds[cert.SeedHash]
	if !exists {
		return false
	}

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
