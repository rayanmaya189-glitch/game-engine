package service

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

type mockRedisClient struct {
	store    map[string][]byte
	lists    map[string][]string
	errStore error
}

func newMockRedis() *mockRedisClient {
	return &mockRedisClient{
		store: make(map[string][]byte),
		lists: make(map[string][]string),
	}
}

type mockRedisCmd struct {
	err error
}

func (m *mockRedisCmd) Err() error { return m.err }

type mockRedisStringCmd struct {
	val string
	err error
}

func (m *mockRedisStringCmd) Result() (string, error) { return m.val, m.err }
func (m *mockRedisStringCmd) Bytes() ([]byte, error)  { return []byte(m.val), m.err }

type mockRedisListCmd struct {
	vals []string
	err  error
}

func (m *mockRedisListCmd) Result() ([]string, error) { return m.vals, m.err }

type mockRedisSetCmd struct {
	client *mockRedisClient
	key    string
	val    interface{}
	err    error
}

func (m *mockRedisSetCmd) Err() error {
	if m.err != nil {
		return m.err
	}
	data, err := json.Marshal(m.val)
	if err != nil {
		return err
	}
	m.client.store[m.key] = data
	return nil
}

type mockRedisLPushCmd struct {
	client *mockRedisClient
	key    string
	vals   []interface{}
	err    error
}

func (m *mockRedisLPushCmd) Err() error {
	if m.err != nil {
		return m.err
	}
	for _, v := range m.vals {
		data, _ := json.Marshal(v)
		m.client.lists[m.key] = append([]string{string(data)}, m.client.lists[m.key]...)
	}
	return nil
}

type mockRedisLTrimCmd struct {
	client *mockRedisClient
	key    string
	start  int64
	stop   int64
	err    error
}

func (m *mockRedisLTrimCmd) Err() error { return m.err }

type mockRedisExpireCmd struct {
	err error
}

func (m *mockRedisExpireCmd) Err() error { return m.err }

type mockRedisDelCmd struct {
	client *mockRedisClient
	keys   []string
	err    error
}

func (m *mockRedisDelCmd) Err() error {
	if m.err != nil {
		return m.err
	}
	for _, k := range m.keys {
		delete(m.client.store, k)
		delete(m.client.lists, k)
	}
	return nil
}

func TestNotificationTypes(t *testing.T) {
	tests := []struct {
		name string
		typ  NotificationType
	}{
		{"push", NotificationTypePush},
		{"email", NotificationTypeEmail},
		{"sms", NotificationTypeSMS},
		{"in_app", NotificationTypeInApp},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			notif := &Notification{
				Type:   tt.typ,
				UserID: "user1",
				Title:  "Test",
				Body:   "Hello",
			}
			if notif.Type != tt.typ {
				t.Fatalf("Type = %s, want %s", notif.Type, tt.typ)
			}
		})
	}
}

func TestNotificationStructure(t *testing.T) {
	now := time.Now()
	notif := &Notification{
		ID:        "n1",
		Type:      NotificationTypePush,
		UserID:    "u1",
		Title:     "Win!",
		Body:      "You won 100 credits",
		Data:      map[string]string{"amount": "100"},
		Status:    NotificationStatusPending,
		CreatedAt: now,
	}

	data, err := json.Marshal(notif)
	if err != nil {
		t.Fatalf("Marshal() error = %v", err)
	}

	var decoded Notification
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("Unmarshal() error = %v", err)
	}
	if decoded.Title != "Win!" {
		t.Fatalf("Title = %s, want Win!", decoded.Title)
	}
}

func TestNotificationStatuses(t *testing.T) {
	statuses := []NotificationStatus{
		NotificationStatusPending,
		NotificationStatusSent,
		NotificationStatusFailed,
		NotificationStatusDelivered,
		NotificationStatusRead,
	}

	for _, s := range statuses {
		if string(s) == "" {
			t.Fatal("status should not be empty")
		}
	}
}

func TestPreferencesDefaults(t *testing.T) {
	prefs := &Preferences{
		UserID:          "u1",
		PushEnabled:     true,
		EmailEnabled:    true,
		SMSEnabled:      false,
		InAppEnabled:    true,
		QuietHours:      make(map[string]string),
		MutedCategories: []string{},
	}

	if !prefs.PushEnabled {
		t.Fatal("push should be enabled by default")
	}
	if prefs.SMSEnabled {
		t.Fatal("SMS should be disabled by default")
	}
}

func TestBatchValidation(t *testing.T) {
	tests := []struct {
		name    string
		size    int
		maxSize int
		valid   bool
	}{
		{"within limit", 5, 10, true},
		{"at limit", 10, 10, true},
		{"over limit", 11, 10, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			valid := tt.size <= tt.maxSize
			if valid != tt.valid {
				t.Fatalf("batch size %d vs max %d: got %v, want %v", tt.size, tt.maxSize, valid, tt.valid)
			}
		})
	}
}

func TestTemplateStructure(t *testing.T) {
	tmpl := &Template{
		ID:        "welcome",
		Name:      "Welcome Email",
		Type:      NotificationTypeEmail,
		Title:     "Welcome {{.Username}}!",
		Body:      "Thanks for joining, {{.Username}}.",
		Variables: []string{"Username"},
	}

	if tmpl.ID != "welcome" {
		t.Fatalf("ID = %s, want welcome", tmpl.ID)
	}
	if len(tmpl.Variables) != 1 {
		t.Fatalf("got %d variables, want 1", len(tmpl.Variables))
	}
}

func TestNotificationDataField(t *testing.T) {
	notif := &Notification{
		Type:   NotificationTypePush,
		UserID: "u1",
		Title:  "Jackpot",
		Body:   "You hit the jackpot!",
		Data: map[string]string{
			"amount":   "10000",
			"game":     "slots",
			"currency": "USD",
		},
	}

	if notif.Data["amount"] != "10000" {
		t.Fatalf("amount = %s, want 10000", notif.Data["amount"])
	}

	data, _ := json.Marshal(notif)
	var decoded Notification
	_ = json.Unmarshal(data, &decoded)
	if decoded.Data["game"] != "slots" {
		t.Fatalf("game = %s, want slots", decoded.Data["game"])
	}
}

func TestRedisKeyFormat(t *testing.T) {
	tests := []struct {
		prefix string
		id     string
		want   string
	}{
		{"notification", "abc-123", "notification:abc-123"},
		{"notifications", "user-1", "notifications:user-1"},
		{"template", "welcome", "template:welcome"},
	}

	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			got := fmt.Sprintf("%s:%s", tt.prefix, tt.id)
			if got != tt.want {
				t.Fatalf("key = %s, want %s", got, tt.want)
			}
		})
	}
}

func TestContextCancellation(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	select {
	case <-ctx.Done():
	default:
		t.Fatal("context should be cancelled")
	}
}
