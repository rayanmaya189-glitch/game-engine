package handler

import (
	"context"

	"github.com/game_engine/notification/internal/service"
)

// NotificationHandler handles notification-related requests
type NotificationHandler struct {
	service *service.NotificationService
}

// NewNotificationHandler creates a new notification handler
func NewNotificationHandler(service *service.NotificationService) *NotificationHandler {
	return &NotificationHandler{
		service: service,
	}
}

// SendNotificationRequest represents a send notification request
type SendNotificationRequest struct {
	Type   service.NotificationType `json:"type"`
	UserID string                   `json:"user_id"`
	Title  string                   `json:"title"`
	Body   string                   `json:"body"`
	Data   map[string]string        `json:"data"`
}

// SendBatchRequest represents a batch notification request
type SendBatchRequest struct {
	Notifications []service.Notification `json:"notifications"`
}

// NotificationResponse represents a notification response
type NotificationResponse struct {
	ID     string                     `json:"id"`
	Type   service.NotificationType   `json:"type"`
	UserID string                     `json:"user_id"`
	Title  string                     `json:"title"`
	Body   string                     `json:"body"`
	Status service.NotificationStatus `json:"status"`
	SentAt string                     `json:"sent_at"`
}

// SendNotification sends a notification
func (h *NotificationHandler) SendNotification(ctx context.Context, req *SendNotificationRequest) (*NotificationResponse, error) {
	notif := &service.Notification{
		Type:   req.Type,
		UserID: req.UserID,
		Title:  req.Title,
		Body:   req.Body,
		Data:   req.Data,
	}

	if err := h.service.SendNotification(ctx, notif); err != nil {
		return nil, err
	}

	return &NotificationResponse{
		ID:     notif.ID,
		Type:   notif.Type,
		UserID: notif.UserID,
		Title:  notif.Title,
		Body:   notif.Body,
		Status: notif.Status,
	}, nil
}

// SendBatch sends batch notifications
func (h *NotificationHandler) SendBatch(ctx context.Context, req *SendBatchRequest) error {
	return h.service.SendBatch(ctx, req.Notifications)
}

// GetNotifications gets user notifications
func (h *NotificationHandler) GetNotifications(ctx context.Context, userID string, limit int) ([]NotificationResponse, error) {
	notifs, err := h.service.GetNotifications(ctx, userID, limit)
	if err != nil {
		return nil, err
	}

	result := make([]NotificationResponse, 0, len(notifs))
	for _, n := range notifs {
		result = append(result, NotificationResponse{
			ID:     n.ID,
			Type:   n.Type,
			UserID: n.UserID,
			Title:  n.Title,
			Body:   n.Body,
			Status: n.Status,
		})
	}

	return result, nil
}

// MarkAsRead marks a notification as read
func (h *NotificationHandler) MarkAsRead(ctx context.Context, notificationID string) error {
	return h.service.MarkAsRead(ctx, notificationID)
}

// GetPreferences gets user preferences
func (h *NotificationHandler) GetPreferences(ctx context.Context, userID string) (*service.Preferences, error) {
	return h.service.GetPreferences(ctx, userID)
}

// SavePreferences saves user preferences
func (h *NotificationHandler) SavePreferences(ctx context.Context, prefs *service.Preferences) error {
	return h.service.SavePreferences(ctx, prefs)
}
