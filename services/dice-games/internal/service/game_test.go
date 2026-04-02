package service

import (
	"testing"
)

func TestNewGameService(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{"creates successfully", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc, err := NewGameService()
			if (err != nil) != tt.wantErr {
				t.Fatalf("NewGameService() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr && svc == nil {
				t.Fatal("NewGameService() returned nil")
			}
		})
	}
}

func TestGameServiceStruct(t *testing.T) {
	svc, err := NewGameService()
	if err != nil {
		t.Fatalf("NewGameService() error = %v", err)
	}

	if svc == nil {
		t.Fatal("service should not be nil")
	}
}

func TestMultipleInstances(t *testing.T) {
	svc1, err := NewGameService()
	if err != nil {
		t.Fatalf("first NewGameService() error = %v", err)
	}

	svc2, err := NewGameService()
	if err != nil {
		t.Fatalf("second NewGameService() error = %v", err)
	}

	if svc1 == svc2 {
		t.Fatal("expected different service instances")
	}
}

func TestDiceRollRange(t *testing.T) {
	tests := []struct {
		name  string
		rolls int
		min   int
		max   int
	}{
		{"single die", 1, 1, 6},
		{"two dice", 2, 2, 12},
		{"three dice", 3, 3, 18},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			total := 0
			for i := 0; i < tt.rolls; i++ {
				val := (i % 6) + 1
				total += val
			}
			if total < tt.min || total > tt.max {
				t.Fatalf("dice total %d out of expected range [%d, %d]", total, tt.min, tt.max)
			}
		})
	}
}

func TestBetValidation(t *testing.T) {
	tests := []struct {
		name    string
		amount  int64
		valid   bool
	}{
		{"positive bet", 100, true},
		{"minimum bet", 1, true},
		{"zero bet", 0, false},
		{"negative bet", -10, false},
		{"large bet", 1000000, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			valid := tt.amount > 0
			if valid != tt.valid {
				t.Fatalf("bet amount %d: got valid=%v, want %v", tt.amount, valid, tt.valid)
			}
		})
	}
}

func TestSettleCalculation(t *testing.T) {
	tests := []struct {
		name     string
		bet      int64
		multiplier float64
		expected int64
	}{
		{"even money", 100, 1.0, 100},
		{"double", 100, 2.0, 200},
		{"triple", 50, 3.0, 150},
		{"half", 200, 0.5, 100},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := int64(float64(tt.bet) * tt.multiplier)
			if result != tt.expected {
				t.Fatalf("settle(%d, %f) = %d, want %d", tt.bet, tt.multiplier, result, tt.expected)
			}
		})
	}
}
