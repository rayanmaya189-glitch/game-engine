import apiClient from './api';

export interface User {
  id: string;
  username: string;
  email: string;
  role: 'ADMIN' | 'SUPER_ADMIN' | 'SUPPORT' | 'AGENT';
  status: 'ACTIVE' | 'INACTIVE' | 'SUSPENDED';
  createdAt: string;
  lastLoginAt?: string;
}

export interface CreateUserRequest {
  username: string;
  email: string;
  password: string;
  role: string;
}

export const userService = {
  getUsers: async (page = 1, limit = 20, filters?: Record<string, string>) => {
    const params = new URLSearchParams({ page: String(page), limit: String(limit) });
    if (filters) {
      Object.entries(filters).forEach(([key, value]) => {
        params.append(key, value);
      });
    }
    const response = await apiClient.get(`/admin/users?${params}`);
    return response.data;
  },

  getUser: async (userId: string) => {
    const response = await apiClient.get(`/admin/users/${userId}`);
    return response.data;
  },

  createUser: async (data: CreateUserRequest) => {
    const response = await apiClient.post('/admin/users', data);
    return response.data;
  },

  updateUser: async (userId: string, data: Partial<User>) => {
    const response = await apiClient.put(`/admin/users/${userId}`, data);
    return response.data;
  },

  deleteUser: async (userId: string) => {
    const response = await apiClient.delete(`/admin/users/${userId}`);
    return response.data;
  },

  suspendUser: async (userId: string, reason: string) => {
    const response = await apiClient.post(`/admin/users/${userId}/suspend`, { reason });
    return response.data;
  },

  activateUser: async (userId: string) => {
    const response = await apiClient.post(`/admin/users/${userId}/activate`);
    return response.data;
  },
};
