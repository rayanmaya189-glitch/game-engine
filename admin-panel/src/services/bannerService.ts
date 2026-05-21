import apiClient from './api';

export interface Banner {
  id: string;
  name: string;
  type: 'HERO' | 'SIDEBAR' | 'POPUP' | 'IN_GAME' | 'FOOTER';
  imageUrl: string;
  mobileImageUrl?: string;
  targetUrl?: string;
  title?: string;
  subtitle?: string;
  ctaText?: string;
  placement: 'HOME' | 'LOBBY' | 'GAME' | 'CHECKOUT' | 'ALL';
  targetingRules: TargetingRule[];
  schedule: ScheduleConfig;
  status: 'DRAFT' | 'ACTIVE' | 'PAUSED' | 'EXPIRED';
  impressions: number;
  clicks: number;
  ctr: number;
  conversions: number;
  createdAt: string;
  updatedAt: string;
}

export interface TargetingRule {
  type: 'COUNTRY' | 'VIP_LEVEL' | 'NEW_PLAYER' | 'SEGMENT' | 'DEVICE';
  operator: 'EQUALS' | 'NOT_EQUALS' | 'IN' | 'NOT_IN';
  values: string[];
}

export interface ScheduleConfig {
  startDate: string;
  endDate?: string;
  daysOfWeek?: number[];
  hoursOfDay?: { start: number; end: number }[];
  timezone: string;
}

export interface Announcement {
  id: string;
  title: string;
  content: string;
  type: 'INFO' | 'WARNING' | 'MAINTENANCE' | 'PROMOTION';
  priority: 'LOW' | 'MEDIUM' | 'HIGH' | 'URGENT';
  channels: ('WEBSITE' | 'MOBILE' | 'EMAIL' | 'SMS' | 'PUSH')[];
  targetingRules: TargetingRule[];
  schedule: ScheduleConfig;
  status: 'DRAFT' | 'SCHEDULED' | 'ACTIVE' | 'EXPIRED';
  createdAt: string;
  sentAt?: string;
}

export const bannerService = {
  getAllBanners: async (page = 1, limit = 20, filters?: Record<string, string>) => {
    const params = new URLSearchParams({ page: String(page), limit: String(limit) });
    if (filters) {
      Object.entries(filters).forEach(([key, value]) => {
        params.append(key, value);
      });
    }
    const response = await apiClient.get(`/admin/banners?${params}`);
    return response.data;
  },

  getBanner: async (bannerId: string) => {
    const response = await apiClient.get(`/admin/banners/${bannerId}`);
    return response.data;
  },

  createBanner: async (banner: Partial<Banner>) => {
    const response = await apiClient.post('/admin/banners', banner);
    return response.data;
  },

  updateBanner: async (bannerId: string, banner: Partial<Banner>) => {
    const response = await apiClient.put(`/admin/banners/${bannerId}`, banner);
    return response.data;
  },

  deleteBanner: async (bannerId: string) => {
    const response = await apiClient.delete(`/admin/banners/${bannerId}`);
    return response.data;
  },

  activateBanner: async (bannerId: string) => {
    const response = await apiClient.post(`/admin/banners/${bannerId}/activate`);
    return response.data;
  },

  pauseBanner: async (bannerId: string) => {
    const response = await apiClient.post(`/admin/banners/${bannerId}/pause`);
    return response.data;
  },

  uploadBannerImage: async (file: File) => {
    const formData = new FormData();
    formData.append('file', file);
    const response = await apiClient.post('/admin/banners/upload', formData, {
      headers: { 'Content-Type': 'multipart/form-data' },
    });
    return response.data;
  },

  getBannerAnalytics: async (bannerId: string, startDate: string, endDate: string) => {
    const params = new URLSearchParams({ startDate, endDate });
    const response = await apiClient.get(`/admin/banners/${bannerId}/analytics?${params}`);
    return response.data;
  },

  getAllAnnouncements: async (page = 1, limit = 20) => {
    const params = new URLSearchParams({ page: String(page), limit: String(limit) });
    const response = await apiClient.get(`/admin/announcements?${params}`);
    return response.data;
  },

  createAnnouncement: async (announcement: Partial<Announcement>) => {
    const response = await apiClient.post('/admin/announcements', announcement);
    return response.data;
  },

  sendAnnouncement: async (announcementId: string) => {
    const response = await apiClient.post(`/admin/announcements/${announcementId}/send`);
    return response.data;
  },

  getAssetLibrary: async (type?: string) => {
    const params = type ? new URLSearchParams({ type }) : '';
    const response = await apiClient.get(`/admin/assets?${params}`);
    return response.data;
  },

  uploadAsset: async (file: File, category: string) => {
    const formData = new FormData();
    formData.append('file', file);
    formData.append('category', category);
    const response = await apiClient.post('/admin/assets', formData, {
      headers: { 'Content-Type': 'multipart/form-data' },
    });
    return response.data;
  },

  deleteAsset: async (assetId: string) => {
    const response = await apiClient.delete(`/admin/assets/${assetId}`);
    return response.data;
  },
};
