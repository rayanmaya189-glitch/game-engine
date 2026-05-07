package rbac

// Role constants
const (
	RoleSuperAdmin   = "superadmin"
	RoleAdmin        = "admin"
	RoleSupport      = "support"
	RoleFinance      = "finance"
	RoleCS           = "cs" // Customer Service
	RoleAudit        = "audit"
	RoleMarketing    = "marketing"
	RoleAgent        = "agent"
	RoleAffiliate    = "affiliate"
	RolePlayer       = "player"
)

// RoleInfo contains role metadata
type RoleInfo struct {
	Key         string `json:"key"`
	Name        string `json:"name"`
	Description string `json:"description"`
	IsAdmin     bool   `json:"is_admin"`
}

// GetRoleInfo returns all role definitions
func GetRoleInfo() []RoleInfo {
	return []RoleInfo{
		{RoleSuperAdmin, "Super Admin", "Full system access", true},
		{RoleAdmin, "Admin", "Administrative access", true},
		{RoleSupport, "Support", "Customer support access", true},
		{RoleFinance, "Finance", "Financial operations access", true},
		{RoleCS, "Customer Service", "Customer service access", true},
		{RoleAudit, "Auditor", "Audit and compliance access", true},
		{RoleMarketing, "Marketing", "Marketing operations access", true},
		{RoleAgent, "Agent", "Agent portal access", false},
		{RoleAffiliate, "Affiliate", "Affiliate portal access", false},
		{RolePlayer, "Player", "Player access", false},
	}
}

// GetAdminRoles returns roles that can access the admin panel
func GetAdminRoles() []string {
	return []string{
		RoleSuperAdmin,
		RoleAdmin,
		RoleSupport,
		RoleFinance,
		RoleCS,
		RoleAudit,
		RoleMarketing,
	}
}

