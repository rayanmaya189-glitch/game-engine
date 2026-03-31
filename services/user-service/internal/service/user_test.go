package service

import (
	"context"
	"testing"

	"github.com/game_engine/user-service/internal/config"
	"github.com/game_engine/user-service/internal/model"
)

func testUserConfig() *config.Config {
	return &config.Config{
		Cache: config.CacheConfig{ProfileTTL: 300},
		RateLimiting: config.RateLimitingConfig{ProfileUpdateMax: 10},
	}
}

func TestGetProfile_NotFound(t *testing.T) {
	svc := &UserService{cfg: testUserConfig()}
	_, err := svc.GetProfile(context.Background(), "nonexistent")
	if err == nil {
		t.Fatal("expected error for missing profile")
	}
	if err.Error() != "profile not found" {
		t.Errorf("expected 'profile not found', got '%s'", err.Error())
	}
}

func TestUpdateProfile_NotFound(t *testing.T) {
	svc := &UserService{cfg: testUserConfig()}
	_, err := svc.UpdateProfile(context.Background(), "missing", &model.Profile{DisplayName: "New"})
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestGetPlayerSettings_Defaults(t *testing.T) {
	svc := &UserService{cfg: testUserConfig()}
	settings, err := svc.GetPlayerSettings(context.Background(), "user-1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if settings == nil {
		t.Fatal("expected default settings")
	}
	if !settings.EmailNotifications {
		t.Error("expected email notifications enabled")
	}
	if settings.SoundVolume != 50 {
		t.Errorf("expected volume 50, got %d", settings.SoundVolume)
	}
	if settings.Theme != "default" {
		t.Errorf("expected theme 'default', got '%s'", settings.Theme)
	}
	if settings.PushNotifications != true {
		t.Error("expected push notifications enabled")
	}
	if settings.ProfilePublic != false {
		t.Error("expected profile private by default")
	}
}

func TestUpdatePlayerSettings_Validation(t *testing.T) {
	tests := []struct {
		name    string
		volume  int
		wantErr bool
	}{
		{"valid 0", 0, false},
		{"valid 50", 50, false},
		{"valid 100", 100, false},
		{"invalid -1", -1, true},
		{"invalid 101", 101, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc := &UserService{cfg: testUserConfig()}
			settings := &model.PlayerSettings{UserID: "u1", SoundVolume: tt.volume}
			_, err := svc.UpdatePlayerSettings(context.Background(), "u1", settings)
			if (err != nil) != tt.wantErr {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetProfile_NilRedis(t *testing.T) {
	svc := &UserService{cfg: testUserConfig()}
	_, err := svc.GetProfile(context.Background(), "u1")
	if err == nil {
		t.Fatal("expected error with nil redis")
	}
}

func TestUpdateProfile_NilRedis(t *testing.T) {
	svc := &UserService{cfg: testUserConfig()}
	_, err := svc.UpdateProfile(context.Background(), "u1", &model.Profile{DisplayName: "X"})
	if err == nil {
		t.Fatal("expected error with nil redis")
	}
}
