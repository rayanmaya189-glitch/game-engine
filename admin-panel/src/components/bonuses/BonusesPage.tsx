import { useState } from 'react';
import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query';
import { bonusService, Bonus } from '../../services/bonusService';
import { PlusIcon, XCircleIcon } from '@heroicons/react/24/outline';

export function BonusesPage() {
  const [page, setPage] = useState(1);
  const [limit] = useState(20);
  const [showCreateModal, setShowCreateModal] = useState(false);
  const queryClient = useQueryClient();

  const { data: bonusesData, isLoading } = useQuery({
    queryKey: ['bonuses', page, limit],
    queryFn: () => bonusService.getAllBonuses(page, limit),
  });

  const createMutation = useMutation({
    mutationFn: (bonus: Partial<Bonus>) => bonusService.createBonus(bonus),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['bonuses'] });
      setShowCreateModal(false);
    },
  });

  const deleteMutation = useMutation({
    mutationFn: (bonusId: string) => bonusService.deleteBonus(bonusId),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['bonuses'] });
    },
  });

  const activateMutation = useMutation({
    mutationFn: (bonusId: string) => bonusService.activateBonus(bonusId),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['bonuses'] });
    },
  });

  const deactivateMutation = useMutation({
    mutationFn: (bonusId: string) => bonusService.deactivateBonus(bonusId),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['bonuses'] });
    },
  });

  const handleCreate = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    const formData = new FormData(e.currentTarget);
    createMutation.mutate({
      name: formData.get('name') as string,
      type: formData.get('type') as Bonus['type'],
      amount: Number(formData.get('amount')),
      currency: 'USD',
      wageringRequirement: Number(formData.get('wageringRequirement')),
      minDeposit: Number(formData.get('minDeposit')) || undefined,
      validFrom: formData.get('validFrom') as string,
      validUntil: formData.get('validUntil') as string,
      targetAudience: formData.get('targetAudience') as Bonus['targetAudience'],
    });
  };

  const handleDelete = (id: string) => {
    if (confirm('Are you sure you want to delete this bonus?')) {
      deleteMutation.mutate(id);
    }
  };

  const bonuses = bonusesData?.data?.bonuses || [];
  const total = bonusesData?.data?.total || 0;

  return (
    <div className="space-y-6">
      <div className="flex justify-between items-center">
        <h1 className="text-2xl font-bold text-gray-900">Bonuses & Promotions</h1>
        <button
          onClick={() => setShowCreateModal(true)}
          className="flex items-center px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700"
        >
          <PlusIcon className="h-5 w-5 mr-2" />
          Create Bonus
        </button>
      </div>

      {/* Bonuses List */}
      <div className="bg-white shadow rounded-lg overflow-hidden">
        <table className="min-w-full divide-y divide-gray-200">
          <thead className="bg-gray-50">
            <tr>
              <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Name</th>
              <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Type</th>
              <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Amount</th>
              <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Wagering</th>
              <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Target</th>
              <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Valid Until</th>
              <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Status</th>
              <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Actions</th>
            </tr>
          </thead>
          <tbody className="bg-white divide-y divide-gray-200">
            {isLoading ? (
              <tr>
                <td colSpan={8} className="px-6 py-4 text-center text-gray-500">Loading...</td>
              </tr>
            ) : bonuses.length === 0 ? (
              <tr>
                <td colSpan={8} className="px-6 py-4 text-center text-gray-500">No bonuses found</td>
              </tr>
            ) : (
              bonuses.map((bonus: Bonus) => (
                <tr key={bonus.id} className="hover:bg-gray-50">
                  <td className="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">
                    {bonus.name}
                  </td>
                  <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                    {bonus.type.replace('_', ' ')}
                  </td>
                  <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                    {bonus.currency} {bonus.amount.toFixed(2)}
                  </td>
                  <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                    {bonus.wageringRequirement}x
                  </td>
                  <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                    {bonus.targetAudience.replace('_', ' ')}
                  </td>
                  <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                    {new Date(bonus.validUntil).toLocaleDateString()}
                  </td>
                  <td className="px-6 py-4 whitespace-nowrap">
                    <span className={`px-2 py-1 text-xs rounded-full ${
                      bonus.status === 'ACTIVE' ? 'bg-green-100 text-green-800' :
                      bonus.status === 'INACTIVE' ? 'bg-gray-100 text-gray-800' :
                      'bg-red-100 text-red-800'
                    }`}>
                      {bonus.status}
                    </span>
                  </td>
                  <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500 space-x-2">
                    {bonus.status === 'ACTIVE' ? (
                      <button
                        onClick={() => deactivateMutation.mutate(bonus.id)}
                        className="text-yellow-600 hover:text-yellow-900"
                      >
                        Deactivate
                      </button>
                    ) : (
                      <button
                        onClick={() => activateMutation.mutate(bonus.id)}
                        className="text-green-600 hover:text-green-900"
                      >
                        Activate
                      </button>
                    )}
                    <button
                      onClick={() => handleDelete(bonus.id)}
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
          disabled={bonuses.length < limit}
          className="px-4 py-2 border rounded-lg disabled:opacity-50"
        >
          Next
        </button>
      </div>

      {/* Create Bonus Modal */}
      {showCreateModal && (
        <div className="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
          <div className="bg-white rounded-lg p-6 max-w-md w-full">
            <div className="flex justify-between items-center mb-4">
              <h2 className="text-xl font-bold">Create Bonus</h2>
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
                  <option value="WELCOME">Welcome Bonus</option>
                  <option value="DEPOSIT">Deposit Bonus</option>
                  <option value="NO_DEPOSIT">No Deposit Bonus</option>
                  <option value="FREE_SPINS">Free Spins</option>
                  <option value="CASHBACK">Cashback</option>
                  <option value="LOYALTY">Loyalty Bonus</option>
                </select>
              </div>
              <div>
                <label className="block text-sm font-medium text-gray-700">Amount</label>
                <input name="amount" type="number" step="0.01" required className="mt-1 block w-full border border-gray-300 rounded-md px-3 py-2" />
              </div>
              <div>
                <label className="block text-sm font-medium text-gray-700">Wagering Requirement (x)</label>
                <input name="wageringRequirement" type="number" step="1" required className="mt-1 block w-full border border-gray-300 rounded-md px-3 py-2" />
              </div>
              <div>
                <label className="block text-sm font-medium text-gray-700">Min Deposit (Optional)</label>
                <input name="minDeposit" type="number" step="0.01" className="mt-1 block w-full border border-gray-300 rounded-md px-3 py-2" />
              </div>
              <div>
                <label className="block text-sm font-medium text-gray-700">Target Audience</label>
                <select name="targetAudience" className="mt-1 block w-full border border-gray-300 rounded-md px-3 py-2">
                  <option value="ALL">All Players</option>
                  <option value="NEW_PLAYERS">New Players Only</option>
                  <option value="VIP">VIP Players</option>
                </select>
              </div>
              <div>
                <label className="block text-sm font-medium text-gray-700">Valid From</label>
                <input name="validFrom" type="date" required className="mt-1 block w-full border border-gray-300 rounded-md px-3 py-2" />
              </div>
              <div>
                <label className="block text-sm font-medium text-gray-700">Valid Until</label>
                <input name="validUntil" type="date" required className="mt-1 block w-full border border-gray-300 rounded-md px-3 py-2" />
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
