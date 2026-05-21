import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query';
import { bonusService, Bonus, BonusCampaign } from '../services/bonusService';

export function useBonuses(page: number, limit: number, filters?: Record<string, string>) {
  return useQuery({
    queryKey: ['bonuses', page, limit, filters],
    queryFn: () => bonusService.getAllBonuses(page, limit, filters),
  });
}

export function useBonus(bonusId: string) {
  return useQuery({
    queryKey: ['bonus', bonusId],
    queryFn: () => bonusService.getBonus(bonusId),
    enabled: !!bonusId,
  });
}

export function useCreateBonus() {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (bonus: Partial<Bonus>) => bonusService.createBonus(bonus),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['bonuses'] });
    },
  });
}

export function useUpdateBonus() {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: ({ bonusId, bonus }: { bonusId: string; bonus: Partial<Bonus> }) =>
      bonusService.updateBonus(bonusId, bonus),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['bonuses'] });
    },
  });
}

export function useDeleteBonus() {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (bonusId: string) => bonusService.deleteBonus(bonusId),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['bonuses'] });
    },
  });
}

export function useActivateBonus() {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (bonusId: string) => bonusService.activateBonus(bonusId),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['bonuses'] });
    },
  });
}

export function useDeactivateBonus() {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (bonusId: string) => bonusService.deactivateBonus(bonusId),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['bonuses'] });
    },
  });
}

export function useBonusCampaigns(page: number, limit: number) {
  return useQuery({
    queryKey: ['bonusCampaigns', page, limit],
    queryFn: () => bonusService.getCampaigns(page, limit),
  });
}

export function useCreateCampaign() {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (campaign: Partial<BonusCampaign>) => bonusService.createCampaign(campaign),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['bonusCampaigns'] });
    },
  });
}

export function usePlayerBonuses(userId: string, page: number, limit: number) {
  return useQuery({
    queryKey: ['playerBonuses', userId, page, limit],
    queryFn: () => bonusService.getPlayerBonuses(userId, page, limit),
    enabled: !!userId,
  });
}

export function useGrantBonus() {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: ({ userId, bonusId }: { userId: string; bonusId: string }) =>
      bonusService.grantBonus(userId, bonusId),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['playerBonuses'] });
    },
  });
}

export function useBonusPerformance(bonusId: string, startDate: string, endDate: string) {
  return useQuery({
    queryKey: ['bonusPerformance', bonusId, startDate, endDate],
    queryFn: () => bonusService.getBonusPerformance(bonusId, startDate, endDate),
    enabled: !!bonusId && !!startDate && !!endDate,
  });
}
