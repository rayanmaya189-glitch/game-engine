package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/redis/go-redis/v9"
)

// RoutePermission maps HTTP method + path to required permission
type RoutePermission struct {
	Method     string
	PathPrefix string
	Permission string
}

// RoutePermissionMap defines which permission is required for each admin route
func GetRoutePermissionMap() []RoutePermission {
	return []RoutePermission{
		// Players
		{"GET", "/api/v1/admin/players", "players:view"},
		{"PUT", "/api/v1/admin/players/", "players:edit"},
		// KYC
		{"GET", "/api/v1/admin/kyc", "kyc:view"},
		{"PUT", "/api/v1/admin/kyc/", "kyc:approve"},
		// Games
		{"GET", "/api/v1/admin/games", "games:view"},
		{"POST", "/api/v1/admin/games", "games:create"},
		{"PUT", "/api/v1/admin/games/", "games:edit"},
		// Wallet
		{"GET", "/api/v1/admin/wallet", "wallet:view"},
		{"POST", "/api/v1/admin/wallet/adjust", "wallet:adjust"},
		// Commission Claims
		{"GET", "/api/v1/admin/claims/commission", "commission:view"},
		{"POST", "/api/v1/admin/claims/commission/", "commission:approve"},
		// Rebet Claims
		{"GET", "/api/v1/admin/claims/rebet", "rebet:view"},
		{"POST", "/api/v1/admin/claims/rebet/", "rebet:approve"},
		// Insurance Claims
		{"GET", "/api/v1/admin/claims/insurance", "insurance:view"},
		{"POST", "/api/v1/admin/claims/insurance/", "insurance:approve"},
		// Settlements
		{"GET", "/api/v1/admin/claims/settlements", "settlements:view"},
		{"GET", "/api/v1/admin/claims/statistics", "claims:view"},
		// Merchants
		{"GET", "/api/v1/admin/merchants", "merchants:view"},
		{"POST", "/api/v1/admin/merchants", "merchants:create"},
		{"PUT", "/api/v1/admin/merchants/", "merchants:edit"},
		{"DELETE", "/api/v1/admin/merchants/", "merchants:delete"},
		// Agents
		{"GET", "/api/v1/admin/agents", "agents:view"},
		{"POST", "/api/v1/admin/agents", "agents:create"},
		{"PUT", "/api/v1/admin/agents/", "agents:edit"},
		{"DELETE", "/api/v1/admin/agents/", "agents:delete"},
		// Tournaments
		{"GET", "/api/v1/admin/tournaments", "tournaments:view"},
		{"POST", "/api/v1/admin/tournaments", "tournaments:create"},
		{"PUT", "/api/v1/admin/tournaments/", "tournaments:edit"},
		{"DELETE", "/api/v1/admin/tournaments/", "tournaments:delete"},
		// Jackpots
		{"GET", "/api/v1/admin/jackpots", "jackpots:view"},
		{"POST", "/api/v1/admin/jackpots", "jackpots:create"},
		{"PUT", "/api/v1/admin/jackpots/", "jackpots:edit"},
		{"DELETE", "/api/v1/admin/jackpots/", "jackpots:delete"},
		// Bonuses
		{"GET", "/api/v1/admin/bonuses", "bonuses:view"},
		{"POST", "/api/v1/admin/bonuses", "bonuses:create"},
		{"PUT", "/api/v1/admin/bonuses/", "bonuses:edit"},
		{"DELETE", "/api/v1/admin/bonuses/", "bonuses:delete"},
		// Payments
		{"GET", "/api/v1/admin/payments", "payments:view"},
		{"PUT", "/api/v1/admin/payments/", "payments:approve"},
		// Reports
		{"GET", "/api/v1/admin/reports", "reports:view"},
	}
}

// RBACMiddleware provides permission-based access control
type RBACMiddleware struct {
	redis          *redis.Client
	routePermMap   []RoutePermission
	permissionTTL  time.Duration
}

// NewRBACMiddleware creates a new RBAC middleware
func NewRBACMiddleware(redisClient *redis.Client) *RBACMiddleware {
	return &RBACMiddleware{
		redis:         redisClient,
		routePermMap:  GetRoutePermissionMap(),
		permissionTTL: 5 * time.Minute,
	}
}

