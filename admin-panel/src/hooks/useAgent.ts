import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query';
import { agentService, Agent } from '../services/agentService';

export function useAgents(page: number, limit: number, filters?: Record<string, string>) {
  return useQuery({
    queryKey: ['agents', page, limit, filters],
    queryFn: () => agentService.getAllAgents(page, limit, filters),
  });
}

export function useAgent(agentId: string) {
  return useQuery({
    queryKey: ['agent', agentId],
    queryFn: () => agentService.getAgent(agentId),
    enabled: !!agentId,
  });
}

export function useCreateAgent() {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (agent: Partial<Agent>) => agentService.createAgent(agent),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['agents'] });
    },
  });
}

export function useUpdateAgent() {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: ({ agentId, agent }: { agentId: string; agent: Partial<Agent> }) =>
      agentService.updateAgent(agentId, agent),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['agents'] });
    },
  });
}

export function useSuspendAgent() {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: ({ agentId, reason }: { agentId: string; reason: string }) =>
      agentService.suspendAgent(agentId, reason),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['agents'] });
    },
  });
}

export function useActivateAgent() {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (agentId: string) => agentService.activateAgent(agentId),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['agents'] });
    },
  });
}

export function useAgentHierarchy(agentId: string) {
  return useQuery({
    queryKey: ['agentHierarchy', agentId],
    queryFn: () => agentService.getAgentHierarchy(agentId),
    enabled: !!agentId,
  });
}

export function useAgentPlayers(agentId: string, page: number, limit: number) {
  return useQuery({
    queryKey: ['agentPlayers', agentId, page, limit],
    queryFn: () => agentService.getAgentPlayers(agentId, page, limit),
    enabled: !!agentId,
  });
}

export function useCommissionPlans() {
  return useQuery({
    queryKey: ['commissionPlans'],
    queryFn: () => agentService.getCommissionPlans(),
  });
}

export function useCreateCommissionPlan() {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (plan: Partial<any>) => agentService.createCommissionPlan(plan),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['commissionPlans'] });
    },
  });
}

export function useUpdateCommissionPlan() {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: ({ planId, plan }: { planId: string; plan: Partial<any> }) =>
      agentService.updateCommissionPlan(planId, plan),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['commissionPlans'] });
    },
  });
}

export function useCommissionReports(agentId?: string, startDate?: string, endDate?: string) {
  return useQuery({
    queryKey: ['commissionReports', agentId, startDate, endDate],
    queryFn: () => agentService.getCommissionReports(agentId, startDate, endDate),
  });
}

export function useApproveCommission() {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (reportId: string) => agentService.approveCommission(reportId),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['commissionReports'] });
    },
  });
}

export function usePayCommission() {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (reportId: string) => agentService.payCommission(reportId),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['commissionReports'] });
    },
  });
}

export function useAgentStatistics(agentId: string, startDate: string, endDate: string) {
  return useQuery({
    queryKey: ['agentStatistics', agentId, startDate, endDate],
    queryFn: () => agentService.getAgentStatistics(agentId, startDate, endDate),
    enabled: !!agentId && !!startDate && !!endDate,
  });
}
