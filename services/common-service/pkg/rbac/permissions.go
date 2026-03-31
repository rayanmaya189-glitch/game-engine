package rbac

// Permission constants for the admin system
// Format: RESOURCE:ACTION
const (
	// Player Management
	PermPlayersView   = "players:view"
	PermPlayersEdit   = "players:edit"
	PermPlayersBan    = "players:ban"
	PermPlayersDelete = "players:delete"

	// KYC Management
	PermKYCView   = "kyc:view"
	PermKYCApprove = "kyc:approve"
	PermKYCReject  = "kyc:reject"

	// Game Management
	PermGamesView   = "games:view"
	PermGamesCreate = "games:create"
	PermGamesEdit   = "games:edit"
	PermGamesDelete = "games:delete"

	// Wallet Management
	PermWalletView    = "wallet:view"
	PermWalletAdjust  = "wallet:adjust"
	PermWalletReverse = "wallet:reverse"

	// Claims Management
	PermClaimsView        = "claims:view"
	PermClaimsApprove     = "claims:approve"
	PermClaimsReject      = "claims:reject"
	PermClaimsPay         = "claims:pay"
	PermCommissionView    = "commission:view"
	PermCommissionApprove = "commission:approve"
	PermCommissionReject  = "commission:reject"
	PermCommissionPay     = "commission:pay"
	PermRebetView         = "rebet:view"
	PermRebetApprove      = "rebet:approve"
	PermRebetReject       = "rebet:reject"
	PermInsuranceView     = "insurance:view"
	PermInsuranceApprove  = "insurance:approve"
	PermInsuranceReject   = "insurance:reject"
	PermInsurancePay      = "insurance:pay"
	PermSettlementsView   = "settlements:view"

	// Merchant Management
	PermMerchantsView   = "merchants:view"
	PermMerchantsCreate = "merchants:create"
	PermMerchantsEdit   = "merchants:edit"
	PermMerchantsDelete = "merchants:delete"

	// Agent Management
	PermAgentsView   = "agents:view"
	PermAgentsCreate = "agents:create"
	PermAgentsEdit   = "agents:edit"
	PermAgentsDelete = "agents:delete"

	// Tournament Management
	PermTournamentsView   = "tournaments:view"
	PermTournamentsCreate = "tournaments:create"
	PermTournamentsEdit   = "tournaments:edit"
	PermTournamentsDelete = "tournaments:delete"

	// Jackpot Management
	PermJackpotsView   = "jackpots:view"
	PermJackpotsCreate = "jackpots:create"
	PermJackpotsEdit   = "jackpots:edit"
	PermJackpotsDelete = "jackpots:delete"

	// Bonus Management
	PermBonusesView   = "bonuses:view"
	PermBonusesCreate = "bonuses:create"
	PermBonusesEdit   = "bonuses:edit"
	PermBonusesDelete = "bonuses:delete"

	// Payment Management
	PermPaymentsView    = "payments:view"
	PermPaymentsApprove = "payments:approve"
	PermPaymentsReject  = "payments:reject"
	PermPaymentsProcess = "payments:process"

	// Reports
	PermReportsView  = "reports:view"
	PermReportsExport = "reports:export"

	// Settings
	PermSettingsView = "settings:view"
	PermSettingsEdit = "settings:edit"

	// User Management (admin users)
	PermAdminUsersView   = "admin_users:view"
	PermAdminUsersCreate = "admin_users:create"
	PermAdminUsersEdit   = "admin_users:edit"
	PermAdminUsersDelete = "admin_users:delete"

	// Role Management
	PermRolesView   = "roles:view"
	PermRolesCreate = "roles:create"
	PermRolesEdit   = "roles:edit"
	PermRolesDelete = "roles:delete"

	// Audit Log
	PermAuditView = "audit:view"
)

