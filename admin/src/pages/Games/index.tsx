import { useState } from 'react';
import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query';
import {
  Box, Card, CardContent, Typography, Grid, TextField, InputAdornment,
  Table, TableBody, TableCell, TableContainer, TableHead, TableRow,
  Button, Chip, IconButton, Tooltip, Switch, FormControlLabel, MenuItem
} from '@mui/material';
import { Search, Add, Edit, Visibility, Delete, Casino } from '@mui/icons-material';
import { useAppDispatch } from '../../store/hooks';
import { showSnackbar } from '../../store/slices/uiSlice';
import { gamesAPI } from '../../services/api';

const Games = () => {
  const dispatch = useAppDispatch();
  const queryClient = useQueryClient();
  const [search, setSearch] = useState('');
  const [category, setCategory] = useState('');
  const [status, setStatus] = useState('');

  const { data, isLoading } = useQuery({
    queryKey: ['games', search, category, status],
    queryFn: () => gamesAPI.getAll({ category, status, page: 1, limit: 50 }),
  });

  const deleteMutation = useMutation({
    mutationFn: (id: string) => gamesAPI.delete(id),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['games'] });
      dispatch(showSnackbar({ message: 'Game deleted successfully', severity: 'success' }));
    },
    onError: (error: any) => {
      dispatch(showSnackbar({ message: error.message, severity: 'error' }));
    },
  });

  const games = data?.data?.games || [
    { id: '1', name: 'Starburst', provider: 'NetEnt', category: 'slot', rtp: 96.1, minBet: 0.1, maxBet: 100, volatility: 'medium', status: 'active' },
    { id: '2', name: 'Mega Moolah', provider: 'Microgaming', category: 'slot', rtp: 88.12, minBet: 0.25, maxBet: 6.25, volatility: 'high', status: 'active' },
    { id: '3', name: 'Blackjack VIP', provider: 'Evolution', category: 'live_casino', rtp: 99.5, minBet: 10, maxBet: 5000, volatility: 'low', status: 'active' },
    { id: '4', name: 'Sic Bo', provider: 'Playtech', category: 'dice', rtp: 97.22, minBet: 1, maxBet: 1000, volatility: 'high', status: 'active' },
    { id: '5', name: 'Texas Hold\'em', provider: 'Evolution', category: 'table_games', rtp: 97.8, minBet: 5, maxBet: 2500, volatility: 'medium', status: 'inactive' },
    { id: '6', name: 'Baccarat', provider: 'Pragmatic', category: 'table_games', rtp: 98.94, minBet: 1, maxBet: 10000, volatility: 'low', status: 'active' },
    { id: '7', name: 'Roulette European', provider: 'NetEnt', category: 'table_games', rtp: 97.3, minBet: 1, maxBet: 500, volatility: 'medium', status: 'maintenance' },
    { id: '8', name: 'Crazy Time', provider: 'Evolution', category: 'live_casino', rtp: 96.23, minBet: 0.1, maxBet: 5000, volatility: 'high', status: 'active' },
  ];

  const getCategoryLabel = (cat: string) => {
    const labels: Record<string, string> = {
      slot: 'Slots',
      live_casino: 'Live Casino',
      table_games: 'Table Games',
      dice: 'Dice',
      card: 'Card Games',
    };
    return labels[cat] || cat;
  };

  const getStatusColor = (status: string) => {
    switch (status) {
      case 'active': return 'success';
      case 'inactive': return 'error';
      case 'maintenance': return 'warning';
      default: return 'default';
    }
  };

  const getVolatilityColor = (vol: string) => {
    switch (vol) {
      case 'high': return 'error';
      case 'medium': return 'warning';
      case 'low': return 'success';
      default: return 'default';
    }
  };

  return (
    <Box>
      <Box sx={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', mb: 3 }}>
        <Typography variant="h4" fontWeight="bold">
          Games Management
        </Typography>
        <Button variant="contained" startIcon={<Add />}>
          Add Game
        </Button>
      </Box>

      <Card sx={{ mb: 3 }}>
        <CardContent>
          <Grid container spacing={2}>
            <Grid item xs={12} md={4}>
              <TextField
                fullWidth
                placeholder="Search games..."
                value={search}
                onChange={(e) => setSearch(e.target.value)}
                InputProps={{
                  startAdornment: (
                    <InputAdornment position="start">
                      <Search />
                    </InputAdornment>
                  ),
                }}
              />
            </Grid>
            <Grid item xs={12} md={3}>
              <TextField
                fullWidth
                select
                label="Category"
                value={category}
                onChange={(e) => setCategory(e.target.value)}
              >
                <MenuItem value="">All</MenuItem>
                <MenuItem value="slot">Slots</MenuItem>
                <MenuItem value="live_casino">Live Casino</MenuItem>
                <MenuItem value="table_games">Table Games</MenuItem>
                <MenuItem value="dice">Dice</MenuItem>
                <MenuItem value="card">Card Games</MenuItem>
              </TextField>
            </Grid>
            <Grid item xs={12} md={3}>
              <TextField
                fullWidth
                select
                label="Status"
                value={status}
                onChange={(e) => setStatus(e.target.value)}
              >
                <MenuItem value="">All</MenuItem>
                <MenuItem value="active">Active</MenuItem>
                <MenuItem value="inactive">Inactive</MenuItem>
                <MenuItem value="maintenance">Maintenance</MenuItem>
              </TextField>
            </Grid>
          </Grid>
        </CardContent>
      </Card>

      <Card>
        <TableContainer>
          <Table>
            <TableHead>
              <TableRow>
                <TableCell>Game</TableCell>
                <TableCell>Provider</TableCell>
                <TableCell>Category</TableCell>
                <TableCell align="right">RTP %</TableCell>
                <TableCell align="right">Min Bet</TableCell>
                <TableCell align="right">Max Bet</TableCell>
                <TableCell>Volatility</TableCell>
                <TableCell>Status</TableCell>
                <TableCell align="center">Actions</TableCell>
              </TableRow>
            </TableHead>
            <TableBody>
              {games.map((game: any) => (
                <TableRow key={game.id} hover>
                  <TableCell>
                    <Box sx={{ display: 'flex', alignItems: 'center', gap: 1.5 }}>
                      <Box sx={{ 
                        width: 40, 
                        height: 40, 
                        borderRadius: 1, 
                        bgcolor: 'primary.light',
                        display: 'flex',
                        alignItems: 'center',
                        justifyContent: 'center'
                      }}>
                        <Casino sx={{ color: 'primary.main', fontSize: 20 }} />
                      </Box>
                      <Typography fontWeight={500}>{game.name}</Typography>
                    </Box>
                  </TableCell>
                  <TableCell>{game.provider}</TableCell>
                  <TableCell>
                    <Chip label={getCategoryLabel(game.category)} size="small" variant="outlined" />
                  </TableCell>
                  <TableCell align="right">
                    <Typography 
                      color={game.rtp >= 96 ? 'success.main' : game.rtp >= 94 ? 'warning.main' : 'error.main'}
                      fontWeight={600}
                    >
                      {game.rtp}%
                    </Typography>
                  </TableCell>
                  <TableCell align="right">${game.minBet}</TableCell>
                  <TableCell align="right">${game.maxBet}</TableCell>
                  <TableCell>
                    <Chip 
                      label={game.volatility} 
                      size="small" 
                      color={getVolatilityColor(game.volatility) as any}
                    />
                  </TableCell>
                  <TableCell>
                    <Chip 
                      label={game.status} 
                      size="small" 
                      color={getStatusColor(game.status) as any}
                    />
                  </TableCell>
                  <TableCell align="center">
                    <Tooltip title="View Details">
                      <IconButton size="small">
                        <Visibility fontSize="small" />
                      </IconButton>
                    </Tooltip>
                    <Tooltip title="Edit">
                      <IconButton size="small">
                        <Edit fontSize="small" />
                      </IconButton>
                    </Tooltip>
                    <Tooltip title="Delete">
                      <IconButton 
                        size="small" 
                        color="error"
                        onClick={() => {
                          if (window.confirm('Are you sure you want to delete this game?')) {
                            deleteMutation.mutate(game.id);
                          }
                        }}
                      >
                        <Delete fontSize="small" />
                      </IconButton>
                    </Tooltip>
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

export default Games;
