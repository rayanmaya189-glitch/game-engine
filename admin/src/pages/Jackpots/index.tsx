import { useState } from 'react';
import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query';
import {
  Box, Card, CardContent, Typography, Grid, TextField, InputAdornment,
  Table, TableBody, TableCell, TableContainer, TableHead, TableRow,
  Button, Chip, IconButton, Tooltip, LinearProgress
} from '@mui/material';
import { Search, Add, Edit, Visibility, Star, Casino } from '@mui/icons-material';
import { useAppDispatch } from '../../store/hooks';
import { showSnackbar } from '../../store/slices/uiSlice';
import { jackpotsAPI } from '../../services/api';

const Jackpots = () => {
  const dispatch = useAppDispatch();
  const queryClient = useQueryClient();
  const [search, setSearch] = useState('');

  const { data, isLoading } = useQuery({
    queryKey: ['jackpots', search],
    queryFn: () => jackpotsAPI.getAll({ search, page: 1, limit: 20 }),
  });

  const jackpots = data?.data?.jackpots || [];

  return (
    <Box>
      <Box sx={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', mb: 3 }}>
        <Typography variant="h4" fontWeight="bold">Jackpots Management</Typography>
        <Button variant="contained" startIcon={<Add />}>Create Jackpot</Button>
      </Box>

      <Grid container spacing={3} sx={{ mb: 3 }}>
        <Grid item xs={12} md={3}>
          <Card sx={{ background: 'linear-gradient(135deg, #667eea 0%, #764ba2 100%)' }}>
            <CardContent>
              <Typography color="white" variant="subtitle2">Total Jackpots</Typography>
              <Typography color="white" variant="h3">4</Typography>
            </CardContent>
          </Card>
        </Grid>
        <Grid item xs={12} md={3}>
          <Card><CardContent><Typography color="text.secondary">Total Paid Out</Typography><Typography variant="h4">$1,545,000</Typography></CardContent></Card>
        </Grid>
        <Grid item xs={12} md={3}>
          <Card><CardContent><Typography color="text.secondary">Total Hits</Typography><Typography variant="h4">1,061</Typography></CardContent></Card>
        </Grid>
        <Grid item xs={12} md={3}>
          <Card><CardContent><Typography color="text.secondary">Active Players</Typography><Typography variant="h4">234</Typography></CardContent></Card>
        </Grid>
      </Grid>

      <Card sx={{ mb: 3 }}>
        <CardContent>
          <TextField
            fullWidth
            placeholder="Search jackpots..."
            value={search}
            onChange={(e) => setSearch(e.target.value)}
            InputProps={{ startAdornment: <InputAdornment position="start"><Search /></InputAdornment> }}
          />
        </CardContent>
      </Card>

      <Card>
        <TableContainer>
          <Table>
            <TableHead>
              <TableRow>
                <TableCell>Jackpot Name</TableCell>
                <TableCell>Game</TableCell>
                <TableCell>Current Amount</TableCell>
                <TableCell>Min Bet</TableCell>
                <TableCell>Max Win</TableCell>
                <TableCell>Hits</TableCell>
                <TableCell>Status</TableCell>
                <TableCell align="right">Actions</TableCell>
              </TableRow>
            </TableHead>
            <TableBody>
              {jackpots.map((jackpot) => (
                <TableRow key={jackpot.id} hover>
                  <TableCell><Box sx={{ display: 'flex', alignItems: 'center', gap: 1 }}><Star color="warning" />{jackpot.name}</Box></TableCell>
                  <TableCell>{jackpot.game}</TableCell>
                  <TableCell><Typography color="success.main" fontWeight="bold">${jackpot.currentAmount.toLocaleString()}</Typography></TableCell>
                  <TableCell>${jackpot.minBet}</TableCell>
                  <TableCell>${jackpot.maxWin.toLocaleString()}</TableCell>
                  <TableCell>{jackpot.hits}</TableCell>
                  <TableCell><Chip label={jackpot.status} color="success" size="small" /></TableCell>
                  <TableCell align="right">
                    <Tooltip title="View"><IconButton size="small"><Visibility /></IconButton></Tooltip>
                    <Tooltip title="Edit"><IconButton size="small"><Edit /></IconButton></Tooltip>
                  </TableCell>
                </TableRow>
              ))}
            </TableBody>
          </Table>
        </TableContainer>
      </Card>
    </Box>
  );
};

export default Jackpots;
