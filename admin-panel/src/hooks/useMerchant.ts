import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query';
import { merchantService, Merchant } from '../services/merchantService';

export function useMerchants(page: number, limit: number, filters?: Record<string, string>) {
  return useQuery({
    queryKey: ['merchants', page, limit, filters],
    queryFn: () => merchantService.getAllMerchants(page, limit, filters),
  });
}

export function useMerchant(merchantId: string) {
  return useQuery({
    queryKey: ['merchant', merchantId],
    queryFn: () => merchantService.getMerchant(merchantId),
    enabled: !!merchantId,
  });
}

export function useCreateMerchant() {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (merchant: Partial<Merchant>) => merchantService.createMerchant(merchant),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['merchants'] });
    },
  });
}

export function useUpdateMerchant() {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: ({ merchantId, merchant }: { merchantId: string; merchant: Partial<Merchant> }) =>
      merchantService.updateMerchant(merchantId, merchant),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['merchants'] });
    },
  });
}

export function useSuspendMerchant() {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: ({ merchantId, reason }: { merchantId: string; reason: string }) =>
      merchantService.suspendMerchant(merchantId, reason),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['merchants'] });
    },
  });
}

export function useActivateMerchant() {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (merchantId: string) => merchantService.activateMerchant(merchantId),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['merchants'] });
    },
  });
}

export function useRegenerateApiKey() {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (merchantId: string) => merchantService.regenerateApiKey(merchantId),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['merchant', merchantId] });
    },
  });
}

export function useMerchantPlayers(merchantId: string, page: number, limit: number) {
  return useQuery({
    queryKey: ['merchantPlayers', merchantId, page, limit],
    queryFn: () => merchantService.getMerchantPlayers(merchantId, page, limit),
    enabled: !!merchantId,
  });
}

export function useMerchantReports(merchantId?: string, startDate?: string, endDate?: string) {
  return useQuery({
    queryKey: ['merchantReports', merchantId, startDate, endDate],
    queryFn: () => merchantService.getMerchantReports(merchantId, startDate, endDate),
  });
}

export function useApproveMerchantReport() {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (reportId: string) => merchantService.approveMerchantReport(reportId),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['merchantReports'] });
    },
  });
}

export function usePayMerchant() {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (reportId: string) => merchantService.payMerchant(reportId),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['merchantReports'] });
    },
  });
}

export function useMerchantStatistics(merchantId: string, startDate: string, endDate: string) {
  return useQuery({
    queryKey: ['merchantStatistics', merchantId, startDate, endDate],
    queryFn: () => merchantService.getMerchantStatistics(merchantId, startDate, endDate),
    enabled: !!merchantId && !!startDate && !!endDate,
  });
}

export function useUpdateWhiteLabelConfig() {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: ({ merchantId, config }: { merchantId: string; config: any }) =>
      merchantService.updateWhiteLabelConfig(merchantId, config),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['merchant'] });
    },
  });
}

export function useMerchantGames(merchantId: string) {
  return useQuery({
    queryKey: ['merchantGames', merchantId],
    queryFn: () => merchantService.getMerchantGames(merchantId),
    enabled: !!merchantId,
  });
}

export function useUpdateMerchantGames() {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: ({ merchantId, gameIds }: { merchantId: string; gameIds: string[] }) =>
      merchantService.updateMerchantGames(merchantId, gameIds),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['merchantGames'] });
    },
  });
}
