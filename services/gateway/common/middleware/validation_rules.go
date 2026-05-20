package middleware

import (
	"context"
	"strconv"
	"strings"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// Common validation rules for all gateways
func GetCommonValidationRules() map[string][]ValidationRule {
	return map[string][]ValidationRule{
		"/api/v1/auth/register": {
			{Field: "username", Type: "required", Message: "username is required"},
			{Field: "username", Type: "min", Min: 3, Message: "username must be at least 3 characters"},
			{Field: "username", Type: "max", Max: 32, Message: "username must be at most 32 characters"},
			{Field: "username", Type: "alphanumeric", Message: "username must be alphanumeric"},
			{Field: "email", Type: "required", Message: "email is required"},
			{Field: "email", Type: "email", Message: "invalid email format"},
			{Field: "password", Type: "required", Message: "password is required"},
			{Field: "password", Type: "min", Min: 8, Message: "password must be at least 8 characters"},
			{Field: "password", Type: "max", Max: 128, Message: "password must be at most 128 characters"},
		},
		"/api/v1/auth/login": {
			{Field: "username", Type: "required", Message: "username or email is required"},
			{Field: "password", Type: "required", Message: "password is required"},
		},
		"/api/v1/auth/refresh": {
			{Field: "refresh_token", Type: "required", Message: "refresh token is required"},
		},
	}
}

// Player validation rules
func GetPlayerValidationRules() map[string][]ValidationRule {
	rules := GetCommonValidationRules()

	rules["/api/v1/users/profile"] = []ValidationRule{
		{Field: "username", Type: "alphanumeric", Message: "username must be alphanumeric"},
		{Field: "username", Type: "min", Min: 3, Message: "username must be at least 3 characters"},
		{Field: "username", Type: "max", Max: 32, Message: "username must be at most 32 characters"},
		{Field: "email", Type: "email", Message: "invalid email format"},
	}

	rules["/api/v1/wallet/deposit"] = []ValidationRule{
		{Field: "amount", Type: "required", Message: "amount is required"},
		{Field: "amount", Type: "min", Min: 1, Message: "amount must be at least 1"},
		{Field: "amount", Type: "max", Max: 1000000, Message: "amount exceeds maximum deposit limit"},
	}

	rules["/api/v1/wallet/withdraw"] = []ValidationRule{
		{Field: "amount", Type: "required", Message: "amount is required"},
		{Field: "amount", Type: "min", Min: 1, Message: "amount must be at least 1"},
		{Field: "amount", Type: "max", Max: 500000, Message: "amount exceeds maximum withdrawal limit"},
	}

	return rules
}

// Admin validation rules
func GetAdminValidationRules() map[string][]ValidationRule {
	rules := GetCommonValidationRules()

	rules["/api/v1/admin/players/:id"] = []ValidationRule{
		{Field: "id", Type: "uuid", Message: "invalid player ID"},
	}

	rules["/api/v1/admin/players/:id/status"] = []ValidationRule{
		{Field: "id", Type: "uuid", Message: "invalid player ID"},
		{Field: "status", Type: "required", Message: "status is required"},
		{Field: "status", Type: "enum", Enum: []string{"active", "suspended", "banned", "inactive"}, Message: "invalid status value"},
	}

	rules["/api/v1/admin/kyc/:id/approve"] = []ValidationRule{
		{Field: "id", Type: "uuid", Message: "invalid KYC ID"},
	}

	rules["/api/v1/admin/kyc/:id/reject"] = []ValidationRule{
		{Field: "id", Type: "uuid", Message: "invalid KYC ID"},
		{Field: "reason", Type: "required", Message: "rejection reason is required"},
		{Field: "reason", Type: "min", Min: 10, Message: "reason must be at least 10 characters"},
	}

	rules["/api/v1/admin/games"] = []ValidationRule{
		{Field: "name", Type: "required", Message: "game name is required"},
		{Field: "category", Type: "required", Message: "game category is required"},
	}

	rules["/api/v1/admin/wallet/adjust"] = []ValidationRule{
		{Field: "player_id", Type: "required", Message: "player ID is required"},
		{Field: "player_id", Type: "uuid", Message: "invalid player ID"},
		{Field: "amount", Type: "required", Message: "adjustment amount is required"},
		{Field: "reason", Type: "required", Message: "adjustment reason is required"},
	}

	return rules
}

// Merchant validation rules
func GetMerchantValidationRules() map[string][]ValidationRule {
	return map[string][]ValidationRule{
		"/api/v1/merchant/webhooks/register": {
			{Field: "url", Type: "required", Message: "webhook URL is required"},
			{Field: "url", Type: "regex", Pattern: `^https://`, Message: "webhook URL must use HTTPS"},
			{Field: "events", Type: "required", Message: "events are required"},
		},
		"/api/v1/merchant/config": {
			{Field: "name", Type: "required", Message: "merchant name is required"},
		},
	}
}

// Agent validation rules
func GetAgentValidationRules() map[string][]ValidationRule {
	return map[string][]ValidationRule{
		"/api/v1/affiliate/tracking/click": {
			{Field: "code", Type: "required", Message: "affiliate code is required"},
			{Field: "code", Type: "min", Min: 3, Message: "affiliate code must be at least 3 characters"},
		},
	}
}

// ValidateRequestBody validates JSON body
func ValidateRequestBody(ctx *app.RequestContext, target interface{}) error {
	if len(ctx.Request.Body()) == 0 {
		return &ValidationError{Field: "body", Message: "request body is required"}
	}

	if err := ctx.BindJSON(target); err != nil {
		return &ValidationError{Field: "body", Message: "invalid JSON body: " + err.Error()}
	}

	return nil
}

// PathParamValidator validates path parameters
func PathParamValidator(paramName, paramType string) app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		value := ctx.Param(paramName)

		switch paramType {
		case "uuid":
			if !isValidUUID(value) {
				ctx.JSON(consts.StatusBadRequest, map[string]interface{}{
					"error": "invalid " + paramName,
					"code":  "VALIDATION_ERROR",
				})
				ctx.Abort()
				return
			}
		case "int":
			if _, err := strconv.Atoi(value); err != nil {
				ctx.JSON(consts.StatusBadRequest, map[string]interface{}{
					"error": "invalid " + paramName,
					"code":  "VALIDATION_ERROR",
				})
				ctx.Abort()
				return
			}
		case "alphanumeric":
			if !isAlphanumeric(value) {
				ctx.JSON(consts.StatusBadRequest, map[string]interface{}{
					"error": paramName + " must be alphanumeric",
					"code":  "VALIDATION_ERROR",
				})
				ctx.Abort()
				return
			}
		}

		ctx.Next(c)
	}
}

