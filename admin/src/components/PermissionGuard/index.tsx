import React from 'react';
import { Navigate } from 'react-router-dom';
import { useAppSelector } from '../store/hooks';
import { hasPermission, hasAnyPermission, isSuperAdmin } from '../utils/permissions';

interface PermissionGuardProps {
  children: React.ReactNode;
  permission?: string;
  permissions?: string[]; // requires ANY of these
  requireAll?: boolean; // if true, requires ALL permissions
  fallback?: React.ReactNode;
}

// Component that conditionally renders children based on user permissions
export const PermissionGuard: React.FC<PermissionGuardProps> = ({
  children,
  permission,
  permissions,
  requireAll = false,
  fallback = null,
}) => {
  const { user } = useAppSelector((state: any) => state.auth);

  if (!user) return <>{fallback}</>;

  // Superadmin always has access
  if (isSuperAdmin(user.role)) return <>{children}</>;

  const userPermissions: string[] = user.permissions || [];

  if (permission) {
    if (!hasPermission(userPermissions, permission)) return <>{fallback}</>;
  }

  if (permissions && permissions.length > 0) {
    if (requireAll) {
      if (!permissions.every(p => userPermissions.includes(p))) return <>{fallback}</>;
    } else {
      if (!hasAnyPermission(userPermissions, permissions)) return <>{fallback}</>;
    }
  }

  return <>{children}</>;
};

// Route-level guard that redirects to dashboard if no permission
interface ProtectedRouteProps {
  children: React.ReactNode;
  permission?: string;
  permissions?: string[];
}

export const ProtectedRoute: React.FC<ProtectedRouteProps> = ({
  children,
  permission,
  permissions,
}) => {
  const { user, isAuthenticated } = useAppSelector((state: any) => state.auth);

  if (!isAuthenticated) return <Navigate to="/login" />;

  if (!user) return <Navigate to="/" />;

  // Superadmin always has access
  if (isSuperAdmin(user.role)) return <>{children}</>;

  const userPermissions: string[] = user.permissions || [];

  if (permission && !hasPermission(userPermissions, permission)) {
    return <Navigate to="/" />;
  }

  if (permissions && permissions.length > 0) {
    if (!hasAnyPermission(userPermissions, permissions)) {
      return <Navigate to="/" />;
    }
  }

  return <>{children}</>;
};

// Hook for checking permissions in components
export const usePermissions = () => {
  const { user } = useAppSelector((state: any) => state.auth);

  const userPermissions: string[] = user?.permissions || [];
  const userRole: string = user?.role || '';

  return {
    hasPermission: (perm: string) => isSuperAdmin(userRole) || hasPermission(userPermissions, perm),
    hasAnyPermission: (perms: string[]) => isSuperAdmin(userRole) || hasAnyPermission(userPermissions, perms),
    isSuperAdmin: isSuperAdmin(userRole),
    isAdmin: ['superadmin', 'admin'].includes(userRole),
    role: userRole,
    permissions: userPermissions,
  };
};
