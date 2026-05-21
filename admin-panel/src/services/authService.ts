import apiClient from './api';

export interface AuthResponse {
  token: string;
  user: {
    id: string;
    username: string;
    email: string;
    role: string;
  };
}

export const authService = {
  login: async (username: string, password: string) => {
    const response = await apiClient.post('/admin/auth/login', { username, password });
    return response.data;
  },

  logout: async () => {
    const response = await apiClient.post('/admin/auth/logout');
    return response.data;
  },

  refreshToken: async () => {
    const response = await apiClient.post('/admin/auth/refresh');
    return response.data;
  },

  getCurrentUser: async () => {
    const response = await apiClient.get('/admin/auth/me');
    return response.data;
  },

  changePassword: async (currentPassword: string, newPassword: string) => {
    const response = await apiClient.post('/admin/auth/change-password', { currentPassword, newPassword });
    return response.data;
  },
};
