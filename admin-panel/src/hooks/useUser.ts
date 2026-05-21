import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query';
import { userService } from '../services/userService';

export function useUsers(page: number, limit: number, filters?: Record<string, string>) {
  return useQuery({
    queryKey: ['users', page, limit, filters],
    queryFn: () => userService.getUsers(page, limit, filters),
  });
}

export function useUser(userId: string) {
  return useQuery({
    queryKey: ['user', userId],
    queryFn: () => userService.getUser(userId),
    enabled: !!userId,
  });
}

export function useCreateUser() {
  const queryClient = useQueryClient();
  
  return useMutation({
    mutationFn: userService.createUser,
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['users'] });
    },
  });
}

export function useUpdateUser(userId: string) {
  const queryClient = useQueryClient();
  
  return useMutation({
    mutationFn: (data: any) => userService.updateUser(userId, data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['users'] });
      queryClient.invalidateQueries({ queryKey: ['user', userId] });
    },
  });
}

export function useDeleteUser() {
  const queryClient = useQueryClient();
  
  return useMutation({
    mutationFn: userService.deleteUser,
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['users'] });
    },
  });
}

export function useSuspendUser() {
  const queryClient = useQueryClient();
  
  return useMutation({
    mutationFn: ({ userId, reason }: { userId: string; reason: string }) => 
      userService.suspendUser(userId, reason),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['users'] });
    },
  });
}

export function useActivateUser() {
  const queryClient = useQueryClient();
  
  return useMutation({
    mutationFn: userService.activateUser,
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['users'] });
    },
  });
}
