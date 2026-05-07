package rbac

import (
	"encoding/json"
	"fmt"
	"log"
)

// Seeder handles permission and role seeding
type Seeder struct {
	store PermissionStore
}

// PermissionStore interface for persisting RBAC data
type PermissionStore interface {
	// Permissions
	CreatePermission(key, name, description, group string) error
	GetAllPermissions() ([]StoredPermission, error)

	// Roles
	CreateRole(key, name, description string, isAdmin bool) error
	GetAllRoles() ([]StoredRole, error)

	// Role-Permission mapping
	AssignPermissionToRole(roleKey, permissionKey string) error
	GetRolePermissions(roleKey) ([]string, error)

	// User-Role mapping
	AssignRoleToUser(userID, roleKey string) error
	GetUserRoles(userID string) ([]string, error)
	GetUserPermissions(userID string) ([]string, error)

	// Admin user creation
	CreateAdminUser(username, email, hashedPassword, roleKey string) error
	GetAdminUserByUsername(username string) (*StoredAdminUser, error)
}

type StoredPermission struct {
	Key         string `json:"key"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Group       string `json:"group"`
}

type StoredRole struct {
	Key         string `json:"key"`
	Name        string `json:"name"`
	Description string `json:"description"`
	IsAdmin     bool   `json:"is_admin"`
}

type StoredAdminUser struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}

// NewSeeder creates a new seeder instance
func NewSeeder(store PermissionStore) *Seeder {
	return &Seeder{store: store}
}

// SeedPermissions creates all permission records
func (s *Seeder) SeedPermissions() error {
	log.Println("Seeding permissions...")

	existingPerms, err := s.store.GetAllPermissions()
	if err != nil {
		return fmt.Errorf("failed to get existing permissions: %w", err)
	}

	existingMap := make(map[string]bool)
	for _, p := range existingPerms {
		existingMap[p.Key] = true
	}

	permissions := GetPermissionInfo()
	created := 0

	for _, perm := range permissions {
		if existingMap[perm.Key] {
			continue
		}
		if err := s.store.CreatePermission(perm.Key, perm.Name, perm.Description, perm.Group); err != nil {
			return fmt.Errorf("failed to create permission %s: %w", perm.Key, err)
		}
		created++
	}

	log.Printf("Permissions seeded: %d created, %d total", created, len(permissions))
	return nil
}

// SeedRoles creates all role records
func (s *Seeder) SeedRoles() error {
	log.Println("Seeding roles...")

	existingRoles, err := s.store.GetAllRoles()
	if err != nil {
		return fmt.Errorf("failed to get existing roles: %w", err)
	}

	existingMap := make(map[string]bool)
	for _, r := range existingRoles {
		existingMap[r.Key] = true
	}

	roles := GetRoleInfo()
	created := 0

	for _, role := range roles {
		if existingMap[role.Key] {
			continue
		}
		if err := s.store.CreateRole(role.Key, role.Name, role.Description, role.IsAdmin); err != nil {
			return fmt.Errorf("failed to create role %s: %w", role.Key, err)
		}
		created++
	}

	log.Printf("Roles seeded: %d created, %d total", created, len(roles))
	return nil
}

// SeedRolePermissions assigns permissions to roles
func (s *Seeder) SeedRolePermissions() error {
	log.Println("Seeding role-permission mappings...")

	allPermissions := GetAllPermissions()

	// Assign ALL permissions to superadmin
	for _, perm := range allPermissions {
		if err := s.store.AssignPermissionToRole(RoleSuperAdmin, perm); err != nil {
			return fmt.Errorf("failed to assign permission %s to superadmin: %w", perm, err)
		}
	}

	// Assign permissions to other roles
	rolePerms := GetRolePermissions()
	for roleKey, perms := range rolePerms {
		for _, perm := range perms {
			if err := s.store.AssignPermissionToRole(roleKey, perm); err != nil {
				return fmt.Errorf("failed to assign permission %s to role %s: %w", perm, roleKey, err)
			}
		}
	}

	log.Printf("Role-permission mappings seeded: superadmin has %d permissions", len(allPermissions))
	return nil
}

// SeedDefaultAdmin creates a default superadmin user
func (s *Seeder) SeedDefaultAdmin(username, email, hashedPassword string) error {
	log.Println("Seeding default superadmin user...")

	existing, err := s.store.GetAdminUserByUsername(username)
	if err == nil && existing != nil {
		log.Printf("Default admin user '%s' already exists", username)
		return nil
	}

	if err := s.store.CreateAdminUser(username, email, hashedPassword, RoleSuperAdmin); err != nil {
		return fmt.Errorf("failed to create default admin user: %w", err)
	}

	log.Printf("Default superadmin user '%s' created", username)
	return nil
}

// SeedAll runs all seeders in order
func (s *Seeder) SeedAll(defaultAdminUser, defaultAdminEmail, defaultAdminHashedPassword string) error {
	if err := s.SeedPermissions(); err != nil {
		return fmt.Errorf("permissions seeding failed: %w", err)
	}

	if err := s.SeedRoles(); err != nil {
		return fmt.Errorf("roles seeding failed: %w", err)
	}

	if err := s.SeedRolePermissions(); err != nil {
		return fmt.Errorf("role-permission seeding failed: %w", err)
	}

	if err := s.SeedDefaultAdmin(defaultAdminUser, defaultAdminEmail, defaultAdminHashedPassword); err != nil {
		return fmt.Errorf("default admin seeding failed: %w", err)
	}

	log.Println("RBAC seeding completed successfully")
	return nil
}

// GetPermissionsJSON returns all permissions as JSON for API responses
func GetPermissionsJSON() ([]byte, error) {
	permissions := GetPermissionInfo()

	grouped := make(map[string][]PermissionInfo)
	for _, p := range permissions {
		grouped[p.Group] = append(grouped[p.Group], p)
	}

	return json.Marshal(grouped)
}

// GetRolesJSON returns all roles with their permissions as JSON
func GetRolesJSON(store PermissionStore) ([]byte, error) {
	roles, err := store.GetAllRoles()
	if err != nil {
		return nil, err
	}

	type RoleWithPermissions struct {
		Role        StoredRole `json:"role"`
		Permissions []string   `json:"permissions"`
	}

	result := make([]RoleWithPermissions, 0, len(roles))
	for _, role := range roles {
		perms, err := store.GetRolePermissions(role.Key)
		if err != nil {
			perms = []string{}
		}
		result = append(result, RoleWithPermissions{
			Role:        role,
			Permissions: perms,
		})
	}

	return json.Marshal(result)
}
