package handler

import (
	"context"
	"fmt"

	"github.com/game-engine/rng-service/internal/rng"
	"github.com/game-engine/rng-service/internal/service"
)

// RNGHandler handles RNG-related requests
type RNGHandler struct {
	rngService *service.RNGService
}

// NewRNGHandler creates a new RNG handler
func NewRNGHandler(rngService *service.RNGService) *RNGHandler {
	return &RNGHandler{
		rngService: rngService,
	}
}

// GenerateRandomInt generates a random integer
func (h *RNGHandler) GenerateRandomInt(ctx context.Context, max int) (int, error) {
	return h.rngService.GenerateInt(max)
}

// GenerateRandomInt64 generates a random int64
func (h *RNGHandler) GenerateRandomInt64(ctx context.Context, max int64) (int64, error) {
	return h.rngService.GenerateInt64(max)
}

// GenerateRandomFloat generates a random float64
func (h *RNGHandler) GenerateRandomFloat(ctx context.Context) (float64, error) {
	return h.rngService.GenerateFloat()
}

// GenerateRandomFloatRange generates a random float64 in range
func (h *RNGHandler) GenerateRandomFloatRange(ctx context.Context, min, max float64) (float64, error) {
	return h.rngService.GenerateFloatRange(min, max)
}

// GenerateCardDeck generates a shuffled card deck
func (h *RNGHandler) GenerateCardDeck(ctx context.Context, deckCount int) ([]int, error) {
	return h.rngService.GenerateCardDeck(deckCount)
}

// GenerateDiceRolls generates random dice rolls
func (h *RNGHandler) GenerateDiceRolls(ctx context.Context, diceCount int) ([]int, error) {
	return h.rngService.GenerateDiceRolls(diceCount)
}

// GenerateSlotReels generates slot reel positions
func (h *RNGHandler) GenerateSlotReels(ctx context.Context, reelCount, symbolCount int) ([]int, error) {
	return h.rngService.GenerateSlotReels(reelCount, symbolCount)
}

// GenerateSeed generates a new seed pair
func (h *RNGHandler) GenerateSeed(ctx context.Context, gameID string) (string, string, error) {
	seedPair, err := h.rngService.GenerateSeed(gameID)
	if err != nil {
		return "", "", err
	}
	return seedPair.ServerSeed, seedPair.Hash, nil
}

// SetClientSeed sets the client seed for a game
func (h *RNGHandler) SetClientSeed(ctx context.Context, gameID, serverSeedHash, clientSeed string) error {
	_, err := h.rngService.SetClientSeed(gameID, serverSeedHash, clientSeed)
	return err
}

// ComputeGameResult computes a game result using the seed pair
func (h *RNGHandler) ComputeGameResult(ctx context.Context, gameID string, max int) (int, error) {
	return h.rngService.ComputeResult(gameID, max)
}

// VerifyResult verifies a game result
func (h *RNGHandler) VerifyResult(ctx context.Context, gameID string, result, max int) (bool, error) {
	return h.rngService.VerifyResult(gameID, result, max)
}

// GenerateCertificate generates a certificate for a game
func (h *RNGHandler) GenerateCertificate(ctx context.Context, gameID string) (*CertificateResponse, error) {
	cert, err := h.rngService.GenerateCertificate(gameID)
	if err != nil {
		return nil, err
	}

	return &CertificateResponse{
		SeedHash:   cert.SeedHash,
		Signature:  cert.Signature,
		Timestamp:  cert.Timestamp.Unix(),
		ValidUntil: cert.ValidUntil.Unix(),
		RNGState:   cert.RNGState,
	}, nil
}

// VerifyCertificate verifies a certificate
func (h *RNGHandler) VerifyCertificate(ctx context.Context, cert *CertificateRequest) bool {
	return h.rngService.VerifyCertificate(&rng.SeedCertificate{
		SeedHash:  cert.SeedHash,
		Signature: cert.Signature,
	})
}

// HealthCheck checks if the service is healthy
func (h *RNGHandler) HealthCheck(ctx context.Context) error {
	return h.rngService.HealthCheck()
}

// Request/Response types (would be generated from protobuf in production)

type GenerateIntRequest struct {
	Max int `json:"max"`
}

type GenerateIntResponse struct {
	Value int `json:"value"`
}

type GenerateFloatResponse struct {
	Value float64 `json:"value"`
}

type GenerateFloatRangeRequest struct {
	Min float64 `json:"min"`
	Max float64 `json:"max"`
}

type GenerateCardDeckRequest struct {
	DeckCount int `json:"deck_count"`
}

type GenerateCardDeckResponse struct {
	Cards []int `json:"cards"`
}

type GenerateDiceRollsRequest struct {
	DiceCount int `json:"dice_count"`
}

type GenerateDiceRollsResponse struct {
	Rolls []int `json:"rolls"`
}

type GenerateSlotReelsRequest struct {
	ReelCount   int `json:"reel_count"`
	SymbolCount int `json:"symbol_count"`
}

type GenerateSlotReelsResponse struct {
	Positions []int `json:"positions"`
}

type GenerateSeedRequest struct {
	GameID string `json:"game_id"`
}

type GenerateSeedResponse struct {
	ServerSeed string `json:"server_seed"`
	SeedHash   string `json:"seed_hash"`
}

type SetClientSeedRequest struct {
	GameID         string `json:"game_id"`
	ServerSeedHash string `json:"server_seed_hash"`
	ClientSeed     string `json:"client_seed"`
}

type ComputeResultRequest struct {
	GameID string `json:"game_id"`
	Max    int    `json:"max"`
}

type ComputeResultResponse struct {
	Result int `json:"result"`
}

type VerifyResultRequest struct {
	GameID string `json:"game_id"`
	Result int    `json:"result"`
	Max    int    `json:"max"`
}

type VerifyResultResponse struct {
	Valid bool `json:"valid"`
}

type CertificateRequest struct {
	SeedHash  string `json:"seed_hash"`
	Signature string `json:"signature"`
}

type CertificateResponse struct {
	SeedHash   string `json:"seed_hash"`
	Signature  string `json:"signature"`
	Timestamp  int64  `json:"timestamp"`
	ValidUntil int64  `json:"valid_until"`
	RNGState   string `json:"rng_state"`
}

// Helper to convert int slice to interface slice for protobuf
func intsToInterfaces(ints []int) []interface{} {
	result := make([]interface{}, len(ints))
	for i, v := range ints {
		result[i] = v
	}
	return result
}

// Format the response as a string for gRPC
func (h *RNGHandler) formatCardsResponse(cards []int) string {
	return fmt.Sprintf("%v", cards)
}
