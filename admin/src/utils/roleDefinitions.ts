// Role definitions with their assigned permissions
// Synced with backend rbac/roles.go GetRolePermissions()

export interface RoleDefinition {
  key: string;
  name: string;
  description: string;
  isAdmin: boolean;
  permissions: string[] | 'all';
}

export const roleDefinitions: RoleDefinition[] = [
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
