import { clsx, type ClassValue } from 'clsx';
import { twMerge } from 'tailwind-merge';

export function cn(...inputs: ClassValue[]) {
  return twMerge(clsx(inputs));
}

export function formatCurrency(amount: number, currency: string = 'USD'): string {
  return new Intl.NumberFormat('en-US', {
    style: 'currency',
    currency,
  }).format(amount);
}

export function formatDate(date: string | Date): string {
  return new Intl.DateTimeFormat('en-US', {
    year: 'numeric',
    month: 'short',
    day: 'numeric',
    hour: '2-digit',
    minute: '2-digit',
  }).format(new Date(date));
}

export function formatRelativeTime(date: string | Date): string {
  const now = new Date();
  const then = new Date(date);
  const diffInSeconds = Math.floor((now.getTime() - then.getTime()) / 1000);

  if (diffInSeconds < 60) return 'just now';
  if (diffInSeconds < 3600) return `${Math.floor(diffInSeconds / 60)}m ago`;
  if (diffInSeconds < 86400) return `${Math.floor(diffInSeconds / 3600)}h ago`;
  if (diffInSeconds < 604800) return `${Math.floor(diffInSeconds / 86400)}d ago`;
  return formatDate(date);
}

export function getStatusColor(status: string): string {
  const colors: Record<string, string> = {
    ACTIVE: 'bg-success-100 text-success-800',
    INACTIVE: 'bg-gray-100 text-gray-800',
    SUSPENDED: 'bg-danger-100 text-danger-800',
    PENDING: 'bg-yellow-100 text-yellow-800',
    COMPLETED: 'bg-success-100 text-success-800',
    FAILED: 'bg-danger-100 text-danger-800',
    CANCELLED: 'bg-gray-100 text-gray-800',
    MAINTENANCE: 'bg-yellow-100 text-yellow-800',
  };
  return colors[status] || 'bg-gray-100 text-gray-800';
}

export function getRoleBadgeColor(role: string): string {
  const colors: Record<string, string> = {
    SUPER_ADMIN: 'bg-purple-100 text-purple-800',
    ADMIN: 'bg-blue-100 text-blue-800',
    SUPPORT: 'bg-green-100 text-green-800',
    AGENT: 'bg-orange-100 text-orange-800',
  };
  return colors[role] || 'bg-gray-100 text-gray-800';
}
