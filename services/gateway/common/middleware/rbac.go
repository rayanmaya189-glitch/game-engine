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

// RBACMiddleware provides permission-based access control
type RBACMiddleware struct {
	redis         *redis.Client
	routePermMap  []RoutePermission
	permissionTTL time.Duration
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

		requiredPerm := m.findRequiredPermission(method, path)
		if requiredPerm == "" {
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
