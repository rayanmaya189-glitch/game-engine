package handler

import (
	"context"
	"fmt"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// NotFoundHandler handles 404 Not Found
func (h *ErrorHandler) NotFoundHandler(c context.Context, ctx *app.RequestContext) {
	requestID := ctx.GetString("request_id")
	ctx.JSON(consts.StatusNotFound, NewErrorResponse(
		ErrCodeNotFound,
		"endpoint not found",
	).WithRequestID(requestID))
}

// MethodNotAllowedHandler handles 405 Method Not Allowed
func (h *ErrorHandler) MethodNotAllowedHandler(c context.Context, ctx *app.RequestContext) {
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
func HandleHealthCheck(c context.Context, ctx *app.RequestContext) {
	ctx.JSON(consts.StatusOK, map[string]interface{}{
		"status":    "healthy",
		"timestamp": time.Now().UTC().Format(time.RFC3339),
	})
}

// HandleReadinessCheck handles readiness check requests
func HandleReadinessCheck(c context.Context, ctx *app.RequestContext) {
	ctx.JSON(consts.StatusOK, map[string]interface{}{
		"status":    "ready",
		"timestamp": time.Now().UTC().Format(time.RFC3339),
	})
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
