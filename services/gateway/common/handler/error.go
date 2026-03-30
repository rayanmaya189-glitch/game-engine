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
	ErrCodeInternalError      ErrorCode = "E1001"
	ErrCodeNotFound           ErrorCode = "E1003"
	ErrCodeUnauthorized       ErrorCode = "E1004"
	ErrCodeForbidden          ErrorCode = "E1005"
	ErrCodeBadRequest         ErrorCode = "E1010"
	ErrCodeValidationError    ErrorCode = "E1011"
	ErrCodeRateLimitError     ErrorCode = "E1007"
	ErrCodeServiceUnavailable ErrorCode = "E1008"

	ErrCodeInvalidCredentials ErrorCode = "E2001"
	ErrCodeTokenExpired       ErrorCode = "E2002"
	ErrCodeTokenInvalid       ErrorCode = "E2003"
	ErrCodeMFARequired        ErrorCode = "E2007"
	ErrCodeMFAInvalid         ErrorCode = "E2008"

	ErrCodeInsufficientFunds ErrorCode = "E4001"
	ErrCodeWalletLocked      ErrorCode = "E4006"

	ErrCodeGameNotFound    ErrorCode = "E5001"
	ErrCodeGameUnavailable ErrorCode = "E5008"
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

// HandleError handles errors and returns appropriate HTTP response
func (h *ErrorHandler) HandleError(c interface{}, ctx *app.RequestContext, err error) {
	requestID := ctx.GetString("request_id")

	fmt.Printf("[%s] Error: %v\n", requestID, err)

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
		ctx.JSON(consts.StatusInternalServerError, NewErrorResponse(
			ErrCodeInternalError,
			"internal server error",
		).WithRequestID(requestID))
	}
}

// HandleServiceError handles gRPC service errors
func (h *ErrorHandler) HandleServiceError(c interface{}, ctx *app.RequestContext, serviceErr error) {
	requestID := ctx.GetString("request_id")

	errResp := NewErrorResponse(
		ErrCodeInternalError,
		"service error",
	).WithRequestID(requestID)

	ctx.JSON(consts.StatusInternalServerError, errResp)
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
