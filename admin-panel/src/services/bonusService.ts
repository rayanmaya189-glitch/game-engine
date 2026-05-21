import apiClient from './api';

export interface Bonus {
  id: string;
  name: string;
  type: 'WELCOME' | 'DEPOSIT' | 'NO_DEPOSIT' | 'FREE_SPINS' | 'CASHBACK' | 'LOYALTY';
  amount: number;
  currency: string;
  wageringRequirement: number;
  minDeposit?: number;
  maxBonus?: number;
  validFrom: string;
  validUntil: string;
  status: 'ACTIVE' | 'INACTIVE' | 'EXPIRED';
  targetAudience: 'ALL' | 'NEW_PLAYERS' | 'VIP' | 'SEGMENT';
  gameRestrictions?: string[];
  createdAt: string;
}

export interface BonusCampaign {
  id: string;
  name: string;
  description: string;
  bonusType: string;
  totalBudget: number;
  usedBudget: number;
  playerCount: number;
  conversionRate: number;
  status: 'DRAFT' | 'ACTIVE' | 'PAUSED' | 'COMPLETED';
  createdAt: string;
}

export const bonusService = {
  getAllBonuses: async (page = 1, limit = 20, filters?: Record<string, string>) => {
    const params = new URLSearchParams({ page: String(page), limit: String(limit) });
    if (filters) {
      Object.entries(filters).forEach(([key, value]) => {
        params.append(key, value);
      });
    }
    const response = await apiClient.get(`/admin/bonuses?${params}`);
    return response.data;
  },

  getBonus: async (bonusId: string) => {
    const response = await apiClient.get(`/admin/bonuses/${bonusId}`);
    return response.data;
  },

  createBonus: async (bonus: Partial<Bonus>) => {
    const response = await apiClient.post('/admin/bonuses', bonus);
    return response.data;
  },

  updateBonus: async (bonusId: string, bonus: Partial<Bonus>) => {
    const response = await apiClient.put(`/admin/bonuses/${bonusId}`, bonus);
    return response.data;
  },

  deleteBonus: async (bonusId: string) => {
    const response = await apiClient.delete(`/admin/bonuses/${bonusId}`);
    return response.data;
  },

  activateBonus: async (bonusId: string) => {
    const response = await apiClient.post(`/admin/bonuses/${bonusId}/activate`);
    return response.data;
  },

  deactivateBonus: async (bonusId: string) => {
    const response = await apiClient.post(`/admin/bonuses/${bonusId}/deactivate`);
    return response.data;
  },

  getCampaigns: async (page = 1, limit = 20) => {
    const params = new URLSearchParams({ page: String(page), limit: String(limit) });
    const response = await apiClient.get(`/admin/bonus-campaigns?${params}`);
    return response.data;
  },

  createCampaign: async (campaign: Partial<BonusCampaign>) => {
    const response = await apiClient.post('/admin/bonus-campaigns', campaign);
    return response.data;
  },

  getPlayerBonuses: async (userId: string, page = 1, limit = 20) => {
    const params = new URLSearchParams({ page: String(page), limit: String(limit) });
    const response = await apiClient.get(`/admin/users/${userId}/bonuses?${params}`);
    return response.data;
  },

  grantBonus: async (userId: string, bonusId: string) => {
    const response = await apiClient.post(`/admin/users/${userId}/bonuses/${bonusId}`);
    return response.data;
  },

  getBonusPerformance: async (bonusId: string, startDate: string, endDate: string) => {
    const params = new URLSearchParams({ startDate, endDate });
    const response = await apiClient.get(`/admin/bonuses/${bonusId}/performance?${params}`);
    return response.data;
  },
};
