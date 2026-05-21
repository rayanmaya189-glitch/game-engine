import { useState } from 'react';
import { useGames, useToggleGameStatus } from '../hooks/useGame';
import { getStatusColor } from '../hooks/useUtils';

export function GamesPage() {
  const [page, setPage] = useState(1);
  const [filters, setFilters] = useState<Record<string, string>>({});
  const { data, isLoading } = useGames(page, 20, filters);
  const toggleGameStatus = useToggleGameStatus();

  const handleToggleStatus = async (gameId: string, currentStatus: string) => {
    const newStatus = currentStatus === 'ACTIVE' ? 'INACTIVE' : 'ACTIVE';
    await toggleGameStatus.mutateAsync({ gameId, status: newStatus as any });
  };

  if (isLoading) {
    return <div className="text-center py-8">Loading...</div>;
  }

  return (
    <div className="space-y-6">
      <div className="flex justify-between items-center">
        <h1 className="text-2xl font-bold text-gray-900">Game Management</h1>
        <button className="btn-primary">Add Game</button>
      </div>

      {/* Filters */}
      <div className="card">
        <div className="grid grid-cols-1 gap-4 sm:grid-cols-4">
          <input
            type="text"
            placeholder="Search by name..."
            className="input-field"
            onChange={(e) => setFilters({ ...filters, search: e.target.value })}
          />
          <select
            className="input-field"
            onChange={(e) => setFilters({ ...filters, type: e.target.value })}
          >
            <option value="">All Types</option>
            <option value="SLOT">Slots</option>
            <option value="CARD">Card Games</option>
            <option value="DICE">Dice Games</option>
            <option value="TABLE">Table Games</option>
            <option value="LIVE_DEALER">Live Dealer</option>
          </select>
          <select
            className="input-field"
            onChange={(e) => setFilters({ ...filters, status: e.target.value })}
          >
            <option value="">All Status</option>
            <option value="ACTIVE">Active</option>
            <option value="INACTIVE">Inactive</option>
            <option value="MAINTENANCE">Maintenance</option>
          </select>
          <select
            className="input-field"
            onChange={(e) => setFilters({ ...filters, provider: e.target.value })}
          >
            <option value="">All Providers</option>
            <option value="PROVIDER_A">Provider A</option>
            <option value="PROVIDER_B">Provider B</option>
          </select>
        </div>
      </div>

      {/* Games Grid */}
      <div className="grid grid-cols-1 gap-6 md:grid-cols-2 lg:grid-cols-3">
        {data?.games?.map((game: any) => (
          <div key={game.id} className="card">
            <div className="flex justify-between items-start mb-4">
              <div>
                <h3 className="text-lg font-semibold text-gray-900">{game.name}</h3>
                <p className="text-sm text-gray-500">{game.provider}</p>
              </div>
              <span className={`px-2 py-1 text-xs rounded-full ${getStatusColor(game.status)}`}>
                {game.status}
              </span>
            </div>
            
            <div className="space-y-2 text-sm">
              <div className="flex justify-between">
                <span className="text-gray-500">Type:</span>
                <span className="font-medium">{game.type}</span>
              </div>
              <div className="flex justify-between">
                <span className="text-gray-500">RTP:</span>
                <span className="font-medium">{game.rtp}%</span>
              </div>
              <div className="flex justify-between">
                <span className="text-gray-500">Bet Range:</span>
                <span className="font-medium">${game.minBet} - ${game.maxBet}</span>
              </div>
              <div className="flex justify-between">
                <span className="text-gray-500">Total Plays:</span>
                <span className="font-medium">{game.totalPlays.toLocaleString()}</span>
              </div>
            </div>

            <div className="mt-4 flex space-x-2">
              <button className="btn-primary flex-1">Configure</button>
              <button 
                className="btn-danger flex-1"
                onClick={() => handleToggleStatus(game.id, game.status)}
              >
                {game.status === 'ACTIVE' ? 'Deactivate' : 'Activate'}
              </button>
            </div>
          </div>
        ))}
      </div>

      {/* Pagination */}
      <div className="flex justify-center space-x-4">
        <button
          onClick={() => setPage(Math.max(1, page - 1))}
          disabled={page === 1}
          className="btn-primary disabled:opacity-50"
        >
          Previous
        </button>
        <span className="text-sm text-gray-600 self-center">Page {page}</span>
        <button
          onClick={() => setPage(page + 1)}
          disabled={!data?.hasMore}
          className="btn-primary disabled:opacity-50"
        >
          Next
        </button>
      </div>
    </div>
  );
}
