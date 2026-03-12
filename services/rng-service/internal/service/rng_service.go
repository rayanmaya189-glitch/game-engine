package service

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/game_engine/rng-service/internal/config"
	"github.com/game_engine/rng-service/internal/rng"
)

// RNGService provides random number generation operations
type RNGService struct {
	mu          sync.RWMutex
	rng         *rng.CryptoRNG
	seedManager *rng.SeedManager
	cfg         *config.Config
}

// NewRNGService creates a new RNG service
func NewRNGService(cfg *config.Config) (*RNGService, error) {
	rngInstance, err := rng.NewCryptoRNG(
		cfg.RNG.HardwareRNG.Enabled,
		cfg.RNG.HardwareRNG.AWSNitro,
		cfg.RNG.SeedGeneration.ServerSeedLength,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create RNG: %w", err)
	}

	seedManager := rng.NewSeedManager(
		cfg.RNG.SeedGeneration.ServerSeedLength,
		cfg.RNG.SeedGeneration.ClientSeedLength,
		cfg.RNG.SeedGeneration.NonceIncrement,
	)

	return &RNGService{
		rng:         rngInstance,
		seedManager: seedManager,
		cfg:         cfg,
	}, nil
}

// GenerateInt generates a random integer in range [0, max)
func (s *RNGService) GenerateInt(max int) (int, error) {
	if max <= 0 || max > int(s.cfg.RNG.Limits.MaxInt) {
		return 0, errors.New("invalid max value")
	}
	return s.rng.Intn(max)
}

// GenerateInt64 generates a random int64 in range [0, max)
func (s *RNGService) GenerateInt64(max int64) (int64, error) {
	if max <= 0 || max > s.cfg.RNG.Limits.MaxInt {
		return 0, errors.New("invalid max value")
	}
	return s.rng.Int63n(max)
}

// GenerateFloat generates a random float64 in range [0, 1)
func (s *RNGService) GenerateFloat() (float64, error) {
	return s.rng.Float64()
}

// GenerateFloatRange generates a random float64 in range [min, max)
func (s *RNGService) GenerateFloatRange(min, max float64) (float64, error) {
	if min >= max {
		return 0, errors.New("min must be less than max")
	}

	f, err := s.rng.Float64()
	if err != nil {
		return 0, err
	}

	return min + f*(max-min), nil
}

// GenerateCardDeck generates a shuffled card deck
func (s *RNGService) GenerateCardDeck(deckCount int) ([]int, error) {
	if deckCount <= 0 || deckCount > int(s.cfg.RNG.Limits.MaxDeckSize) {
		return nil, errors.New("invalid deck count")
	}
	return s.rng.GenerateCardDeck(deckCount)
}

// GenerateDiceRolls generates random dice rolls
func (s *RNGService) GenerateDiceRolls(diceCount int) ([]int, error) {
	if diceCount <= 0 || diceCount > int(s.cfg.RNG.Limits.MaxDiceCount) {
		return nil, errors.New("invalid dice count")
	}
	return s.rng.GenerateDiceRolls(diceCount)
}

// GenerateSlotReels generates slot reel positions
func (s *RNGService) GenerateSlotReels(reelCount, symbolCount int) ([]int, error) {
	if reelCount <= 0 || reelCount > int(s.cfg.RNG.Limits.MaxSlotReels) {
		return nil, errors.New("invalid reel count")
	}
	return s.rng.GenerateSlotReels(reelCount, symbolCount)
}

// Seed management

// GenerateSeed generates a new seed pair for a game
func (s *RNGService) GenerateSeed(gameID string) (*rng.SeedPair, error) {
	return s.seedManager.GenerateSeedPair(gameID)
}

// SetClientSeed sets the client seed for a game
func (s *RNGService) SetClientSeed(gameID, serverSeedHash, clientSeed string) (*rng.SeedPair, error) {
	return s.seedManager.SetClientSeed(gameID, serverSeedHash, clientSeed)
}

// GetSeed gets the seed pair for a game
func (s *RNGService) GetSeed(gameID string) (*rng.SeedPair, error) {
	return s.seedManager.GetSeed(gameID)
}

// ComputeResult computes a result from the seed pair
func (s *RNGService) ComputeResult(gameID string, max int) (int, error) {
	return s.seedManager.ComputeResult(gameID, max)
}

// VerifyResult verifies that a result was generated fairly
func (s *RNGService) VerifyResult(gameID string, result, max int) (bool, error) {
	return s.seedManager.VerifyResult(gameID, result, max)
}

// GenerateCertificate generates a certificate for the current RNG state
func (s *RNGService) GenerateCertificate(gameID string) (*rng.SeedCertificate, error) {
	return s.seedManager.GenerateCertificate(gameID)
}

// VerifyCertificate verifies a certificate
func (s *RNGService) VerifyCertificate(cert *rng.SeedCertificate) bool {
	return s.seedManager.VerifyCertificate(cert)
}

// Cleanup cleans up old seed pairs
func (s *RNGService) Cleanup() int {
	return s.seedManager.Cleanup(24 * 60 * 60) // 24 hours
}

// Verification

// Verify verifies that output was generated with given seed
func (s *RNGService) Verify(seed, output []byte) bool {
	return s.rng.Verify(seed, output)
}

// HealthCheck checks if the RNG is working correctly
func (s *RNGService) HealthCheck() error {
	// Generate a random number to verify RNG is working
	_, err := s.rng.Intn(100)
	return err
}

// Context-based methods for gRPC

type rngKey string

const rngServiceKey rngKey = "rng_service"

// NewContext creates a new context with the RNG service
func NewContext(ctx context.Context, service *RNGService) context.Context {
	return context.WithValue(ctx, rngServiceKey, service)
}

// FromContext gets the RNG service from context
func FromContext(ctx context.Context) (*RNGService, error) {
	service, ok := ctx.Value(rngServiceKey).(*RNGService)
	if !ok || service == nil {
		return nil, errors.New("RNG service not found in context")
	}
	return service, nil
}
