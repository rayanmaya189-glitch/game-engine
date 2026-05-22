import apiClient from './api';

export interface ReportData {
  type: 'FINANCIAL' | 'PLAYER' | 'GAME' | 'COMPLIANCE';
  startDate: string;
  endDate: string;
  data: Record<string, any>;
  generatedAt: string;
}

export const reportService = {
  getReport: async (type: string, startDate: string, endDate: string) => {
    const params = new URLSearchParams({ type, startDate, endDate });
    const response = await apiClient.get(`/admin/reports?${params}`);
    return response.data;
  },

  exportReport: async (type: string, startDate: string, endDate: string, format: 'PDF' | 'CSV' | 'XLSX') => {
    const params = new URLSearchParams({ type, startDate, endDate, format });
    const response = await apiClient.get(`/admin/reports/export?${params}`, {
      responseType: 'blob',
    });
    // Trigger download
    const url = window.URL.createObjectURL(new Blob([response.data]));
    const link = document.createElement('a');
    link.href = url;
    link.setAttribute('download', `report_${type}_${startDate}_${endDate}.${format.toLowerCase()}`);
    document.body.appendChild(link);
    link.click();
    link.remove();
    return response.data;
  },

  getScheduledReports: async () => {
    const response = await apiClient.get('/admin/reports/scheduled');
    return response.data;
  },

  scheduleReport: async (reportConfig: {
    type: string;
    frequency: 'DAILY' | 'WEEKLY' | 'MONTHLY';
    recipients: string[];
    format: 'PDF' | 'CSV' | 'XLSX';
  }) => {
    const response = await apiClient.post('/admin/reports/schedule', reportConfig);
    return response.data;
  },

  deleteScheduledReport: async (reportId: string) => {
    const response = await apiClient.delete(`/admin/reports/scheduled/${reportId}`);
    return response.data;
  },
};