// GetRolePermissions returns the permissions for each role
// SuperAdmin gets ALL permissions automatically (handled in seeder)
func GetRolePermissions() map[string][]string {
	return map[string][]string{
		RoleAdmin: {
			// Players
			PermPlayersView, PermPlayersEdit, PermPlayersBan,
			// KYC
			PermKYCView, PermKYCApprove, PermKYCReject,
			// Games
			PermGamesView, PermGamesCreate, PermGamesEdit,
			// Wallet
			PermWalletView, PermWalletAdjust,
			// Claims
			PermClaimsView, PermClaimsApprove, PermClaimsReject, PermClaimsPay,
			PermCommissionView, PermCommissionApprove, PermCommissionReject, PermCommissionPay,
			PermRebetView, PermRebetApprove, PermRebetReject,
			PermInsuranceView, PermInsuranceApprove, PermInsuranceReject, PermInsurancePay,
			PermSettlementsView,
			// Merchants
			PermMerchantsView, PermMerchantsCreate, PermMerchantsEdit,
			// Agents
			PermAgentsView, PermAgentsCreate, PermAgentsEdit,
			// Tournaments
			PermTournamentsView, PermTournamentsCreate, PermTournamentsEdit,
			// Jackpots
			PermJackpotsView, PermJackpotsCreate, PermJackpotsEdit,
			// Bonuses
			PermBonusesView, PermBonusesCreate, PermBonusesEdit,
			// Payments
			PermPaymentsView, PermPaymentsApprove, PermPaymentsReject, PermPaymentsProcess,
			// Reports
			PermReportsView, PermReportsExport,
			// Settings
			PermSettingsView,
			// Audit
			PermAuditView,
		},

		RoleSupport: {
			// Players (read-only + ban)
			PermPlayersView, PermPlayersBan,
			// KYC
			PermKYCView, PermKYCApprove, PermKYCReject,
			// Wallet (view only)
			PermWalletView,
			// Claims (view only)
			PermClaimsView, PermCommissionView, PermRebetView,
			PermInsuranceView, PermSettlementsView,
			// Merchants (view only)
			PermMerchantsView,
			// Agents (view only)
			PermAgentsView,
			// Games (view only)
			PermGamesView,
			// Tournaments (view only)
			PermTournamentsView,
			// Jackpots (view only)
			PermJackpotsView,
			// Bonuses (view only)
			PermBonusesView,
			// Payments (view only)
			PermPaymentsView,
			// Reports (view only)
			PermReportsView,
		},

		RoleFinance: {
			// Players (view only)
			PermPlayersView,
			// Wallet
			PermWalletView, PermWalletAdjust, PermWalletReverse,
			// Claims
			PermClaimsView, PermClaimsApprove, PermClaimsReject, PermClaimsPay,
			PermCommissionView, PermCommissionApprove, PermCommissionReject, PermCommissionPay,
			PermRebetView, PermRebetApprove, PermRebetReject,
			PermInsuranceView, PermInsuranceApprove, PermInsuranceReject, PermInsurancePay,
			PermSettlementsView,
			// Payments
			PermPaymentsView, PermPaymentsApprove, PermPaymentsReject, PermPaymentsProcess,
			// Reports
			PermReportsView, PermReportsExport,
			// Merchants (view only)
			PermMerchantsView,
			// Agents (view only)
			PermAgentsView,
		},

		RoleCS: {
			// Players (view + edit, no ban)
			PermPlayersView, PermPlayersEdit,
			// KYC
			PermKYCView,
			// Wallet (view only)
			PermWalletView,
			// Claims (view only)
			PermClaimsView, PermCommissionView, PermRebetView,
			PermInsuranceView, PermSettlementsView,
			// Games (view only)
			PermGamesView,
			// Tournaments (view only)
			PermTournamentsView,
			// Bonuses (view only)
			PermBonusesView,
			// Payments (view only)
			PermPaymentsView,
		},

		RoleAudit: {
			// Players (view only)
			PermPlayersView,
			// KYC (view only)
			PermKYCView,
			// Wallet (view only)
			PermWalletView,
			// Claims (view only)
			PermClaimsView, PermCommissionView, PermRebetView,
			PermInsuranceView, PermSettlementsView,
			// Merchants (view only)
			PermMerchantsView,
			// Agents (view only)
			PermAgentsView,
			// Games (view only)
			PermGamesView,
			// Tournaments (view only)
			PermTournamentsView,
			// Jackpots (view only)
			PermJackpotsView,
			// Bonuses (view only)
			PermBonusesView,
			// Payments (view only)
			PermPaymentsView,
			// Reports
			PermReportsView, PermReportsExport,
			// Audit
			PermAuditView,
			// Admin Users (view only)
			PermAdminUsersView,
			// Roles (view only)
			PermRolesView,
		},

		RoleMarketing: {
			// Players (view only)
			PermPlayersView,
			// Games (view only)
			PermGamesView,
			// Tournaments
			PermTournamentsView, PermTournamentsCreate, PermTournamentsEdit,
			// Jackpots
			PermJackpotsView, PermJackpotsCreate, PermJackpotsEdit,
			// Bonuses
			PermBonusesView, PermBonusesCreate, PermBonusesEdit,
			// Reports
			PermReportsView, PermReportsExport,
			// Merchants (view only)
			PermMerchantsView,
			// Agents (view only)
			PermAgentsView,
		},

		RoleAgent: {
			// Limited to agent-specific operations
			PermPlayersView,
			PermCommissionView,
			PermReportsView,
		},

		RoleAffiliate: {
			// Limited to affiliate-specific operations
			PermPlayersView,
			PermCommissionView,
			PermReportsView,
		},

		RolePlayer: {
			// No admin permissions
		},
	}
}

// GetAllRoles returns all role keys
func GetAllRoles() []string {
	roles := make([]string, 0, len(GetRoleInfo()))
	for _, r := range GetRoleInfo() {
		roles = append(roles, r.Key)
	}
	return roles
}
