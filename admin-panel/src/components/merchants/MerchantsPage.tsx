import { useState } from 'react';
import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query';
import { merchantService, Merchant } from '../../services/merchantService';
import { PlusIcon, XCircleIcon } from '@heroicons/react/24/outline';

export function MerchantsPage() {
  const [page, setPage] = useState(1);
  const [limit] = useState(20);
  const [showCreateModal, setShowCreateModal] = useState(false);
  const queryClient = useQueryClient();

  const { data: merchantsData, isLoading } = useQuery({
    queryKey: ['merchants', page, limit],
    queryFn: () => merchantService.getAllMerchants(page, limit),
  });

  const createMutation = useMutation({
    mutationFn: (merchant: Partial<Merchant>) => merchantService.createMerchant(merchant),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['merchants'] });
      setShowCreateModal(false);
    },
  });

  const deleteMutation = useMutation({
    mutationFn: (merchantId: string) => merchantService.deleteMerchant(merchantId),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['merchants'] });
    },
  });

  const suspendMutation = useMutation({
    mutationFn: ({ id, reason }: { id: string; reason: string }) => 
      merchantService.suspendMerchant(id, reason),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['merchants'] });
    },
  });

  const activateMutation = useMutation({
    mutationFn: (merchantId: string) => merchantService.activateMerchant(merchantId),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['merchants'] });
    },
  });

  const handleCreate = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    const formData = new FormData(e.currentTarget);
    createMutation.mutate({
      name: formData.get('name') as string,
      email: formData.get('email') as string,
      phone: formData.get('phone') as string,
      currency: 'USD',
      revenueShare: Number(formData.get('revenueShare')),
      whiteLabelDomain: formData.get('whiteLabelDomain') as string,
    });
  };

  const handleDelete = (id: string) => {
    if (confirm('Are you sure you want to delete this merchant?')) {
      deleteMutation.mutate(id);
    }
  };

  const handleSuspend = (id: string) => {
    const reason = prompt('Enter suspension reason:');
    if (reason) {
      suspendMutation.mutate({ id, reason });
    }
  };

  const merchants = merchantsData?.data?.merchants || [];
  const total = merchantsData?.data?.total || 0;

  return (
    <div className="space-y-6">
      <div className="flex justify-between items-center">
        <h1 className="text-2xl font-bold text-gray-900">Merchants</h1>
        <button
          onClick={() => setShowCreateModal(true)}
          className="flex items-center px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700"
        >
          <PlusIcon className="h-5 w-5 mr-2" />
          Create Merchant
        </button>
      </div>

      {/* Merchants List */}
      <div className="bg-white shadow rounded-lg overflow-hidden">
        <table className="min-w-full divide-y divide-gray-200">
          <thead className="bg-gray-50">
            <tr>
              <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Name</th>
              <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Email</th>
              <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Domain</th>
              <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Players</th>
              <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Revenue Share</th>
              <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">GGR</th>
              <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Status</th>
              <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Actions</th>
            </tr>
          </thead>
          <tbody className="bg-white divide-y divide-gray-200">
            {isLoading ? (
              <tr>
                <td colSpan={8} className="px-6 py-4 text-center text-gray-500">Loading...</td>
              </tr>
            ) : merchants.length === 0 ? (
              <tr>
                <td colSpan={8} className="px-6 py-4 text-center text-gray-500">No merchants found</td>
              </tr>
            ) : (
              merchants.map((merchant: Merchant) => (
                <tr key={merchant.id} className="hover:bg-gray-50">
                  <td className="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">
                    {merchant.name}
                  </td>
                  <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                    {merchant.email}
                  </td>
                  <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                    {merchant.whiteLabelDomain}
                  </td>
                  <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                    {merchant.totalPlayers}
                  </td>
                  <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                    {merchant.revenueShare}%
                  </td>
                  <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                    ${merchant.totalGGR.toFixed(2)}
                  </td>
                  <td className="px-6 py-4 whitespace-nowrap">
                    <span className={`px-2 py-1 text-xs rounded-full ${
                      merchant.status === 'ACTIVE' ? 'bg-green-100 text-green-800' :
                      merchant.status === 'SUSPENDED' ? 'bg-red-100 text-red-800' :
                      'bg-yellow-100 text-yellow-800'
                    }`}>
                      {merchant.status}
                    </span>
                  </td>
                  <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500 space-x-2">
                    {merchant.status === 'ACTIVE' ? (
                      <button
                        onClick={() => handleSuspend(merchant.id)}
                        className="text-yellow-600 hover:text-yellow-900"
                      >
                        Suspend
                      </button>
                    ) : (
                      <button
                        onClick={() => activateMutation.mutate(merchant.id)}
                        className="text-green-600 hover:text-green-900"
                      >
                        Activate
                      </button>
                    )}
                    <button
                      onClick={() => handleDelete(merchant.id)}
                      className="text-red-600 hover:text-red-900"
                    >
                      Delete
                    </button>
                  </td>
                </tr>
              ))
            )}
          </tbody>
        </table>
      </div>

      {/* Pagination */}
      <div className="flex justify-between items-center">
        <button
          onClick={() => setPage(p => Math.max(1, p - 1))}
          disabled={page === 1}
          className="px-4 py-2 border rounded-lg disabled:opacity-50"
        >
          Previous
        </button>
        <span className="text-sm text-gray-600">Page {page} - Total: {total}</span>
        <button
          onClick={() => setPage(p => p + 1)}
          disabled={merchants.length < limit}
          className="px-4 py-2 border rounded-lg disabled:opacity-50"
        >
          Next
        </button>
      </div>

      {/* Create Merchant Modal */}
      {showCreateModal && (
        <div className="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
          <div className="bg-white rounded-lg p-6 max-w-md w-full">
            <div className="flex justify-between items-center mb-4">
              <h2 className="text-xl font-bold">Create Merchant</h2>
              <button onClick={() => setShowCreateModal(false)} className="text-gray-500 hover:text-gray-700">
                <XCircleIcon className="h-6 w-6" />
              </button>
            </div>
            <form onSubmit={handleCreate} className="space-y-4">
              <div>
                <label className="block text-sm font-medium text-gray-700">Name</label>
                <input name="name" required className="mt-1 block w-full border border-gray-300 rounded-md px-3 py-2" />
              </div>
              <div>
                <label className="block text-sm font-medium text-gray-700">Email</label>
                <input name="email" type="email" required className="mt-1 block w-full border border-gray-300 rounded-md px-3 py-2" />
              </div>
              <div>
                <label className="block text-sm font-medium text-gray-700">Phone</label>
                <input name="phone" type="tel" className="mt-1 block w-full border border-gray-300 rounded-md px-3 py-2" />
              </div>
              <div>
                <label className="block text-sm font-medium text-gray-700">White Label Domain</label>
                <input name="whiteLabelDomain" placeholder="casino.example.com" className="mt-1 block w-full border border-gray-300 rounded-md px-3 py-2" />
              </div>
              <div>
                <label className="block text-sm font-medium text-gray-700">Revenue Share (%)</label>
                <input name="revenueShare" type="number" step="0.1" defaultValue={10} className="mt-1 block w-full border border-gray-300 rounded-md px-3 py-2" />
              </div>
              <div className="flex justify-end space-x-2">
                <button
                  type="button"
                  onClick={() => setShowCreateModal(false)}
                  className="px-4 py-2 border rounded-lg"
                >
                  Cancel
                </button>
                <button
                  type="submit"
                  disabled={createMutation.isPending}
                  className="px-4 py-2 bg-blue-600 text-white rounded-lg disabled:opacity-50"
                >
                  Create
                </button>
              </div>
            </form>
          </div>
        </div>
      )}
    </div>
  );
}
