import apiClient from './api';

export interface ComplianceAlert {
  id: string;
  type: 'AML' | 'FRAUD' | 'SANCTIONS' | 'PEP' | 'LARGE_TRANSACTION' | 'SUSPICIOUS_PATTERN';
  severity: 'LOW' | 'MEDIUM' | 'HIGH' | 'CRITICAL';
  userId: string;
  userName: string;
  description: string;
  status: 'NEW' | 'UNDER_REVIEW' | 'ESCALATED' | 'RESOLVED' | 'FALSE_POSITIVE';
  assignedTo?: string;
  createdAt: string;
  updatedAt?: string;
  resolvedAt?: string;
  metadata: Record<string, any>;
}

export interface ComplianceCase {
  id: string;
  caseNumber: string;
  type: 'AML_INVESTIGATION' | 'FRAUD_INVESTIGATION' | 'SAR_FILING' | 'SANCTIONS_REVIEW';
  relatedAlerts: string[];
  userId?: string;
  title: string;
  description: string;
  status: 'OPEN' | 'IN_PROGRESS' | 'PENDING_REVIEW' | 'CLOSED' | 'FILED';
  priority: 'LOW' | 'MEDIUM' | 'HIGH' | 'URGENT';
  assignedTo: string;
  createdAt: string;
  updatedAt: string;
  closedAt?: string;
  notes: CaseNote[];
  attachments: Attachment[];
}

export interface CaseNote {
  id: string;
  content: string;
  authorId: string;
  authorName: string;
  createdAt: string;
}

export interface Attachment {
  id: string;
  fileName: string;
  fileType: string;
  fileSize: number;
  url: string;
  uploadedAt: string;
}

export interface RiskScore {
  userId: string;
  score: number;
  level: 'LOW' | 'MEDIUM' | 'HIGH' | 'CRITICAL';
  factors: RiskFactor[];
  lastUpdated: string;
}

export interface RiskFactor {
  category: string;
  factor: string;
  weight: number;
  score: number;
  description: string;
}

export const complianceService = {
  getAlerts: async (page = 1, limit = 20, filters?: Record<string, string>) => {
    const params = new URLSearchParams({ page: String(page), limit: String(limit) });
    if (filters) {
      Object.entries(filters).forEach(([key, value]) => {
        params.append(key, value);
      });
    }
    const response = await apiClient.get(`/admin/compliance/alerts?${params}`);
    return response.data;
  },

  getAlert: async (alertId: string) => {
    const response = await apiClient.get(`/admin/compliance/alerts/${alertId}`);
    return response.data;
  },

  updateAlertStatus: async (alertId: string, status: ComplianceAlert['status'], notes?: string) => {
    const response = await apiClient.put(`/admin/compliance/alerts/${alertId}/status`, { status, notes });
    return response.data;
  },

  assignAlert: async (alertId: string, assignedTo: string) => {
    const response = await apiClient.post(`/admin/compliance/alerts/${alertId}/assign`, { assignedTo });
    return response.data;
  },

  escalateAlert: async (alertId: string, reason: string) => {
    const response = await apiClient.post(`/admin/compliance/alerts/${alertId}/escalate`, { reason });
    return response.data;
  },

  getCases: async (page = 1, limit = 20, filters?: Record<string, string>) => {
    const params = new URLSearchParams({ page: String(page), limit: String(limit) });
    if (filters) {
      Object.entries(filters).forEach(([key, value]) => {
        params.append(key, value);
      });
    }
    const response = await apiClient.get(`/admin/compliance/cases?${params}`);
    return response.data;
  },

  createCase: async (caseData: Partial<ComplianceCase>) => {
    const response = await apiClient.post('/admin/compliance/cases', caseData);
    return response.data;
  },

  getCase: async (caseId: string) => {
    const response = await apiClient.get(`/admin/compliance/cases/${caseId}`);
    return response.data;
  },

  updateCase: async (caseId: string, caseData: Partial<ComplianceCase>) => {
    const response = await apiClient.put(`/admin/compliance/cases/${caseId}`, caseData);
    return response.data;
  },

  closeCase: async (caseId: string, resolution: string) => {
    const response = await apiClient.post(`/admin/compliance/cases/${caseId}/close`, { resolution });
    return response.data;
  },

  addCaseNote: async (caseId: string, content: string) => {
    const response = await apiClient.post(`/admin/compliance/cases/${caseId}/notes`, { content });
    return response.data;
  },

  uploadCaseAttachment: async (caseId: string, file: File) => {
    const formData = new FormData();
    formData.append('file', file);
    const response = await apiClient.post(`/admin/compliance/cases/${caseId}/attachments`, formData, {
      headers: { 'Content-Type': 'multipart/form-data' },
    });
    return response.data;
  },

  getRiskScores: async (page = 1, limit = 20, filters?: Record<string, string>) => {
    const params = new URLSearchParams({ page: String(page), limit: String(limit) });
    if (filters) {
      Object.entries(filters).forEach(([key, value]) => {
        params.append(key, value);
      });
    }
    const response = await apiClient.get(`/admin/compliance/risk-scores?${params}`);
    return response.data;
  },

  getRiskScore: async (userId: string) => {
    const response = await apiClient.get(`/admin/compliance/risk-scores/${userId}`);
    return response.data;
  },

  recalculateRiskScore: async (userId: string) => {
    const response = await apiClient.post(`/admin/compliance/risk-scores/${userId}/recalculate`);
    return response.data;
  },

  getComplianceDashboard: async () => {
    const response = await apiClient.get('/admin/compliance/dashboard');
    return response.data;
  },

  exportSAR: async (caseId: string) => {
    const response = await apiClient.post(`/admin/compliance/cases/${caseId}/export-sar`);
    return response.data;
  },

  getSanctionsScreening: async (userId: string) => {
    const response = await apiClient.get(`/admin/compliance/sanctions/${userId}`);
    return response.data;
  },

  runSanctionsScreening: async (userId: string) => {
    const response = await apiClient.post(`/admin/compliance/sanctions/${userId}/screen`);
    return response.data;
  },
};
