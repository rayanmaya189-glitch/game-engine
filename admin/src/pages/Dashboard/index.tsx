import { useQuery } from '@tanstack/react-query';
import { Grid, Card, CardContent, Typography, Box, LinearProgress } from '@mui/material';
import { TrendingUp, People, Games as GamesIcon, AttachMoney, TrendingDown } from '@mui/icons-material';
import { reportsAPI } from '../../services/api';

const Dashboard = () => {
  const { data: stats, isLoading } = useQuery({
    queryKey: ['dashboard-stats'],
    queryFn: () => reportsAPI.getDashboardStats(),
  });

  const dashboardData = stats?.data || {
    totalRevenue: 125430,
    activeUsers: 1234,
    gamesPlayed: 5678,
    conversionRate: 12.5,
    revenueChange: 5.2,
    usersChange: 8.1,
    gamesChange: -2.3,
    conversionChange: 1.5
  };

  const statCards = [
    { 
      title: 'Total Revenue', 
      value: `${dashboardData.totalRevenue.toLocaleString()}`, 
      icon: <AttachMoney />, 
      color: '#22c55e',
      change: dashboardData.revenueChange,
      trend: dashboardData.revenueChange >= 0 ? 'up' : 'down'
    },
    { 
      title: 'Active Users', 
      value: dashboardData.activeUsers.toLocaleString(), 
      icon: <People />, 
      color: '#3b82f6',
      change: dashboardData.usersChange,
      trend: dashboardData.usersChange >= 0 ? 'up' : 'down'
    },
    { 
      title: 'Games Played', 
      value: dashboardData.gamesPlayed.toLocaleString(), 
      icon: <GamesIcon />, 
      color: '#f97316',
      change: dashboardData.gamesChange,
      trend: dashboardData.gamesChange >= 0 ? 'up' : 'down'
    },
    { 
      title: 'Conversion Rate', 
      value: `${dashboardData.conversionRate}%`, 
      icon: <TrendingUp />, 
      color: '#8b5cf6',
      change: dashboardData.conversionChange,
      trend: dashboardData.conversionChange >= 0 ? 'up' : 'down'
    },
  ];

  return (
    <Box>
      <Typography variant="h4" fontWeight="bold" gutterBottom>
        Dashboard
      </Typography>

      {isLoading && <LinearProgress sx={{ mb: 2 }} />}

      <Grid container spacing={3}>
        {statCards.map((stat, index) => {
          const isPositive = stat.change >= 0;
          return (
            <Grid item xs={12} sm={6} md={3} key={index}>
              <Card>
                <CardContent>
                  <Box sx={{ display: 'flex', justifyContent: 'space-between', alignItems: 'flex-start' }}>
                    <Box>
                      <Typography color="text.secondary" variant="body2">
                        {stat.title}
                      </Typography>
                      <Typography variant="h4" fontWeight="bold">
                        {stat.value}
                      </Typography>
                      <Box sx={{ display: 'flex', alignItems: 'center', mt: 1 }}>
                        {isPositive ? 
                          <TrendingUp sx={{ fontSize: 16, color: '#22c55e', mr: 0.5 }} /> : 
                          <TrendingDown sx={{ fontSize: 16, color: '#ef4444', mr: 0.5 }} />
                        }
                        <Typography 
                          variant="caption" 
                          sx={{ color: isPositive ? '#22c55e' : '#ef4444', fontWeight: 600 }}
                        >
                          {isPositive ? '+' : ''}{stat.change}%
                        </Typography>
                        <Typography variant="caption" color="text.secondary" sx={{ ml: 0.5 }}>
                          vs last week
                        </Typography>
                      </Box>
                    </Box>
                    <Box sx={{ 
                      p: 1.5, 
                      borderRadius: 2, 
                      bgcolor: `${stat.color}20`,
                      color: stat.color 
                    }}>
                      {stat.icon}
                    </Box>
                  </Box>
                </CardContent>
              </Card>
            </Grid>
          );
        })}
      </Grid>
    </Box>
  );
};

export default Dashboard;
