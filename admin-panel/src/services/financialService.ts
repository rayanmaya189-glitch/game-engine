import apiClient from './api';

export interface Transaction {
  id: string;
  userId: string;
  type: 'DEPOSIT' | 'WITHDRAWAL' | 'BET' | 'WIN' | 'BONUS' | 'REFUND';
  amount: number;
  currency: string;
  status: 'PENDING' | 'COMPLETED' | 'FAILED' | 'CANCELLED';
  paymentMethod?: string;
  transactionHash?: string;
  createdAt: string;
  processedAt?: string;
}

export interface Wallet {
  userId: string;
  balance: number;
  bonusBalance: number;
  currency: string;
  frozenAmount: number;
}

export const financialService = {
  getTransactions: async (page = 1, limit = 20, filters?: Record<string, string>) => {
    const params = new URLSearchParams({ page: String(page), limit: String(limit) });
    if (filters) {
      Object.entries(filters).forEach(([key, value]) => {
        params.append(key, value);
      });
    }
    const response = await apiClient.get(`/admin/transactions?${params}`);
    return response.data;
  },

  getTransaction: async (transactionId: string) => {
    const response = await apiClient.get(`/admin/transactions/${transactionId}`);
    return response.data;
  },

  approveWithdrawal: async (transactionId: string) => {
    const response = await apiClient.post(`/admin/transactions/${transactionId}/approve`);
    return response.data;
  },

  rejectWithdrawal: async (transactionId: string, reason: string) => {
    const response = await apiClient.post(`/admin/transactions/${transactionId}/reject`, { reason });
    return response.data;
  },

  getWallet: async (userId: string) => {
    const response = await apiClient.get(`/admin/wallets/${userId}`);
    return response.data;
  },

  adjustBalance: async (userId: string, amount: number, reason: string, type: 'CREDIT' | 'DEBIT') => {
    const response = await apiClient.post(`/admin/wallets/${userId}/adjust`, { amount, reason, type });
    return response.data;
  },

  getFinancialReport: async (startDate: string, endDate: string) => {
    const params = new URLSearchParams({ startDate, endDate });
    const response = await apiClient.get(`/admin/reports/financial?${params}`);
    return response.data;
  },

  getRevenueStats: async (period: 'DAY' | 'WEEK' | 'MONTH' | 'YEAR') => {
    const response = await apiClient.get(`/admin/reports/revenue?period=${period}`);
    return response.data;
  },
};
