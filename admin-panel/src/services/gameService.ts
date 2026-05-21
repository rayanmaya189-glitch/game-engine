import apiClient from './api';

export interface Game {
  id: string;
  name: string;
  type: 'SLOT' | 'CARD' | 'DICE' | 'TABLE' | 'LIVE_DEALER';
  provider: string;
  status: 'ACTIVE' | 'INACTIVE' | 'MAINTENANCE';
  rtp: number;
  minBet: number;
  maxBet: number;
  totalPlays: number;
  createdAt: string;
}

export interface GameStats {
  gameId: string;
  totalBets: number;
  totalWins: number;
  totalPayout: number;
  houseEdge: number;
  activePlayers: number;
}

export const gameService = {
  getGames: async (page = 1, limit = 20, filters?: Record<string, string>) => {
    const params = new URLSearchParams({ page: String(page), limit: String(limit) });
    if (filters) {
      Object.entries(filters).forEach(([key, value]) => {
        params.append(key, value);
      });
    }
    const response = await apiClient.get(`/admin/games?${params}`);
    return response.data;
  },

  getGame: async (gameId: string) => {
    const response = await apiClient.get(`/admin/games/${gameId}`);
    return response.data;
  },

  createGame: async (data: Partial<Game>) => {
    const response = await apiClient.post('/admin/games', data);
    return response.data;
  },

  updateGame: async (gameId: string, data: Partial<Game>) => {
    const response = await apiClient.put(`/admin/games/${gameId}`, data);
    return response.data;
  },

  deleteGame: async (gameId: string) => {
    const response = await apiClient.delete(`/admin/games/${gameId}`);
    return response.data;
  },

  toggleGameStatus: async (gameId: string, status: 'ACTIVE' | 'INACTIVE' | 'MAINTENANCE') => {
    const response = await apiClient.patch(`/admin/games/${gameId}/status`, { status });
    return response.data;
  },

  getGameStats: async (gameId: string, startDate?: string, endDate?: string) => {
    const params = new URLSearchParams();
    if (startDate) params.append('startDate', startDate);
    if (endDate) params.append('endDate', endDate);
    const response = await apiClient.get(`/admin/games/${gameId}/stats?${params}`);
    return response.data;
  },

  getAllGameStats: async (startDate?: string, endDate?: string) => {
    const params = new URLSearchParams();
    if (startDate) params.append('startDate', startDate);
    if (endDate) params.append('endDate', endDate);
    const response = await apiClient.get(`/admin/games/stats?${params}`);
    return response.data;
  },
};
