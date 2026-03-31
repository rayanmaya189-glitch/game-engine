package middleware

// RoutePermission maps HTTP method + path to required permission
type RoutePermission struct {
	Method     string
	PathPrefix string
	Permission string
}

// GetRoutePermissionMap defines which permission is required for each admin route
func GetRoutePermissionMap() []RoutePermission {
	return []RoutePermission{
		// Players
		{"GET", "/api/v1/admin/players", "players:view"},
		{"PUT", "/api/v1/admin/players/", "players:edit"},
		// KYC
		{"GET", "/api/v1/admin/kyc", "kyc:view"},
		{"PUT", "/api/v1/admin/kyc/", "kyc:approve"},
		// Games
		{"GET", "/api/v1/admin/games", "games:view"},
		{"POST", "/api/v1/admin/games", "games:create"},
		{"PUT", "/api/v1/admin/games/", "games:edit"},
		// Wallet
		{"GET", "/api/v1/admin/wallet", "wallet:view"},
		{"POST", "/api/v1/admin/wallet/adjust", "wallet:adjust"},
		// Commission Claims
		{"GET", "/api/v1/admin/claims/commission", "commission:view"},
		{"POST", "/api/v1/admin/claims/commission/", "commission:approve"},
		// Rebet Claims
		{"GET", "/api/v1/admin/claims/rebet", "rebet:view"},
		{"POST", "/api/v1/admin/claims/rebet/", "rebet:approve"},
		// Insurance Claims
		{"GET", "/api/v1/admin/claims/insurance", "insurance:view"},
		{"POST", "/api/v1/admin/claims/insurance/", "insurance:approve"},
		// Settlements
		{"GET", "/api/v1/admin/claims/settlements", "settlements:view"},
		{"GET", "/api/v1/admin/claims/statistics", "claims:view"},
		// Merchants
		{"GET", "/api/v1/admin/merchants", "merchants:view"},
		{"POST", "/api/v1/admin/merchants", "merchants:create"},
		{"PUT", "/api/v1/admin/merchants/", "merchants:edit"},
		{"DELETE", "/api/v1/admin/merchants/", "merchants:delete"},
		// Agents
		{"GET", "/api/v1/admin/agents", "agents:view"},
		{"POST", "/api/v1/admin/agents", "agents:create"},
		{"PUT", "/api/v1/admin/agents/", "agents:edit"},
		{"DELETE", "/api/v1/admin/agents/", "agents:delete"},
		// Tournaments
		{"GET", "/api/v1/admin/tournaments", "tournaments:view"},
		{"POST", "/api/v1/admin/tournaments", "tournaments:create"},
		{"PUT", "/api/v1/admin/tournaments/", "tournaments:edit"},
		{"DELETE", "/api/v1/admin/tournaments/", "tournaments:delete"},
		// Jackpots
		{"GET", "/api/v1/admin/jackpots", "jackpots:view"},
		{"POST", "/api/v1/admin/jackpots", "jackpots:create"},
		{"PUT", "/api/v1/admin/jackpots/", "jackpots:edit"},
		{"DELETE", "/api/v1/admin/jackpots/", "jackpots:delete"},
		// Bonuses
		{"GET", "/api/v1/admin/bonuses", "bonuses:view"},
		{"POST", "/api/v1/admin/bonuses", "bonuses:create"},
		{"PUT", "/api/v1/admin/bonuses/", "bonuses:edit"},
		{"DELETE", "/api/v1/admin/bonuses/", "bonuses:delete"},
		// Payments
		{"GET", "/api/v1/admin/payments", "payments:view"},
		{"PUT", "/api/v1/admin/payments/", "payments:approve"},
		// Reports
		{"GET", "/api/v1/admin/reports", "reports:view"},
	}
}

// GetRolePermissionsStatic returns permissions for a role (static mapping)
func GetRolePermissionsStatic(role string) []string {
	rolePerms := map[string][]string{
		"admin": {
			"players:view", "players:edit", "players:ban",
			"kyc:view", "kyc:approve", "kyc:reject",
			"games:view", "games:create", "games:edit",
			"wallet:view", "wallet:adjust",
			"claims:view", "claims:approve", "claims:reject", "claims:pay",
			"commission:view", "commission:approve", "commission:reject", "commission:pay",
			"rebet:view", "rebet:approve", "rebet:reject",
			"insurance:view", "insurance:approve", "insurance:reject", "insurance:pay",
			"settlements:view",
			"merchants:view", "merchants:create", "merchants:edit",
			"agents:view", "agents:create", "agents:edit",
			"tournaments:view", "tournaments:create", "tournaments:edit",
			"jackpots:view", "jackpots:create", "jackpots:edit",
			"bonuses:view", "bonuses:create", "bonuses:edit",
			"payments:view", "payments:approve", "payments:reject", "payments:process",
			"reports:view", "reports:export",
			"settings:view",
			"audit:view",
		},
		"support": {
			"players:view", "players:ban",
			"kyc:view", "kyc:approve", "kyc:reject",
			"wallet:view",
			"claims:view", "commission:view", "rebet:view",
			"insurance:view", "settlements:view",
			"merchants:view", "agents:view", "games:view",
			"tournaments:view", "jackpots:view", "bonuses:view",
			"payments:view", "reports:view",
		},
		"finance": {
			"players:view",
			"wallet:view", "wallet:adjust", "wallet:reverse",
			"claims:view", "claims:approve", "claims:reject", "claims:pay",
			"commission:view", "commission:approve", "commission:reject", "commission:pay",
			"rebet:view", "rebet:approve", "rebet:reject",
			"insurance:view", "insurance:approve", "insurance:reject", "insurance:pay",
			"settlements:view",
			"payments:view", "payments:approve", "payments:reject", "payments:process",
			"reports:view", "reports:export",
			"merchants:view", "agents:view",
		},
		"cs": {
			"players:view", "players:edit",
			"kyc:view",
			"wallet:view",
			"claims:view", "commission:view", "rebet:view",
			"insurance:view", "settlements:view",
			"games:view", "tournaments:view", "bonuses:view",
			"payments:view",
		},
		"audit": {
			"players:view", "kyc:view", "wallet:view",
			"claims:view", "commission:view", "rebet:view",
			"insurance:view", "settlements:view",
			"merchants:view", "agents:view", "games:view",
			"tournaments:view", "jackpots:view", "bonuses:view",
			"payments:view",
			"reports:view", "reports:export",
			"audit:view",
			"admin_users:view", "roles:view",
		},
		"marketing": {
			"players:view", "games:view",
			"tournaments:view", "tournaments:create", "tournaments:edit",
			"jackpots:view", "jackpots:create", "jackpots:edit",
			"bonuses:view", "bonuses:create", "bonuses:edit",
			"reports:view", "reports:export",
			"merchants:view", "agents:view",
		},
	}

	if perms, ok := rolePerms[role]; ok {
		return perms
	}
	return []string{}
}
