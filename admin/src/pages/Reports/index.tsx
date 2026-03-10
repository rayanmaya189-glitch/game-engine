import { useState } from 'react';
import { useQuery } from '@tanstack/react-query';
import { 
  Box, Typography, Card, CardContent, Grid, TextField, Button, 
  FormControl, InputLabel, Select, MenuItem, Table, TableBody, 
  TableCell, TableContainer, TableHead, TableRow, Chip, Tabs, Tab,
  LinearProgress, Paper
} from '@mui/material';
import { Download, TrendingUp, TrendingDown, People, Games, AttachMoney, MonetizationOn } from '@mui/icons-material';
import { reportsAPI } from '../../services/api';

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
          <Grid container spacing={2} alignItems="center">
            <Grid item xs={12} md={3}>
              <TextField 
                type="date" 
                label="From" 
                fullWidth 
                size="small"
                value={dateRange.from}
                onChange={(e) => setDateRange({ ...dateRange, from: e.target.value })}
                InputLabelProps={{ shrink: true }}
              />
            </Grid>
            <Grid item xs={12} md={3}>
              <TextField 
                type="date" 
                label="To" 
                fullWidth 
                size="small"
                value={dateRange.to}
                onChange={(e) => setDateRange({ ...dateRange, to: e.target.value })}
                InputLabelProps={{ shrink: true }}
              />
            </Grid>
            {activeTab === 1 && (
              <Grid item xs={12} md={3}>
                <FormControl fullWidth size="small">
                  <InputLabel>Period</InputLabel>
                  <Select value={period} label="Period" onChange={(e) => setPeriod(e.target.value)}>
                    <MenuItem value="7">Last 7 days</MenuItem>
                    <MenuItem value="30">Last 30 days</MenuItem>
                    <MenuItem value="90">Last 90 days</MenuItem>
                    <MenuItem value="365">Last year</MenuItem>
                  </Select>
                </FormControl>
              </Grid>
            )}
            {activeTab === 2 && (
              <Grid item xs={12} md={3}>
                <FormControl fullWidth size="small">
                  <InputLabel>Category</InputLabel>
                  <Select value={category} label="Category" onChange={(e) => setCategory(e.target.value)}>
                    <MenuItem value="">All Games</MenuItem>
                    <MenuItem value="slot">Slots</MenuItem>
                    <MenuItem value="live_casino">Live Casino</MenuItem>
                    <MenuItem value="table_games">Table Games</MenuItem>
                    <MenuItem value="dice">Dice</MenuItem>
                  </Select>
                </FormControl>
              </Grid>
            )}
            <Grid item xs={12} md={activeTab === 0 ? 6 : 3}>
              <Button variant="contained" startIcon={<Download />} onClick={handleExport}>
                Export
              </Button>
            </Grid>
          </Grid>
        </CardContent>
      </Card>

      {isLoading && <LinearProgress sx={{ mb: 2 }} />}

      {/* Revenue Report */}
      <TabPanel value={activeTab} index={0}>
        <Card>
          <CardContent>
            <Typography variant="h6" gutterBottom>Revenue Overview</Typography>
            <Grid container spacing={2} sx={{ mb: 3 }}>
              <Grid item xs={12} md={3}>
                <Paper sx={{ p: 2, display: 'flex', alignItems: 'center', gap: 2 }}>
                  <AttachMoney sx={{ color: 'success.main', fontSize: 32 }} />
                  <Box>
                    <Typography variant="body2" color="text.secondary">Total Deposits</Typography>
                    <Typography variant="h6" fontWeight="bold">$72,900</Typography>
                  </Box>
                </Paper>
              </Grid>
              <Grid item xs={12} md={3}>
                <Paper sx={{ p: 2, display: 'flex', alignItems: 'center', gap: 2 }}>
                  <MonetizationOn sx={{ color: 'error.main', fontSize: 32 }} />
                  <Box>
                    <Typography variant="body2" color="text.secondary">Total Withdrawals</Typography>
                    <Typography variant="h6" fontWeight="bold">$47,600</Typography>
                  </Box>
                </Paper>
              </Grid>
              <Grid item xs={12} md={3}>
                <Paper sx={{ p: 2, display: 'flex', alignItems: 'center', gap: 2 }}>
                  <TrendingUp sx={{ color: 'primary.main', fontSize: 32 }} />
                  <Box>
                    <Typography variant="body2" color="text.secondary">Net Revenue</Typography>
                    <Typography variant="h6" fontWeight="bold">$25,300</Typography>
                  </Box>
                </Paper>
              </Grid>
              <Grid item xs={12} md={3}>
                <Paper sx={{ p: 2, display: 'flex', alignItems: 'center', gap: 2 }}>
                  <Games sx={{ color: 'warning.main', fontSize: 32 }} />
                  <Box>
                    <Typography variant="body2" color="text.secondary">Total Rake</Typography>
                    <Typography variant="h6" fontWeight="bold">$17,800</Typography>
                  </Box>
                </Paper>
              </Grid>
            </Grid>
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
            <Grid container spacing={2} sx={{ mb: 3 }}>
              <Grid item xs={12} md={3}>
                <Paper sx={{ p: 2, display: 'flex', alignItems: 'center', gap: 2 }}>
                  <People sx={{ color: 'info.main', fontSize: 32 }} />
                  <Box>
                    <Typography variant="body2" color="text.secondary">New Users</Typography>
                    <Typography variant="h6" fontWeight="bold">251</Typography>
                  </Box>
                </Paper>
              </Grid>
              <Grid item xs={12} md={3}>
                <Paper sx={{ p: 2, display: 'flex', alignItems: 'center', gap: 2 }}>
                  <TrendingUp sx={{ color: 'success.main', fontSize: 32 }} />
                  <Box>
                    <Typography variant="body2" color="text.secondary">Active Users</Typography>
                    <Typography variant="h6" fontWeight="bold">6,447</Typography>
                  </Box>
                </Paper>
              </Grid>
              <Grid item xs={12} md={3}>
                <Paper sx={{ p: 2, display: 'flex', alignItems: 'center', gap: 2 }}>
                  <TrendingDown sx={{ color: 'warning.main', fontSize: 32 }} />
                  <Box>
                    <Typography variant="body2" color="text.secondary">KYC Approved</Typography>
                    <Typography variant="h6" fontWeight="bold">67</Typography>
                  </Box>
                </Paper>
              </Grid>
              <Grid item xs={12} md={3}>
                <Paper sx={{ p: 2, display: 'flex', alignItems: 'center', gap: 2 }}>
                  <AttachMoney sx={{ color: 'primary.main', fontSize: 32 }} />
                  <Box>
                    <Typography variant="body2" color="text.secondary">Conversions</Typography>
                    <Typography variant="h6" fontWeight="bold">65.9%</Typography>
                  </Box>
                </Paper>
              </Grid>
            </Grid>
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
