package service

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

// NotificationType represents the type of notification
type NotificationType string

const (
	NotificationTypePush  NotificationType = "push"
	NotificationTypeEmail NotificationType = "email"
	NotificationTypeSMS   NotificationType = "sms"
	NotificationTypeInApp NotificationType = "in_app"
)

// NotificationStatus represents the status of a notification
type NotificationStatus string

const (
	NotificationStatusPending   NotificationStatus = "pending"
	NotificationStatusSent      NotificationStatus = "sent"
	NotificationStatusFailed    NotificationStatus = "failed"
	NotificationStatusDelivered NotificationStatus = "delivered"
	NotificationStatusRead      NotificationStatus = "read"
)

// Notification represents a notification
type Notification struct {
	ID          string             `json:"id"`
	Type        NotificationType   `json:"type"`
	UserID      string             `json:"user_id"`
	Title       string             `json:"title"`
	Body        string             `json:"body"`
	Data        map[string]string  `json:"data"`
	Status      NotificationStatus `json:"status"`
	SentAt      time.Time          `json:"sent_at"`
	DeliveredAt time.Time          `json:"delivered_at"`
	ReadAt      time.Time          `json:"read_at"`
	CreatedAt   time.Time          `json:"created_at"`
}

// Template represents a notification template
type Template struct {
	ID        string           `json:"id"`
	Name      string           `json:"name"`
	Type      NotificationType `json:"type"`
	Title     string           `json:"title"`
	Body      string           `json:"body"`
	Variables []string         `json:"variables"`
}

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

// NotificationService provides notification business logic
type NotificationService struct {
	redisClient *redis.Client
	config      *Config
}

// NewNotificationService creates a new notification service
func NewNotificationService(config *Config, redisClient *redis.Client) (*NotificationService, error) {
	return &NotificationService{
		config:      config,
		redisClient: redisClient,
	}, nil
}

// SendNotification sends a notification
func (s *NotificationService) SendNotification(ctx context.Context, notif *Notification) error {
	notif.ID = uuid.New().String()
	notif.Status = NotificationStatusPending
	notif.CreatedAt = time.Now()

	// Save to Redis
	data, err := json.Marshal(notif)
	if err != nil {
		return err
	}

	key := fmt.Sprintf("notification:%s", notif.ID)
	if err := s.redisClient.Set(ctx, key, data, 0).Err(); err != nil {
		return err
	}

	// Send based on type
	switch notif.Type {
	case NotificationTypePush:
		return s.sendPush(ctx, notif)
	case NotificationTypeEmail:
		return s.sendEmail(ctx, notif)
	case NotificationTypeSMS:
		return s.sendSMS(ctx, notif)
	case NotificationTypeInApp:
		return s.sendInApp(ctx, notif)
	default:
		return fmt.Errorf("unknown notification type: %s", notif.Type)
	}
}

// sendPush sends a push notification
func (s *NotificationService) sendPush(ctx context.Context, notif *Notification) error {
	if !s.config.Notification.Push.Enabled {
		return fmt.Errorf("push notifications are disabled")
	}

	// Simulate sending push notification
	// In production, integrate with FCM/APNS
	notif.Status = NotificationStatusSent
	notif.SentAt = time.Now()

	return s.updateNotification(ctx, notif)
}

// sendEmail sends an email notification
func (s *NotificationService) sendEmail(ctx context.Context, notif *Notification) error {
	if !s.config.Notification.Email.Enabled {
		return fmt.Errorf("email notifications are disabled")
	}

	// Simulate sending email
	// In production, integrate with SMTP
	notif.Status = NotificationStatusSent
	notif.SentAt = time.Now()

	return s.updateNotification(ctx, notif)
}

// sendSMS sends an SMS notification
func (s *NotificationService) sendSMS(ctx context.Context, notif *Notification) error {
	if !s.config.Notification.SMS.Enabled {
		return fmt.Errorf("SMS notifications are disabled")
	}

	// Simulate sending SMS
	// In production, integrate with Twilio
	notif.Status = NotificationStatusSent
	notif.SentAt = time.Now()

	return s.updateNotification(ctx, notif)
}

// sendInApp sends an in-app notification
func (s *NotificationService) sendInApp(ctx context.Context, notif *Notification) error {
	if !s.config.Notification.InApp.Enabled {
		return fmt.Errorf("in-app notifications are disabled")
	}

	// Save to user's notification list
	key := fmt.Sprintf("notifications:%s", notif.UserID)
	data, err := json.Marshal(notif)
	if err != nil {
		return err
	}

	s.redisClient.LPush(ctx, key, data)
	s.redisClient.LTrim(ctx, key, 0, int64(s.config.Notification.InApp.MaxPerUser))
	s.redisClient.Expire(ctx, key, time.Duration(s.config.Notification.InApp.RetentionDays)*24*time.Hour)

	notif.Status = NotificationStatusSent
	notif.SentAt = time.Now()

	return s.updateNotification(ctx, notif)
}

