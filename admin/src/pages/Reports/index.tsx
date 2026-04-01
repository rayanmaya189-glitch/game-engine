import { useState } from 'react';
import { useQuery } from '@tanstack/react-query';
import { 
  Box, Typography, Card, CardContent,
  Table, TableBody, TableCell, TableContainer, TableHead, 
  TableRow, Chip, Tabs, Tab, LinearProgress
} from '@mui/material';
import { reportsAPI } from '../../services/api';
import { RevenueSummaryCards, UserSummaryCards } from './ReportCharts';
import ReportFilters from './ReportFilters';

const Reports = () => {
  const [activeTab, setActiveTab] = useState(0);
  const [dateRange, setDateRange] = useState({ from: '', to: '' });
  const [period, setPeriod] = useState('30');
  const [category, setCategory] = useState('');

  const { data: revenueData, isLoading: revenueLoading } = useQuery({
    queryKey: ['revenue-report', dateRange],
    queryFn: () => reportsAPI.getRevenueReport({ 
      startDate: dateRange.from, 
      endDate: dateRange.to 
    }),
    enabled: activeTab === 0,
  });

  const { data: userData, isLoading: userLoading } = useQuery({
    queryKey: ['user-report', period],
    queryFn: () => reportsAPI.getUserReport({}),
    enabled: activeTab === 1,
  });

  const { data: gameData, isLoading: gameLoading } = useQuery({
    queryKey: ['game-report', category],
    queryFn: () => reportsAPI.getGameReport({}),
    enabled: activeTab === 2,
  });

  const revenueReport = revenueData?.data || [
    { date: '2024-01-15', deposits: 12500, withdrawals: 8500, netRevenue: 4000, bonus: 1200, rake: 2800 },
    { date: '2024-01-16', deposits: 15800, withdrawals: 9200, netRevenue: 6600, bonus: 1500, rake: 5100 },
    { date: '2024-01-17', deposits: 11200, withdrawals: 7800, netRevenue: 3400, bonus: 900, rake: 2500 },
    { date: '2024-01-18', deposits: 18900, withdrawals: 12300, netRevenue: 6600, bonus: 2100, rake: 4500 },
    { date: '2024-01-19', deposits: 14500, withdrawals: 9800, netRevenue: 4700, bonus: 1800, rake: 2900 },
  ];

  const userReport = userData?.data || [
    { date: '2024-01-15', newUsers: 45, activeUsers: 1234, kycApproved: 12, deposits: 28, conversions: 62.5 },
    { date: '2024-01-16', newUsers: 52, activeUsers: 1280, kycApproved: 15, deposits: 35, conversions: 67.3 },
    { date: '2024-01-17', newUsers: 38, activeUsers: 1198, kycApproved: 8, deposits: 22, conversions: 57.8 },
    { date: '2024-01-18', newUsers: 61, activeUsers: 1345, kycApproved: 18, deposits: 42, conversions: 68.8 },
    { date: '2024-01-19', newUsers: 55, activeUsers: 1390, kycApproved: 14, deposits: 38, conversions: 69.1 },
  ];

  const gameReport = gameData?.data || [
    { game: 'Starburst', provider: 'NetEnt', category: 'slot', revenue: 45200, plays: 12500, rtp: 96.1 },
    { game: 'Mega Moolah', provider: 'Microgaming', category: 'slot', revenue: 38500, plays: 4200, rtp: 88.12 },
    { game: 'Blackjack VIP', provider: 'Evolution', category: 'live_casino', revenue: 32100, plays: 890, rtp: 99.5 },
    { game: 'Sic Bo', provider: 'Playtech', category: 'dice', revenue: 18500, plays: 5600, rtp: 97.22 },
    { game: 'Crazy Time', provider: 'Evolution', category: 'live_casino', revenue: 52800, plays: 7200, rtp: 96.23 },
  ];

  const handleExport = () => {
    // Export functionality would be implemented here
    alert('Export functionality would download the report data');
  };

  const TabPanel = ({ children, value, index }: { children: React.ReactNode; value: number; index: number }) => (
    <Box role="tabpanel" hidden={value !== index} sx={{ py: 3 }}>
      {value === index && children}
    </Box>
  );

  const isLoading = revenueLoading || userLoading || gameLoading;

  return (
    <Box>
      <Typography variant="h4" fontWeight="bold" gutterBottom>
        Reports
      </Typography>

      <Card sx={{ mb: 3 }}>
        <Box sx={{ borderBottom: 1, borderColor: 'divider' }}>
          <Tabs value={activeTab} onChange={(_, v) => setActiveTab(v)}>
            <Tab label="Revenue" />
            <Tab label="Users" />
            <Tab label="Games" />
          </Tabs>
        </Box>
        <CardContent>
          <ReportFilters
            activeTab={activeTab}
            dateRange={dateRange}
            setDateRange={setDateRange}
            period={period}
            setPeriod={setPeriod}
            category={category}
            setCategory={setCategory}
            onExport={handleExport}
          />
        </CardContent>
      </Card>

      {isLoading && <LinearProgress sx={{ mb: 2 }} />}

      {/* Revenue Report */}
      <TabPanel value={activeTab} index={0}>
        <Card>
          <CardContent>
            <Typography variant="h6" gutterBottom>Revenue Overview</Typography>
            <RevenueSummaryCards />
            <TableContainer>
              <Table>
                <TableHead>
                  <TableRow>
                    <TableCell>Date</TableCell>
                    <TableCell align="right">Deposits</TableCell>
                    <TableCell align="right">Withdrawals</TableCell>
                    <TableCell align="right">Net Revenue</TableCell>
                    <TableCell align="right">Bonus</TableCell>
                    <TableCell align="right">Rake</TableCell>
                  </TableRow>
                </TableHead>
                <TableBody>
                  {revenueReport.map((row: any, index: number) => (
                    <TableRow key={index}>
                      <TableCell>{row.date}</TableCell>
                      <TableCell align="right">${row.deposits.toLocaleString()}</TableCell>
                      <TableCell align="right">${row.withdrawals.toLocaleString()}</TableCell>
                      <TableCell align="right">
                        <Chip 
                          label={`$${row.netRevenue.toLocaleString()}`} 
                          color="success" 
                          size="small"
                        />
                      </TableCell>
                      <TableCell align="right">${row.bonus.toLocaleString()}</TableCell>
                      <TableCell align="right">${row.rake.toLocaleString()}</TableCell>
                    </TableRow>
                  ))}
                </TableBody>
              </Table>
            </TableContainer>
          </CardContent>
        </Card>
      </TabPanel>

      {/* User Report */}
      <TabPanel value={activeTab} index={1}>
        <Card>
          <CardContent>
            <Typography variant="h6" gutterBottom>User Statistics</Typography>
            <UserSummaryCards />
            <TableContainer>
              <Table>
                <TableHead>
                  <TableRow>
                    <TableCell>Date</TableCell>
                    <TableCell align="right">New Users</TableCell>
                    <TableCell align="right">Active Users</TableCell>
                    <TableCell align="right">KYC Approved</TableCell>
                    <TableCell align="right">Deposits</TableCell>
                    <TableCell align="right">Conversion</TableCell>
                  </TableRow>
                </TableHead>
                <TableBody>
                  {userReport.map((row: any, index: number) => (
                    <TableRow key={index}>
                      <TableCell>{row.date}</TableCell>
                      <TableCell align="right">
                        <Chip label={row.newUsers} color="info" size="small" />
                      </TableCell>
                      <TableCell align="right">{row.activeUsers.toLocaleString()}</TableCell>
                      <TableCell align="right">{row.kycApproved}</TableCell>
                      <TableCell align="right">{row.deposits}</TableCell>
                      <TableCell align="right">
                        <Chip 
                          label={`${row.conversions}%`} 
                          color={row.conversions > 60 ? 'success' : 'warning'} 
                          size="small"
                        />
                      </TableCell>
                    </TableRow>
                  ))}
                </TableBody>
              </Table>
            </TableContainer>
          </CardContent>
        </Card>
      </TabPanel>

      {/* Game Report */}
      <TabPanel value={activeTab} index={2}>
        <Card>
          <CardContent>
            <Typography variant="h6" gutterBottom>Game Performance</Typography>
            <TableContainer>
              <Table>
                <TableHead>
                  <TableRow>
                    <TableCell>Game</TableCell>
                    <TableCell>Provider</TableCell>
                    <TableCell>Category</TableCell>
                    <TableCell align="right">Revenue</TableCell>
                    <TableCell align="right">Total Plays</TableCell>
                    <TableCell align="right">RTP</TableCell>
                  </TableRow>
                </TableHead>
                <TableBody>
                  {gameReport.map((row: any, index: number) => (
                    <TableRow key={index}>
                      <TableCell>
                        <Typography fontWeight={500}>{row.game}</Typography>
                      </TableCell>
                      <TableCell>{row.provider}</TableCell>
                      <TableCell>
                        <Chip label={row.category} size="small" variant="outlined" />
                      </TableCell>
                      <TableCell align="right">
                        <Typography fontWeight="bold" color="success.main">
                          ${row.revenue.toLocaleString()}
                        </Typography>
                      </TableCell>
                      <TableCell align="right">{row.plays.toLocaleString()}</TableCell>
                      <TableCell align="right">
                        <Chip 
                          label={`${row.rtp}%`}
                          color={row.rtp >= 96 ? 'success' : row.rtp >= 94 ? 'warning' : 'error'}
                          size="small"
                        />
                      </TableCell>
                    </TableRow>
                  ))}
                </TableBody>
              </Table>
            </TableContainer>
          </CardContent>
        </Card>
      </TabPanel>
    </Box>
  );
};

export default Reports;
