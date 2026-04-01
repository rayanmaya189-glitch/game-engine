import axios from 'axios';

const API_BASE_URL = import.meta.env.VITE_API_URL || 'http://localhost:8081/api/v1';

const api = axios.create({
  baseURL: API_BASE_URL,
  headers: {
    'Content-Type': 'application/json',
  },
});

api.interceptors.request.use((config) => {
  const token = sessionStorage.getItem('adminToken');
  if (token) {
    config.headers.Authorization = `Bearer ${token}`;
  }
  return config;
});

api.interceptors.response.use(
  (response) => response,
  (error) => {
    if (error.response?.status === 401) {
      sessionStorage.removeItem('adminToken');
      window.location.href = '/login';
    }
    return Promise.reject(error);
  }
);

// Auth API
export const authAPI = {
  login: (data: { username: string; password: string }) =>
    api.post('/admin/auth/login', data),
  logout: () => api.post('/admin/auth/logout'),
  refreshToken: (data: { refreshToken: string }) =>
    api.post('/admin/auth/refresh', data),
};

// Players API (matches backend /api/v1/admin/players)
export const playersAPI = {
  getAll: (params?: { page?: number; limit?: number; search?: string; status?: string }) =>
    api.get('/admin/players', { params }),
  getById: (id: string) => api.get(`/admin/players/${id}`),
  updateStatus: (id: string, data: { status: string; reason?: string }) =>
    api.put(`/admin/players/${id}/status`, data),
  getStats: (id: string) => api.get(`/admin/players/${id}/stats`),
};

// KYC API (matches backend /api/v1/admin/kyc)
export const kycAPI = {
  getAll: (params?: { status?: string; level?: string; page?: number; limit?: number }) =>
    api.get('/admin/kyc', { params }),
  getById: (id: string) => api.get(`/admin/kyc/${id}`),
  approve: (id: string, data?: { adminNote?: string }) =>
    api.put(`/admin/kyc/${id}/approve`, data),
  reject: (id: string, data: { reason: string; adminNote?: string }) =>
    api.put(`/admin/kyc/${id}/reject`, data),
};

// Claims API (matches backend /api/v1/admin/claims/*)
export const claimsAPI = {
  // Commission Claims
  getAllCommissionClaims: (params?: { status?: string; page?: number; limit?: number }) =>
    api.get('/admin/claims/commission', { params }),
  getCommissionClaimById: (id: string) => api.get(`/admin/claims/commission/${id}`),
  approveCommissionClaim: (id: string, data?: { adminNote?: string }) =>
    api.post(`/admin/claims/commission/${id}/approve`, data),
  rejectCommissionClaim: (id: string, data: { reason: string; adminNote?: string }) =>
    api.post(`/admin/claims/commission/${id}/reject`, data),
  payCommissionClaim: (id: string) =>
    api.post(`/admin/claims/commission/${id}/pay`),

  // Rebet Claims
  getAllRebetClaims: (params?: { status?: string; page?: number; limit?: number }) =>
    api.get('/admin/claims/rebet', { params }),
  getRebetClaimById: (id: string) => api.get(`/admin/claims/rebet/${id}`),
  approveRebetClaim: (id: string, data?: { adminNote?: string }) =>
    api.post(`/admin/claims/rebet/${id}/approve`, data),
  rejectRebetClaim: (id: string, data: { reason: string; adminNote?: string }) =>
    api.post(`/admin/claims/rebet/${id}/reject`, data),

  // Insurance Claims
  getAllInsuranceClaims: (params?: { status?: string; page?: number; limit?: number }) =>
    api.get('/admin/claims/insurance', { params }),
  getInsuranceClaimById: (id: string) => api.get(`/admin/claims/insurance/${id}`),
  approveInsuranceClaim: (id: string, data?: { adminNote?: string }) =>
    api.post(`/admin/claims/insurance/${id}/approve`, data),
  rejectInsuranceClaim: (id: string, data: { reason: string; adminNote?: string }) =>
    api.post(`/admin/claims/insurance/${id}/reject`, data),
  payInsuranceClaim: (id: string) =>
    api.post(`/admin/claims/insurance/${id}/pay`),

  // Settlements
  getAllSettlements: (params?: { status?: string; type?: string; page?: number; limit?: number }) =>
    api.get('/admin/claims/settlements', { params }),
  getSettlementById: (id: string) => api.get(`/admin/claims/settlements/${id}`),

  // Statistics
  getClaimStatistics: () => api.get('/admin/claims/statistics'),
};

