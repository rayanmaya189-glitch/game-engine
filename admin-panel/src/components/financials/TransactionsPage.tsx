import { useState } from 'react';
import { useTransactions, useApproveWithdrawal, useRejectWithdrawal } from '../hooks/useFinancial';
import { formatDate, formatCurrency, getStatusColor } from '../hooks/useUtils';

export function TransactionsPage() {
  const [page, setPage] = useState(1);
  const [filters, setFilters] = useState<Record<string, string>>({});
  const { data, isLoading } = useTransactions(page, 20, filters);
  const approveWithdrawal = useApproveWithdrawal();
  const rejectWithdrawal = useRejectWithdrawal();

  const handleApprove = async (transactionId: string) => {
    if (confirm('Are you sure you want to approve this withdrawal?')) {
      await approveWithdrawal.mutateAsync(transactionId);
    }
  };

  const handleReject = async (transactionId: string) => {
    const reason = prompt('Enter rejection reason:');
    if (reason) {
      await rejectWithdrawal.mutateAsync({ transactionId, reason });
    }
  };

  if (isLoading) {
    return <div className="text-center py-8">Loading...</div>;
  }

  return (
    <div className="space-y-6">
      <h1 className="text-2xl font-bold text-gray-900">Transaction Management</h1>

      {/* Filters */}
      <div className="card">
        <div className="grid grid-cols-1 gap-4 sm:grid-cols-4">
          <input
            type="text"
            placeholder="Search by ID or user..."
            className="input-field"
            onChange={(e) => setFilters({ ...filters, search: e.target.value })}
          />
          <select
            className="input-field"
            onChange={(e) => setFilters({ ...filters, type: e.target.value })}
          >
            <option value="">All Types</option>
            <option value="DEPOSIT">Deposit</option>
            <option value="WITHDRAWAL">Withdrawal</option>
            <option value="BET">Bet</option>
            <option value="WIN">Win</option>
            <option value="BONUS">Bonus</option>
          </select>
          <select
            className="input-field"
            onChange={(e) => setFilters({ ...filters, status: e.target.value })}
          >
            <option value="">All Status</option>
            <option value="PENDING">Pending</option>
            <option value="COMPLETED">Completed</option>
            <option value="FAILED">Failed</option>
            <option value="CANCELLED">Cancelled</option>
          </select>
          <button className="btn-primary">Apply Filters</button>
        </div>
      </div>

      {/* Transactions Table */}
      <div className="card table-container">
        <table className="table">
          <thead className="table-header">
            <tr>
              <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">ID</th>
              <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">User</th>
              <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Type</th>
              <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Amount</th>
              <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Status</th>
              <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Date</th>
              <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Actions</th>
            </tr>
          </thead>
          <tbody className="bg-white divide-y divide-gray-200">
            {data?.transactions?.map((tx: any) => (
              <tr key={tx.id} className="table-row">
                <td className="px-6 py-4 whitespace-nowrap text-sm font-mono text-gray-600">{tx.id.slice(0, 8)}...</td>
                <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{tx.userId}</td>
                <td className="px-6 py-4 whitespace-nowrap">
                  <span className={`px-2 py-1 text-xs rounded-full ${
                    tx.type === 'DEPOSIT' ? 'bg-green-100 text-green-800' :
                    tx.type === 'WITHDRAWAL' ? 'bg-red-100 text-red-800' :
                    tx.type === 'BET' ? 'bg-blue-100 text-blue-800' :
                    'bg-purple-100 text-purple-800'
                  }`}>
                    {tx.type}
                  </span>
                </td>
                <td className={`px-6 py-4 whitespace-nowrap text-sm font-medium ${
                  tx.type === 'WITHDRAWAL' || tx.type === 'BET' ? 'text-red-600' : 'text-green-600'
                }`}>
                  {tx.type === 'WITHDRAWAL' || tx.type === 'BET' ? '-' : '+'}{formatCurrency(tx.amount, tx.currency)}
                </td>
                <td className="px-6 py-4 whitespace-nowrap">
                  <span className={`px-2 py-1 text-xs rounded-full ${getStatusColor(tx.status)}`}>
                    {tx.status}
                  </span>
                </td>
                <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                  {formatDate(tx.createdAt)}
                </td>
                <td className="px-6 py-4 whitespace-nowrap text-sm space-x-2">
                  <button className="text-primary-600 hover:text-primary-900">View</button>
                  {tx.type === 'WITHDRAWAL' && tx.status === 'PENDING' && (
                    <>
                      <button 
                        className="text-success-600 hover:text-success-900"
                        onClick={() => handleApprove(tx.id)}
                      >
                        Approve
                      </button>
                      <button 
                        className="text-danger-600 hover:text-danger-900"
                        onClick={() => handleReject(tx.id)}
                      >
                        Reject
                      </button>
                    </>
                  )}
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
