import apiClient from './api';

export interface Tournament {
  id: string;
  name: string;
  type: 'SCHEDULED' | 'SIT_AND_GO' | 'FREEROLL' | 'SATELLITE' | 'REBUY' | 'BOUNTY';
  gameId: string;
  gameName: string;
  entryFee: number;
  currency: string;
  guaranteedPrizePool: number;
  currentPrizePool: number;
  maxPlayers: number;
  registeredPlayers: number;
  status: 'UPCOMING' | 'REGISTERING' | 'IN_PROGRESS' | 'COMPLETED' | 'CANCELLED';
  startTime: string;
  endTime?: string;
  blindStructure?: BlindLevel[];
  prizeDistribution: PrizeDistribution[];
  createdAt: string;
}

export interface BlindLevel {
  level: number;
  smallBlind: number;
  bigBlind: number;
  ante?: number;
  durationMinutes: number;
}

export interface PrizeDistribution {
  rank: number;
  percentage: number;
  amount?: number;
}

export interface TournamentPlayer {
  playerId: string;
  playerName: string;
  rank?: number;
  chips: number;
  status: 'ACTIVE' | 'ELIMINATED' | 'REGISTERED';
  bountyAmount?: number;
}

export const tournamentService = {
  getAllTournaments: async (page = 1, limit = 20, filters?: Record<string, string>) => {
    const params = new URLSearchParams({ page: String(page), limit: String(limit) });
    if (filters) {
      Object.entries(filters).forEach(([key, value]) => {
        params.append(key, value);
      });
    }
    const response = await apiClient.get(`/admin/tournaments?${params}`);
    return response.data;
  },

  getTournament: async (tournamentId: string) => {
    const response = await apiClient.get(`/admin/tournaments/${tournamentId}`);
    return response.data;
  },

  createTournament: async (tournament: Partial<Tournament>) => {
    const response = await apiClient.post('/admin/tournaments', tournament);
    return response.data;
  },

  updateTournament: async (tournamentId: string, tournament: Partial<Tournament>) => {
    const response = await apiClient.put(`/admin/tournaments/${tournamentId}`, tournament);
    return response.data;
  },

  cancelTournament: async (tournamentId: string, reason: string) => {
    const response = await apiClient.post(`/admin/tournaments/${tournamentId}/cancel`, { reason });
    return response.data;
  },

  startTournament: async (tournamentId: string) => {
    const response = await apiClient.post(`/admin/tournaments/${tournamentId}/start`);
    return response.data;
  },

  getTournamentPlayers: async (tournamentId: string) => {
    const response = await apiClient.get(`/admin/tournaments/${tournamentId}/players`);
    return response.data;
  },

  registerPlayer: async (tournamentId: string, userId: string) => {
    const response = await apiClient.post(`/admin/tournaments/${tournamentId}/register`, { userId });
    return response.data;
  },

  getLeaderboard: async (tournamentId: string) => {
    const response = await apiClient.get(`/admin/tournaments/${tournamentId}/leaderboard`);
    return response.data;
  },

  distributePrizes: async (tournamentId: string) => {
    const response = await apiClient.post(`/admin/tournaments/${tournamentId}/distribute-prizes`);
    return response.data;
  },

  getTournamentHistory: async (page = 1, limit = 20, filters?: Record<string, string>) => {
    const params = new URLSearchParams({ page: String(page), limit: String(limit) });
    if (filters) {
      Object.entries(filters).forEach(([key, value]) => {
        params.append(key, value);
      });
    }
    const response = await apiClient.get(`/admin/tournaments/history?${params}`);
    return response.data;
  },

  getTournamentStatistics: async (tournamentId: string) => {
    const response = await apiClient.get(`/admin/tournaments/${tournamentId}/statistics`);
    return response.data;
  },
};
