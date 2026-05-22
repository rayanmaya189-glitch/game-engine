import { useState } from 'react';
import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query';
import { jackpotService, Jackpot } from '../../services/jackpotService';
import { PlusIcon, XCircleIcon, CurrencyDollarIcon } from '@heroicons/react/24/outline';

export function JackpotsPage() {
  const [page, setPage] = useState(1);
  const [limit] = useState(20);
  const [showCreateModal, setShowCreateModal] = useState(false);
  const [selectedJackpot, setSelectedJackpot] = useState<Jackpot | null>(null);
  const queryClient = useQueryClient();

  const { data: jackpotsData, isLoading } = useQuery({
    queryKey: ['jackpots', page, limit],
    queryFn: () => jackpotService.getAllJackpots(page, limit),
  });

  const createMutation = useMutation({
    mutationFn: (jackpot: Partial<Jackpot>) => jackpotService.createJackpot(jackpot),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['jackpots'] });
      setShowCreateModal(false);
    },
  });

  const deleteMutation = useMutation({
    mutationFn: (jackpotId: string) => jackpotService.deleteJackpot(jackpotId),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['jackpots'] });
    },
  });

  const activateMutation = useMutation({
    mutationFn: (jackpotId: string) => jackpotService.activateJackpot(jackpotId),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['jackpots'] });
    },
  });

  const deactivateMutation = useMutation({
    mutationFn: (jackpotId: string) => jackpotService.deactivateJackpot(jackpotId),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['jackpots'] });
    },
  });

  const handleCreate = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    const formData = new FormData(e.currentTarget);
    createMutation.mutate({
      name: formData.get('name') as string,
      type: formData.get('type') as Jackpot['type'],
      currentAmount: Number(formData.get('currentAmount')),
      seedAmount: Number(formData.get('seedAmount')),
      maxAmount: Number(formData.get('maxAmount')) || undefined,
      contributionRate: Number(formData.get('contributionRate')),
      currency: 'USD',
      triggerType: formData.get('triggerType') as Jackpot['triggerType'],
      linkedGames: formData.get('linkedGames')?.toString().split(',').filter(Boolean) || [],
    });
  };

  const handleDelete = (id: string) => {
    if (confirm('Are you sure you want to delete this jackpot?')) {
      deleteMutation.mutate(id);
    }
  };

  const handleViewHits = (id: string) => {
    // Navigate to hits view or show modal
    console.log('View hits for:', id);
  };

  const jackpots = jackpotsData?.data?.jackpots || [];
  const total = jackpotsData?.data?.total || 0;

  return (
    <div className="space-y-6">
      <div className="flex justify-between items-center">
        <h1 className="text-2xl font-bold text-gray-900">Jackpot Management</h1>
        <button
          onClick={() => setShowCreateModal(true)}
          className="flex items-center px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700"
        >
          <PlusIcon className="h-5 w-5 mr-2" />
          Create Jackpot
        </button>
      </div>

      {/* Stats */}
      <div className="grid grid-cols-1 md:grid-cols-4 gap-4">
        <div className="bg-white p-4 rounded-lg shadow">
          <div className="text-sm text-gray-600">Total Jackpots</div>
          <div className="text-2xl font-bold text-blue-600">{total}</div>
        </div>
        <div className="bg-white p-4 rounded-lg shadow">
          <div className="text-sm text-gray-600">Active Jackpots</div>
          <div className="text-2xl font-bold text-green-600">
            {jackpots.filter(j => j.status === 'ACTIVE').length}
          </div>
        </div>
        <div className="bg-white p-4 rounded-lg shadow">
          <div className="text-sm text-gray-600">Total Pool Value</div>
          <div className="text-2xl font-bold text-purple-600">
            ${jackpots.reduce((sum, j) => sum + j.currentAmount, 0).toLocaleString()}
          </div>
        </div>
        <div className="bg-white p-4 rounded-lg shadow">
          <div className="text-sm text-gray-600">Hits Today</div>
          <div className="text-2xl font-bold text-yellow-600">
            {jackpotsData?.data?.stats?.hitsToday || 0}
          </div>
        </div>
      </div>

      {/* Jackpots List */}
      <div className="bg-white shadow rounded-lg overflow-hidden">
        <table className="min-w-full divide-y divide-gray-200">
          <thead className="bg-gray-50">
            <tr>
              <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Name</th>
              <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Type</th>
              <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Current Amount</th>
              <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Seed Amount</th>
              <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Contribution</th>
              <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Trigger</th>
              <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Status</th>
              <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Actions</th>
            </tr>
          </thead>
          <tbody className="bg-white divide-y divide-gray-200">
            {isLoading ? (
              <tr>
                <td colSpan={8} className="px-6 py-4 text-center text-gray-500">Loading...</td>
              </tr>
            ) : jackpots.length === 0 ? (
              <tr>
                <td colSpan={8} className="px-6 py-4 text-center text-gray-500">No jackpots found</td>
              </tr>
            ) : (
              jackpots.map((jackpot: Jackpot) => (
                <tr key={jackpot.id} className="hover:bg-gray-50">
                  <td className="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">
                    {jackpot.name}
                  </td>
                  <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                    {jackpot.type.replace('_', ' ')}
                  </td>
                  <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-900 font-semibold">
                    ${jackpot.currentAmount.toLocaleString()}
                  </td>
                  <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                    ${jackpot.seedAmount.toLocaleString()}
                  </td>
                  <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                    {jackpot.contributionRate}%
                  </td>
                  <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                    {jackpot.triggerType.replace('_', ' ')}
                  </td>
                  <td className="px-6 py-4 whitespace-nowrap">
                    <span className={`px-2 py-1 text-xs rounded-full ${
                      jackpot.status === 'ACTIVE' ? 'bg-green-100 text-green-800' :
                      jackpot.status === 'INACTIVE' ? 'bg-gray-100 text-gray-800' :
                      'bg-red-100 text-red-800'
                    }`}>
                      {jackpot.status}
                    </span>
                  </td>
                  <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500 space-x-2">
                    <button
                      onClick={() => setSelectedJackpot(jackpot)}
                      className="text-blue-600 hover:text-blue-900"
                    >
                      View
                    </button>
                    {jackpot.status === 'ACTIVE' ? (
                      <button
                        onClick={() => deactivateMutation.mutate(jackpot.id)}
                        className="text-yellow-600 hover:text-yellow-900"
                      >
                        Deactivate
                      </button>
                    ) : (
                      <button
                        onClick={() => activateMutation.mutate(jackpot.id)}
                        className="text-green-600 hover:text-green-900"
                      >
                        Activate
                      </button>
                    )}
                    <button
                      onClick={() => handleDelete(jackpot.id)}
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
          disabled={jackpots.length < limit}
          className="px-4 py-2 border rounded-lg disabled:opacity-50"
        >
          Next
        </button>
      </div>

      {/* Create Jackpot Modal */}
      {showCreateModal && (
        <div className="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
          <div className="bg-white rounded-lg p-6 max-w-md w-full">
            <div className="flex justify-between items-center mb-4">
              <h2 className="text-xl font-bold">Create Jackpot</h2>
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
                <label className="block text-sm font-medium text-gray-700">Type</label>
                <select name="type" className="mt-1 block w-full border border-gray-300 rounded-md px-3 py-2">
                  <option value="FIXED">Fixed</option>
                  <option value="PROGRESSIVE_LOCAL">Progressive Local</option>
                  <option value="PROGRESSIVE_NETWORK">Progressive Network</option>
                  <option value="MYSTERY">Mystery</option>
                  <option value="MULTI_TIER">Multi-Tier</option>
                </select>
              </div>
              <div>
                <label className="block text-sm font-medium text-gray-700">Current Amount</label>
                <input name="currentAmount" type="number" step="0.01" required className="mt-1 block w-full border border-gray-300 rounded-md px-3 py-2" />
              </div>
              <div>
                <label className="block text-sm font-medium text-gray-700">Seed Amount</label>
                <input name="seedAmount" type="number" step="0.01" required className="mt-1 block w-full border border-gray-300 rounded-md px-3 py-2" />
              </div>
              <div>
                <label className="block text-sm font-medium text-gray-700">Max Amount (Optional)</label>
                <input name="maxAmount" type="number" step="0.01" className="mt-1 block w-full border border-gray-300 rounded-md px-3 py-2" />
              </div>
              <div>
                <label className="block text-sm font-medium text-gray-700">Contribution Rate (%)</label>
                <input name="contributionRate" type="number" step="0.01" required className="mt-1 block w-full border border-gray-300 rounded-md px-3 py-2" />
              </div>
              <div>
                <label className="block text-sm font-medium text-gray-700">Trigger Type</label>
                <select name="triggerType" className="mt-1 block w-full border border-gray-300 rounded-md px-3 py-2">
                  <option value="SYMBOL_COMBINATION">Symbol Combination</option>
                  <option value="RANDOM">Random</option>
                  <option value="THRESHOLD">Threshold</option>
                </select>
              </div>
              <div>
                <label className="block text-sm font-medium text-gray-700">Linked Games (comma-separated IDs)</label>
                <input name="linkedGames" placeholder="game1,game2,game3" className="mt-1 block w-full border border-gray-300 rounded-md px-3 py-2" />
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

      {/* Jackpot Detail Modal */}
      {selectedJackpot && (
        <div className="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
          <div className="bg-white rounded-lg p-6 max-w-2xl w-full">
            <div className="flex justify-between items-center mb-4">
              <h2 className="text-xl font-bold">Jackpot Details</h2>
              <button onClick={() => setSelectedJackpot(null)} className="text-gray-500 hover:text-gray-700">
                <XCircleIcon className="h-6 w-6" />
              </button>
            </div>
            <div className="space-y-4">
              <div className="grid grid-cols-2 gap-4">
                <div>
                  <p className="text-sm text-gray-600">Name</p>
                  <p className="font-semibold">{selectedJackpot.name}</p>
                </div>
                <div>
                  <p className="text-sm text-gray-600">Type</p>
                  <p className="font-semibold">{selectedJackpot.type}</p>
                </div>
                <div>
                  <p className="text-sm text-gray-600">Current Amount</p>
                  <p className="font-semibold text-green-600">${selectedJackpot.currentAmount.toLocaleString()}</p>
                </div>
                <div>
                  <p className="text-sm text-gray-600">Seed Amount</p>
                  <p className="font-semibold">${selectedJackpot.seedAmount.toLocaleString()}</p>
                </div>
                <div>
                  <p className="text-sm text-gray-600">Contribution Rate</p>
                  <p className="font-semibold">{selectedJackpot.contributionRate}%</p>
                </div>
                <div>
                  <p className="text-sm text-gray-600">Trigger Type</p>
                  <p className="font-semibold">{selectedJackpot.triggerType}</p>
                </div>
              </div>
              <div className="flex justify-end space-x-2 pt-4">
                <button
                  onClick={() => handleViewHits(selectedJackpot.id)}
                  className="px-4 py-2 bg-purple-600 text-white rounded-lg hover:bg-purple-700"
                >
                  View Hits History
                </button>
                <button
                  onClick={() => setSelectedJackpot(null)}
                  className="px-4 py-2 border rounded-lg"
                >
                  Close
                </button>
              </div>
            </div>
          </div>
        </div>
      )}
    </div>
  );
}
