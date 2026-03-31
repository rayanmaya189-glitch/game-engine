// Permission constants - synced with backend rbac/permissions.go
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

  // Claims
  CLAIMS_VIEW: 'claims:view',
  CLAIMS_APPROVE: 'claims:approve',
  CLAIMS_REJECT: 'claims:reject',
  CLAIMS_PAY: 'claims:pay',
  COMMISSION_VIEW: 'commission:view',
  COMMISSION_APPROVE: 'commission:approve',
  COMMISSION_REJECT: 'commission:reject',
  COMMISSION_PAY: 'commission:pay',
  REBET_VIEW: 'rebet:view',
  REBET_APPROVE: 'rebet:approve',
  REBET_REJECT: 'rebet:reject',
  INSURANCE_VIEW: 'insurance:view',
  INSURANCE_APPROVE: 'insurance:approve',
  INSURANCE_REJECT: 'insurance:reject',
  INSURANCE_PAY: 'insurance:pay',
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
} as const;

export type Permission = typeof Permissions[keyof typeof Permissions];

// Roles that can access the admin panel
export const AdminRoles = ['superadmin', 'admin', 'support', 'finance', 'cs', 'audit', 'marketing'] as const;
export type AdminRole = typeof AdminRoles[number];

// Check if user has a specific permission
export function hasPermission(userPermissions: string[], permission: string): boolean {
  return userPermissions.includes(permission);
}

// Check if user has any of the given permissions
export function hasAnyPermission(userPermissions: string[], permissions: string[]): boolean {
  return permissions.some(p => userPermissions.includes(p));
}

