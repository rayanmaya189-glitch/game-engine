package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

type LoggerConfig struct {
	OutputPath string
	LogLevel   string
	Format     string // json or text
}

type LoggerMiddleware struct {
	config *LoggerConfig
}

// RequestLog represents the structured log format
type RequestLog struct {
	Timestamp    string                 `json:"timestamp"`
	RequestID    string                 `json:"request_id"`
	Method       string                 `json:"method"`
	Path         string                 `json:"path"`
	Query        string                 `json:"query,omitempty"`
	IP           string                 `json:"ip"`
	UserAgent    string                 `json:"user_agent"`
	StatusCode   int                    `json:"status_code"`
	Latency      float64                `json:"latency_ms"`
	Error        string                 `json:"error,omitempty"`
	RequestSize  int64                  `json:"request_size"`
	ResponseSize int64                  `json:"response_size"`
	UserID       string                 `json:"user_id,omitempty"`
	Role         string                 `json:"role,omitempty"`
	Extra        map[string]interface{} `json:"extra,omitempty"`
}

func NewLoggerMiddleware(config *LoggerConfig) *LoggerMiddleware {
	// Set up custom logger
	if config.OutputPath != "" {
		f, err := os.OpenFile(config.OutputPath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		if err == nil {
			hlog.SetOutput(f)
		}
	}

	// Set log level
	switch config.LogLevel {
	case "debug":
		hlog.SetLevel(hlog.LevelDebug)
	case "info":
		hlog.SetLevel(hlog.LevelInfo)
	case "warn":
		hlog.SetLevel(hlog.LevelWarn)
	case "error":
		hlog.SetLevel(hlog.LevelError)
	default:
		hlog.SetLevel(hlog.LevelInfo)
	}

	return &LoggerMiddleware{config: config}
}

// StructuredLogger returns a structured logging middleware
func (m *LoggerMiddleware) StructuredLogger() app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		start := time.Now()
		requestID := string(ctx.Request.Header.Get("X-Request-ID"))
		if requestID == "" {
			requestID = generateRequestID()
		}

		// Add request ID to context
		ctx.Set("request_id", requestID)
		ctx.Response.Header.Set("X-Request-ID", requestID)

		// Get user info if available
		userID := ctx.GetString("user_id")
		role := ctx.GetString("role")

		// Process request
		ctx.Next(c)

		// Build log entry
		latency := time.Since(start).Milliseconds()
		statusCode := ctx.Response.StatusCode()
		path := string(ctx.Request.URI().Path())
		query := string(ctx.Request.URI().QueryString())
		method := string(ctx.Request.Method())
		ip := ctx.RemoteAddr().String()
		userAgent := string(ctx.Request.Header.UserAgent())

		logEntry := RequestLog{
			Timestamp:    time.Now().UTC().Format(time.RFC3339),
			RequestID:    requestID,
			Method:       method,
			Path:         path,
			Query:        query,
			IP:           ip,
			UserAgent:    userAgent,
			StatusCode:   statusCode,
			Latency:      float64(latency),
			RequestSize:  int64(ctx.Request.Header.ContentLength()),
			ResponseSize: int64(ctx.Response.Header.ContentLength()),
			UserID:       userID,
			Role:         role,
		}

		// Add error if status code >= 400
		if statusCode >= 400 {
			logEntry.Error = string(ctx.Response.Body())
		}

		// Log based on format
		if m.config.Format == "json" {
			logData, _ := json.Marshal(logEntry)
			hlog.DefaultLogger().Info(string(logData))
		} else {
			hlog.DefaultLogger().Info(
				fmt.Sprintf("[%s] %s %s %d %dms - User: %s, IP: %s",
					requestID,
					method,
					path,
					statusCode,
					latency,
					userID,
					ip,
				),
			)
		}

		// Log slow requests
		if latency > 1000 {
			hlog.Warnf("Slow request: %s %s took %dms", method, path, latency)
		}

		// Log errors
		if statusCode >= 500 {
			hlog.Errorf("Server error: %s %s returned %d", method, path, statusCode)
		}
	}
}

// RequestID returns a middleware that adds a request ID to each request
func (m *LoggerMiddleware) RequestID() app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		requestID := string(ctx.Request.Header.Get("X-Request-ID"))
		if requestID == "" {
			requestID = generateRequestID()
		}

		ctx.Set("request_id", requestID)
		ctx.Response.Header.Set("X-Request-ID", requestID)

		ctx.Next(c)
	}
}

// generateRequestID generates a unique request ID
func generateRequestID() string {
	return fmt.Sprintf("req-%d-%s", time.Now().UnixNano(), generateRandomString(8))
}

func generateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[time.Now().UnixNano()%int64(len(charset))]
	}
	return string(b)
}

// PanicRecovery returns a middleware that recovers from panics
func (m *LoggerMiddleware) PanicRecovery() app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		defer func() {
			if err := recover(); err != nil {
				requestID := ctx.GetString("request_id")
				hlog.Errorf("[%s] Panic recovered: %v", requestID, err)

				ctx.JSON(consts.StatusInternalServerError, map[string]interface{}{
					"error":      "internal server error",
					"code":       "INTERNAL_ERROR",
					"request_id": requestID,
				})
				ctx.Abort()
			}
		}()

		ctx.Next(c)
	}
}
