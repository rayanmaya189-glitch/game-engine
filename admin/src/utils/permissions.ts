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
// Synced with backend GetPermissionInfo() groups - must match 67 total permissions
export const permissionGroups = [
  {
    name: 'Players',
    permissions: [
      { key: 'players:view', name: 'View Players', description: 'View player list and details' },
      { key: 'players:edit', name: 'Edit Players', description: 'Edit player profiles and settings' },
      { key: 'players:ban', name: 'Ban Players', description: 'Ban or suspend player accounts' },
      { key: 'players:delete', name: 'Delete Players', description: 'Delete player accounts' },
    ],
  },
  {
    name: 'KYC',
    permissions: [
      { key: 'kyc:view', name: 'View KYC', description: 'View KYC verification requests' },
      { key: 'kyc:approve', name: 'Approve KYC', description: 'Approve KYC verification' },
      { key: 'kyc:reject', name: 'Reject KYC', description: 'Reject KYC verification' },
    ],
  },
  {
    name: 'Games',
    permissions: [
      { key: 'games:view', name: 'View Games', description: 'View game list and details' },
      { key: 'games:create', name: 'Create Games', description: 'Create new games' },
      { key: 'games:edit', name: 'Edit Games', description: 'Edit game configurations' },
      { key: 'games:delete', name: 'Delete Games', description: 'Delete games' },
    ],
  },
  {
    name: 'Wallet',
    permissions: [
      { key: 'wallet:view', name: 'View Wallet', description: 'View wallet transactions' },
      { key: 'wallet:adjust', name: 'Adjust Balance', description: 'Adjust player balances' },
      { key: 'wallet:reverse', name: 'Reverse Transaction', description: 'Reverse wallet transactions' },
    ],
  },
  {
    name: 'Claims',
    permissions: [
      { key: 'claims:view', name: 'View Claims', description: 'View all claim types' },
      { key: 'claims:approve', name: 'Approve Claims', description: 'Approve claim requests' },
      { key: 'claims:reject', name: 'Reject Claims', description: 'Reject claim requests' },
      { key: 'claims:pay', name: 'Pay Claims', description: 'Process claim payments' },
    ],
  },
  {
    name: 'Commission',
    permissions: [
      { key: 'commission:view', name: 'View Commission', description: 'View commission claims' },
      { key: 'commission:approve', name: 'Approve Commission', description: 'Approve commission claims' },
      { key: 'commission:reject', name: 'Reject Commission', description: 'Reject commission claims' },
      { key: 'commission:pay', name: 'Pay Commission', description: 'Process commission payments' },
    ],
  },
  {
    name: 'Rebet',
    permissions: [
      { key: 'rebet:view', name: 'View Rebet', description: 'View rebet claims' },
      { key: 'rebet:approve', name: 'Approve Rebet', description: 'Approve rebet claims' },
      { key: 'rebet:reject', name: 'Reject Rebet', description: 'Reject rebet claims' },
    ],
  },
  {
    name: 'Insurance',
    permissions: [
      { key: 'insurance:view', name: 'View Insurance', description: 'View insurance claims' },
      { key: 'insurance:approve', name: 'Approve Insurance', description: 'Approve insurance claims' },
      { key: 'insurance:reject', name: 'Reject Insurance', description: 'Reject insurance claims' },
      { key: 'insurance:pay', name: 'Pay Insurance', description: 'Process insurance payments' },
    ],
  },
  {
    name: 'Settlements',
    permissions: [
      { key: 'settlements:view', name: 'View Settlements', description: 'View settlement records' },
    ],
  },
  {
    name: 'Merchants',
    permissions: [
      { key: 'merchants:view', name: 'View Merchants', description: 'View merchant list and details' },
      { key: 'merchants:create', name: 'Create Merchants', description: 'Create new merchant accounts' },
      { key: 'merchants:edit', name: 'Edit Merchants', description: 'Edit merchant configurations' },
      { key: 'merchants:delete', name: 'Delete Merchants', description: 'Delete merchant accounts' },
    ],
  },
  {
    name: 'Agents',
    permissions: [
      { key: 'agents:view', name: 'View Agents', description: 'View agent list and details' },
      { key: 'agents:create', name: 'Create Agents', description: 'Create new agent accounts' },
      { key: 'agents:edit', name: 'Edit Agents', description: 'Edit agent configurations' },
      { key: 'agents:delete', name: 'Delete Agents', description: 'Delete agent accounts' },
    ],
  },
  {
    name: 'Tournaments',
    permissions: [
      { key: 'tournaments:view', name: 'View Tournaments', description: 'View tournament list' },
      { key: 'tournaments:create', name: 'Create Tournaments', description: 'Create new tournaments' },
      { key: 'tournaments:edit', name: 'Edit Tournaments', description: 'Edit tournament settings' },
      { key: 'tournaments:delete', name: 'Delete Tournaments', description: 'Delete tournaments' },
    ],
  },
  {
    name: 'Jackpots',
    permissions: [
      { key: 'jackpots:view', name: 'View Jackpots', description: 'View jackpot list' },
      { key: 'jackpots:create', name: 'Create Jackpots', description: 'Create new jackpots' },
      { key: 'jackpots:edit', name: 'Edit Jackpots', description: 'Edit jackpot settings' },
      { key: 'jackpots:delete', name: 'Delete Jackpots', description: 'Delete jackpots' },
    ],
  },
  {
    name: 'Bonuses',
    permissions: [
      { key: 'bonuses:view', name: 'View Bonuses', description: 'View bonus list' },
      { key: 'bonuses:create', name: 'Create Bonuses', description: 'Create new bonuses' },
      { key: 'bonuses:edit', name: 'Edit Bonuses', description: 'Edit bonus settings' },
      { key: 'bonuses:delete', name: 'Delete Bonuses', description: 'Delete bonuses' },
    ],
  },
  {
    name: 'Payments',
    permissions: [
      { key: 'payments:view', name: 'View Payments', description: 'View payment list' },
      { key: 'payments:approve', name: 'Approve Payments', description: 'Approve payment requests' },
      { key: 'payments:reject', name: 'Reject Payments', description: 'Reject payment requests' },
      { key: 'payments:process', name: 'Process Payments', description: 'Process pending payments' },
    ],
  },
  {
    name: 'Reports',
    permissions: [
      { key: 'reports:view', name: 'View Reports', description: 'View reports and analytics' },
      { key: 'reports:export', name: 'Export Reports', description: 'Export reports to files' },
    ],
  },
  {
    name: 'Settings',
    permissions: [
      { key: 'settings:view', name: 'View Settings', description: 'View system settings' },
      { key: 'settings:edit', name: 'Edit Settings', description: 'Edit system settings' },
    ],
  },
  {
    name: 'Admin Users',
    permissions: [
      { key: 'admin_users:view', name: 'View Admin Users', description: 'View admin user list' },
      { key: 'admin_users:create', name: 'Create Admin Users', description: 'Create new admin accounts' },
      { key: 'admin_users:edit', name: 'Edit Admin Users', description: 'Edit admin user accounts' },
      { key: 'admin_users:delete', name: 'Delete Admin Users', description: 'Delete admin accounts' },
    ],
  },
  {
    name: 'Roles',
    permissions: [
      { key: 'roles:view', name: 'View Roles', description: 'View roles and permissions' },
      { key: 'roles:create', name: 'Create Roles', description: 'Create new roles' },
      { key: 'roles:edit', name: 'Edit Roles', description: 'Edit role permissions' },
      { key: 'roles:delete', name: 'Delete Roles', description: 'Delete roles' },
    ],
  },
  {
    name: 'Audit',
    permissions: [
      { key: 'audit:view', name: 'View Audit Log', description: 'View audit trail' },
    ],
  },
];

