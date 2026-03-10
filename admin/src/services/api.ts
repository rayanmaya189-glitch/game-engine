import axios from 'axios';

const API_BASE_URL = import.meta.env.VITE_API_URL || 'http://localhost:8080/api/v1';

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
};

// Claims API (Admin endpoints)
export const claimsAPI = {
  // Commission Claims
  getAllCommissionClaims: (params?: { status?: string; page?: number; limit?: number }) =>
    api.get('/admin/claims/commission', { params }),
  getCommissionClaimById: (id: string) => api.get(`/admin/claims/commission/${id}`),
  approveCommissionClaim: (id: string, data: { adminNote: string }) =>
    api.post(`/admin/claims/commission/${id}/approve`, data),
  rejectCommissionClaim: (id: string, data: { adminNote: string }) =>
    api.post(`/admin/claims/commission/${id}/reject`, data),
  payCommissionClaim: (id: string) => 
    api.post(`/admin/claims/commission/${id}/pay`),

  // Rebet Claims
  getAllRebetClaims: (params?: { status?: string; page?: number; limit?: number }) =>
    api.get('/admin/claims/rebet', { params }),
  getRebetClaimById: (id: string) => api.get(`/admin/claims/rebet/${id}`),
  approveRebetClaim: (id: string, data: { adminNote: string }) =>
    api.post(`/admin/claims/rebet/${id}/approve`, data),
  rejectRebetClaim: (id: string, data: { adminNote: string }) =>
    api.post(`/admin/claims/rebet/${id}/reject`, data),

  // Insurance Claims
  getAllInsuranceClaims: (params?: { status?: string; page?: number; limit?: number }) =>
    api.get('/admin/claims/insurance', { params }),
  getInsuranceClaimById: (id: string) => api.get(`/admin/claims/insurance/${id}`),
  approveInsuranceClaim: (id: string, data: { adminNote: string }) =>
    api.post(`/admin/claims/insurance/${id}/approve`, data),
  rejectInsuranceClaim: (id: string, data: { adminNote: string }) =>
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

// Users API
export const usersAPI = {
  getAll: (params?: { page?: number; limit?: number; search?: string }) =>
    api.get('/admin/users', { params }),
  getById: (id: string) => api.get(`/admin/users/${id}`),
  updateStatus: (id: string, data: { status: string }) =>
    api.patch(`/admin/users/${id}/status`, data),
  updateKYC: (id: string, data: { kycLevel: string }) =>
    api.patch(`/admin/users/${id}/kyc`, data),
};

// Games API
export const gamesAPI = {
  getAll: (params?: { category?: string; status?: string; page?: number; limit?: number; search?: string }) =>
    api.get('/admin/games', { params }),
  getById: (id: string) => api.get(`/admin/games/${id}`),
  create: (data: any) => api.post('/admin/games', data),
  update: (id: string, data: any) => api.put(`/admin/games/${id}`, data),
  delete: (id: string) => api.delete(`/admin/games/${id}`),
};

// Reports API
export const reportsAPI = {
  getDashboardStats: () => api.get('/admin/reports/dashboard'),
  getRevenueReport: (params?: { startDate?: string; endDate?: string }) =>
    api.get('/admin/reports/revenue', { params }),
  getUserReport: (params?: { startDate?: string; endDate?: string }) =>
    api.get('/admin/reports/users', { params }),
  getGameReport: (params?: { startDate?: string; endDate?: string }) =>
    api.get('/admin/reports/games', { params }),
};

// Merchants API
export const merchantsAPI = {
  getAll: (params?: { page?: number; limit?: number; search?: string }) =>
    api.get('/admin/merchants', { params }),
  getById: (id: string) => api.get(`/admin/merchants/${id}`),
  create: (data: any) => api.post('/admin/merchants', data),
  update: (id: string, data: any) => api.put(`/admin/merchants/${id}`, data),
  updateStatus: (id: string, data: { status: string }) =>
    api.put(`/admin/merchants/${id}/status`, data),
  delete: (id: string) => api.delete(`/admin/merchants/${id}`),
};

// Agents API
export const agentsAPI = {
  getAll: (params?: { page?: number; limit?: number; search?: string }) =>
    api.get('/admin/agents', { params }),
  getById: (id: string) => api.get(`/admin/agents/${id}`),
  create: (data: any) => api.post('/admin/agents', data),
  update: (id: string, data: any) => api.put(`/admin/agents/${id}`, data),
  updateStatus: (id: string, data: { status: string }) =>
    api.put(`/admin/agents/${id}/status`, data),
  delete: (id: string) => api.delete(`/admin/agents/${id}`),
};

// Tournaments API
export const tournamentsAPI = {
  getAll: (params?: { page?: number; limit?: number; search?: string }) =>
    api.get('/admin/tournaments', { params }),
  getById: (id: string) => api.get(`/admin/tournaments/${id}`),
  create: (data: any) => api.post('/admin/tournaments', data),
  update: (id: string, data: any) => api.put(`/admin/tournaments/${id}`, data),
  updateStatus: (id: string, data: { status: string }) =>
    api.put(`/admin/tournaments/${id}/status`, data),
  delete: (id: string) => api.delete(`/admin/tournaments/${id}`),
  getLeaderboard: (id: string) => api.get(`/admin/tournaments/${id}/leaderboard`),
};

// Jackpots API
export const jackpotsAPI = {
  getAll: (params?: { page?: number; limit?: number; search?: string }) =>
    api.get('/admin/jackpots', { params }),
  getById: (id: string) => api.get(`/admin/jackpots/${id}`),
  create: (data: any) => api.post('/admin/jackpots', data),
  update: (id: string, data: any) => api.put(`/admin/jackpots/${id}`, data),
  updateStatus: (id: string, data: { status: string }) =>
    api.put(`/admin/jackpots/${id}/status`, data),
  delete: (id: string) => api.delete(`/admin/jackpots/${id}`),
  getHits: (id: string) => api.get(`/admin/jackpots/${id}/hits`),
};

// Bonuses API
export const bonusesAPI = {
  getAll: (params?: { page?: number; limit?: number; search?: string }) =>
    api.get('/admin/bonuses', { params }),
  getById: (id: string) => api.get(`/admin/bonuses/${id}`),
  create: (data: any) => api.post('/admin/bonuses', data),
  update: (id: string, data: any) => api.put(`/admin/bonuses/${id}`, data),
  updateStatus: (id: string, data: { status: string }) =>
    api.put(`/admin/bonuses/${id}/status`, data),
  delete: (id: string) => api.delete(`/admin/bonuses/${id}`),
};

// Payments API
export const paymentsAPI = {
  getAll: (params?: { page?: number; limit?: number; status?: string; type?: string }) =>
    api.get('/admin/payments', { params }),
  getById: (id: string) => api.get(`/admin/payments/${id}`),
  approve: (id: string, data?: { adminNote?: string }) =>
    api.put(`/admin/payments/${id}/approve`, data),
  reject: (id: string, data?: { adminNote?: string }) =>
    api.put(`/admin/payments/${id}/reject`, data),
  process: (id: string) => api.put(`/admin/payments/${id}/process`),
};

export default api;
