import { Box, Paper, Typography } from '@mui/material';
import { TrendingUp, TrendingDown, People, Games, AttachMoney, MonetizationOn } from '@mui/icons-material';

export const RevenueSummaryCards = () => (
  <Box sx={{ display: 'flex', gap: 2, flexWrap: 'wrap', mb: 3 }}>
    <Paper sx={{ p: 2, display: 'flex', alignItems: 'center', gap: 2, flex: 1, minWidth: 200 }}>
      <AttachMoney sx={{ color: 'success.main', fontSize: 32 }} />
      <Box>
        <Typography variant="body2" color="text.secondary">Total Deposits</Typography>
        <Typography variant="h6" fontWeight="bold">$72,900</Typography>
      </Box>
    </Paper>
    <Paper sx={{ p: 2, display: 'flex', alignItems: 'center', gap: 2, flex: 1, minWidth: 200 }}>
      <MonetizationOn sx={{ color: 'error.main', fontSize: 32 }} />
      <Box>
        <Typography variant="body2" color="text.secondary">Total Withdrawals</Typography>
        <Typography variant="h6" fontWeight="bold">$47,600</Typography>
      </Box>
    </Paper>
    <Paper sx={{ p: 2, display: 'flex', alignItems: 'center', gap: 2, flex: 1, minWidth: 200 }}>
      <TrendingUp sx={{ color: 'primary.main', fontSize: 32 }} />
      <Box>
        <Typography variant="body2" color="text.secondary">Net Revenue</Typography>
        <Typography variant="h6" fontWeight="bold">$25,300</Typography>
      </Box>
    </Paper>
    <Paper sx={{ p: 2, display: 'flex', alignItems: 'center', gap: 2, flex: 1, minWidth: 200 }}>
      <Games sx={{ color: 'warning.main', fontSize: 32 }} />
      <Box>
        <Typography variant="body2" color="text.secondary">Total Rake</Typography>
        <Typography variant="h6" fontWeight="bold">$17,800</Typography>
      </Box>
    </Paper>
  </Box>
);

export const UserSummaryCards = () => (
  <Box sx={{ display: 'flex', gap: 2, flexWrap: 'wrap', mb: 3 }}>
    <Paper sx={{ p: 2, display: 'flex', alignItems: 'center', gap: 2, flex: 1, minWidth: 200 }}>
      <People sx={{ color: 'info.main', fontSize: 32 }} />
      <Box>
        <Typography variant="body2" color="text.secondary">New Users</Typography>
        <Typography variant="h6" fontWeight="bold">251</Typography>
      </Box>
    </Paper>
    <Paper sx={{ p: 2, display: 'flex', alignItems: 'center', gap: 2, flex: 1, minWidth: 200 }}>
      <TrendingUp sx={{ color: 'success.main', fontSize: 32 }} />
      <Box>
        <Typography variant="body2" color="text.secondary">Active Users</Typography>
        <Typography variant="h6" fontWeight="bold">6,447</Typography>
      </Box>
    </Paper>
    <Paper sx={{ p: 2, display: 'flex', alignItems: 'center', gap: 2, flex: 1, minWidth: 200 }}>
      <TrendingDown sx={{ color: 'warning.main', fontSize: 32 }} />
      <Box>
        <Typography variant="body2" color="text.secondary">KYC Approved</Typography>
        <Typography variant="h6" fontWeight="bold">67</Typography>
      </Box>
    </Paper>
    <Paper sx={{ p: 2, display: 'flex', alignItems: 'center', gap: 2, flex: 1, minWidth: 200 }}>
      <AttachMoney sx={{ color: 'primary.main', fontSize: 32 }} />
      <Box>
        <Typography variant="body2" color="text.secondary">Conversions</Typography>
        <Typography variant="h6" fontWeight="bold">65.9%</Typography>
      </Box>
    </Paper>
  </Box>
);
