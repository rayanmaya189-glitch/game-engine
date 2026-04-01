import {
  Box, Typography, Card, CardContent, Grid,
  Table, TableBody, TableCell, TableContainer, TableHead, TableRow, Chip, Paper
} from '@mui/material';
import { TrendingUp, TrendingDown, People, Games, AttachMoney, MonetizationOn } from '@mui/icons-material';
import 
interface ReportChartsProps {
  activeTab: number;
  revenueReport: any[];
  userReport: any[];
  gameReport: any[];
}

const TabPanel = ({ children, value, index }: { children: React.ReactNode; value: number; index: number }) => (
  <Box role="tabpanel" hidden={value !== index} sx={{ py: 3 }}>
    {value === index && children}
  </Box>
);

const ReportCharts = ({ activeTab, revenueReport, userReport, gameReport }: ReportChartsProps) => {
  return (
    <>
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
    </>
  );
};

export default ReportCharts;