// GetAllPermissions returns all defined permissions
func GetAllPermissions() []string {
	return []string{
		// Players
		PermPlayersView, PermPlayersEdit, PermPlayersBan, PermPlayersDelete,
		// KYC
		PermKYCView, PermKYCApprove, PermKYCReject,
		// Games
		PermGamesView, PermGamesCreate, PermGamesEdit, PermGamesDelete,
		// Wallet
		PermWalletView, PermWalletAdjust, PermWalletReverse,
		// Claims
		PermClaimsView, PermClaimsApprove, PermClaimsReject, PermClaimsPay,
		PermCommissionView, PermCommissionApprove, PermCommissionReject, PermCommissionPay,
		PermRebetView, PermRebetApprove, PermRebetReject,
		PermInsuranceView, PermInsuranceApprove, PermInsuranceReject, PermInsurancePay,
		PermSettlementsView,
		// Merchants
		PermMerchantsView, PermMerchantsCreate, PermMerchantsEdit, PermMerchantsDelete,
		// Agents
		PermAgentsView, PermAgentsCreate, PermAgentsEdit, PermAgentsDelete,
		// Tournaments
		PermTournamentsView, PermTournamentsCreate, PermTournamentsEdit, PermTournamentsDelete,
		// Jackpots
		PermJackpotsView, PermJackpotsCreate, PermJackpotsEdit, PermJackpotsDelete,
		// Bonuses
		PermBonusesView, PermBonusesCreate, PermBonusesEdit, PermBonusesDelete,
		// Payments
		PermPaymentsView, PermPaymentsApprove, PermPaymentsReject, PermPaymentsProcess,
		// Reports
		PermReportsView, PermReportsExport,
		// Settings
		PermSettingsView, PermSettingsEdit,
		// Admin Users
		PermAdminUsersView, PermAdminUsersCreate, PermAdminUsersEdit, PermAdminUsersDelete,
		// Roles
		PermRolesView, PermRolesCreate, PermRolesEdit, PermRolesDelete,
		// Audit
		PermAuditView,
	}
}