// Games API (matches backend /api/v1/admin/games)
export const gamesAPI = {
  getAll: (params?: { category?: string; status?: string; page?: number; limit?: number; search?: string }) =>
    api.get('/admin/games', { params }),
  getById: (id: string) => api.get(`/admin/games/${id}`),
  create: (data: { name: string; category: string; provider: string; description?: string; minBet?: number; maxBet?: number }) =>
    api.post('/admin/games', data),
  update: (id: string, data: { name?: string; status?: string; minBet?: number; maxBet?: number }) =>
    api.put(`/admin/games/${id}`, data),
  delete: (id: string) => api.delete(`/admin/games/${id}`),
};

// Merchants API (matches backend /api/v1/admin/merchants)
export const merchantsAPI = {
  getAll: (params?: { page?: number; limit?: number; search?: string; status?: string }) =>
    api.get('/admin/merchants', { params }),
  getById: (id: string) => api.get(`/admin/merchants/${id}`),
  create: (data: { name: string; email: string; webhookUrl?: string; configName?: string }) =>
    api.post('/admin/merchants', data),
  update: (id: string, data: { name?: string; webhookUrl?: string; configName?: string }) =>
    api.put(`/admin/merchants/${id}`, data),
  updateStatus: (id: string, data: { status: string }) =>
    api.put(`/admin/merchants/${id}/status`, data),
  delete: (id: string) => api.delete(`/admin/merchants/${id}`),
};

// Agents API (matches backend /api/v1/admin/agents)
export const agentsAPI = {
  getAll: (params?: { page?: number; limit?: number; search?: string; status?: string }) =>
    api.get('/admin/agents', { params }),
  getById: (id: string) => api.get(`/admin/agents/${id}`),
  create: (data: { name: string; email: string; affiliateCode?: string; commissionRate?: number }) =>
    api.post('/admin/agents', data),
  update: (id: string, data: { name?: string; commissionRate?: number }) =>
    api.put(`/admin/agents/${id}`, data),
  updateStatus: (id: string, data: { status: string }) =>
    api.put(`/admin/agents/${id}/status`, data),
  delete: (id: string) => api.delete(`/admin/agents/${id}`),
};

// Tournaments API (matches backend /api/v1/admin/tournaments)
export const tournamentsAPI = {
  getAll: (params?: { page?: number; limit?: number; search?: string; status?: string }) =>
    api.get('/admin/tournaments', { params }),
  getById: (id: string) => api.get(`/admin/tournaments/${id}`),
  create: (data: { name: string; type: string; buyIn?: number; prizePool?: number; startTime?: string; endTime?: string }) =>
    api.post('/admin/tournaments', data),
  update: (id: string, data: { name?: string; prizePool?: number; startTime?: string; endTime?: string }) =>
    api.put(`/admin/tournaments/${id}`, data),
  updateStatus: (id: string, data: { status: string }) =>
    api.put(`/admin/tournaments/${id}/status`, data),
  delete: (id: string) => api.delete(`/admin/tournaments/${id}`),
  getLeaderboard: (id: string) => api.get(`/admin/tournaments/${id}/leaderboard`),
};

// Jackpots API (matches backend /api/v1/admin/jackpots)
export const jackpotsAPI = {
  getAll: (params?: { page?: number; limit?: number; search?: string; status?: string }) =>
    api.get('/admin/jackpots', { params }),
  getById: (id: string) => api.get(`/admin/jackpots/${id}`),
  create: (data: { name: string; type: string; seedAmount?: number; contributionRate?: number }) =>
    api.post('/admin/jackpots', data),
  update: (id: string, data: { name?: string; contributionRate?: number }) =>
    api.put(`/admin/jackpots/${id}`, data),
  updateStatus: (id: string, data: { status: string }) =>
    api.put(`/admin/jackpots/${id}/status`, data),
  delete: (id: string) => api.delete(`/admin/jackpots/${id}`),
  getHits: (id: string, params?: { page?: number; limit?: number }) =>
    api.get(`/admin/jackpots/${id}/hits`, { params }),
};

