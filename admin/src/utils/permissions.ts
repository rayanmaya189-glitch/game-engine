// Permission constants - synced 1:1 with backend rbac/permissions.go GetAllPermissions()
export const Permissions = {
  // Players
  PLAYERS_VIEW: 'players:view',
  PLAYERS_EDIT: 'players:edit',
  PLAYERS_BAN: 'players:ban',
  PLAYERS_DELETE: 'players:delete',

  // KYC
  KYC_VIEW: 'kyc:view',
  KYC_APPROVE: 'kyc:approve',
  KYC_REJECT: 'kyc:reject',

  // Games
  GAMES_VIEW: 'games:view',
  GAMES_CREATE: 'games:create',
  GAMES_EDIT: 'games:edit',
  GAMES_DELETE: 'games:delete',

  // Wallet
  WALLET_VIEW: 'wallet:view',
  WALLET_ADJUST: 'wallet:adjust',
  WALLET_REVERSE: 'wallet:reverse',

  // Claims (core)
  CLAIMS_VIEW: 'claims:view',
  CLAIMS_APPROVE: 'claims:approve',
  CLAIMS_REJECT: 'claims:reject',
  CLAIMS_PAY: 'claims:pay',

  // Commission
  COMMISSION_VIEW: 'commission:view',
  COMMISSION_APPROVE: 'commission:approve',
  COMMISSION_REJECT: 'commission:reject',
  COMMISSION_PAY: 'commission:pay',

  // Rebet
  REBET_VIEW: 'rebet:view',
  REBET_APPROVE: 'rebet:approve',
  REBET_REJECT: 'rebet:reject',

  // Insurance
  INSURANCE_VIEW: 'insurance:view',
  INSURANCE_APPROVE: 'insurance:approve',
  INSURANCE_REJECT: 'insurance:reject',
  INSURANCE_PAY: 'insurance:pay',

  // Settlements
  SETTLEMENTS_VIEW: 'settlements:view',

  // Merchants
  MERCHANTS_VIEW: 'merchants:view',
  MERCHANTS_CREATE: 'merchants:create',
  MERCHANTS_EDIT: 'merchants:edit',
  MERCHANTS_DELETE: 'merchants:delete',

  // Agents
  AGENTS_VIEW: 'agents:view',
  AGENTS_CREATE: 'agents:create',
  AGENTS_EDIT: 'agents:edit',
  AGENTS_DELETE: 'agents:delete',

  // Tournaments
  TOURNAMENTS_VIEW: 'tournaments:view',
  TOURNAMENTS_CREATE: 'tournaments:create',
  TOURNAMENTS_EDIT: 'tournaments:edit',
  TOURNAMENTS_DELETE: 'tournaments:delete',

  // Jackpots
  JACKPOTS_VIEW: 'jackpots:view',
  JACKPOTS_CREATE: 'jackpots:create',
  JACKPOTS_EDIT: 'jackpots:edit',
  JACKPOTS_DELETE: 'jackpots:delete',

  // Bonuses
  BONUSES_VIEW: 'bonuses:view',
  BONUSES_CREATE: 'bonuses:create',
  BONUSES_EDIT: 'bonuses:edit',
  BONUSES_DELETE: 'bonuses:delete',

  // Payments
  PAYMENTS_VIEW: 'payments:view',
  PAYMENTS_APPROVE: 'payments:approve',
  PAYMENTS_REJECT: 'payments:reject',
  PAYMENTS_PROCESS: 'payments:process',

  // Reports
  REPORTS_VIEW: 'reports:view',
  REPORTS_EXPORT: 'reports:export',

  // Settings
  SETTINGS_VIEW: 'settings:view',
  SETTINGS_EDIT: 'settings:edit',

  // Admin Users
  ADMIN_USERS_VIEW: 'admin_users:view',
  ADMIN_USERS_CREATE: 'admin_users:create',
  ADMIN_USERS_EDIT: 'admin_users:edit',
  ADMIN_USERS_DELETE: 'admin_users:delete',

  // Roles
  ROLES_VIEW: 'roles:view',
  ROLES_CREATE: 'roles:create',
  ROLES_EDIT: 'roles:edit',
  ROLES_DELETE: 'roles:delete',

  // Audit
  AUDIT_VIEW: 'audit:view',

  // Banners
  BANNERS_VIEW: 'banners:view',
  BANNERS_CREATE: 'banners:create',
  BANNERS_EDIT: 'banners:edit',
  BANNERS_DELETE: 'banners:delete',

  // Referrals
  REFERRALS_VIEW: 'referrals:view',
  REFERRALS_EDIT: 'referrals:edit',

  // Live Dealer
  LIVE_DEALER_VIEW: 'live_dealer:view',
  LIVE_DEALER_MANAGE: 'live_dealer:manage',

  // Chat
  CHAT_VIEW: 'chat:view',
  CHAT_MODERATE: 'chat:moderate',

  // Notifications
  NOTIFICATIONS_VIEW: 'notifications:view',
  NOTIFICATIONS_SEND: 'notifications:send',
} as const;

