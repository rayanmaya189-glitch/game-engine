import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query';
import { financialService } from '../services/financialService';

export function useTransactions(page: number, limit: number, filters?: Record<string, string>) {
  return useQuery({
    queryKey: ['transactions', page, limit, filters],
    queryFn: () => financialService.getTransactions(page, limit, filters),
  });
}

export function useTransaction(transactionId: string) {
  return useQuery({
    queryKey: ['transaction', transactionId],
    queryFn: () => financialService.getTransaction(transactionId),
    enabled: !!transactionId,
  });
}

export function useApproveWithdrawal() {
  const queryClient = useQueryClient();
  
  return useMutation({
    mutationFn: financialService.approveWithdrawal,
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['transactions'] });
    },
  });
}

export function useRejectWithdrawal() {
  const queryClient = useQueryClient();
  
  return useMutation({
    mutationFn: ({ transactionId, reason }: { transactionId: string; reason: string }) => 
      financialService.rejectWithdrawal(transactionId, reason),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['transactions'] });
    },
  });
}

export function useWallet(userId: string) {
  return useQuery({
    queryKey: ['wallet', userId],
    queryFn: () => financialService.getWallet(userId),
    enabled: !!userId,
  });
}

export function useAdjustBalance() {
  const queryClient = useQueryClient();
  
  return useMutation({
    mutationFn: ({ userId, amount, reason, type }: { userId: string; amount: number; reason: string; type: 'CREDIT' | 'DEBIT' }) => 
      financialService.adjustBalance(userId, amount, reason, type),
    onSuccess: (_, variables) => {
      queryClient.invalidateQueries({ queryKey: ['wallet', variables.userId] });
      queryClient.invalidateQueries({ queryKey: ['transactions'] });
    },
  });
}

export function useRevenueStats(period: 'DAY' | 'WEEK' | 'MONTH' | 'YEAR') {
  return useQuery({
    queryKey: ['revenueStats', period],
    queryFn: () => financialService.getRevenueStats(period),
  });
}
