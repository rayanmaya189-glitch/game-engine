package service

import (
	"context"
	"encoding/json"
	"fmt"
)

// Template represents a notification template
type Template struct {
	ID        string           `json:"id"`
	Name      string           `json:"name"`
	Type      NotificationType `json:"type"`
	Title     string           `json:"title"`
	Body      string           `json:"body"`
	Variables []string         `json:"variables"`
}

// GetTemplate gets a notification template
func (s *NotificationService) GetTemplate(ctx context.Context, templateID string) (*Template, error) {
	key := fmt.Sprintf("template:%s", templateID)
	data, err := s.redisClient.Get(ctx, key).Bytes()
	if err != nil {
		return nil, err
	}

	var template Template
	if err := json.Unmarshal(data, &template); err != nil {
		return nil, err
	}

	return &template, nil
}

// SaveTemplate saves a notification template
func (s *NotificationService) SaveTemplate(ctx context.Context, template *Template) error {
	data, err := json.Marshal(template)
	if err != nil {
		return err
	}

	key := fmt.Sprintf("template:%s", template.ID)
	return s.redisClient.Set(ctx, key, data, 0).Err()
}
