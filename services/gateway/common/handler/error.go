package handler

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// ErrorCode defines error codes for the API
type ErrorCode string

const (
	// Common error codes
	ErrCodeInternalError      ErrorCode = "INTERNAL_ERROR"
	ErrCodeNotFound           ErrorCode = "NOT_FOUND"
	ErrCodeUnauthorized       ErrorCode = "UNAUTHORIZED"
	ErrCodeForbidden          ErrorCode = "FORBIDDEN"
	ErrCodeBadRequest         ErrorCode = "BAD_REQUEST"
	ErrCodeValidationError    ErrorCode = "VALIDATION_ERROR"
	ErrCodeRateLimitError     ErrorCode = "RATE_LIMIT_EXCEEDED"
	ErrCodeServiceUnavailable ErrorCode = "SERVICE_UNAVAILABLE"

	// Auth error codes
	ErrCodeInvalidCredentials ErrorCode = "INVALID_CREDENTIALS"
	ErrCodeTokenExpired       ErrorCode = "TOKEN_EXPIRED"
	ErrCodeTokenInvalid       ErrorCode = "TOKEN_INVALID"
	ErrCodeMFARequired        ErrorCode = "MFA_REQUIRED"
	ErrCodeMFAInvalid         ErrorCode = "MFA_INVALID"

	// Wallet error codes
	ErrCodeInsufficientFunds ErrorCode = "INSUFFICIENT_FUNDS"
	ErrCodeWalletLocked      ErrorCode = "WALLET_LOCKED"

	// Game error codes
	ErrCodeGameNotFound    ErrorCode = "GAME_NOT_FOUND"
	ErrCodeGameUnavailable ErrorCode = "GAME_UNAVAILABLE"
)

// ErrorResponse represents a standard error response
type ErrorResponse struct {
	Error     string                 `json:"error"`
	Code      ErrorCode              `json:"code"`
	Message   string                 `json:"message,omitempty"`
	RequestID string                 `json:"request_id,omitempty"`
	Timestamp string                 `json:"timestamp"`
	Details   map[string]interface{} `json:"details,omitempty"`
}

// NewErrorResponse creates a new error response
func NewErrorResponse(code ErrorCode, message string) ErrorResponse {
	return ErrorResponse{
		Error:     string(code),
		Code:      code,
		Message:   message,
		Timestamp: time.Now().UTC().Format(time.RFC3339),
	}
}

// WithRequestID adds request ID to error response
func (e ErrorResponse) WithRequestID(requestID string) ErrorResponse {
	e.RequestID = requestID
	return e
}

// WithDetails adds additional details to error response
func (e ErrorResponse) WithDetails(details map[string]interface{}) ErrorResponse {
	e.Details = details
	return e
}

// ErrorHandler is the global error handler
type ErrorHandler struct{}

func NewErrorHandler() *ErrorHandler {
	return &ErrorHandler{}
}

// HandleError handles errors and returns appropriate HTTP response
func (h *ErrorHandler) HandleError(c interface{}, ctx *app.RequestContext, err error) {
	requestID := ctx.GetString("request_id")

	// Log the error
	fmt.Printf("[%s] Error: %v\n", requestID, err)

	// Determine error type and return appropriate response
	switch e := err.(type) {
	case *ValidationError:
		ctx.JSON(consts.StatusBadRequest, NewErrorResponse(
			ErrCodeValidationError,
			e.Message,
		).WithRequestID(requestID))
	case *NotFoundError:
		ctx.JSON(consts.StatusNotFound, NewErrorResponse(
			ErrCodeNotFound,
			e.Message,
		).WithRequestID(requestID))
	case *UnauthorizedError:
		ctx.JSON(consts.StatusUnauthorized, NewErrorResponse(
			ErrCodeUnauthorized,
			e.Message,
		).WithRequestID(requestID))
	case *ForbiddenError:
		ctx.JSON(consts.StatusForbidden, NewErrorResponse(
			ErrCodeForbidden,
			e.Message,
		).WithRequestID(requestID))
	case *RateLimitError:
		ctx.JSON(consts.StatusTooManyRequests, NewErrorResponse(
			ErrCodeRateLimitError,
			e.Message,
		).WithRequestID(requestID))
	default:
		// Default to internal server error
		ctx.JSON(consts.StatusInternalServerError, NewErrorResponse(
			ErrCodeInternalError,
			"internal server error",
		).WithRequestID(requestID))
	}
}

// ValidationError represents a validation error
type ValidationError struct {
	Message string
	Field   string
}

func (e *ValidationError) Error() string {
	if e.Field != "" {
		return fmt.Sprintf("validation error on field %s: %s", e.Field, e.Message)
	}
	return "validation error: " + e.Message
}

// NotFoundError represents a not found error
type NotFoundError struct {
	Message  string
	Resource string
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("%s not found: %s", e.Resource, e.Message)
}

// UnauthorizedError represents an unauthorized error
type UnauthorizedError struct {
	Message string
}

func (e *UnauthorizedError) Error() string {
	return e.Message
}

// ForbiddenError represents a forbidden error
type ForbiddenError struct {
	Message string
}

