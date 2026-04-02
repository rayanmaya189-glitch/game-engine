package service

import (
	"testing"

	"github.com/game_engine/rng-service/internal/config"
	"github.com/game_engine/rng-service/internal/rng"
)

func newTestRNGConfig() *config.Config {
	return &config.Config{
		RNG: config.RNGConfig{
			HardwareRNG: config.HardwareRNGConfig{
				Enabled:  false,
				AWSNitro: false,
			},
			SeedGeneration: config.SeedGenerationConfig{
				ServerSeedLength: 32,
				ClientSeedLength: 16,
				NonceIncrement:   true,
			},
			Limits: config.LimitsConfig{
				MaxInt:       1000000,
				MaxDeckSize:  8,
				MaxDiceCount: 10,
				MaxSlotReels: 10,
			},
		},
	}
}

func TestNewRNGService(t *testing.T) {
	cfg := newTestRNGConfig()
	svc, err := NewRNGService(cfg)
	if err != nil {
		t.Fatalf("NewRNGService() error = %v", err)
	}
	if svc == nil {
		t.Fatal("NewRNGService() returned nil")
	}
}

func TestGenerateInt(t *testing.T) {
	cfg := newTestRNGConfig()
	svc, err := NewRNGService(cfg)
	if err != nil {
		t.Fatalf("NewRNGService() error = %v", err)
	}

	tests := []struct {
		name    string
		max     int
		wantErr bool
	}{
		{"small range", 10, false},
		{"medium range", 100, false},
		{"large range", 1000, false},
		{"zero max", 0, true},
		{"negative max", -5, true},
		{"over limit", 2000000, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			val, err := svc.GenerateInt(tt.max)
			if (err != nil) != tt.wantErr {
				t.Fatalf("GenerateInt(%d) error = %v, wantErr %v", tt.max, err, tt.wantErr)
			}
			if !tt.wantErr && (val < 0 || val >= tt.max) {
				t.Fatalf("GenerateInt(%d) = %d, want [0, %d)", tt.max, val, tt.max)
			}
		})
	}
}

func TestGenerateInt64(t *testing.T) {
	cfg := newTestRNGConfig()
	svc, _ := NewRNGService(cfg)

	val, err := svc.GenerateInt64(1000)
	if err != nil {
		t.Fatalf("GenerateInt64() error = %v", err)
	}
	if val < 0 || val >= 1000 {
		t.Fatalf("GenerateInt64() = %d, want [0, 1000)", val)
	}
}

func TestGenerateFloat(t *testing.T) {
	cfg := newTestRNGConfig()
	svc, _ := NewRNGService(cfg)

	for i := 0; i < 100; i++ {
		val, err := svc.GenerateFloat()
		if err != nil {
			t.Fatalf("GenerateFloat() error = %v", err)
		}
		if val < 0.0 || val >= 1.0 {
			t.Fatalf("GenerateFloat() = %f, want [0, 1)", val)
		}
	}
}

func TestGenerateFloatRange(t *testing.T) {
	cfg := newTestRNGConfig()
	svc, _ := NewRNGService(cfg)

	tests := []struct {
		name    string
		min     float64
		max     float64
		wantErr bool
	}{
		{"valid range", 0.0, 100.0, false},
		{"decimal range", 1.5, 2.5, false},
		{"min equals max", 5.0, 5.0, true},
		{"min exceeds max", 10.0, 5.0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			val, err := svc.GenerateFloatRange(tt.min, tt.max)
			if (err != nil) != tt.wantErr {
				t.Fatalf("GenerateFloatRange(%f, %f) error = %v, wantErr %v", tt.min, tt.max, err, tt.wantErr)
			}
			if !tt.wantErr && (val < tt.min || val >= tt.max) {
				t.Fatalf("GenerateFloatRange(%f, %f) = %f, out of range", tt.min, tt.max, val)
			}
		})
	}
}

func TestGenerateDiceRolls(t *testing.T) {
	cfg := newTestRNGConfig()
	svc, _ := NewRNGService(cfg)

	tests := []struct {
		name    string
		count   int
		wantErr bool
	}{
		{"single die", 1, false},
		{"two dice", 2, false},
		{"max dice", 10, false},
		{"zero dice", 0, true},
		{"too many dice", 11, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rolls, err := svc.GenerateDiceRolls(tt.count)
			if (err != nil) != tt.wantErr {
				t.Fatalf("GenerateDiceRolls(%d) error = %v, wantErr %v", tt.count, err, tt.wantErr)
			}
			if !tt.wantErr {
				if len(rolls) != tt.count {
					t.Fatalf("got %d rolls, want %d", len(rolls), tt.count)
				}
				for i, r := range rolls {
					if r < 1 || r > 6 {
						t.Fatalf("roll[%d] = %d, want [1,6]", i, r)
					}
				}
			}
		})
	}
}

func TestGenerateCardDeck(t *testing.T) {
	cfg := newTestRNGConfig()
	svc, _ := NewRNGService(cfg)

	deck, err := svc.GenerateCardDeck(1)
	if err != nil {
		t.Fatalf("GenerateCardDeck() error = %v", err)
	}
	if len(deck) == 0 {
		t.Fatal("deck should not be empty")
	}
}

func TestGenerateSlotReels(t *testing.T) {
	cfg := newTestRNGConfig()
	svc, _ := NewRNGService(cfg)

	reels, err := svc.GenerateSlotReels(5, 20)
	if err != nil {
		t.Fatalf("GenerateSlotReels() error = %v", err)
	}
	if len(reels) != 5 {
		t.Fatalf("got %d reels, want 5", len(reels))
	}
	for i, r := range reels {
		if r < 0 || r >= 20 {
			t.Fatalf("reel[%d] = %d, want [0, 20)", i, r)
		}
	}
}

func TestSeedManagement(t *testing.T) {
	sm := rng.NewSeedManager(32, 16, true)

	pair, err := sm.GenerateSeedPair("game-001")
	if err != nil {
		t.Fatalf("GenerateSeedPair() error = %v", err)
	}
	if pair.ServerSeed == "" {
		t.Fatal("server seed should not be empty")
	}
	if pair.GameID != "game-001" {
		t.Fatalf("GameID = %s, want game-001", pair.GameID)
	}

	got, err := sm.GetSeed("game-001")
	if err != nil {
		t.Fatalf("GetSeed() error = %v", err)
	}
	if got.Hash != pair.Hash {
		t.Fatal("hash mismatch")
	}
}

func TestComputeAndVerifyResult(t *testing.T) {
	sm := rng.NewSeedManager(32, 16, true)

	_, err := sm.GenerateSeedPair("game-compute")
	if err != nil {
		t.Fatalf("GenerateSeedPair() error = %v", err)
	}

	result, err := sm.ComputeResult("game-compute", 100)
	if err != nil {
		t.Fatalf("ComputeResult() error = %v", err)
	}
	if result < 0 || result >= 100 {
		t.Fatalf("ComputeResult() = %d, want [0, 100)", result)
	}
}

func TestHealthCheck(t *testing.T) {
	cfg := newTestRNGConfig()
	svc, _ := NewRNGService(cfg)

	err := svc.HealthCheck()
	if err != nil {
		t.Fatalf("HealthCheck() error = %v", err)
	}
}
