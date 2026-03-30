package games

import (
	"crypto/sha256"
	"fmt"
)

// Seed hash for provably fair
func megawaysSeedHash(seed string) string {
	if seed == "" {
		return ""
	}
	h := sha256.Sum256([]byte(seed))
	return fmt.Sprintf("%x", h[:8])
}
