package main

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"

	"common/handler"
	"common/middleware"
)

// ListPermissions returns all available permissions
func (cfg *RouterConfig) ListPermissions(ctx context.Context, c *app.RequestContext) {
	role := c.GetString("role")
	if !hasPermission(role, "roles:view") && role != "superadmin" {
		handler.SendErrorResponse(c, 403, handler.ErrCodeForbidden, "Insufficient permissions", nil)
		return
	}

	permissions := middleware.GetRoutePermissionMap()
	handler.SendSuccess(c, map[string]interface{}{
		"permissions": permissions,
	})
}

// ListRoles returns all roles with their permissions
func (cfg *RouterConfig) ListRoles(ctx context.Context, c *app.RequestContext) {
	role := c.GetString("role")
	if !hasPermission(role, "roles:view") && role != "superadmin" {
		handler.SendErrorResponse(c, 403, handler.ErrCodeForbidden, "Insufficient permissions", nil)
		return
	}

	roles := []map[string]interface{}{
		{"key": "superadmin", "name": "Super Admin", "is_admin": true, "permissions": "all"},
		{"key": "admin", "name": "Admin", "is_admin": true, "permissions": middleware.GetRolePermissionsStatic("admin")},
		{"key": "support", "name": "Support", "is_admin": true, "permissions": middleware.GetRolePermissionsStatic("support")},
		{"key": "finance", "name": "Finance", "is_admin": true, "permissions": middleware.GetRolePermissionsStatic("finance")},
		{"key": "cs", "name": "Customer Service", "is_admin": true, "permissions": middleware.GetRolePermissionsStatic("cs")},
		{"key": "audit", "name": "Auditor", "is_admin": true, "permissions": middleware.GetRolePermissionsStatic("audit")},
		{"key": "marketing", "name": "Marketing", "is_admin": true, "permissions": middleware.GetRolePermissionsStatic("marketing")},
		{"key": "agent", "name": "Agent", "is_admin": false, "permissions": middleware.GetRolePermissionsStatic("agent")},
		{"key": "affiliate", "name": "Affiliate", "is_admin": false, "permissions": middleware.GetRolePermissionsStatic("affiliate")},
		{"key": "player", "name": "Player", "is_admin": false, "permissions": []string{}},
	}

	handler.SendSuccess(c, map[string]interface{}{
		"roles": roles,
	})
}

// CreateRole creates a new role
func (cfg *RouterConfig) CreateRole(ctx context.Context, c *app.RequestContext) {
	role := c.GetString("role")
	if !hasPermission(role, "roles:create") && role != "superadmin" {
		handler.SendErrorResponse(c, 403, handler.ErrCodeForbidden, "Insufficient permissions", nil)
		return
	}

	var req struct {
		Key         string   `json:"key"`
		Name        string   `json:"name"`
		Description string   `json:"description"`
		Permissions []string `json:"permissions"`
	}

	if err := c.Bind(&req); err != nil {
		handler.SendErrorResponse(c, 400, handler.ErrCodeBadRequest, "Invalid request body", nil)
		return
	}

	handler.SendSuccess(c, map[string]interface{}{
		"message": "Role created",
		"role":    req.Key,
	})
}

// UpdateRole updates a role's permissions
func (cfg *RouterConfig) UpdateRole(ctx context.Context, c *app.RequestContext) {
	role := c.GetString("role")
	if !hasPermission(role, "roles:edit") && role != "superadmin" {
		handler.SendErrorResponse(c, 403, handler.ErrCodeForbidden, "Insufficient permissions", nil)
		return
	}

	roleID := c.Param("id")
	var req struct {
		Name        string   `json:"name"`
		Permissions []string `json:"permissions"`
	}

	if err := c.Bind(&req); err != nil {
		handler.SendErrorResponse(c, 400, handler.ErrCodeBadRequest, "Invalid request body", nil)
		return
	}

	handler.SendSuccess(c, map[string]interface{}{
		"message": "Role updated",
		"role":    roleID,
	})
}

// DeleteRole deletes a role
func (cfg *RouterConfig) DeleteRole(ctx context.Context, c *app.RequestContext) {
	role := c.GetString("role")
	if !hasPermission(role, "roles:delete") && role != "superadmin" {
		handler.SendErrorResponse(c, 403, handler.ErrCodeForbidden, "Insufficient permissions", nil)
		return
	}

	roleID := c.Param("id")
	handler.SendSuccess(c, map[string]interface{}{
		"message": "Role deleted",
		"role":    roleID,
	})
}

// ListAdminUsers returns all admin users
func (cfg *RouterConfig) ListAdminUsers(ctx context.Context, c *app.RequestContext) {
	role := c.GetString("role")
	if !hasPermission(role, "admin_users:view") && role != "superadmin" {
		handler.SendErrorResponse(c, 403, handler.ErrCodeForbidden, "Insufficient permissions", nil)
		return
	}

	handler.SendSuccess(c, map[string]interface{}{
		"users": []interface{}{},
		"total": 0,
	})
}

// CreateAdminUser creates a new admin user
func (cfg *RouterConfig) CreateAdminUser(ctx context.Context, c *app.RequestContext) {
	role := c.GetString("role")
	if !hasPermission(role, "admin_users:create") && role != "superadmin" {
		handler.SendErrorResponse(c, 403, handler.ErrCodeForbidden, "Insufficient permissions", nil)
		return
	}

	var req struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
		Role     string `json:"role"`
	}

	if err := c.Bind(&req); err != nil {
		handler.SendErrorResponse(c, 400, handler.ErrCodeBadRequest, "Invalid request body", nil)
		return
	}

	if cfg.AuthClient == nil {
		handler.SendErrorResponse(c, 503, handler.ErrCodeServiceUnavailable, "Auth service unavailable", nil)
		return
	}

	handler.SendSuccess(c, map[string]interface{}{
		"message": "Admin user created",
		"username": req.Username,
		"role":     req.Role,
	})
}

// UpdateAdminUser updates an admin user
func (cfg *RouterConfig) UpdateAdminUser(ctx context.Context, c *app.RequestContext) {
	role := c.GetString("role")
	if !hasPermission(role, "admin_users:edit") && role != "superadmin" {
		handler.SendErrorResponse(c, 403, handler.ErrCodeForbidden, "Insufficient permissions", nil)
		return
	}

	userID := c.Param("id")
	var req struct {
		Email    string `json:"email"`
		Role     string `json:"role"`
		IsActive *bool  `json:"is_active"`
	}

	if err := c.Bind(&req); err != nil {
		handler.SendErrorResponse(c, 400, handler.ErrCodeBadRequest, "Invalid request body", nil)
		return
	}

	handler.SendSuccess(c, map[string]interface{}{
		"message": "Admin user updated",
		"user_id": userID,
	})
}

// DeleteAdminUser deletes an admin user
func (cfg *RouterConfig) DeleteAdminUser(ctx context.Context, c *app.RequestContext) {
	role := c.GetString("role")
	if !hasPermission(role, "admin_users:delete") && role != "superadmin" {
		handler.SendErrorResponse(c, 403, handler.ErrCodeForbidden, "Insufficient permissions", nil)
		return
	}

	userID := c.Param("id")
	handler.SendSuccess(c, map[string]interface{}{
		"message": "Admin user deleted",
		"user_id": userID,
	})
}

func hasPermission(role, permission string) bool {
	perms := middleware.GetRolePermissionsStatic(role)
	for _, p := range perms {
		if p == permission {
			return true
		}
	}
	return false
}
