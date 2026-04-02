// Permission groups for the permissions management page
// Synced with backend rbac/permissions.go GetPermissionInfo() - must match 67 total permissions

export interface PermissionItem {
  key: string;
  name: string;
  description: string;
}

export interface PermissionGroup {
  name: string;
  permissions: PermissionItem[];
}

export const permissionGroups: PermissionGroup[] = [
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
  {
    name: 'Live Dealer',
    permissions: [
      { key: 'live_dealer:view', name: 'View Live Dealer', description: 'View live dealer tables' },
      { key: 'live_dealer:manage', name: 'Manage Live Dealer', description: 'Create tables and control sessions' },
    ],
  },
  {
    name: 'Chat',
    permissions: [
      { key: 'chat:view', name: 'View Chat', description: 'View chat rooms and messages' },
      { key: 'chat:moderate', name: 'Moderate Chat', description: 'Delete messages and ban users' },
    ],
  },
  {
    name: 'Notifications',
    permissions: [
      { key: 'notifications:view', name: 'View Notifications', description: 'View notification history' },
      { key: 'notifications:send', name: 'Send Notifications', description: 'Create and send notifications' },
    ],
  },
];
