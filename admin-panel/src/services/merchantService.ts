import apiClient from './api';

export interface Merchant {
  id: string;
  merchantCode: string;
  name: string;
  email: string;
  phone: string;
  status: 'ACTIVE' | 'INACTIVE' | 'SUSPENDED' | 'PENDING';
  whiteLabelDomain: string;
  logoUrl: string;
  themeConfig: ThemeConfig;
  allowedCountries: string[];
  restrictedCountries: string[];
  currency: string;
  languages: string[];
  apiKey: string;
  webhookUrl: string;
  totalPlayers: number;
  activePlayers: number;
  totalGGR: number;
  revenueShare: number;
  createdAt: string;
}

export interface ThemeConfig {
  primaryColor: string;
  secondaryColor: string;
  backgroundColor: string;
  textColor: string;
  logoPosition: 'LEFT' | 'CENTER' | 'RIGHT';
  showFooter: boolean;
  customCSS?: string;
}

export interface MerchantReport {
  merchantId: string;
  merchantName: string;
  period: string;
  totalPlayers: number;
  newPlayers: number;
  activePlayers: number;
  totalBets: number;
  totalWins: number;
  ggr: number;
  ngr: number;
  revenueShare: number;
  payout: number;
  status: 'PENDING' | 'APPROVED' | 'PAID';
}

export const merchantService = {
  getAllMerchants: async (page = 1, limit = 20, filters?: Record<string, string>) => {
    const params = new URLSearchParams({ page: String(page), limit: String(limit) });
    if (filters) {
      Object.entries(filters).forEach(([key, value]) => {
        params.append(key, value);
      });
    }
    const response = await apiClient.get(`/admin/merchants?${params}`);
    return response.data;
  },

  getMerchant: async (merchantId: string) => {
    const response = await apiClient.get(`/admin/merchants/${merchantId}`);
    return response.data;
  },

  createMerchant: async (merchant: Partial<Merchant>) => {
    const response = await apiClient.post('/admin/merchants', merchant);
    return response.data;
  },

  updateMerchant: async (merchantId: string, merchant: Partial<Merchant>) => {
    const response = await apiClient.put(`/admin/merchants/${merchantId}`, merchant);
    return response.data;
  },

  suspendMerchant: async (merchantId: string, reason: string) => {
    const response = await apiClient.post(`/admin/merchants/${merchantId}/suspend`, { reason });
    return response.data;
  },

  activateMerchant: async (merchantId: string) => {
    const response = await apiClient.post(`/admin/merchants/${merchantId}/activate`);
    return response.data;
  },

  regenerateApiKey: async (merchantId: string) => {
    const response = await apiClient.post(`/admin/merchants/${merchantId}/regenerate-key`);
    return response.data;
  },

  getMerchantPlayers: async (merchantId: string, page = 1, limit = 20) => {
    const params = new URLSearchParams({ page: String(page), limit: String(limit) });
    const response = await apiClient.get(`/admin/merchants/${merchantId}/players?${params}`);
    return response.data;
  },

  getMerchantReports: async (merchantId?: string, startDate?: string, endDate?: string) => {
    const params = new URLSearchParams();
    if (merchantId) params.append('merchantId', merchantId);
    if (startDate) params.append('startDate', startDate);
    if (endDate) params.append('endDate', endDate);
    const response = await apiClient.get(`/admin/merchant-reports?${params}`);
    return response.data;
  },

  approveMerchantReport: async (reportId: string) => {
    const response = await apiClient.post(`/admin/merchant-reports/${reportId}/approve`);
    return response.data;
  },

  payMerchant: async (reportId: string) => {
    const response = await apiClient.post(`/admin/merchant-reports/${reportId}/pay`);
    return response.data;
  },

  getMerchantStatistics: async (merchantId: string, startDate: string, endDate: string) => {
    const params = new URLSearchParams({ startDate, endDate });
    const response = await apiClient.get(`/admin/merchants/${merchantId}/statistics?${params}`);
    return response.data;
  },

  updateWhiteLabelConfig: async (merchantId: string, config: Partial<ThemeConfig>) => {
    const response = await apiClient.put(`/admin/merchants/${merchantId}/whitelabel`, config);
    return response.data;
  },

  getMerchantGames: async (merchantId: string) => {
    const response = await apiClient.get(`/admin/merchants/${merchantId}/games`);
    return response.data;
  },

  updateMerchantGames: async (merchantId: string, gameIds: string[]) => {
    const response = await apiClient.put(`/admin/merchants/${merchantId}/games`, { gameIds });
    return response.data;
  },
};
