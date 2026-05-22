import { useState } from 'react';
import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query';
import { tournamentService, Tournament } from '../../services/tournamentService';
import { PlusIcon, XCircleIcon, TrophyIcon } from '@heroicons/react/24/outline';

export function TournamentsPage() {
  const [page, setPage] = useState(1);
  const [limit] = useState(20);
  const [selectedTournament, setSelectedTournament] = useState<string | null>(null);
  const [showCreateModal, setShowCreateModal] = useState(false);
  const queryClient = useQueryClient();

  const { data: tournamentsData, isLoading } = useQuery({
    queryKey: ['tournaments', page, limit],
    queryFn: () => tournamentService.getAllTournaments(page, limit),
  });

  const { data: selectedTournamentData } = useQuery({
    queryKey: ['tournament', selectedTournament],
    queryFn: () => tournamentService.getTournament(selectedTournament!),
    enabled: !!selectedTournament,
  });

  const { data: leaderboardData } = useQuery({
    queryKey: ['tournamentLeaderboard', selectedTournament],
    queryFn: () => tournamentService.getLeaderboard(selectedTournament!),
    enabled: !!selectedTournament,
  });

  const createMutation = useMutation({
    mutationFn: (tournament: Partial<Tournament>) => tournamentService.createTournament(tournament),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['tournaments'] });
      setShowCreateModal(false);
    },
  });

  const startMutation = useMutation({
    mutationFn: (tournamentId: string) => tournamentService.startTournament(tournamentId),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['tournaments'] });
    },
  });

  const cancelMutation = useMutation({
    mutationFn: ({ tournamentId, reason }: { tournamentId: string; reason: string }) =>
      tournamentService.cancelTournament(tournamentId, reason),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['tournaments'] });
    },
  });

  const deleteMutation = useMutation({
    mutationFn: (tournamentId: string) => tournamentService.deleteTournament(tournamentId),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['tournaments'] });
      setSelectedTournament(null);
    },
  });

  const handleCreate = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    const formData = new FormData(e.currentTarget);
    createMutation.mutate({
      name: formData.get('name') as string,
      type: formData.get('type') as Tournament['type'],
      gameId: formData.get('gameId') as string,
      entryFee: Number(formData.get('entryFee')),
      guaranteedPrizePool: Number(formData.get('guaranteedPrizePool')),
      maxPlayers: Number(formData.get('maxPlayers')),
      startTime: formData.get('startTime') as string,
    });
  };

  const handleStart = (id: string) => {
    if (confirm('Are you sure you want to start this tournament?')) {
      startMutation.mutate(id);
    }
  };

  const handleCancel = (id: string) => {
    const reason = prompt('Enter cancellation reason:');
    if (reason) {
      cancelMutation.mutate({ tournamentId: id, reason });
    }
  };

  const handleDelete = (id: string) => {
    if (confirm('Are you sure you want to delete this tournament?')) {
      deleteMutation.mutate(id);
    }
  };

  const tournaments = tournamentsData?.data?.tournaments || [];
  const total = tournamentsData?.data?.total || 0;

  return (
    <div className="space-y-6">
      <div className="flex justify-between items-center">
        <h1 className="text-2xl font-bold text-gray-900">Tournaments</h1>
        <button
          onClick={() => setShowCreateModal(true)}
          className="flex items-center px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700"
        >
          <PlusIcon className="h-5 w-5 mr-2" />
          Create Tournament
        </button>
      </div>

      {/* Tournaments List */}
      <div className="bg-white shadow rounded-lg overflow-hidden">
        <table className="min-w-full divide-y divide-gray-200">
          <thead className="bg-gray-50">
            <tr>
              <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Name</th>
              <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Type</th>
              <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Game</th>
              <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Entry Fee</th>
              <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Prize Pool</th>
              <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Players</th>
              <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Status</th>
              <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Actions</th>
            </tr>
          </thead>
          <tbody className="bg-white divide-y divide-gray-200">
            {isLoading ? (
              <tr>
                <td colSpan={8} className="px-6 py-4 text-center text-gray-500">Loading...</td>
              </tr>
            ) : tournaments.length === 0 ? (
              <tr>
                <td colSpan={8} className="px-6 py-4 text-center text-gray-500">No tournaments found</td>
              </tr>
            ) : (
              tournaments.map((tournament: Tournament) => (
                <tr key={tournament.id} className="hover:bg-gray-50">
                  <td className="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">
                    {tournament.name}
                  </td>
                  <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                    {tournament.type.replace('_', ' ')}
                  </td>
                  <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                    {tournament.gameName}
                  </td>
                  <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                    {tournament.currency} {tournament.entryFee.toFixed(2)}
                  </td>
                  <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                    {tournament.currency} {tournament.currentPrizePool.toFixed(2)}
                  </td>
                  <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                    {tournament.registeredPlayers}/{tournament.maxPlayers}
                  </td>
                  <td className="px-6 py-4 whitespace-nowrap">
                    <span className={`px-2 py-1 text-xs rounded-full ${
                      tournament.status === 'IN_PROGRESS' ? 'bg-green-100 text-green-800' :
                      tournament.status === 'REGISTERING' ? 'bg-blue-100 text-blue-800' :
                      tournament.status === 'COMPLETED' ? 'bg-gray-100 text-gray-800' :
                      'bg-yellow-100 text-yellow-800'
                    }`}>
                      {tournament.status.replace('_', ' ')}
                    </span>
                  </td>
                  <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500 space-x-2">
                    <button
                      onClick={() => setSelectedTournament(tournament.id)}
                      className="text-blue-600 hover:text-blue-900"
                    >
                      View
                    </button>
                    {tournament.status === 'REGISTERING' && (
                      <button
                        onClick={() => handleStart(tournament.id)}
                        className="text-green-600 hover:text-green-900"
                      >
                        Start
                      </button>
                    )}
                    {tournament.status === 'REGISTERING' && (
                      <button
                        onClick={() => handleCancel(tournament.id)}
                        className="text-red-600 hover:text-red-900"
                      >
                        Cancel
                      </button>
                    )}
                    <button
                      onClick={() => handleDelete(tournament.id)}
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
          disabled={tournaments.length < limit}
          className="px-4 py-2 border rounded-lg disabled:opacity-50"
        >
          Next
        </button>
      </div>

      {/* Tournament Details Modal */}
      {selectedTournament && selectedTournamentData && (
        <div className="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
          <div className="bg-white rounded-lg p-6 max-w-4xl w-full max-h-[90vh] overflow-y-auto">
            <div className="flex justify-between items-center mb-4">
              <h2 className="text-xl font-bold">{selectedTournamentData.data.name}</h2>
              <button onClick={() => setSelectedTournament(null)} className="text-gray-500 hover:text-gray-700">
                <XCircleIcon className="h-6 w-6" />
              </button>
            </div>
            
            <div className="grid grid-cols-2 gap-4 mb-6">
              <div>
                <p className="text-sm text-gray-500">Type</p>
                <p className="font-medium">{selectedTournamentData.data.type}</p>
              </div>
              <div>
                <p className="text-sm text-gray-500">Status</p>
                <p className="font-medium">{selectedTournamentData.data.status}</p>
              </div>
              <div>
                <p className="text-sm text-gray-500">Start Time</p>
                <p className="font-medium">{new Date(selectedTournamentData.data.startTime).toLocaleString()}</p>
              </div>
              <div>
                <p className="text-sm text-gray-500">Prize Pool</p>
                <p className="font-medium">{selectedTournamentData.data.currency} {selectedTournamentData.data.currentPrizePool}</p>
              </div>
            </div>

            {leaderboardData && leaderboardData.data?.leaderboard && (
              <div>
                <h3 className="text-lg font-semibold mb-2 flex items-center">
                  <TrophyIcon className="h-5 w-5 mr-2" />
                  Leaderboard
                </h3>
                <table className="min-w-full divide-y divide-gray-200">
                  <thead>
                    <tr>
                      <th className="px-4 py-2 text-left text-xs font-medium text-gray-500">Rank</th>
                      <th className="px-4 py-2 text-left text-xs font-medium text-gray-500">Player</th>
                      <th className="px-4 py-2 text-left text-xs font-medium text-gray-500">Chips</th>
                      <th className="px-4 py-2 text-left text-xs font-medium text-gray-500">Status</th>
                    </tr>
                  </thead>
                  <tbody>
                    {leaderboardData.data.leaderboard.slice(0, 10).map((entry: any, idx: number) => (
                      <tr key={entry.playerId}>
                        <td className="px-4 py-2">{idx + 1}</td>
                        <td className="px-4 py-2">{entry.playerName}</td>
                        <td className="px-4 py-2">{entry.chips}</td>
                        <td className="px-4 py-2">{entry.status}</td>
                      </tr>
                    ))}
                  </tbody>
                </table>
              </div>
            )}
          </div>
        </div>
      )}

      {/* Create Tournament Modal */}
      {showCreateModal && (
        <div className="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
          <div className="bg-white rounded-lg p-6 max-w-md w-full">
            <h2 className="text-xl font-bold mb-4">Create Tournament</h2>
            <form onSubmit={handleCreate} className="space-y-4">
              <div>
                <label className="block text-sm font-medium text-gray-700">Name</label>
                <input name="name" required className="mt-1 block w-full border border-gray-300 rounded-md px-3 py-2" />
              </div>
              <div>
                <label className="block text-sm font-medium text-gray-700">Type</label>
                <select name="type" className="mt-1 block w-full border border-gray-300 rounded-md px-3 py-2">
                  <option value="SCHEDULED">Scheduled</option>
                  <option value="SIT_AND_GO">Sit & Go</option>
                  <option value="FREEROLL">Freeroll</option>
                  <option value="REBUY">Rebuy</option>
                  <option value="BOUNTY">Bounty</option>
                </select>
              </div>
              <div>
                <label className="block text-sm font-medium text-gray-700">Game</label>
                <select name="gameId" className="mt-1 block w-full border border-gray-300 rounded-md px-3 py-2">
                  <option value="poker-texas">Texas Hold'em Poker</option>
                  <option value="blackjack">Blackjack</option>
                  <option value="baccarat">Baccarat</option>
                </select>
              </div>
              <div>
                <label className="block text-sm font-medium text-gray-700">Entry Fee</label>
                <input name="entryFee" type="number" step="0.01" required className="mt-1 block w-full border border-gray-300 rounded-md px-3 py-2" />
              </div>
              <div>
                <label className="block text-sm font-medium text-gray-700">Guaranteed Prize Pool</label>
                <input name="guaranteedPrizePool" type="number" step="0.01" required className="mt-1 block w-full border border-gray-300 rounded-md px-3 py-2" />
              </div>
              <div>
                <label className="block text-sm font-medium text-gray-700">Max Players</label>
                <input name="maxPlayers" type="number" required className="mt-1 block w-full border border-gray-300 rounded-md px-3 py-2" />
              </div>
              <div>
                <label className="block text-sm font-medium text-gray-700">Start Time</label>
                <input name="startTime" type="datetime-local" required className="mt-1 block w-full border border-gray-300 rounded-md px-3 py-2" />
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