// Role definitions with their assigned permissions
// Synced with backend rbac/roles.go GetRolePermissions()
export const roleDefinitions = [
  {
    key: 'superadmin',
    name: 'Super Admin',
    description: 'Full system access - all 67 permissions',
    isAdmin: true,
    permissions: 'all',
  },
  {
    key: 'admin',
    name: 'Admin',
    description: 'Administrative access',
    isAdmin: true,
    permissions: [
      'players:view', 'players:edit', 'players:ban',
      'kyc:view', 'kyc:approve', 'kyc:reject',
      'games:view', 'games:create', 'games:edit',
      'wallet:view', 'wallet:adjust',
      'claims:view', 'claims:approve', 'claims:reject', 'claims:pay',
      'commission:view', 'commission:approve', 'commission:reject', 'commission:pay',
      'rebet:view', 'rebet:approve', 'rebet:reject',
      'insurance:view', 'insurance:approve', 'insurance:reject', 'insurance:pay',
      'settlements:view',
      'merchants:view', 'merchants:create', 'merchants:edit',
      'agents:view', 'agents:create', 'agents:edit',
      'tournaments:view', 'tournaments:create', 'tournaments:edit',
      'jackpots:view', 'jackpots:create', 'jackpots:edit',
      'bonuses:view', 'bonuses:create', 'bonuses:edit',
      'payments:view', 'payments:approve', 'payments:reject', 'payments:process',
      'reports:view', 'reports:export',
      'settings:view',
      'audit:view',
    ],
  },
  {
    key: 'support',
    name: 'Support',
    description: 'Customer support access',
    isAdmin: true,
    permissions: [
      'players:view', 'players:ban',
      'kyc:view', 'kyc:approve', 'kyc:reject',
      'wallet:view',
      'claims:view', 'commission:view', 'rebet:view',
      'insurance:view', 'settlements:view',
      'merchants:view', 'agents:view', 'games:view',
      'tournaments:view', 'jackpots:view', 'bonuses:view',
      'payments:view', 'reports:view',
    ],
  },
  {
    key: 'finance',
    name: 'Finance',
    description: 'Financial operations access',
    isAdmin: true,
    permissions: [
      'players:view',
      'wallet:view', 'wallet:adjust', 'wallet:reverse',
      'claims:view', 'claims:approve', 'claims:reject', 'claims:pay',
      'commission:view', 'commission:approve', 'commission:reject', 'commission:pay',
      'rebet:view', 'rebet:approve', 'rebet:reject',
      'insurance:view', 'insurance:approve', 'insurance:reject', 'insurance:pay',
      'settlements:view',
      'payments:view', 'payments:approve', 'payments:reject', 'payments:process',
      'reports:view', 'reports:export',
      'merchants:view', 'agents:view',
    ],
  },
  {
    key: 'cs',
    name: 'Customer Service',
    description: 'Customer service access',
    isAdmin: true,
    permissions: [
      'players:view', 'players:edit',
      'kyc:view',
      'wallet:view',
      'claims:view', 'commission:view', 'rebet:view',
      'insurance:view', 'settlements:view',
      'games:view', 'tournaments:view', 'bonuses:view',
      'payments:view',
    ],
  },
  {
    key: 'audit',
    name: 'Auditor',
    description: 'Audit and compliance access',
    isAdmin: true,
    permissions: [
      'players:view', 'kyc:view', 'wallet:view',
      'claims:view', 'commission:view', 'rebet:view',
      'insurance:view', 'settlements:view',
      'merchants:view', 'agents:view', 'games:view',
      'tournaments:view', 'jackpots:view', 'bonuses:view',
      'payments:view',
      'reports:view', 'reports:export',
      'audit:view',
      'admin_users:view', 'roles:view',
    ],
  },
  {
    key: 'marketing',
    name: 'Marketing',
    description: 'Marketing operations access',
    isAdmin: true,
    permissions: [
      'players:view', 'games:view',
      'tournaments:view', 'tournaments:create', 'tournaments:edit',
      'jackpots:view', 'jackpots:create', 'jackpots:edit',
      'bonuses:view', 'bonuses:create', 'bonuses:edit',
      'reports:view', 'reports:export',
      'merchants:view', 'agents:view',
    ],
  },
  {
    key: 'agent',
    name: 'Agent',
    description: 'Agent portal access',
    isAdmin: false,
    permissions: ['players:view', 'commission:view', 'reports:view'],
  },
  {
    key: 'affiliate',
    name: 'Affiliate',
    description: 'Affiliate portal access',
    isAdmin: false,
    permissions: ['players:view', 'commission:view', 'reports:view'],
  },
  {
    key: 'player',
    name: 'Player',
    description: 'Player access',
    isAdmin: false,
    permissions: [],
  },
];