func (e *ForbiddenError) Error() string {
	return e.Message
}

// RateLimitError represents a rate limit error
type RateLimitError struct {
	Message    string
	RetryAfter int
}

func (e *RateLimitError) Error() string {
	return e.Message
}

// ServiceError represents a service error
type ServiceError struct {
	Message    string
	Code       ErrorCode
	StatusCode int
}

func (e *ServiceError) Error() string {
	return e.Message
}

// HandleServiceError handles gRPC service errors
func (h *ErrorHandler) HandleServiceError(c interface{}, ctx *app.RequestContext, serviceErr error) {
	requestID := ctx.GetString("request_id")

	// Parse gRPC error
	// In production, you'd parse the actual gRPC status

	errResp := NewErrorResponse(
		ErrCodeInternalError,
		"service error",
	).WithRequestID(requestID)

	ctx.JSON(consts.StatusInternalServerError, errResp)
}

// NotFoundHandler handles 404 Not Found
func (h *ErrorHandler) NotFoundHandler(ctx *app.RequestContext) {
	requestID := ctx.GetString("request_id")
	ctx.JSON(consts.StatusNotFound, NewErrorResponse(
		ErrCodeNotFound,
		"endpoint not found",
	).WithRequestID(requestID))
}

// MethodNotAllowedHandler handles 405 Method Not Allowed
func (h *ErrorHandler) MethodNotAllowedHandler(ctx *app.RequestContext) {
	requestID := ctx.GetString("request_id")
	ctx.JSON(consts.StatusMethodNotAllowed, NewErrorResponse(
		ErrCodeBadRequest,
		"method not allowed",
	).WithRequestID(requestID))
}

// SendJSONError sends a JSON error response
func SendJSONError(ctx *app.RequestContext, statusCode int, code ErrorCode, message string) {
	requestID := ctx.GetString("request_id")
	ctx.JSON(statusCode, NewErrorResponse(code, message).WithRequestID(requestID))
}

// SendSuccess sends a success response
func SendSuccess(ctx *app.RequestContext, data interface{}) {
	ctx.JSON(consts.StatusOK, map[string]interface{}{
		"success": true,
		"data":    data,
	})
}

// SendCreated sends a 201 Created response
func SendCreated(ctx *app.RequestContext, data interface{}) {
	ctx.JSON(consts.StatusCreated, map[string]interface{}{
		"success": true,
		"data":    data,
	})
}

// SendNoContent sends a 204 No Content response
func SendNoContent(ctx *app.RequestContext) {
	ctx.Status(consts.StatusNoContent)
}

// HandleHealthCheck handles health check requests
func HandleHealthCheck(ctx *app.RequestContext) {
	ctx.JSON(consts.StatusOK, map[string]interface{}{
		"status":    "healthy",
		"timestamp": time.Now().UTC().Format(time.RFC3339),
	})
}

// HandleReadinessCheck handles readiness check requests
func HandleReadinessCheck(ctx *app.RequestContext) {
	// In production, check dependencies (Redis, PostgreSQL, gRPC services)
	ctx.JSON(consts.StatusOK, map[string]interface{}{
		"status":    "ready",
		"timestamp": time.Now().UTC().Format(time.RFC3339),
	})
}

// ParseRequestBody parses JSON request body
func ParseRequestBody(ctx *app.RequestContext, target interface{}) error {
	if len(ctx.Request.Body()) == 0 {
		return &ValidationError{Message: "request body is required"}
	}

	err := json.Unmarshal(ctx.Request.Body(), target)
	if err != nil {
		return &ValidationError{Message: "invalid JSON body"}
	}

	return nil
}

// SendErrorResponse sends a standardized error response
func SendErrorResponse(ctx *app.RequestContext, statusCode int, errCode ErrorCode, message string, details map[string]interface{}) {
	requestID := ctx.GetString("request_id")

	response := NewErrorResponse(errCode, message).WithRequestID(requestID)
	if details != nil {
		response = response.WithDetails(details)
	}

	ctx.JSON(statusCode, response)
}

// HandlePanic recovers from panics and returns error response
func HandlePanic(c interface{}, ctx *app.RequestContext) {
	if err := recover(); err != nil {
		requestID := ctx.GetString("request_id")
		fmt.Printf("[%s] Panic recovered: %v\n", requestID, err)

		ctx.JSON(consts.StatusInternalServerError, NewErrorResponse(
			ErrCodeInternalError,
			"internal server error",
		).WithRequestID(requestID))
	}
}

// TimeoutHandler handles request timeouts
func TimeoutHandler(ctx *app.RequestContext) {
	requestID := ctx.GetString("request_id")
	ctx.JSON(consts.StatusGatewayTimeout, NewErrorResponse(
		ErrCodeServiceUnavailable,
		"request timeout",
	).WithRequestID(requestID))
}

// ContextCanceledHandler handles context canceled errors
func ContextCanceledHandler(ctx *app.RequestContext) {
	requestID := ctx.GetString("request_id")
	ctx.JSON(consts.StatusRequestTimeout, NewErrorResponse(
		ErrCodeServiceUnavailable,
		"request cancelled",
	).WithRequestID(requestID))
}
