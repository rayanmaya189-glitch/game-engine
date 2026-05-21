import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query';
import { complianceService, ComplianceAlert, ComplianceCase } from '../services/complianceService';

export function useComplianceAlerts(page: number, limit: number, filters?: Record<string, string>) {
  return useQuery({
    queryKey: ['complianceAlerts', page, limit, filters],
    queryFn: () => complianceService.getAlerts(page, limit, filters),
  });
}

export function useComplianceAlert(alertId: string) {
  return useQuery({
    queryKey: ['complianceAlert', alertId],
    queryFn: () => complianceService.getAlert(alertId),
    enabled: !!alertId,
  });
}

export function useUpdateAlertStatus() {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: ({ alertId, status, notes }: { alertId: string; status: ComplianceAlert['status']; notes?: string }) =>
      complianceService.updateAlertStatus(alertId, status, notes),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['complianceAlerts'] });
    },
  });
}

export function useAssignAlert() {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: ({ alertId, assignedTo }: { alertId: string; assignedTo: string }) =>
      complianceService.assignAlert(alertId, assignedTo),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['complianceAlerts'] });
    },
  });
}

export function useEscalateAlert() {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: ({ alertId, reason }: { alertId: string; reason: string }) =>
      complianceService.escalateAlert(alertId, reason),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['complianceAlerts'] });
    },
  });
}

export function useComplianceCases(page: number, limit: number, filters?: Record<string, string>) {
  return useQuery({
    queryKey: ['complianceCases', page, limit, filters],
    queryFn: () => complianceService.getCases(page, limit, filters),
  });
}

export function useCreateCase() {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (caseData: Partial<ComplianceCase>) => complianceService.createCase(caseData),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['complianceCases'] });
    },
  });
}

export function useComplianceCase(caseId: string) {
  return useQuery({
    queryKey: ['complianceCase', caseId],
    queryFn: () => complianceService.getCase(caseId),
    enabled: !!caseId,
  });
}

export function useUpdateCase() {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: ({ caseId, caseData }: { caseId: string; caseData: Partial<ComplianceCase> }) =>
      complianceService.updateCase(caseId, caseData),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['complianceCases'] });
    },
  });
}

export function useCloseCase() {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: ({ caseId, resolution }: { caseId: string; resolution: string }) =>
      complianceService.closeCase(caseId, resolution),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['complianceCases'] });
    },
  });
}

export function useAddCaseNote() {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: ({ caseId, content }: { caseId: string; content: string }) =>
      complianceService.addCaseNote(caseId, content),
    onSuccess: (_, variables) => {
      queryClient.invalidateQueries({ queryKey: ['complianceCase', variables.caseId] });
    },
  });
}

export function useRiskScores(page: number, limit: number, filters?: Record<string, string>) {
  return useQuery({
    queryKey: ['riskScores', page, limit, filters],
    queryFn: () => complianceService.getRiskScores(page, limit, filters),
  });
}

export function useRiskScore(userId: string) {
  return useQuery({
    queryKey: ['riskScore', userId],
    queryFn: () => complianceService.getRiskScore(userId),
    enabled: !!userId,
  });
}

export function useRecalculateRiskScore() {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (userId: string) => complianceService.recalculateRiskScore(userId),
    onSuccess: (_, userId) => {
      queryClient.invalidateQueries({ queryKey: ['riskScore', userId] });
    },
  });
}

export function useComplianceDashboard() {
  return useQuery({
    queryKey: ['complianceDashboard'],
    queryFn: () => complianceService.getComplianceDashboard(),
  });
}

export function useExportSAR() {
  return useMutation({
    mutationFn: (caseId: string) => complianceService.exportSAR(caseId),
  });
}
