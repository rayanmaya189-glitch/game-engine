import { useState } from 'react';
import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query';
import { agentService, Agent } from '../../services/agentService';
import { PlusIcon, XCircleIcon } from '@heroicons/react/24/outline';

export function AgentsPage() {
  const [page, setPage] = useState(1);
  const [limit] = useState(20);
  const [showCreateModal, setShowCreateModal] = useState(false);
  const queryClient = useQueryClient();

  const { data: agentsData, isLoading } = useQuery({
    queryKey: ['agents', page, limit],
    queryFn: () => agentService.getAllAgents(page, limit),
  });

  const createMutation = useMutation({
    mutationFn: (agent: Partial<Agent>) => agentService.createAgent(agent),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['agents'] });
      setShowCreateModal(false);
    },
  });

  const suspendMutation = useMutation({
    mutationFn: ({ id, reason }: { id: string; reason: string }) => 
      agentService.suspendAgent(id, reason),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['agents'] });
    },
  });

  const activateMutation = useMutation({
    mutationFn: (agentId: string) => agentService.activateAgent(agentId),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['agents'] });
    },
  });

  const handleCreate = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    const formData = new FormData(e.currentTarget);
    createMutation.mutate({
      name: formData.get('name') as string,
      email: formData.get('email') as string,
      phone: formData.get('phone') as string,
      tier: formData.get('tier') as Agent['tier'],
    });
  };

  const handleSuspend = (id: string) => {
    const reason = prompt('Enter suspension reason:');
    if (reason) {
      suspendMutation.mutate({ id, reason });
    }
  };

  const agents = agentsData?.data?.agents || [];
  const total = agentsData?.data?.total || 0;

  return (
    <div className="space-y-6">
      <div className="flex justify-between items-center">
        <h1 className="text-2xl font-bold text-gray-900">Agents & Affiliates</h1>
        <button
          onClick={() => setShowCreateModal(true)}
          className="flex items-center px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700"
        >
          <PlusIcon className="h-5 w-5 mr-2" />
          Create Agent
        </button>
      </div>

      {/* Agents List */}
      <div className="bg-white shadow rounded-lg overflow-hidden">
        <table className="min-w-full divide-y divide-gray-200">
          <thead className="bg-gray-50">
            <tr>
              <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Name</th>
              <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Email</th>
              <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Tier</th>
              <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Players</th>
              <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Total Commission</th>
              <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Pending</th>
              <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Status</th>
              <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Actions</th>
            </tr>
          </thead>
          <tbody className="bg-white divide-y divide-gray-200">
            {isLoading ? (
              <tr>
                <td colSpan={8} className="px-6 py-4 text-center text-gray-500">Loading...</td>
              </tr>
            ) : agents.length === 0 ? (
              <tr>
                <td colSpan={8} className="px-6 py-4 text-center text-gray-500">No agents found</td>
              </tr>
            ) : (
              agents.map((agent: Agent) => (
                <tr key={agent.id} className="hover:bg-gray-50">
                  <td className="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">
                    {agent.name}
                  </td>
                  <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                    {agent.email}
                  </td>
                  <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                    <span className={`px-2 py-1 text-xs rounded-full ${
                      agent.tier === 'DIAMOND' ? 'bg-purple-100 text-purple-800' :
                      agent.tier === 'PLATINUM' ? 'bg-blue-100 text-blue-800' :
                      agent.tier === 'GOLD' ? 'bg-yellow-100 text-yellow-800' :
                      agent.tier === 'SILVER' ? 'bg-gray-100 text-gray-800' :
                      'bg-orange-100 text-orange-800'
                    }`}>
                      {agent.tier}
                    </span>
                  </td>
                  <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                    {agent.totalPlayers}
                  </td>
                  <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                    ${agent.totalCommission.toFixed(2)}
                  </td>
                  <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                    ${agent.pendingCommission.toFixed(2)}
                  </td>
                  <td className="px-6 py-4 whitespace-nowrap">
                    <span className={`px-2 py-1 text-xs rounded-full ${
                      agent.status === 'ACTIVE' ? 'bg-green-100 text-green-800' :
                      agent.status === 'SUSPENDED' ? 'bg-red-100 text-red-800' :
                      'bg-gray-100 text-gray-800'
                    }`}>
                      {agent.status}
                    </span>
                  </td>
                  <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500 space-x-2">
                    {agent.status === 'ACTIVE' ? (
                      <button
                        onClick={() => handleSuspend(agent.id)}
                        className="text-yellow-600 hover:text-yellow-900"
                      >
                        Suspend
                      </button>
                    ) : (
                      <button
                        onClick={() => activateMutation.mutate(agent.id)}
                        className="text-green-600 hover:text-green-900"
                      >
                        Activate
                      </button>
                    )}
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
          disabled={agents.length < limit}
          className="px-4 py-2 border rounded-lg disabled:opacity-50"
        >
          Next
        </button>
      </div>

      {/* Create Agent Modal */}
      {showCreateModal && (
        <div className="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
          <div className="bg-white rounded-lg p-6 max-w-md w-full">
            <div className="flex justify-between items-center mb-4">
              <h2 className="text-xl font-bold">Create Agent</h2>
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
                <label className="block text-sm font-medium text-gray-700">Tier</label>
                <select name="tier" className="mt-1 block w-full border border-gray-300 rounded-md px-3 py-2">
                  <option value="BRONZE">Bronze</option>
                  <option value="SILVER">Silver</option>
                  <option value="GOLD">Gold</option>
                  <option value="PLATINUM">Platinum</option>
                  <option value="DIAMOND">Diamond</option>
                </select>
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
