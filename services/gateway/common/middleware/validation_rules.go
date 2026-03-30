package middleware

import (
	"strconv"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// Common validation rules for all gateways
func GetCommonValidationRules() map[string][]ValidationRule {
	return map[string][]ValidationRule{
		"/api/v1/auth/register": {
			{Field: "username", Type: "required", Message: "username is required"},
			{Field: "email", Type: "required", Message: "email is required"},
			{Field: "email", Type: "email", Message: "invalid email format"},
			{Field: "password", Type: "required", Message: "password is required"},
			{Field: "password", Type: "min", Min: 8, Message: "password must be at least 8 characters"},
		},
		"/api/v1/auth/login": {
			{Field: "username", Type: "required", Message: "username is required"},
			{Field: "password", Type: "required", Message: "password is required"},
		},
	}
}

// Player validation rules
func GetPlayerValidationRules() map[string][]ValidationRule {
	rules := GetCommonValidationRules()

	rules["/api/v1/users/profile"] = []ValidationRule{
		{Field: "username", Type: "alphanumeric", Message: "username must be alphanumeric"},
	}

	rules["/api/v1/wallet/deposit"] = []ValidationRule{
		{Field: "amount", Type: "required", Message: "amount is required"},
		{Field: "amount", Type: "min", Min: 1, Message: "amount must be at least 1"},
	}

	rules["/api/v1/wallet/withdraw"] = []ValidationRule{
		{Field: "amount", Type: "required", Message: "amount is required"},
		{Field: "amount", Type: "min", Min: 1, Message: "amount must be at least 1"},
	}

	return rules
}

// Admin validation rules
func GetAdminValidationRules() map[string][]ValidationRule {
	rules := GetCommonValidationRules()

	rules["/api/v1/admin/players/:id"] = []ValidationRule{
		{Field: "id", Type: "uuid", Message: "invalid player ID"},
	}

	rules["/api/v1/admin/kyc/:id/approve"] = []ValidationRule{
		{Field: "id", Type: "uuid", Message: "invalid KYC ID"},
	}

	rules["/api/v1/admin/kyc/:id/reject"] = []ValidationRule{
		{Field: "id", Type: "uuid", Message: "invalid KYC ID"},
		{Field: "reason", Type: "required", Message: "rejection reason is required"},
	}

	return rules
}

// Merchant validation rules
func GetMerchantValidationRules() map[string][]ValidationRule {
	return map[string][]ValidationRule{
		"/api/v1/merchant/webhooks/register": {
			{Field: "url", Type: "required", Message: "webhook URL is required"},
			{Field: "events", Type: "required", Message: "events are required"},
		},
	}
}

// Agent validation rules
func GetAgentValidationRules() map[string][]ValidationRule {
	return map[string][]ValidationRule{
		"/api/v1/affiliate/tracking/click": {
			{Field: "code", Type: "required", Message: "affiliate code is required"},
		},
	}
}

// ValidateRequestBody validates JSON body
func ValidateRequestBody(ctx *app.RequestContext, target interface{}) error {
	if len(ctx.Request.Body()) == 0 {
		return &ValidationError{Field: "body", Message: "request body is required"}
	}

	// Note: In Hertz, body parsing would need to use ctx.Bind() or similar
	// This is a placeholder for the validation logic
	return nil
}

// PathParamValidator validates path parameters
func PathParamValidator(paramName, paramType string) app.HandlerFunc {
	return func(c interface{}, ctx *app.RequestContext) {
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
	return func(c interface{}, ctx *app.RequestContext) {
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
			// Additional enum validation would be handled separately
		}

		ctx.Next(c)
	}
}