// Check if user has all of the given permissions
export function hasAllPermissions(userPermissions: string[], permissions: string[]): boolean {
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
  { text: 'Claims Management', icon: 'Assignment', path: '/claims', permission: 'claims:view' },
  { text: 'Users', icon: 'People', path: '/users', permission: 'players:view' },
  { text: 'Merchants', icon: 'Business', path: '/merchants', permission: 'merchants:view' },
  { text: 'Agents', icon: 'SupervisorAccount', path: '/agents', permission: 'agents:view' },
  { text: 'Games', icon: 'Games', path: '/games', permission: 'games:view' },
  { text: 'Tournaments', icon: 'EmojiEvents', path: '/tournaments', permission: 'tournaments:view' },
  { text: 'Jackpots', icon: 'EmojiEvents', path: '/jackpots', permission: 'jackpots:view' },
  { text: 'Bonuses', icon: 'CardGiftcard', path: '/bonuses', permission: 'bonuses:view' },
  { text: 'Payments', icon: 'AccountBalance', path: '/payments', permission: 'payments:view' },
  { text: 'Reports', icon: 'Assessment', path: '/reports', permission: 'reports:view' },
  { text: 'Settings', icon: 'Settings', path: '/settings', permission: 'settings:view' },
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

// Permission groups for the permissions management page
export const permissionGroups = [
  {
    name: 'Players',
    permissions: [
      { key: 'players:view', name: 'View Players', description: 'View player list and details' },
      { key: 'players:edit', name: 'Edit Players', description: 'Edit player profiles' },
      { key: 'players:ban', name: 'Ban Players', description: 'Ban/suspend player accounts' },
      { key: 'players:delete', name: 'Delete Players', description: 'Delete player accounts' },
    ],
  },
  {
    name: 'KYC',
    permissions: [
      { key: 'kyc:view', name: 'View KYC', description: 'View KYC requests' },
      { key: 'kyc:approve', name: 'Approve KYC', description: 'Approve KYC verification' },
      { key: 'kyc:reject', name: 'Reject KYC', description: 'Reject KYC verification' },
    ],
  },
  {
    name: 'Games',
    permissions: [
      { key: 'games:view', name: 'View Games', description: 'View game list' },
      { key: 'games:create', name: 'Create Games', description: 'Create new games' },
      { key: 'games:edit', name: 'Edit Games', description: 'Edit game settings' },
      { key: 'games:delete', name: 'Delete Games', description: 'Delete games' },
    ],
  },
  {
    name: 'Wallet',
    permissions: [
      { key: 'wallet:view', name: 'View Wallet', description: 'View transactions' },
      { key: 'wallet:adjust', name: 'Adjust Balance', description: 'Adjust player balances' },
      { key: 'wallet:reverse', name: 'Reverse Transaction', description: 'Reverse transactions' },
    ],
  },
  {
    name: 'Claims',
    permissions: [
      { key: 'claims:view', name: 'View Claims', description: 'View all claims' },
      { key: 'commission:view', name: 'View Commission', description: 'View commission claims' },
      { key: 'commission:approve', name: 'Approve Commission', description: 'Approve commission' },
      { key: 'commission:pay', name: 'Pay Commission', description: 'Pay commission' },
      { key: 'rebet:view', name: 'View Rebet', description: 'View rebet claims' },
      { key: 'rebet:approve', name: 'Approve Rebet', description: 'Approve rebet' },
      { key: 'insurance:view', name: 'View Insurance', description: 'View insurance claims' },
      { key: 'insurance:approve', name: 'Approve Insurance', description: 'Approve insurance' },
      { key: 'insurance:pay', name: 'Pay Insurance', description: 'Pay insurance' },
      { key: 'settlements:view', name: 'View Settlements', description: 'View settlements' },
    ],
  },
  {
    name: 'Merchants',
    permissions: [
      { key: 'merchants:view', name: 'View Merchants', description: 'View merchant list' },
      { key: 'merchants:create', name: 'Create Merchants', description: 'Create merchants' },
      { key: 'merchants:edit', name: 'Edit Merchants', description: 'Edit merchants' },
      { key: 'merchants:delete', name: 'Delete Merchants', description: 'Delete merchants' },
    ],
  },
  {
    name: 'Agents',
    permissions: [
      { key: 'agents:view', name: 'View Agents', description: 'View agent list' },
      { key: 'agents:create', name: 'Create Agents', description: 'Create agents' },
      { key: 'agents:edit', name: 'Edit Agents', description: 'Edit agents' },
      { key: 'agents:delete', name: 'Delete Agents', description: 'Delete agents' },
    ],
  },
  {
    name: 'Tournaments',
    permissions: [
      { key: 'tournaments:view', name: 'View Tournaments', description: 'View tournaments' },
      { key: 'tournaments:create', name: 'Create Tournaments', description: 'Create tournaments' },
      { key: 'tournaments:edit', name: 'Edit Tournaments', description: 'Edit tournaments' },
      { key: 'tournaments:delete', name: 'Delete Tournaments', description: 'Delete tournaments' },
    ],
  },
  {
    name: 'Jackpots',
    permissions: [
      { key: 'jackpots:view', name: 'View Jackpots', description: 'View jackpots' },
      { key: 'jackpots:create', name: 'Create Jackpots', description: 'Create jackpots' },
      { key: 'jackpots:edit', name: 'Edit Jackpots', description: 'Edit jackpots' },
      { key: 'jackpots:delete', name: 'Delete Jackpots', description: 'Delete jackpots' },
    ],
  },
  {
    name: 'Bonuses',
    permissions: [
      { key: 'bonuses:view', name: 'View Bonuses', description: 'View bonuses' },
      { key: 'bonuses:create', name: 'Create Bonuses', description: 'Create bonuses' },
      { key: 'bonuses:edit', name: 'Edit Bonuses', description: 'Edit bonuses' },
      { key: 'bonuses:delete', name: 'Delete Bonuses', description: 'Delete bonuses' },
    ],
  },
  {
    name: 'Payments',
    permissions: [
      { key: 'payments:view', name: 'View Payments', description: 'View payments' },
      { key: 'payments:approve', name: 'Approve Payments', description: 'Approve payments' },
      { key: 'payments:reject', name: 'Reject Payments', description: 'Reject payments' },
      { key: 'payments:process', name: 'Process Payments', description: 'Process payments' },
    ],
  },
  {
    name: 'Admin',
    permissions: [
      { key: 'admin_users:view', name: 'View Admin Users', description: 'View admin users' },
      { key: 'admin_users:create', name: 'Create Admin Users', description: 'Create admin users' },
      { key: 'admin_users:edit', name: 'Edit Admin Users', description: 'Edit admin users' },
      { key: 'admin_users:delete', name: 'Delete Admin Users', description: 'Delete admin users' },
      { key: 'roles:view', name: 'View Roles', description: 'View roles' },
      { key: 'roles:create', name: 'Create Roles', description: 'Create roles' },
      { key: 'roles:edit', name: 'Edit Roles', description: 'Edit roles' },
      { key: 'roles:delete', name: 'Delete Roles', description: 'Delete roles' },
      { key: 'settings:view', name: 'View Settings', description: 'View settings' },
      { key: 'settings:edit', name: 'Edit Settings', description: 'Edit settings' },
      { key: 'audit:view', name: 'View Audit Log', description: 'View audit log' },
    ],
  },
];
