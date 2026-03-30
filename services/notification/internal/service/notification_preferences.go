package service

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/redis/go-redis/v9"
)

// Preferences represents user notification preferences
type Preferences struct {
	UserID          string            `json:"user_id"`
	PushEnabled     bool              `json:"push_enabled"`
	EmailEnabled    bool              `json:"email_enabled"`
	SMSEnabled      bool              `json:"sms_enabled"`
	InAppEnabled    bool              `json:"in_app_enabled"`
	QuietHours      map[string]string `json:"quiet_hours"`
	MutedCategories []string          `json:"muted_categories"`
}

// GetPreferences gets user notification preferences
func (s *NotificationService) GetPreferences(ctx context.Context, userID string) (*Preferences, error) {
	key := fmt.Sprintf("preferences:%s", userID)
	data, err := s.redisClient.Get(ctx, key).Bytes()
	if err == redis.Nil {
		return &Preferences{
			UserID:          userID,
			PushEnabled:     true,
			EmailEnabled:    true,
			SMSEnabled:      false,
			InAppEnabled:    true,
			QuietHours:      make(map[string]string),
			MutedCategories: []string{},
		}, nil
	}

	if err != nil {
		return nil, err
	}

	var prefs Preferences
	if err := json.Unmarshal(data, &prefs); err != nil {
		return nil, err
	}

	return &prefs, nil
}

// SavePreferences saves user notification preferences
func (s *NotificationService) SavePreferences(ctx context.Context, prefs *Preferences) error {
	data, err := json.Marshal(prefs)
	if err != nil {
		return err
	}

	key := fmt.Sprintf("preferences:%s", prefs.UserID)
	return s.redisClient.Set(ctx, key, data, 0).Err()
}