// PermissionMetadata maps permissions to human-readable names and groups
type PermissionInfo struct {
	Key         string `json:"key"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Group       string `json:"group"`
}

func GetPermissionInfo() []PermissionInfo {
	return []PermissionInfo{
		// Players
		{PermPlayersView, "View Players", "View player list and details", "Players"},
		{PermPlayersEdit, "Edit Players", "Edit player profiles and settings", "Players"},
		{PermPlayersBan, "Ban Players", "Ban or suspend player accounts", "Players"},
		{PermPlayersDelete, "Delete Players", "Delete player accounts", "Players"},

		// KYC
		{PermKYCView, "View KYC", "View KYC verification requests", "KYC"},
		{PermKYCApprove, "Approve KYC", "Approve KYC verification", "KYC"},
		{PermKYCReject, "Reject KYC", "Reject KYC verification", "KYC"},

		// Games
		{PermGamesView, "View Games", "View game list and details", "Games"},
		{PermGamesCreate, "Create Games", "Create new games", "Games"},
		{PermGamesEdit, "Edit Games", "Edit game configurations", "Games"},
		{PermGamesDelete, "Delete Games", "Delete games", "Games"},

		// Wallet
		{PermWalletView, "View Wallet", "View wallet transactions", "Wallet"},
		{PermWalletAdjust, "Adjust Balance", "Adjust player balances", "Wallet"},
		{PermWalletReverse, "Reverse Transaction", "Reverse wallet transactions", "Wallet"},

		// Claims
		{PermClaimsView, "View Claims", "View all claim types", "Claims"},
		{PermClaimsApprove, "Approve Claims", "Approve claim requests", "Claims"},
		{PermClaimsReject, "Reject Claims", "Reject claim requests", "Claims"},
		{PermClaimsPay, "Pay Claims", "Process claim payments", "Claims"},
		{PermCommissionView, "View Commission Claims", "View commission claims", "Commission"},
		{PermCommissionApprove, "Approve Commission", "Approve commission claims", "Commission"},
		{PermCommissionReject, "Reject Commission", "Reject commission claims", "Commission"},
		{PermCommissionPay, "Pay Commission", "Process commission payments", "Commission"},
		{PermRebetView, "View Rebet Claims", "View rebet claims", "Rebet"},
		{PermRebetApprove, "Approve Rebet", "Approve rebet claims", "Rebet"},
		{PermRebetReject, "Reject Rebet", "Reject rebet claims", "Rebet"},
		{PermInsuranceView, "View Insurance Claims", "View insurance claims", "Insurance"},
		{PermInsuranceApprove, "Approve Insurance", "Approve insurance claims", "Insurance"},
		{PermInsuranceReject, "Reject Insurance", "Reject insurance claims", "Insurance"},
		{PermInsurancePay, "Pay Insurance", "Process insurance payments", "Insurance"},
		{PermSettlementsView, "View Settlements", "View settlement records", "Settlements"},

		// Merchants
		{PermMerchantsView, "View Merchants", "View merchant list and details", "Merchants"},
		{PermMerchantsCreate, "Create Merchants", "Create new merchant accounts", "Merchants"},
		{PermMerchantsEdit, "Edit Merchants", "Edit merchant configurations", "Merchants"},
		{PermMerchantsDelete, "Delete Merchants", "Delete merchant accounts", "Merchants"},

		// Agents
		{PermAgentsView, "View Agents", "View agent list and details", "Agents"},
		{PermAgentsCreate, "Create Agents", "Create new agent accounts", "Agents"},
		{PermAgentsEdit, "Edit Agents", "Edit agent configurations", "Agents"},
		{PermAgentsDelete, "Delete Agents", "Delete agent accounts", "Agents"},

		// Tournaments
		{PermTournamentsView, "View Tournaments", "View tournament list", "Tournaments"},
		{PermTournamentsCreate, "Create Tournaments", "Create new tournaments", "Tournaments"},
		{PermTournamentsEdit, "Edit Tournaments", "Edit tournament settings", "Tournaments"},
		{PermTournamentsDelete, "Delete Tournaments", "Delete tournaments", "Tournaments"},

		// Jackpots
		{PermJackpotsView, "View Jackpots", "View jackpot list", "Jackpots"},
		{PermJackpotsCreate, "Create Jackpots", "Create new jackpots", "Jackpots"},
		{PermJackpotsEdit, "Edit Jackpots", "Edit jackpot settings", "Jackpots"},
		{PermJackpotsDelete, "Delete Jackpots", "Delete jackpots", "Jackpots"},

		// Bonuses
		{PermBonusesView, "View Bonuses", "View bonus list", "Bonuses"},
		{PermBonusesCreate, "Create Bonuses", "Create new bonuses", "Bonuses"},
		{PermBonusesEdit, "Edit Bonuses", "Edit bonus settings", "Bonuses"},
		{PermBonusesDelete, "Delete Bonuses", "Delete bonuses", "Bonuses"},

		// Payments
		{PermPaymentsView, "View Payments", "View payment list", "Payments"},
		{PermPaymentsApprove, "Approve Payments", "Approve payment requests", "Payments"},
		{PermPaymentsReject, "Reject Payments", "Reject payment requests", "Payments"},
		{PermPaymentsProcess, "Process Payments", "Process pending payments", "Payments"},

		// Reports
		{PermReportsView, "View Reports", "View reports and analytics", "Reports"},
		{PermReportsExport, "Export Reports", "Export reports to files", "Reports"},

		// Settings
		{PermSettingsView, "View Settings", "View system settings", "Settings"},
		{PermSettingsEdit, "Edit Settings", "Edit system settings", "Settings"},

		// Admin Users
		{PermAdminUsersView, "View Admin Users", "View admin user list", "Admin Users"},
		{PermAdminUsersCreate, "Create Admin Users", "Create new admin accounts", "Admin Users"},
		{PermAdminUsersEdit, "Edit Admin Users", "Edit admin user accounts", "Admin Users"},
		{PermAdminUsersDelete, "Delete Admin Users", "Delete admin accounts", "Admin Users"},

		// Roles
		{PermRolesView, "View Roles", "View roles and permissions", "Roles"},
		{PermRolesCreate, "Create Roles", "Create new roles", "Roles"},
		{PermRolesEdit, "Edit Roles", "Edit role permissions", "Roles"},
		{PermRolesDelete, "Delete Roles", "Delete roles", "Roles"},

		// Audit
		{PermAuditView, "View Audit Log", "View audit trail", "Audit"},
	}
}