// QueryParamValidator validates query parameters
func QueryParamValidator(paramName, paramType string, required bool) app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		value := string(ctx.Request.URI().QueryArgs().Peek(paramName))

		if required && value == "" {
			ctx.JSON(consts.StatusBadRequest, map[string]interface{}{
				"error": paramName + " is required",
				"code":  "VALIDATION_ERROR",
			})
			ctx.Abort()
			return
		}

		if value == "" {
			ctx.Next(c)
			return
		}

		switch paramType {
		case "int":
			if _, err := strconv.Atoi(value); err != nil {
				ctx.JSON(consts.StatusBadRequest, map[string]interface{}{
					"error": "invalid " + paramName,
					"code":  "VALIDATION_ERROR",
				})
				ctx.Abort()
				return
			}
		case "email":
			if !isValidEmail(value) {
				ctx.JSON(consts.StatusBadRequest, map[string]interface{}{
					"error": "invalid email format",
					"code":  "VALIDATION_ERROR",
				})
				ctx.Abort()
				return
			}
		case "enum":
			// Enum validation requires allowed values passed separately via QueryParamEnumValidator
			break
		}

		ctx.Next(c)
	}
}

// QueryParamEnumValidator validates that a query parameter value is within allowed values
func QueryParamEnumValidator(paramName string, allowedValues []string, required bool) app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		value := string(ctx.Request.URI().QueryArgs().Peek(paramName))

		if required && value == "" {
			ctx.JSON(consts.StatusBadRequest, map[string]interface{}{
				"error": paramName + " is required",
				"code":  "VALIDATION_ERROR",
			})
			ctx.Abort()
			return
		}

		if value == "" {
			ctx.Next(c)
			return
		}

		for _, allowed := range allowedValues {
			if value == allowed {
				ctx.Next(c)
				return
			}
		}

		ctx.JSON(consts.StatusBadRequest, map[string]interface{}{
			"error": paramName + " must be one of: " + strings.Join(allowedValues, ", "),
			"code":  "VALIDATION_ERROR",
		})
		ctx.Abort()
	}
}
