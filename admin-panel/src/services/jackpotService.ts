import apiClient from './api';

export interface Jackpot {
  id: string;
  name: string;
  type: 'FIXED' | 'PROGRESSIVE_LOCAL' | 'PROGRESSIVE_NETWORK' | 'MYSTERY' | 'MULTI_TIER';
  currentAmount: number;
  seedAmount: number;
  maxAmount?: number;
  contributionRate: number;
  currency: string;
  triggerType: 'SYMBOL_COMBINATION' | 'RANDOM' | 'THRESHOLD';
  linkedGames: string[];
  status: 'ACTIVE' | 'INACTIVE' | 'EXHAUSTED';
  hitCount: number;
  lastHitAt?: string;
  lastHitAmount?: number;
  createdAt: string;
  updatedAt: string;
}

export interface JackpotHit {
  id: string;
  jackpotId: string;
  userId: string;
  userName: string;
  gameId: string;
  gameName: string;
  hitAmount: number;
  triggeredBy: string;
  hitAt: string;
}

export const jackpotService = {
  getAllJackpots: async (page = 1, limit = 20) => {
    const params = new URLSearchParams({ page: String(page), limit: String(limit) });
    const response = await apiClient.get(`/admin/jackpots?${params}`);
    return response.data;
  },

  getJackpot: async (jackpotId: string) => {
    const response = await apiClient.get(`/admin/jackpots/${jackpotId}`);
    return response.data;
  },

  createJackpot: async (jackpot: Partial<Jackpot>) => {
    const response = await apiClient.post('/admin/jackpots', jackpot);
    return response.data;
  },

  updateJackpot: async (jackpotId: string, jackpot: Partial<Jackpot>) => {
    const response = await apiClient.put(`/admin/jackpots/${jackpotId}`, jackpot);
    return response.data;
  },

  deleteJackpot: async (jackpotId: string) => {
    const response = await apiClient.delete(`/admin/jackpots/${jackpotId}`);
    return response.data;
  },

  activateJackpot: async (jackpotId: string) => {
    const response = await apiClient.put(`/admin/jackpots/${jackpotId}/activate`);
    return response.data;
  },

  deactivateJackpot: async (jackpotId: string) => {
    const response = await apiClient.put(`/admin/jackpots/${jackpotId}/deactivate`);
    return response.data;
  },

  getJackpotHits: async (jackpotId: string, page = 1, limit = 20) => {
    const params = new URLSearchParams({ page: String(page), limit: String(limit) });
    const response = await apiClient.get(`/admin/jackpots/${jackpotId}/hits?${params}`);
    return response.data;
  },

  resetJackpot: async (jackpotId: string, amount?: number) => {
    const response = await apiClient.post(`/admin/jackpots/${jackpotId}/reset`, { amount });
    return response.data;
  },

  getJackpotStatistics: async (startDate: string, endDate: string) => {
    const params = new URLSearchParams({ startDate, endDate });
    const response = await apiClient.get(`/admin/jackpots/statistics?${params}`);
    return response.data;
  },
};
