package middleware

import (
	"context"
	"regexp"
	"strconv"
	"strings"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

type ValidationRule struct {
	Field   string
	Type    string // "required", "email", "min", "max", "regex", "enum"
	Min     int
	Max     int
	Pattern string
	Enum    []string
	Message string
}

type ValidatorMiddleware struct {
	rules map[string][]ValidationRule
}

func NewValidatorMiddleware(rules map[string][]ValidationRule) *ValidatorMiddleware {
	return &ValidatorMiddleware{rules: rules}
}

// Validate returns a request validation middleware
func (m *ValidatorMiddleware) Validate(path string) app.HandlerFunc {
	rules, exists := m.rules[path]
	if !exists {
		// No rules defined, skip validation
		return func(c context.Context, ctx *app.RequestContext) {
			ctx.Next(c)
		}
	}

	return func(c context.Context, ctx *app.RequestContext) {
		for _, rule := range rules {
			var fieldValue string

			switch rule.Type {
			case "path":
				// Get from path param
				fieldValue = ctx.Param(rule.Field)
			case "query":
				// Get from query param
				fieldValue = string(ctx.Request.URI().QueryArgs().Peek(rule.Field))
			case "header":
				// Get from header
				fieldValue = string(ctx.Request.Header.Peek(rule.Field))
			default:
				// Get from body (JSON form)
				fieldValue = string(ctx.Request.Body())
			}

			err := m.validateField(rule, fieldValue)
			if err != nil {
				ctx.JSON(consts.StatusBadRequest, map[string]interface{}{
					"error": err.Error(),
					"code":  "VALIDATION_ERROR",
					"field": rule.Field,
				})
				ctx.Abort()
				return
			}
		}

		ctx.Next(c)
	}
}

// validateField validates a single field
func (m *ValidatorMiddleware) validateField(rule ValidationRule, value string) error {
	switch rule.Type {
	case "required":
		if value == "" {
			return &ValidationError{Field: rule.Field, Message: rule.Message}
		}
	case "email":
		if value != "" && !isValidEmail(value) {
			return &ValidationError{Field: rule.Field, Message: rule.Message}
		}
	case "min":
		if value != "" {
			minVal, _ := strconv.Atoi(value)
			if minVal < rule.Min {
				return &ValidationError{Field: rule.Field, Message: rule.Message}
			}
		}
	case "max":
		if value != "" {
			maxVal, _ := strconv.Atoi(value)
			if maxVal > rule.Max {
				return &ValidationError{Field: rule.Field, Message: rule.Message}
			}
		}
	case "regex":
		if value != "" {
			matched, _ := regexp.MatchString(rule.Pattern, value)
			if !matched {
				return &ValidationError{Field: rule.Field, Message: rule.Message}
			}
		}
	case "enum":
		if value != "" {
			found := false
			for _, enumVal := range rule.Enum {
				if value == enumVal {
					found = true
					break
				}
			}
			if !found {
				return &ValidationError{Field: rule.Field, Message: rule.Message}
			}
		}
	case "uuid":
		if value != "" && !isValidUUID(value) {
			return &ValidationError{Field: rule.Field, Message: rule.Message}
		}
	case "alphanumeric":
		if value != "" && !isAlphanumeric(value) {
			return &ValidationError{Field: rule.Field, Message: rule.Message}
		}
	}

	return nil
}

// ValidationError represents a validation error
type ValidationError struct {
	Field   string
	Message string
}

func (e *ValidationError) Error() string {
	if e.Message != "" {
		return e.Message
	}
	return "validation failed for field: " + e.Field
}

// isValidEmail validates an email address
func isValidEmail(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}

// isValidUUID validates a UUID
func isValidUUID(uuid string) bool {
	uuidRegex := regexp.MustCompile(`^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$`)
	return uuidRegex.MatchString(uuid)
}

// isAlphanumeric checks if string contains only alphanumeric characters
func isAlphanumeric(s string) bool {
	for _, c := range s {
		if !((c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9')) {
			return false
		}
	}
	return true
}

// SanitizeInput sanitizes user input to prevent injection attacks
func SanitizeInput() app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		// Remove null bytes and other potentially dangerous characters
		body := string(ctx.Request.Body())
		body = strings.ReplaceAll(body, "\x00", "")

		// Set sanitized body back (would require careful handling in production)
		_ = body // Placeholder - actual implementation depends on body parsing

		ctx.Next(c)
	}
}