// RequirePermission creates middleware that checks for a specific permission
func (m *RBACMiddleware) RequirePermission(requiredPermission string) app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		userID := string(ctx.GetString("user_id"))
		if userID == "" {
			ctx.JSON(consts.StatusUnauthorized, map[string]interface{}{
				"error": "not authenticated",
				"code":  "E1004",
			})
			ctx.Abort()
			return
		}

		role := string(ctx.GetString("role"))

		// Superadmin has all permissions
		if role == "superadmin" {
			ctx.Next(c)
			return
		}

		// Get user permissions from cache or compute
		permissions, err := m.getUserPermissions(userID, role)
		if err != nil {
			ctx.JSON(consts.StatusInternalServerError, map[string]interface{}{
				"error": "failed to check permissions",
				"code":  "E1001",
			})
			ctx.Abort()
			return
		}

		// Check if user has the required permission
		hasPermission := false
		for _, perm := range permissions {
			if perm == requiredPermission {
				hasPermission = true
				break
			}
		}

		if !hasPermission {
			ctx.JSON(consts.StatusForbidden, map[string]interface{}{
				"error":              "insufficient permissions",
				"code":               "E1005",
				"required_permission": requiredPermission,
			})
			ctx.Abort()
			return
		}

		ctx.Next(c)
	}
}

// RequireAnyPermission checks if user has at least one of the given permissions
func (m *RBACMiddleware) RequireAnyPermission(requiredPermissions ...string) app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		userID := string(ctx.GetString("user_id"))
		if userID == "" {
			ctx.JSON(consts.StatusUnauthorized, map[string]interface{}{
				"error": "not authenticated",
				"code":  "E1004",
			})
			ctx.Abort()
			return
		}

		role := string(ctx.GetString("role"))

		if role == "superadmin" {
			ctx.Next(c)
			return
		}

		permissions, err := m.getUserPermissions(userID, role)
		if err != nil {
			ctx.JSON(consts.StatusInternalServerError, map[string]interface{}{
				"error": "failed to check permissions",
				"code":  "E1001",
			})
			ctx.Abort()
			return
		}

		permMap := make(map[string]bool)
		for _, p := range permissions {
			permMap[p] = true
		}

		for _, required := range requiredPermissions {
			if permMap[required] {
				ctx.Next(c)
				return
			}
		}

		ctx.JSON(consts.StatusForbidden, map[string]interface{}{
			"error":               "insufficient permissions",
			"code":                "E1005",
			"required_permissions": requiredPermissions,
		})
		ctx.Abort()
	}
}

// RequireRole checks if user has one of the specified roles
func (m *RBACMiddleware) RequireRole(allowedRoles ...string) app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		role := string(ctx.GetString("role"))

		for _, allowed := range allowedRoles {
			if role == allowed {
				ctx.Next(c)
				return
			}
		}

		ctx.JSON(consts.StatusForbidden, map[string]interface{}{
			"error": "insufficient role",
			"code":  "E1005",
		})
		ctx.Abort()
	}
}

// AutoRoutePermission checks permission based on the route being accessed
func (m *RBACMiddleware) AutoRoutePermission() app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		userID := string(ctx.GetString("user_id"))
		if userID == "" {
			ctx.JSON(consts.StatusUnauthorized, map[string]interface{}{
				"error": "not authenticated",
				"code":  "E1004",
			})
			ctx.Abort()
			return
		}

		role := string(ctx.GetString("role"))

		if role == "superadmin" {
			ctx.Next(c)
			return
		}

		method := string(ctx.Method())
		path := string(ctx.Request.URI().Path())

		// Find required permission for this route
		requiredPerm := m.findRequiredPermission(method, path)
		if requiredPerm == "" {
			// No specific permission required, allow access
			ctx.Next(c)
			return
		}

		permissions, err := m.getUserPermissions(userID, role)
		if err != nil {
			ctx.JSON(consts.StatusInternalServerError, map[string]interface{}{
				"error": "failed to check permissions",
				"code":  "E1001",
			})
			ctx.Abort()
			return
		}

		for _, perm := range permissions {
			if perm == requiredPerm {
				ctx.Next(c)
				return
			}
		}

		ctx.JSON(consts.StatusForbidden, map[string]interface{}{
			"error":              "insufficient permissions",
			"code":               "E1005",
			"required_permission": requiredPerm,
			"route":              fmt.Sprintf("%s %s", method, path),
		})
		ctx.Abort()
	}
}

