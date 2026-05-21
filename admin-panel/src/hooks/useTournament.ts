import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query';
import { tournamentService, Tournament } from '../services/tournamentService';

export function useTournaments(page: number, limit: number, filters?: Record<string, string>) {
  return useQuery({
    queryKey: ['tournaments', page, limit, filters],
    queryFn: () => tournamentService.getAllTournaments(page, limit, filters),
  });
}

export function useTournament(tournamentId: string) {
  return useQuery({
    queryKey: ['tournament', tournamentId],
    queryFn: () => tournamentService.getTournament(tournamentId),
    enabled: !!tournamentId,
  });
}

export function useCreateTournament() {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (tournament: Partial<Tournament>) => tournamentService.createTournament(tournament),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['tournaments'] });
    },
  });
}

export function useUpdateTournament() {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: ({ tournamentId, tournament }: { tournamentId: string; tournament: Partial<Tournament> }) =>
      tournamentService.updateTournament(tournamentId, tournament),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['tournaments'] });
    },
  });
}

export function useCancelTournament() {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: ({ tournamentId, reason }: { tournamentId: string; reason: string }) =>
      tournamentService.cancelTournament(tournamentId, reason),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['tournaments'] });
    },
  });
}

export function useStartTournament() {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (tournamentId: string) => tournamentService.startTournament(tournamentId),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['tournaments'] });
    },
  });
}

export function useTournamentPlayers(tournamentId: string) {
  return useQuery({
    queryKey: ['tournamentPlayers', tournamentId],
    queryFn: () => tournamentService.getTournamentPlayers(tournamentId),
    enabled: !!tournamentId,
  });
}

export function useRegisterPlayer() {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: ({ tournamentId, userId }: { tournamentId: string; userId: string }) =>
      tournamentService.registerPlayer(tournamentId, userId),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['tournamentPlayers'] });
    },
  });
}

export function useTournamentLeaderboard(tournamentId: string) {
  return useQuery({
    queryKey: ['tournamentLeaderboard', tournamentId],
    queryFn: () => tournamentService.getLeaderboard(tournamentId),
    enabled: !!tournamentId,
  });
}

export function useDistributePrizes() {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (tournamentId: string) => tournamentService.distributePrizes(tournamentId),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['tournaments'] });
    },
  });
}

export function useTournamentHistory(page: number, limit: number, filters?: Record<string, string>) {
  return useQuery({
    queryKey: ['tournamentHistory', page, limit, filters],
    queryFn: () => tournamentService.getTournamentHistory(page, limit, filters),
  });
}

export function useTournamentStatistics(tournamentId: string) {
  return useQuery({
    queryKey: ['tournamentStatistics', tournamentId],
    queryFn: () => tournamentService.getTournamentStatistics(tournamentId),
    enabled: !!tournamentId,
  });
}
