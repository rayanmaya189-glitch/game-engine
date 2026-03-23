package rng

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
)

// ProvablyFair manages provably fair RNG for games
type ProvablyFair struct {
	serverSeed     string
	serverSeedHash string
	clientSeed     string
	nonce          int
}

// NewProvablyFair creates a new provably fair RNG
func NewProvablyFair() (*ProvablyFair, error) {
	serverSeed, err := generateSecureSeed(32)
	if err != nil {
		return nil, err
	}

	hash := sha256.Sum256([]byte(serverSeed))

	return &ProvablyFair{
		serverSeed:     serverSeed,
		serverSeedHash: hex.EncodeToString(hash[:]),
		clientSeed:     "",
		nonce:          0,
	}, nil
}

// NewProvablyFairWithSeed creates provably fair RNG with existing seeds
func NewProvablyFairWithSeed(serverSeed, clientSeed string, nonce int) (*ProvablyFair, error) {
	if serverSeed == "" {
		return nil, errors.New("server seed is required")
	}

	hash := sha256.Sum256([]byte(serverSeed))

	return &ProvablyFair{
		serverSeed:     serverSeed,
		serverSeedHash: hex.EncodeToString(hash[:]),
		clientSeed:     clientSeed,
		nonce:          nonce,
	}, nil
}

// GenerateSecureSeed generates a cryptographically secure seed
func generateSecureSeed(length int) (string, error) {
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

// GenerateBytes generates random bytes using HMAC-SHA256
func (pf *ProvablyFair) GenerateBytes(count int) ([]byte, error) {
	if pf.serverSeed == "" || pf.clientSeed == "" {
		return nil, errors.New("seeds not set")
	}

	hmacData := fmt.Sprintf("%s%s%d", pf.serverSeed, pf.clientSeed, pf.nonce)
	hash := sha256.Sum256([]byte(hmacData))

	// Expand hash to required length
	result := make([]byte, count)
	copy(result, hash[:])

	// For more bytes, hash again with counter
	for i := 1; i*32 < count; i++ {
		hmacData := fmt.Sprintf("%s%s%d-%d", pf.serverSeed, pf.clientSeed, pf.nonce, i)
		hash := sha256.Sum256([]byte(hmacData))
		copy(result[i*32:], hash[:])
	}

	pf.nonce++
	return result[:count], nil
}

// GenerateInt generates a random int in range [0, max)
func (pf *ProvablyFair) GenerateInt(max int) (int, error) {
	if max <= 0 {
		return 0, errors.New("max must be positive")
	}

	bytes, err := pf.GenerateBytes(8)
	if err != nil {
		return 0, err
	}

	n := new(big.Int).SetBytes(bytes)
	return int(n.Mod(n, big.NewInt(int64(max))).Int64()), nil
}

// GenerateInt64 generates a random int64 in range [0, max)
func (pf *ProvablyFair) GenerateInt64(max int64) (int64, error) {
	if max <= 0 {
		return 0, errors.New("max must be positive")
	}

	bytes, err := pf.GenerateBytes(8)
	if err != nil {
		return 0, err
	}

	n := new(big.Int).SetBytes(bytes)
	return n.Mod(n, big.NewInt(max)).Int64(), nil
}

// GenerateDice generates random dice rolls (1-6)
func (pf *ProvablyFair) GenerateDice(count int) ([]int, error) {
	rolls := make([]int, count)
	for i := 0; i < count; i++ {
		val, err := pf.GenerateInt(6)
		if err != nil {
			return nil, err
		}
		rolls[i] = val + 1 // Convert 0-5 to 1-6
	}
	return rolls, nil
}

// GenerateSlotSymbols generates random slot symbol indices
func (pf *ProvablyFair) GenerateSlotSymbols(reelCount, symbolCount int) ([]int, error) {
	symbols := make([]int, reelCount)
	for i := 0; i < reelCount; i++ {
		val, err := pf.GenerateInt(symbolCount)
		if err != nil {
			return nil, err
		}
		symbols[i] = val
	}
	return symbols, nil
}

// Shuffle shuffles items using Fisher-Yates algorithm
func (pf *ProvablyFair) Shuffle(items []interface{}) error {
	n := len(items)
	for i := n - 1; i > 0; i-- {
		j, err := pf.GenerateInt(i + 1)
		if err != nil {
			return err
		}
		items[i], items[j] = items[j], items[i]
	}
	return nil
}

// ShuffleInts shuffles integer slice
func (pf *ProvablyFair) ShuffleInts(items []int) error {
	n := len(items)
	for i := n - 1; i > 0; i-- {
		j, err := pf.GenerateInt(i + 1)
		if err != nil {
			return err
		}
		items[i], items[j] = items[j], items[i]
	}
	return nil
}

// SetClientSeed sets the client seed (called from client)
func (pf *ProvablyFair) SetClientSeed(seed string) {
	pf.clientSeed = seed
	pf.nonce = 0 // Reset nonce when client provides seed
}

// GetServerSeedHash returns the hash of the server seed
func (pf *ProvablyFair) GetServerSeedHash() string {
	return pf.serverSeedHash
}

// GetClientSeed returns the client seed
func (pf *ProvablyFair) GetClientSeed() string {
	return pf.clientSeed
}

// GetNonce returns the current nonce
func (pf *ProvablyFair) GetNonce() int {
	return pf.nonce
}

// Verify verifies a previous result was generated fairly
func (pf *ProvablyFair) Verify(serverSeed, clientSeed string, nonce int, resultIndex int, result []byte) (bool, error) {
	testPF, err := NewProvablyFairWithSeed(serverSeed, clientSeed, nonce)
	if err != nil {
		return false, err
	}

	// Generate the same number of results as the original
	for i := 0; i <= resultIndex; i++ {
		generated, err := testPF.GenerateBytes(len(result))
		if err != nil {
			return false, err
		}

		if i == resultIndex {
			for j := range result {
				if generated[j] != result[j] {
					return false, nil
				}
			}
			return true, nil
		}
	}

	return false, errors.New("result index out of range")
}

// Result represents a game result for verification
type Result struct {
	ServerSeed string
	ClientSeed string
	Nonce      int
	Data       []byte
}

// VerifyResults verifies multiple game results
func VerifyResults(serverSeed, clientSeed string, results []Result) (bool, error) {
	nonce := 0

	for _, result := range results {
		testPF, err := NewProvablyFairWithSeed(serverSeed, clientSeed, nonce)
		if err != nil {
			return false, err
		}

		generated, err := testPF.GenerateBytes(len(result.Data))
		if err != nil {
			return false, err
		}

		for j := range result.Data {
			if generated[j] != result.Data[j] {
				return false, nil
			}
		}

		nonce++
	}

	return true, nil
}

// Reseed generates a new server seed (called periodically by server)
func (pf *ProvablyFair) Reseed() error {
	newSeed, err := generateSecureSeed(32)
	if err != nil {
		return err
	}

	pf.serverSeed = newSeed
	hash := sha256.Sum256([]byte(newSeed))
	pf.serverSeedHash = hex.EncodeToString(hash[:])
	pf.nonce = 0

	return nil
}

// GetServerSeed returns the current server seed (for logging/debugging)
func (pf *ProvablyFair) GetServerSeed() string {
	return pf.serverSeed
}