func (m *RBACMiddleware) getUserPermissions(userID, role string) ([]string, error) {
	if m.redis != nil {
		cacheKey := fmt.Sprintf("rbac:perms:%s", userID)
		cached, err := m.redis.Get(context.Background(), cacheKey).Result()
		if err == nil {
			var perms []string
			if json.Unmarshal([]byte(cached), &perms) == nil {
				return perms, nil
			}
		}
	}

	// Get permissions from role mapping
	rolePerms := GetRolePermissionsStatic(role)

	if m.redis != nil {
		data, _ := json.Marshal(rolePerms)
		m.redis.Set(context.Background(), fmt.Sprintf("rbac:perms:%s", userID), data, m.permissionTTL)
	}

	return rolePerms, nil
}

func (m *RBACMiddleware) findRequiredPermission(method, path string) string {
	for _, rp := range m.routePermMap {
		if rp.Method == method && hasPrefix(path, rp.PathPrefix) {
			return rp.Permission
		}
	}
	return ""
}

func hasPrefix(s, prefix string) bool {
	return len(s) >= len(prefix) && s[:len(prefix)] == prefix
}

// GetRolePermissionsStatic returns permissions for a role (static mapping)
func GetRolePermissionsStatic(role string) []string {
	rolePerms := map[string][]string{
		"admin": {
			"players:view", "players:edit", "players:ban",
			"kyc:view", "kyc:approve", "kyc:reject",
			"games:view", "games:create", "games:edit",
			"wallet:view", "wallet:adjust",
			"claims:view", "claims:approve", "claims:reject", "claims:pay",
			"commission:view", "commission:approve", "commission:reject", "commission:pay",
			"rebet:view", "rebet:approve", "rebet:reject",
			"insurance:view", "insurance:approve", "insurance:reject", "insurance:pay",
			"settlements:view",
			"merchants:view", "merchants:create", "merchants:edit",
			"agents:view", "agents:create", "agents:edit",
			"tournaments:view", "tournaments:create", "tournaments:edit",
			"jackpots:view", "jackpots:create", "jackpots:edit",
			"bonuses:view", "bonuses:create", "bonuses:edit",
			"payments:view", "payments:approve", "payments:reject", "payments:process",
			"reports:view", "reports:export",
			"settings:view",
			"audit:view",
		},
		"support": {
			"players:view", "players:ban",
			"kyc:view", "kyc:approve", "kyc:reject",
			"wallet:view",
			"claims:view", "commission:view", "rebet:view",
			"insurance:view", "settlements:view",
			"merchants:view", "agents:view", "games:view",
			"tournaments:view", "jackpots:view", "bonuses:view",
			"payments:view", "reports:view",
		},
		"finance": {
			"players:view",
			"wallet:view", "wallet:adjust", "wallet:reverse",
			"claims:view", "claims:approve", "claims:reject", "claims:pay",
			"commission:view", "commission:approve", "commission:reject", "commission:pay",
			"rebet:view", "rebet:approve", "rebet:reject",
			"insurance:view", "insurance:approve", "insurance:reject", "insurance:pay",
			"settlements:view",
			"payments:view", "payments:approve", "payments:reject", "payments:process",
			"reports:view", "reports:export",
			"merchants:view", "agents:view",
		},
		"cs": {
			"players:view", "players:edit",
			"kyc:view",
			"wallet:view",
			"claims:view", "commission:view", "rebet:view",
			"insurance:view", "settlements:view",
			"games:view", "tournaments:view", "bonuses:view",
			"payments:view",
		},
		"audit": {
			"players:view", "kyc:view", "wallet:view",
			"claims:view", "commission:view", "rebet:view",
			"insurance:view", "settlements:view",
			"merchants:view", "agents:view", "games:view",
			"tournaments:view", "jackpots:view", "bonuses:view",
			"payments:view",
			"reports:view", "reports:export",
			"audit:view",
			"admin_users:view", "roles:view",
		},
		"marketing": {
			"players:view", "games:view",
			"tournaments:view", "tournaments:create", "tournaments:edit",
			"jackpots:view", "jackpots:create", "jackpots:edit",
			"bonuses:view", "bonuses:create", "bonuses:edit",
			"reports:view", "reports:export",
			"merchants:view", "agents:view",
		},
	}

	if perms, ok := rolePerms[role]; ok {
		return perms
	}
	return []string{}
}
