package middleware

import (
	"context"
	"fmt"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/redis/go-redis/v9"
)

type RateLimiterConfig struct {
	RedisClient       *redis.Client
	RequestsPerMinute int
	BurstSize         int
	KeyPrefix         string
}

type RateLimiterMiddleware struct {
	config *RateLimiterConfig
}

func NewRateLimiterMiddleware(config *RateLimiterConfig) *RateLimiterMiddleware {
	if config.RequestsPerMinute == 0 {
		config.RequestsPerMinute = 100 // default
	}
	if config.KeyPrefix == "" {
		config.KeyPrefix = "ratelimit"
	}

	return &RateLimiterMiddleware{config: config}
}

// RateLimiter returns a rate limiting middleware using token bucket algorithm
func (m *RateLimiterMiddleware) RateLimiter() app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		// Get client identifier (user ID or IP)
		var key string
		if userID := ctx.GetString("user_id"); userID != "" {
			key = fmt.Sprintf("%s:user:%s", m.config.KeyPrefix, userID)
		} else if merchantID := ctx.GetString("merchant_id"); merchantID != "" {
			key = fmt.Sprintf("%s:merchant:%s", m.config.KeyPrefix, merchantID)
		} else {
			// Fallback to IP
			ip := ctx.RemoteAddr().String()
			key = fmt.Sprintf("%s:ip:%s", m.config.KeyPrefix, ip)
		}

		// Check rate limit
		allowed, remaining, resetTime, err := m.checkRateLimit(key)
		if err != nil {
			// If Redis is down, allow the request but log warning
			ctx.Next(c)
			return
		}

		// Set rate limit headers
		ctx.Response.Header.Set("X-RateLimit-Limit", fmt.Sprintf("%d", m.config.RequestsPerMinute))
		ctx.Response.Header.Set("X-RateLimit-Remaining", fmt.Sprintf("%d", remaining))
		ctx.Response.Header.Set("X-RateLimit-Reset", fmt.Sprintf("%d", resetTime))

		if !allowed {
			ctx.JSON(consts.StatusTooManyRequests, map[string]interface{}{
				"error":       "rate limit exceeded",
				"code":        "RATE_LIMIT_EXCEEDED",
				"retry_after": resetTime - time.Now().Unix(),
			})
			ctx.Abort()
			return
		}

		ctx.Next(c)
	}
}

// checkRateLimit checks if the request is within rate limits
func (m *RateLimiterMiddleware) checkRateLimit(key string) (bool, int, int64, error) {
	ctx := context.Background()
	now := time.Now()
	windowStart := now.Truncate(time.Minute).Unix()
	resetTime := now.Truncate(time.Minute).Add(time.Minute).Unix()

	// Use sliding window rate limiting with Redis
	rateKey := fmt.Sprintf("%s:%d", key, windowStart)

	// Increment counter
	count, err := m.config.RedisClient.Incr(ctx, rateKey).Result()
	if err != nil {
		return true, 0, 0, err
	}

	// Set expiry for the key
	m.config.RedisClient.Expire(ctx, rateKey, time.Minute*2)

	// Check if within limit
	allowed := count <= int64(m.config.RequestsPerMinute)
	remaining := m.config.RequestsPerMinute - int(count)
	if remaining < 0 {
		remaining = 0
	}

	return allowed, remaining, resetTime, nil
}

// IPBasedRateLimiter returns a rate limiting middleware for IP-based limiting
func (m *RateLimiterMiddleware) IPBasedRateLimiter(requestsPerMinute int) app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		ip := ctx.RemoteAddr().String()
		key := fmt.Sprintf("%s:ip:%s", m.config.KeyPrefix, ip)

		allowed, remaining, resetTime, err := m.checkRateLimit(key)
		if err != nil {
			ctx.Next(c)
			return
		}

		ctx.Response.Header.Set("X-RateLimit-Limit", fmt.Sprintf("%d", requestsPerMinute))
		ctx.Response.Header.Set("X-RateLimit-Remaining", fmt.Sprintf("%d", remaining))
		ctx.Response.Header.Set("X-RateLimit-Reset", fmt.Sprintf("%d", resetTime))

		if !allowed {
			ctx.JSON(consts.StatusTooManyRequests, map[string]interface{}{
				"error":       "rate limit exceeded",
				"code":        "RATE_LIMIT_EXCEEDED",
				"retry_after": resetTime - time.Now().Unix(),
			})
			ctx.Abort()
			return
		}

		ctx.Next(c)
	}
}

// CustomRateLimiter allows custom rate limits based on path patterns
func (m *RateLimiterMiddleware) CustomRateLimiter(pathRateLimits map[string]int) app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		path := string(ctx.Request.URI().Path())

		// Find matching rate limit for path
		rateLimit := m.config.RequestsPerMinute // default
		for pattern, limit := range pathRateLimits {
			if matchPath(path, pattern) {
				rateLimit = limit
				break
			}
		}

		// Get client identifier
		var key string
		if userID := ctx.GetString("user_id"); userID != "" {
			key = fmt.Sprintf("%s:user:%s", m.config.KeyPrefix, userID)
		} else {
			ip := ctx.RemoteAddr().String()
			key = fmt.Sprintf("%s:ip:%s", m.config.KeyPrefix, ip)
		}

		// Check rate limit
		allowed, remaining, resetTime, err := m.checkRateLimitWithLimit(key, rateLimit)
		if err != nil {
			ctx.Next(c)
			return
		}

		ctx.Response.Header.Set("X-RateLimit-Limit", fmt.Sprintf("%d", rateLimit))
		ctx.Response.Header.Set("X-RateLimit-Remaining", fmt.Sprintf("%d", remaining))
		ctx.Response.Header.Set("X-RateLimit-Reset", fmt.Sprintf("%d", resetTime))

		if !allowed {
			ctx.JSON(consts.StatusTooManyRequests, map[string]interface{}{
				"error":       "rate limit exceeded",
				"code":        "RATE_LIMIT_EXCEEDED",
				"retry_after": resetTime - time.Now().Unix(),
			})
			ctx.Abort()
			return
		}

		ctx.Next(c)
	}
}

// checkRateLimitWithLimit checks rate limit with custom limit
func (m *RateLimiterMiddleware) checkRateLimitWithLimit(key string, limit int) (bool, int, int64, error) {
	ctx := context.Background()
	now := time.Now()
	windowStart := now.Truncate(time.Minute).Unix()
	resetTime := now.Truncate(time.Minute).Add(time.Minute).Unix()

	rateKey := fmt.Sprintf("%s:%d", key, windowStart)

	count, err := m.config.RedisClient.Incr(ctx, rateKey).Result()
	if err != nil {
		return true, 0, 0, err
	}

	m.config.RedisClient.Expire(ctx, rateKey, time.Minute*2)

	allowed := count <= int64(limit)
	remaining := limit - int(count)
	if remaining < 0 {
		remaining = 0
	}

	return allowed, remaining, resetTime, nil
}

// matchPath checks if a path matches a pattern (supports wildcards)
func matchPath(path, pattern string) bool {
	if path == pattern {
		return true
	}
	// Simple wildcard matching
	if len(pattern) > 0 && pattern[len(pattern)-1] == '*' {
		prefix := pattern[:len(pattern)-1]
		if len(path) >= len(prefix) && path[:len(prefix)] == prefix {
			return true
		}
	}
	return false
}
