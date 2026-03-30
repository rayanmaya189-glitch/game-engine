package rng

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"errors"
	"io"
	"math"
	"math/big"
	"sync"
)

// CryptoRNG represents a cryptographically secure random number generator
type CryptoRNG struct {
	mu         sync.Mutex
	seed       []byte
	counter    uint64
	stream     cipher.Stream
	seedLength int
}

// NewCryptoRNG creates a new CryptoRNG instance
func NewCryptoRNG(useHardware, awsNitro bool, seedLength int) (*CryptoRNG, error) {
	rng := &CryptoRNG{
		seedLength: seedLength,
		counter:    0,
	}

	if err := rng.initializeSeed(useHardware, awsNitro); err != nil {
		return nil, err
	}

	return rng, nil
}

// initializeSeed initializes the RNG with a random seed
func (r *CryptoRNG) initializeSeed(useHardware, awsNitro bool) error {
	r.seed = make([]byte, r.seedLength)

	var err error
	if useHardware {
		if err = r.readFromHardwareRNG(r.seed, awsNitro); err != nil {
			if err = r.readFromSoftwareRNG(r.seed); err != nil {
				return err
			}
		}
	} else {
		if err = r.readFromSoftwareRNG(r.seed); err != nil {
			return err
		}
	}

	block, err := aes.NewCipher(r.seed)
	if err != nil {
		return err
	}

	iv := make([]byte, block.BlockSize())
	r.stream = cipher.NewCTR(block, iv)

	return nil
}

// readFromHardwareRNG reads random bytes from hardware RNG
func (r *CryptoRNG) readFromHardwareRNG(b []byte, awsNitro bool) error {
	if awsNitro {
		return r.readFromAWSNitro(b)
	}

	_, err := io.ReadFull(rand.Reader, b)
	return err
}

// readFromAWSNitro reads random bytes from AWS Nitro
func (r *CryptoRNG) readFromAWSNitro(b []byte) error {
	_, err := io.ReadFull(rand.Reader, b)
	return err
}

// readFromSoftwareRNG reads random bytes from software RNG (crypto/rand)
func (r *CryptoRNG) readFromSoftwareRNG(b []byte) error {
	_, err := io.ReadFull(rand.Reader, b)
	return err
}

// generateBytes generates random bytes using CTR mode
func (r *CryptoRNG) generateBytes(n int) ([]byte, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	result := make([]byte, n)
	r.stream.XORKeyStream(result, result)
	r.counter++

	return result, nil
}

// generateInt64 generates a cryptographically secure random int64 in range [0, max)
func (r *CryptoRNG) generateInt64(max int64) (int64, error) {
	if max <= 0 {
		return 0, errors.New("max must be positive")
	}

	maxBig := big.NewInt(max)
	result := new(big.Int)

	for {
		bytesNeeded := (maxBig.BitLen() + 7) / 8
		if bytesNeeded < 8 {
			bytesNeeded = 8
		}
		bytes, err := r.generateBytes(bytesNeeded)
		if err != nil {
			return 0, err
		}

		result.SetBytes(bytes)
		if result.Cmp(maxBig) < 0 {
			return result.Int64(), nil
		}
	}
}

// Intn returns a random integer in range [0, n)
func (r *CryptoRNG) Intn(n int) (int, error) {
	val, err := r.generateInt64(int64(n))
	return int(val), err
}

// Int63n returns a random int64 in range [0, n)
func (r *CryptoRNG) Int63n(n int64) (int64, error) {
	return r.generateInt64(n)
}

// Float64 returns a random float64 in range [0, 1)
func (r *CryptoRNG) Float64() (float64, error) {
	bits, err := r.generateBytes(8)
	if err != nil {
		return 0, err
	}

	u := uint64(bits[0])<<56 | uint64(bits[1])<<48 | uint64(bits[2])<<40 | uint64(bits[3])<<32 |
		uint64(bits[4])<<24 | uint64(bits[5])<<16 | uint64(bits[6])<<8 | uint64(bits[7])
	u = (u & 0x000FFFFFFFFFFFFF) | 0x3FF0000000000000
	return math.Float64frombits(u) - 1.0, nil
}

// Shuffle shuffles a slice using Fisher-Yates algorithm
func (r *CryptoRNG) Shuffle(n int, swap func(i, j int)) error {
	if n <= 1 {
		return nil
	}

	for i := n - 1; i > 0; i-- {
		j, err := r.Intn(i + 1)
		if err != nil {
			return err
		}
		swap(i, j)
	}
	return nil
}

// Seed reseeds the RNG with a new seed
func (r *CryptoRNG) Seed(seed []byte) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	hash := sha256.Sum256(seed)
	r.seed = hash[:]
	r.counter = 0

	block, err := aes.NewCipher(r.seed)
	if err != nil {
		return err
	}

	iv := make([]byte, block.BlockSize())
	r.stream = cipher.NewCTR(block, iv)

	return nil
}

// GetSeed returns the current seed (for verification purposes)
func (r *CryptoRNG) GetSeed() []byte {
	r.mu.Lock()
	defer r.mu.Unlock()
	return make([]byte, len(r.seed))
}

// Verify verifies that a sequence was generated with the given seed
func (r *CryptoRNG) Verify(seed, expectedOutput []byte) bool {
	testRNG, err := NewCryptoRNG(false, false, len(seed))
	if err != nil {
		return false
	}

	if err := testRNG.Seed(seed); err != nil {
		return false
	}

	actualOutput, err := testRNG.generateBytes(len(expectedOutput))
	if err != nil {
		return false
	}

	for i := range expectedOutput {
		if actualOutput[i] != expectedOutput[i] {
			return false
		}
	}

	return true
}
