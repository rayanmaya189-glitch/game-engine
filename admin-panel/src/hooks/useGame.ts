import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query';
import { gameService } from '../services/gameService';

export function useGames(page: number, limit: number, filters?: Record<string, string>) {
  return useQuery({
    queryKey: ['games', page, limit, filters],
    queryFn: () => gameService.getGames(page, limit, filters),
  });
}

export function useGame(gameId: string) {
  return useQuery({
    queryKey: ['game', gameId],
    queryFn: () => gameService.getGame(gameId),
    enabled: !!gameId,
  });
}

export function useCreateGame() {
  const queryClient = useQueryClient();
  
  return useMutation({
    mutationFn: gameService.createGame,
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['games'] });
    },
  });
}

export function useUpdateGame(gameId: string) {
  const queryClient = useQueryClient();
  
  return useMutation({
    mutationFn: (data: any) => gameService.updateGame(gameId, data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['games'] });
      queryClient.invalidateQueries({ queryKey: ['game', gameId] });
    },
  });
}

export function useDeleteGame() {
  const queryClient = useQueryClient();
  
  return useMutation({
    mutationFn: gameService.deleteGame,
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['games'] });
    },
  });
}

export function useToggleGameStatus() {
  const queryClient = useQueryClient();
  
  return useMutation({
    mutationFn: ({ gameId, status }: { gameId: string; status: 'ACTIVE' | 'INACTIVE' | 'MAINTENANCE' }) => 
      gameService.toggleGameStatus(gameId, status),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['games'] });
    },
  });
}

export function useGameStats(gameId: string, startDate?: string, endDate?: string) {
  return useQuery({
    queryKey: ['gameStats', gameId, startDate, endDate],
    queryFn: () => gameService.getGameStats(gameId, startDate, endDate),
    enabled: !!gameId,
  });
}