export type Permission = typeof Permissions[keyof typeof Permissions];

// Roles that can access the admin panel
export const AdminRoles = ['superadmin', 'admin', 'support', 'finance', 'cs', 'audit', 'marketing'] as const;
export type AdminRole = typeof AdminRoles[number];

// Check if user has a specific permission
export function hasPermission(userPermissions: string[], permission: string): boolean {
  if (!userPermissions) return false;
  return userPermissions.includes(permission);
}

// Check if user has any of the given permissions
export function hasAnyPermission(userPermissions: string[], permissions: string[]): boolean {
  if (!userPermissions) return false;
  return permissions.some(p => userPermissions.includes(p));
}

// Check if user has all of the given permissions
export function hasAllPermissions(userPermissions: string[], permissions: string[]): boolean {
  if (!userPermissions) return false;
  return permissions.every(p => userPermissions.includes(p));
}

// Check if user role is superadmin (has all permissions)
export function isSuperAdmin(role: string): boolean {
  return role === 'superadmin';
}

// Check if user can access admin panel
export function canAccessAdmin(role: string): boolean {
  return (AdminRoles as readonly string[]).includes(role);
}

// Get all permission values as array
export function getAllPermissions(): string[] {
  return Object.values(Permissions);
}

// Menu item configuration with required permissions
export interface MenuItemConfig {
  text: string;
  icon: string;
  path: string;
  permission?: string;
  permissions?: string[]; // requires ANY of these
}

// Menu items with their required permissions (synced with backend routes)
export const menuItemsConfig: MenuItemConfig[] = [
  { text: 'Dashboard', icon: 'Dashboard', path: '/' },
  { text: 'Claims Management', icon: 'Assignment', path: '/claims', permission: Permissions.CLAIMS_VIEW },
  { text: 'Users', icon: 'People', path: '/users', permission: Permissions.PLAYERS_VIEW },
  { text: 'Merchants', icon: 'Business', path: '/merchants', permission: Permissions.MERCHANTS_VIEW },
  { text: 'Agents', icon: 'SupervisorAccount', path: '/agents', permission: Permissions.AGENTS_VIEW },
  { text: 'Games', icon: 'Games', path: '/games', permission: Permissions.GAMES_VIEW },
  { text: 'Tournaments', icon: 'EmojiEvents', path: '/tournaments', permission: Permissions.TOURNAMENTS_VIEW },
  { text: 'Jackpots', icon: 'EmojiEvents', path: '/jackpots', permission: Permissions.JACKPOTS_VIEW },
  { text: 'Bonuses', icon: 'CardGiftcard', path: '/bonuses', permission: Permissions.BONUSES_VIEW },
  { text: 'Payments', icon: 'AccountBalance', path: '/payments', permission: Permissions.PAYMENTS_VIEW },
  { text: 'Reports', icon: 'Assessment', path: '/reports', permission: Permissions.REPORTS_VIEW },
  { text: 'Settings', icon: 'Settings', path: '/settings', permission: Permissions.SETTINGS_VIEW },
  { text: 'KYC', icon: 'VerifiedUser', path: '/kyc', permission: Permissions.KYC_VIEW },
  { text: 'Banners', icon: 'Image', path: '/banners', permission: Permissions.BANNERS_VIEW },
  { text: 'Referrals', icon: 'Share', path: '/referrals', permission: Permissions.REFERRALS_VIEW },
  { text: 'Live Dealer', icon: 'Videocam', path: '/live-dealer', permission: Permissions.LIVE_DEALER_VIEW },
  { text: 'Chat Moderation', icon: 'Chat', path: '/chat', permission: Permissions.CHAT_VIEW },
  { text: 'Notifications', icon: 'Notifications', path: '/notifications', permission: Permissions.NOTIFICATIONS_VIEW },
];

// Filter menu items based on user permissions
export function getFilteredMenuItems(userRole: string, userPermissions: string[]): MenuItemConfig[] {
  // Superadmin sees everything
  if (isSuperAdmin(userRole)) {
    return menuItemsConfig;
  }

  return menuItemsConfig.filter(item => {
    if (!item.permission) return true; // No permission required
    return hasPermission(userPermissions, item.permission);
  });
}

// Re-export permission groups
export { permissionGroups } from './permissionGroups';
export type { PermissionGroup, PermissionItem } from './permissionGroups';

// Re-export role definitions
export { roleDefinitions } from './roleDefinitions';
export type { RoleDefinition } from './roleDefinitions';
