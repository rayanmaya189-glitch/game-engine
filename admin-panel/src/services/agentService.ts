import apiClient from './api';

export interface Agent {
  id: string;
  agentCode: string;
  name: string;
  email: string;
  phone: string;
  status: 'ACTIVE' | 'INACTIVE' | 'SUSPENDED';
  tier: 'BRONZE' | 'SILVER' | 'GOLD' | 'PLATINUM' | 'DIAMOND';
  parentId?: string;
  parentName?: string;
  commissionPlan: CommissionPlan;
  totalPlayers: number;
  activePlayers: number;
  totalCommission: number;
  pendingCommission: number;
  paidCommission: number;
  createdAt: string;
}

export interface CommissionPlan {
  id: string;
  name: string;
  type: 'PERCENTAGE' | 'FIXED' | 'HYBRID' | 'TIERED';
  rates: CommissionRate[];
  minPayout: number;
  payoutSchedule: 'DAILY' | 'WEEKLY' | 'MONTHLY';
  negativeCarryover: boolean;
  subAgentCommission: number;
}

export interface CommissionRate {
  gameType: string;
  rate: number;
  minAmount?: number;
  maxAmount?: number;
}

export interface CommissionReport {
  agentId: string;
  agentName: string;
  period: string;
  totalBets: number;
  totalGGR: number;
  commissionRate: number;
  grossCommission: number;
  adjustments: number;
  netCommission: number;
  status: 'PENDING' | 'APPROVED' | 'PAID';
  paidAt?: string;
}

export const agentService = {
  getAllAgents: async (page = 1, limit = 20, filters?: Record<string, string>) => {
    const params = new URLSearchParams({ page: String(page), limit: String(limit) });
    if (filters) {
      Object.entries(filters).forEach(([key, value]) => {
        params.append(key, value);
      });
    }
    const response = await apiClient.get(`/admin/agents?${params}`);
    return response.data;
  },

  getAgent: async (agentId: string) => {
    const response = await apiClient.get(`/admin/agents/${agentId}`);
    return response.data;
  },

  createAgent: async (agent: Partial<Agent>) => {
    const response = await apiClient.post('/admin/agents', agent);
    return response.data;
  },

  updateAgent: async (agentId: string, agent: Partial<Agent>) => {
    const response = await apiClient.put(`/admin/agents/${agentId}`, agent);
    return response.data;
  },

  suspendAgent: async (agentId: string, reason: string) => {
    const response = await apiClient.post(`/admin/agents/${agentId}/suspend`, { reason });
    return response.data;
  },

  activateAgent: async (agentId: string) => {
    const response = await apiClient.post(`/admin/agents/${agentId}/activate`);
    return response.data;
  },

  getAgentHierarchy: async (agentId: string) => {
    const response = await apiClient.get(`/admin/agents/${agentId}/hierarchy`);
    return response.data;
  },

  getAgentPlayers: async (agentId: string, page = 1, limit = 20) => {
    const params = new URLSearchParams({ page: String(page), limit: String(limit) });
    const response = await apiClient.get(`/admin/agents/${agentId}/players?${params}`);
    return response.data;
  },

  getCommissionPlans: async () => {
    const response = await apiClient.get('/admin/commission-plans');
    return response.data;
  },

  createCommissionPlan: async (plan: Partial<CommissionPlan>) => {
    const response = await apiClient.post('/admin/commission-plans', plan);
    return response.data;
  },

  updateCommissionPlan: async (planId: string, plan: Partial<CommissionPlan>) => {
    const response = await apiClient.put(`/admin/commission-plans/${planId}`, plan);
    return response.data;
  },

  getCommissionReports: async (agentId?: string, startDate?: string, endDate?: string) => {
    const params = new URLSearchParams();
    if (agentId) params.append('agentId', agentId);
    if (startDate) params.append('startDate', startDate);
    if (endDate) params.append('endDate', endDate);
    const response = await apiClient.get(`/admin/commission-reports?${params}`);
    return response.data;
  },

  approveCommission: async (reportId: string) => {
    const response = await apiClient.post(`/admin/commission-reports/${reportId}/approve`);
    return response.data;
  },

  payCommission: async (reportId: string) => {
    const response = await apiClient.post(`/admin/commission-reports/${reportId}/pay`);
    return response.data;
  },

  getAgentStatistics: async (agentId: string, startDate: string, endDate: string) => {
    const params = new URLSearchParams({ startDate, endDate });
    const response = await apiClient.get(`/admin/agents/${agentId}/statistics?${params}`);
    return response.data;
  },
};