// Bonuses API (matches backend /api/v1/admin/bonuses)
export const bonusesAPI = {
  getAll: (params?: { page?: number; limit?: number; search?: string; status?: string }) =>
    api.get('/admin/bonuses', { params }),
  getById: (id: string) => api.get(`/admin/bonuses/${id}`),
  create: (data: { name: string; type: string; amount?: number; percentage?: number; wageringMultiplier?: number; maxAmount?: number }) =>
    api.post('/admin/bonuses', data),
  update: (id: string, data: { name?: string; amount?: number; percentage?: number; wageringMultiplier?: number }) =>
    api.put(`/admin/bonuses/${id}`, data),
  updateStatus: (id: string, data: { status: string }) =>
    api.put(`/admin/bonuses/${id}/status`, data),
  delete: (id: string) => api.delete(`/admin/bonuses/${id}`),
};

// Payments API (matches backend /api/v1/admin/payments)
export const paymentsAPI = {
  getAll: (params?: { page?: number; limit?: number; status?: string; type?: string }) =>
    api.get('/admin/payments', { params }),
  getById: (id: string) => api.get(`/admin/payments/${id}`),
  approve: (id: string, data?: { adminNote?: string }) =>
    api.put(`/admin/payments/${id}/approve`, data),
  reject: (id: string, data: { reason: string; adminNote?: string }) =>
    api.put(`/admin/payments/${id}/reject`, data),
  process: (id: string) =>
    api.put(`/admin/payments/${id}/process`),
};

// Reports API (matches backend /api/v1/admin/reports/*)
export const reportsAPI = {
  getDashboardStats: () => api.get('/admin/reports/dashboard'),
  getRevenueReport: (params?: { startDate?: string; endDate?: string; period?: string }) =>
    api.get('/admin/reports/revenue', { params }),
  getUserReport: (params?: { startDate?: string; endDate?: string }) =>
    api.get('/admin/reports/users', { params }),
  getGameReport: (params?: { startDate?: string; endDate?: string }) =>
    api.get('/admin/reports/games', { params }),
  getFinancialReport: (params?: { startDate?: string; endDate?: string }) =>
    api.get('/admin/reports/financial', { params }),
};

// Wallet API (matches backend /api/v1/admin/wallet/*)
export const walletAPI = {
  getTransactions: (params?: { userId?: string; type?: string; status?: string; page?: number; limit?: number }) =>
    api.get('/admin/wallet/transactions', { params }),
  adjustBalance: (data: { userId: string; amount: number; reason: string; currency?: string }) =>
    api.post('/admin/wallet/adjust', data),
};

// RBAC API
export const rbacAPI = {
  getPermissions: () => api.get('/admin/permissions'),
  getRoles: () => api.get('/admin/roles'),
  createRole: (data: { key: string; name: string; description: string; permissions: string[] }) =>
    api.post('/admin/roles', data),
  updateRole: (id: string, data: { name?: string; permissions?: string[] }) =>
    api.put(`/admin/roles/${id}`, data),
  deleteRole: (id: string) => api.delete(`/admin/roles/${id}`),
  getAdminUsers: (params?: { page?: number; limit?: number }) =>
    api.get('/admin/admin-users', { params }),
  createAdminUser: (data: { username: string; email: string; password: string; role: string }) =>
    api.post('/admin/admin-users', data),
  updateAdminUser: (id: string, data: { email?: string; role?: string; isActive?: boolean }) =>
    api.put(`/admin/admin-users/${id}`, data),
  deleteAdminUser: (id: string) => api.delete(`/admin/admin-users/${id}`),
};

// Banner API (matches backend /api/v1/admin/banners)
export const bannerAPI = {
  getAll: (params?: { type?: string; status?: string; page?: number; limit?: number }) =>
    api.get('/admin/banners', { params }),
  create: (data: Record<string, any>) => api.post('/admin/banners', data),
  update: (id: string, data: Record<string, any>) => api.put(`/admin/banners/${id}`, data),
  delete: (id: string) => api.delete(`/admin/banners/${id}`),
  updateStatus: (id: string, data: { status: string }) =>
    api.put(`/admin/banners/${id}/status`, data),
};

// Referral API (matches backend /api/v1/admin/referrals)
export const referralAPI = {
  getCodes: (params?: { status?: string; page?: number; limit?: number }) =>
    api.get('/admin/referrals/codes', { params }),
  getStats: () => api.get('/admin/referrals/stats'),
  getRewards: () => api.get('/admin/referrals/rewards'),
  updateReward: (id: string, data: Record<string, any>) =>
    api.put(`/admin/referrals/rewards/${id}`, data),
};

export default api;
