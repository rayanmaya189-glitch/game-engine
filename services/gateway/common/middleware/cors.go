package middleware

import (
	"strings"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

type CORSConfig struct {
	AllowOrigins     []string
	AllowMethods     []string
	AllowHeaders     []string
	ExposeHeaders    []string
	AllowCredentials bool
	MaxAge           int
}

var defaultCORSConfig = CORSConfig{
	AllowOrigins:     []string{"*"},
	AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
	AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization", "X-API-Key", "X-Request-ID", "X-MFA-Code"},
	ExposeHeaders:    []string{"X-Request-ID", "X-RateLimit-Limit", "X-RateLimit-Remaining", "X-RateLimit-Reset"},
	AllowCredentials: false,
	MaxAge:           86400,
}

type CORSMiddleware struct {
	config *CORSConfig
}

func NewCORSMiddleware(config *CORSConfig) *CORSMiddleware {
	if config == nil {
		config = &defaultCORSConfig
	}
	// Use defaults for empty slices
	if len(config.AllowOrigins) == 0 {
		config.AllowOrigins = defaultCORSConfig.AllowOrigins
	}
	if len(config.AllowMethods) == 0 {
		config.AllowMethods = defaultCORSConfig.AllowMethods
	}
	if len(config.AllowHeaders) == 0 {
		config.AllowHeaders = defaultCORSConfig.AllowHeaders
	}

	return &CORSMiddleware{config: config}
}

// CORS returns a CORS middleware
func (m *CORSMiddleware) CORS() app.HandlerFunc {
	return func(c interface{}, ctx *app.RequestContext) {
		origin := string(ctx.Request.Header.Get("Origin"))

		// Check if origin is allowed
		if !m.isOriginAllowed(origin) {
			ctx.Next(c)
			return
		}

		// Set CORS headers
		if len(m.config.AllowOrigins) == 1 && m.config.AllowOrigins[0] == "*" {
			ctx.Response.Header.Set("Access-Control-Allow-Origin", "*")
		} else {
			ctx.Response.Header.Set("Access-Control-Allow-Origin", origin)
		}

		ctx.Response.Header.Set("Access-Control-Allow-Methods", strings.Join(m.config.AllowMethods, ", "))
		ctx.Response.Header.Set("Access-Control-Allow-Headers", strings.Join(m.config.AllowHeaders, ", "))

		if len(m.config.ExposeHeaders) > 0 {
			ctx.Response.Header.Set("Access-Control-Expose-Headers", strings.Join(m.config.ExposeHeaders, ", "))
		}

		if m.config.AllowCredentials {
			ctx.Response.Header.Set("Access-Control-Allow-Credentials", "true")
		}

		if m.config.MaxAge > 0 {
			ctx.Response.Header.Set("Access-Control-Max-Age", string(rune(m.config.MaxAge)))
		}

		// Handle preflight requests
		if string(ctx.Request.Method) == "OPTIONS" {
			ctx.Response.SetStatusCode(consts.StatusNoContent)
			ctx.Abort()
			return
		}

		ctx.Next(c)
	}
}

// isOriginAllowed checks if the origin is allowed
func (m *CORSMiddleware) isOriginAllowed(origin string) bool {
	if origin == "" {
		return true // Allow non-CORS requests
	}

	for _, allowed := range m.config.AllowOrigins {
		if allowed == "*" {
			return true
		}
		if allowed == origin {
			return true
		}
		// Support wildcard subdomains (e.g., "*.example.com")
		if strings.HasPrefix(allowed, "*.") {
			domain := allowed[2:]
			if strings.HasSuffix(origin, domain) {
				return true
			}
		}
	}

	return false
}

// AdminCORS returns a CORS middleware for admin gateway (stricter)
func (m *CORSMiddleware) AdminCORS(allowedDomains []string) app.HandlerFunc {
	return func(c interface{}, ctx *app.RequestContext) {
		origin := string(ctx.Request.Header.Get("Origin"))

		// Strict origin check for admin
		if !m.isOriginAllowed(origin) {
			ctx.JSON(consts.StatusForbidden, map[string]interface{}{
				"error": "origin not allowed",
				"code":  "ORIGIN_NOT_ALLOWED",
			})
			ctx.Abort()
			return
		}

		ctx.Response.Header.Set("Access-Control-Allow-Origin", origin)
		ctx.Response.Header.Set("Access-Control-Allow-Methods", strings.Join(m.config.AllowMethods, ", "))
		ctx.Response.Header.Set("Access-Control-Allow-Headers", strings.Join(m.config.AllowHeaders, ", "))
		ctx.Response.Header.Set("Access-Control-Allow-Credentials", "true")

		if string(ctx.Request.Method) == "OPTIONS" {
			ctx.Response.SetStatusCode(consts.StatusNoContent)
			ctx.Abort()
			return
		}

		ctx.Next(c)
	}
}

// CORSWithConfig returns a CORS middleware with custom configuration
func CORSWithConfig(config *CORSConfig) app.HandlerFunc {
	middleware := NewCORSMiddleware(config)
	return middleware.CORS()
}