// GetNotifications gets user notifications
func (s *NotificationService) GetNotifications(ctx context.Context, userID string, limit int) ([]Notification, error) {
	if limit <= 0 {
		limit = 20
	}

	key := fmt.Sprintf("notifications:%s", userID)
	messages, err := s.redisClient.LRange(ctx, key, 0, int64(limit-1)).Result()
	if err != nil {
		return nil, err
	}

	result := make([]Notification, 0, len(messages))
	for _, msg := range messages {
		var notif Notification
		if err := json.Unmarshal([]byte(msg), &notif); err != nil {
			continue
		}
		result = append(result, notif)
	}

	return result, nil
}

// MarkAsRead marks a notification as read
func (s *NotificationService) MarkAsRead(ctx context.Context, notificationID string) error {
	key := fmt.Sprintf("notification:%s", notificationID)
	data, err := s.redisClient.Get(ctx, key).Bytes()
	if err != nil {
		return err
	}

	var notif Notification
	if err := json.Unmarshal(data, &notif); err != nil {
		return err
	}

	notif.Status = NotificationStatusRead
	notif.ReadAt = time.Now()

	return s.updateNotification(ctx, &notif)
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

// GetPreferences gets user notification preferences
func (s *NotificationService) GetPreferences(ctx context.Context, userID string) (*Preferences, error) {
	key := fmt.Sprintf("preferences:%s", userID)
	data, err := s.redisClient.Get(ctx, key).Bytes()
	if err == redis.Nil {
		// Return default preferences
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

// SendBatch sends batch notifications
func (s *NotificationService) SendBatch(ctx context.Context, notifs []Notification) error {
	if !s.config.Notification.Batch.Enabled {
		return fmt.Errorf("batch notifications are disabled")
	}

	if len(notifs) > s.config.Notification.Batch.MaxBatchSize {
		return fmt.Errorf("batch size exceeds maximum: %d", s.config.Notification.Batch.MaxBatchSize)
	}

	for _, notif := range notifs {
		if err := s.SendNotification(ctx, &notif); err != nil {
			// Log error but continue with others
			continue
		}
	}

	return nil
}

// updateNotification updates a notification in Redis
func (s *NotificationService) updateNotification(ctx context.Context, notif *Notification) error {
	data, err := json.Marshal(notif)
	if err != nil {
		return err
	}

	key := fmt.Sprintf("notification:%s", notif.ID)
	return s.redisClient.Set(ctx, key, data, 0).Err()
}

// Config holds notification service configuration
type Config struct {
	Server struct {
		GRPCPort int `yaml:"grpc_port"`
		HTTPPort int `yaml:"http_port"`
	} `yaml:"server"`

	Redis struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
		DB   int    `yaml:"db"`
	} `yaml:"redis"`

	NATS struct {
		URLs      []string `yaml:"urls"`
		ClusterID string   `yaml:"cluster_id"`
		ClientID  string   `yaml:"client_id"`
	} `yaml:"nats"`

	Database struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Name     string `yaml:"name"`
		SSLMode  string `yaml:"ssl_mode"`
	} `yaml:"database"`

	Notification struct {
		Push struct {
			Enabled  bool   `yaml:"enabled"`
			Provider string `yaml:"provider"`
			Timeout  int    `yaml:"timeout"`
		} `yaml:"push"`
		Email struct {
			Enabled      bool   `yaml:"enabled"`
			SMTPHost     string `yaml:"smtp_host"`
			SMTPPort     int    `yaml:"smtp_port"`
			SMTPUser     string `yaml:"smtp_user"`
			SMTPPassword string `yaml:"smtp_password"`
			FromAddress  string `yaml:"from_address"`
			FromName     string `yaml:"from_name"`
		} `yaml:"email"`
		SMS struct {
			Enabled    bool   `yaml:"enabled"`
			Provider   string `yaml:"provider"`
			AccountSID string `yaml:"account_sid"`
			AuthToken  string `yaml:"auth_token"`
			FromNumber string `yaml:"from_number"`
		} `yaml:"sms"`
		InApp struct {
			Enabled       bool `yaml:"enabled"`
			MaxPerUser    int  `yaml:"max_per_user"`
			RetentionDays int  `yaml:"retention_days"`
		} `yaml:"in_app"`
		Batch struct {
			Enabled       bool `yaml:"enabled"`
			MaxBatchSize  int  `yaml:"max_batch_size"`
			RetryAttempts int  `yaml:"retry_attempts"`
		} `yaml:"batch"`
	} `yaml:"notification"`

	Logging struct {
		Level  string `yaml:"level"`
		Format string `yaml:"format"`
	} `yaml:"logging"`
}
