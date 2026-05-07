// Package common provides shared utilities, error codes, and constants
// used across all services in the casino game engine.
package common

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// ===========================================
// Error Types
// ===========================================

// Error represents a structured error in the game engine
type Error struct {
	Code       string `json:"code"`
	Message    string `json:"message"`
	Details    string `json:"details,omitempty"`
	HTTPStatus int    `json:"-"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}

// NewError creates a new Error with the given code and message
func NewError(code, message string) *Error {
	return &Error{
		Code:       code,
		Message:    message,
		HTTPStatus: http.StatusInternalServerError,
	}
}

// NewErrorWithDetails creates a new Error with the given code, message, and details
func NewErrorWithDetails(code, message, details string) *Error {
	return &Error{
		Code:       code,
		Message:    message,
		Details:    details,
		HTTPStatus: http.StatusInternalServerError,
	}
}

// WithHTTPStatus sets the HTTP status code for the error
func (e *Error) WithHTTPStatus(status int) *Error {
	e.HTTPStatus = status
	return e
}

// ToJSON converts the error to JSON
func (e *Error) ToJSON() string {
	data, _ := json.Marshal(e)
	return string(data)
}

// ===========================================
// gRPC Status Error Mapping
// ===========================================

// ToGRPCStatus converts the error to a gRPC status
func (e *Error) ToGRPCStatus() error {
	code := codes.Internal
	switch e.Code {
	case ErrCodeNotFound:
		code = codes.NotFound
	case ErrCodeUnauthorized:
		code = codes.Unauthenticated
	case ErrCodeForbidden:
		code = codes.PermissionDenied
	case ErrCodeInvalidRequest, ErrCodeInvalidCredentials, ErrCodeInvalidAmount:
		code = codes.InvalidArgument
	case ErrCodeRateLimitExceeded:
		code = codes.ResourceExhausted
	case ErrCodeInternalError:
		code = codes.Internal
	}
	return status.Error(code, e.Message)
}

// ===========================================
// Predefined Errors
// ===========================================

var (
	// General errors
	ErrInternalError     = NewError(ErrCodeInternalError, "An internal error occurred").WithHTTPStatus(http.StatusInternalServerError)
	ErrInvalidRequest    = NewError(ErrCodeInvalidRequest, "Invalid request").WithHTTPStatus(http.StatusBadRequest)
	ErrNotFound          = NewError(ErrCodeNotFound, "Resource not found").WithHTTPStatus(http.StatusNotFound)
	ErrUnauthorized      = NewError(ErrCodeUnauthorized, "Unauthorized").WithHTTPStatus(http.StatusUnauthorized)
	ErrForbidden         = NewError(ErrCodeForbidden, "Forbidden").WithHTTPStatus(http.StatusForbidden)
	ErrConflict          = NewError(ErrCodeConflict, "Resource conflict").WithHTTPStatus(http.StatusConflict)
	ErrRateLimitExceeded = NewError(ErrCodeRateLimitExceeded, "Rate limit exceeded").WithHTTPStatus(http.StatusTooManyRequests)

	// Authentication errors
	ErrInvalidCredentials = NewError(ErrCodeInvalidCredentials, "Invalid credentials").WithHTTPStatus(http.StatusUnauthorized)
	ErrTokenExpired       = NewError(ErrCodeTokenExpired, "Token expired").WithHTTPStatus(http.StatusUnauthorized)
	ErrTokenInvalid       = NewError(ErrCodeTokenInvalid, "Invalid token").WithHTTPStatus(http.StatusUnauthorized)
	ErrAccountLocked      = NewError(ErrCodeAccountLocked, "Account locked").WithHTTPStatus(http.StatusForbidden)
	ErrAccountDisabled    = NewError(ErrCodeAccountDisabled, "Account disabled").WithHTTPStatus(http.StatusForbidden)

	// User errors
	ErrUserNotFound      = NewError(ErrCodeUserNotFound, "User not found").WithHTTPStatus(http.StatusNotFound)
	ErrUserAlreadyExists = NewError(ErrCodeUserAlreadyExists, "User already exists").WithHTTPStatus(http.StatusConflict)
	ErrInvalidEmail      = NewError(ErrCodeInvalidEmail, "Invalid email address").WithHTTPStatus(http.StatusBadRequest)
	ErrInvalidUsername   = NewError(ErrCodeInvalidUsername, "Invalid username").WithHTTPStatus(http.StatusBadRequest)
	ErrWeakPassword      = NewError(ErrCodeWeakPassword, "Password is too weak").WithHTTPStatus(http.StatusBadRequest)

	// Wallet errors
	ErrInsufficientFunds = NewError(ErrCodeInsufficientFunds, "Insufficient funds").WithHTTPStatus(http.StatusBadRequest)
	ErrInvalidAmount     = NewError(ErrCodeInvalidAmount, "Invalid amount").WithHTTPStatus(http.StatusBadRequest)
	ErrTransactionFailed = NewError(ErrCodeTransactionFailed, "Transaction failed").WithHTTPStatus(http.StatusInternalServerError)

	// Game errors
	ErrGameNotFound = NewError(ErrCodeGameNotFound, "Game not found").WithHTTPStatus(http.StatusNotFound)
	ErrGameClosed   = NewError(ErrCodeGameClosed, "Game is closed").WithHTTPStatus(http.StatusBadRequest)
	ErrTableFull    = NewError(ErrCodeTableFull, "Table is full").WithHTTPStatus(http.StatusBadRequest)
)

// ===========================================
// Error Handler
// ===========================================

// ErrorHandler handles errors in the application
type ErrorHandler struct {
	enableDetails bool
}

// NewErrorHandler creates a new ErrorHandler
func NewErrorHandler(enableDetails bool) *ErrorHandler {
	return &ErrorHandler{
		enableDetails: enableDetails,
	}
}

// HandleError handles an error and returns an appropriate response
func (h *ErrorHandler) HandleError(ctx context.Context, err error) (int, interface{}) {
	if ge, ok := err.(*Error); ok {
		if !h.enableDetails && ge.Details != "" {
			ge.Details = ""
		}
		return ge.HTTPStatus, ge
	}
	// Unknown error
	return http.StatusInternalServerError, ErrInternalError
}

// ===========================================
// Context Keys
// ===========================================

type contextKey string

const (
	// RequestIDKey is the key for the request ID in the context
	RequestIDKey contextKey = "request_id"

	// UserIDKey is the key for the user ID in the context
	UserIDKey contextKey = "user_id"

	// ServiceNameKey is the key for the service name in the context
	ServiceNameKey contextKey = "service_name"
)

// GetRequestID gets the request ID from the context
func GetRequestID(ctx context.Context) string {
	if id, ok := ctx.Value(RequestIDKey).(string); ok {
		return id
	}
	return ""
}

// GetUserID gets the user ID from the context
func GetUserID(ctx context.Context) string {
	if id, ok := ctx.Value(UserIDKey).(string); ok {
		return id
	}
	return ""
}

// GetServiceName gets the service name from the context
func GetServiceName(ctx context.Context) string {
	if name, ok := ctx.Value(ServiceNameKey).(string); ok {
		return name
	}
	return ""
}
