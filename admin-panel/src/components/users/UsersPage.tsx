import { useState } from 'react';
import { useUsers, useDeleteUser, useSuspendUser, useActivateUser } from '../hooks/useUser';
import { formatDate, getStatusColor, getRoleBadgeColor } from '../hooks/useUtils';

export function UsersPage() {
  const [page, setPage] = useState(1);
  const [filters, setFilters] = useState<Record<string, string>>({});
  const { data, isLoading } = useUsers(page, 20, filters);
  const deleteUser = useDeleteUser();
  const suspendUser = useSuspendUser();
  const activateUser = useActivateUser();

  const handleDelete = async (userId: string) => {
    if (confirm('Are you sure you want to delete this user?')) {
      await deleteUser.mutateAsync(userId);
    }
  };

  const handleSuspend = async (userId: string) => {
    const reason = prompt('Enter suspension reason:');
    if (reason) {
      await suspendUser.mutateAsync({ userId, reason });
    }
  };

  const handleActivate = async (userId: string) => {
    await activateUser.mutateAsync(userId);
  };

  if (isLoading) {
    return <div className="text-center py-8">Loading...</div>;
  }

  return (
    <div className="space-y-6">
      <div className="flex justify-between items-center">
        <h1 className="text-2xl font-bold text-gray-900">User Management</h1>
        <button className="btn-primary">Add User</button>
      </div>

      {/* Filters */}
      <div className="card">
        <div className="grid grid-cols-1 gap-4 sm:grid-cols-3">
          <input
            type="text"
            placeholder="Search by username..."
            className="input-field"
            onChange={(e) => setFilters({ ...filters, search: e.target.value })}
          />
          <select
            className="input-field"
            onChange={(e) => setFilters({ ...filters, status: e.target.value })}
          >
            <option value="">All Status</option>
            <option value="ACTIVE">Active</option>
            <option value="INACTIVE">Inactive</option>
            <option value="SUSPENDED">Suspended</option>
          </select>
          <select
            className="input-field"
            onChange={(e) => setFilters({ ...filters, role: e.target.value })}
          >
            <option value="">All Roles</option>
            <option value="ADMIN">Admin</option>
            <option value="SUPER_ADMIN">Super Admin</option>
            <option value="SUPPORT">Support</option>
            <option value="AGENT">Agent</option>
          </select>
        </div>
      </div>

      {/* Users Table */}
      <div className="card table-container">
        <table className="table">
          <thead className="table-header">
            <tr>
              <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Username</th>
              <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Email</th>
              <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Role</th>
              <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Status</th>
              <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Last Login</th>
              <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Actions</th>
            </tr>
          </thead>
          <tbody className="bg-white divide-y divide-gray-200">
            {data?.users?.map((user: any) => (
              <tr key={user.id} className="table-row">
                <td className="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">{user.username}</td>
                <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">{user.email}</td>
                <td className="px-6 py-4 whitespace-nowrap">
                  <span className={`px-2 py-1 text-xs rounded-full ${getRoleBadgeColor(user.role)}`}>
                    {user.role}
                  </span>
                </td>
                <td className="px-6 py-4 whitespace-nowrap">
                  <span className={`px-2 py-1 text-xs rounded-full ${getStatusColor(user.status)}`}>
                    {user.status}
                  </span>
                </td>
                <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                  {user.lastLoginAt ? formatDate(user.lastLoginAt) : 'Never'}
                </td>
                <td className="px-6 py-4 whitespace-nowrap text-sm space-x-2">
                  <button className="text-primary-600 hover:text-primary-900">Edit</button>
                  {user.status === 'ACTIVE' ? (
                    <button 
                      className="text-danger-600 hover:text-danger-900"
                      onClick={() => handleSuspend(user.id)}
                    >
                      Suspend
                    </button>
                  ) : (
                    <button 
                      className="text-success-600 hover:text-success-900"
                      onClick={() => handleActivate(user.id)}
                    >
                      Activate
                    </button>
                  )}
                  <button 
                    className="text-danger-600 hover:text-danger-900"
                    onClick={() => handleDelete(user.id)}
                  >
                    Delete
                  </button>
                </td>
              </tr>
            ))}
          </tbody>
        </table>

        {/* Pagination */}
        <div className="flex justify-between items-center mt-4">
          <button
            onClick={() => setPage(Math.max(1, page - 1))}
            disabled={page === 1}
            className="btn-primary disabled:opacity-50"
          >
            Previous
          </button>
          <span className="text-sm text-gray-600">Page {page}</span>
          <button
            onClick={() => setPage(page + 1)}
            disabled={!data?.hasMore}
            className="btn-primary disabled:opacity-50"
          >
            Next
          </button>
        </div>
      </div>
    </div>
  );
}
