package rng

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/big"
	"strconv"

	"github.com/game_engine/game-engine/internal/config"
)

type RNGService struct {
	pf *ProvablyFair
}

func NewRNGService(cfg config.RNGConfig) (*RNGService, error) {
	pf, err := NewProvablyFair()
	if err != nil {
		return nil, err
	}
	return &RNGService{pf: pf}, nil
}

func (s *RNGService) GenerateBytes(count int) ([]byte, error) {
	return s.pf.GenerateBytes(count)
}

func (s *RNGService) GenerateReelResult(gameConfig interface{}) (interface{}, error) {
	symbols, err := s.pf.GenerateSlotSymbols(5, 10)
	if err != nil {
		return nil, err
	}
	return symbols, nil
}

type VerifyResultStruct struct {
	Valid       bool
	ResultHash  string
	ResultValue int64
}

func (s *RNGService) VerifyProvablyFair(serverSeed, clientSeed, nonceStr string) (*VerifyResultStruct, error) {
	nonce, err := strconv.Atoi(nonceStr)
	if err != nil {
		return nil, fmt.Errorf("invalid nonce: %w", err)
	}

	hmacData := fmt.Sprintf("%s%s%d", serverSeed, clientSeed, nonce)
	hash := sha256.Sum256([]byte(hmacData))
	hashStr := hex.EncodeToString(hash[:])

	n := new(big.Int).SetBytes(hash[:])
	val := n.Mod(n, big.NewInt(1000)).Int64()

	return &VerifyResultStruct{
		Valid:       true,
		ResultHash:  hashStr,
		ResultValue: val,
	}, nil
}
