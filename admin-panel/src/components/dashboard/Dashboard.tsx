import { useRevenueStats } from '../hooks/useFinancial';
import { useUsers } from '../hooks/useUser';
import { useGames } from '../hooks/useGame';
import { formatCurrency } from '../hooks/useUtils';
import { BarChart, Bar, XAxis, YAxis, CartesianGrid, Tooltip, ResponsiveContainer } from 'recharts';

export function Dashboard() {
  const { data: revenueData } = useRevenueStats('MONTH');
  const { data: usersData } = useUsers(1, 100);
  const { data: gamesData } = useGames(1, 100);

  const stats = [
    { name: 'Total Revenue', value: formatCurrency(revenueData?.totalRevenue || 0), change: '+12.5%' },
    { name: 'Active Users', value: usersData?.total || 0, change: '+8.2%' },
    { name: 'Total Games', value: gamesData?.total || 0, change: '+2.1%' },
    { name: 'House Edge', value: `${revenueData?.houseEdge || 0}%`, change: '-0.3%' },
  ];

  const chartData = revenueData?.dailyRevenue || [];

  return (
    <div className="space-y-6">
      <h1 className="text-2xl font-bold text-gray-900">Dashboard</h1>

      {/* Stats Grid */}
      <div className="grid grid-cols-1 gap-6 sm:grid-cols-2 lg:grid-cols-4">
        {stats.map((stat) => (
          <div key={stat.name} className="bg-white rounded-lg shadow p-6">
            <dt className="text-sm font-medium text-gray-500 truncate">{stat.name}</dt>
            <dd className="mt-1 text-3xl font-semibold text-gray-900">{stat.value}</dd>
            <dd className={`mt-1 text-sm ${stat.change.startsWith('+') ? 'text-success-600' : 'text-danger-600'}`}>
              {stat.change} from last month
            </dd>
          </div>
        ))}
      </div>

      {/* Revenue Chart */}
      <div className="bg-white rounded-lg shadow p-6">
        <h2 className="text-lg font-medium text-gray-900 mb-4">Revenue Overview</h2>
        <div className="h-80">
          <ResponsiveContainer width="100%" height="100%">
            <BarChart data={chartData}>
              <CartesianGrid strokeDasharray="3 3" />
              <XAxis dataKey="date" />
              <YAxis />
              <Tooltip />
              <Bar dataKey="revenue" fill="#0ea5e9" />
            </BarChart>
          </ResponsiveContainer>
        </div>
      </div>

      {/* Recent Activity */}
      <div className="bg-white rounded-lg shadow p-6">
        <h2 className="text-lg font-medium text-gray-900 mb-4">Quick Actions</h2>
        <div className="grid grid-cols-1 gap-4 sm:grid-cols-2 lg:grid-cols-3">
          <a href="/users" className="block p-4 border rounded-lg hover:bg-gray-50">
            <h3 className="font-medium text-gray-900">Manage Users</h3>
            <p className="text-sm text-gray-500">View and manage user accounts</p>
          </a>
          <a href="/games" className="block p-4 border rounded-lg hover:bg-gray-50">
            <h3 className="font-medium text-gray-900">Game Management</h3>
            <p className="text-sm text-gray-500">Configure game settings</p>
          </a>
          <a href="/transactions" className="block p-4 border rounded-lg hover:bg-gray-50">
            <h3 className="font-medium text-gray-900">Transactions</h3>
            <p className="text-sm text-gray-500">Review and approve transactions</p>
          </a>
        </div>
      </div>
    </div>
  );
}
