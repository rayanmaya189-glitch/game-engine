import { useState } from 'react';
import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query';
import {
  Box, Card, CardContent, Typography, Grid, Table, TableBody, TableCell,
  TableContainer, TableHead, TableRow, Button, Chip, IconButton, Tooltip,
  Dialog, DialogTitle, DialogContent, DialogActions, TextField, MenuItem,
  Alert, Snackbar
} from '@mui/material';
import {
  Add, PlayArrow, Stop, Visibility, Videocam
} from '@mui/icons-material';
import { liveDealerAPI } from '../../services/api';

interface DealerTable {
  id: string;
  dealerName: string;
  gameType: string;
  status: 'active' | 'occupied' | 'empty';
  playersCount: number;
  minBet: number;
  maxBet: number;
  sessionId?: string;
}

interface CreateTableForm {
  dealerName: string;
  gameType: string;
  minBet: number;
  maxBet: number;
}

const GAME_TYPES = ['Blackjack', 'Roulette', 'Baccarat', 'Poker', 'Dragon Tiger', 'Sic Bo'];

const statusColor = (status: string) => {
  switch (status) {
    case 'active': return 'success';
    case 'occupied': return 'warning';
    case 'empty': return 'default';
    default: return 'default';
  }
};

const LiveDealer = () => {
  const queryClient = useQueryClient();
  const [createDialogOpen, setCreateDialogOpen] = useState(false);
  const [snackbar, setSnackbar] = useState({ open: false, message: '', severity: 'success' as 'success' | 'error' });
  const [form, setForm] = useState<CreateTableForm>({ dealerName: '', gameType: '', minBet: 10, maxBet: 1000 });

  const { data, isLoading } = useQuery({
    queryKey: ['live-dealer-tables'],
    queryFn: () => liveDealerAPI.getTables(),
  });

  const tables: DealerTable[] = data?.data?.tables || [];

  const createMutation = useMutation({
    mutationFn: liveDealerAPI.createTable,
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['live-dealer-tables'] });
      setCreateDialogOpen(false);
      setForm({ dealerName: '', gameType: '', minBet: 10, maxBet: 1000 });
      setSnackbar({ open: true, message: 'Table created successfully', severity: 'success' });
    },
    onError: () => setSnackbar({ open: true, message: 'Failed to create table', severity: 'error' }),
  });

  const startMutation = useMutation({
    mutationFn: liveDealerAPI.startSession,
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['live-dealer-tables'] });
      setSnackbar({ open: true, message: 'Session started', severity: 'success' });
    },
    onError: () => setSnackbar({ open: true, message: 'Failed to start session', severity: 'error' }),
  });

  const endMutation = useMutation({
    mutationFn: liveDealerAPI.endSession,
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['live-dealer-tables'] });
      setSnackbar({ open: true, message: 'Session ended', severity: 'success' });
    },
    onError: () => setSnackbar({ open: true, message: 'Failed to end session', severity: 'error' }),
  });

  const handleCreate = () => {
    if (!form.dealerName || !form.gameType) return;
    createMutation.mutate(form);
  };

  const activeTables = tables.filter(t => t.status === 'active').length;
  const occupiedTables = tables.filter(t => t.status === 'occupied').length;
  const totalPlayers = tables.reduce((sum, t) => sum + t.playersCount, 0);

  return (
    <Box>
      <Box sx={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', mb: 3 }}>
        <Typography variant="h4" fontWeight="bold">Live Dealer</Typography>
        <Button variant="contained" startIcon={<Add />} onClick={() => setCreateDialogOpen(true)}>
          New Table
        </Button>
      </Box>

      <Grid container spacing={3} sx={{ mb: 3 }}>
        <Grid item xs={12} md={3}>
          <Card><CardContent>
            <Typography color="text.secondary">Total Tables</Typography>
            <Typography variant="h4">{tables.length}</Typography>
          </CardContent></Card>
        </Grid>
        <Grid item xs={12} md={3}>
          <Card><CardContent>
            <Typography color="text.secondary">Active Sessions</Typography>
            <Typography variant="h4" color="success.main">{activeTables}</Typography>
          </CardContent></Card>
        </Grid>
        <Grid item xs={12} md={3}>
          <Card><CardContent>
            <Typography color="text.secondary">Occupied Tables</Typography>
            <Typography variant="h4" color="warning.main">{occupiedTables}</Typography>
          </CardContent></Card>
        </Grid>
        <Grid item xs={12} md={3}>
          <Card><CardContent>
            <Typography color="text.secondary">Total Players</Typography>
            <Typography variant="h4">{totalPlayers}</Typography>
          </CardContent></Card>
        </Grid>
      </Grid>

      <Card>
        <TableContainer>
          <Table>
            <TableHead>
              <TableRow>
                <TableCell>Table ID</TableCell>
                <TableCell>Dealer</TableCell>
                <TableCell>Game Type</TableCell>
                <TableCell>Status</TableCell>
                <TableCell>Players</TableCell>
                <TableCell>Bet Range</TableCell>
                <TableCell align="right">Actions</TableCell>
              </TableRow>
            </TableHead>
            <TableBody>
              {isLoading ? (
                <TableRow><TableCell colSpan={7} align="center">Loading...</TableCell></TableRow>
              ) : tables.length === 0 ? (
                <TableRow><TableCell colSpan={7} align="center">No tables found</TableCell></TableRow>
              ) : tables.map((table) => (
                <TableRow key={table.id} hover>
                  <TableCell><Typography variant="body2" fontFamily="monospace">{table.id.slice(0, 8)}</Typography></TableCell>
                  <TableCell><Box sx={{ display: 'flex', alignItems: 'center', gap: 1 }}><Videocam color="primary" fontSize="small" />{table.dealerName}</Box></TableCell>
                  <TableCell><Chip label={table.gameType} size="small" /></TableCell>
                  <TableCell><Chip label={table.status} color={statusColor(table.status)} size="small" /></TableCell>
                  <TableCell>{table.playersCount}</TableCell>
                  <TableCell>${table.minBet} - ${table.maxBet}</TableCell>
                  <TableCell align="right">
                    {table.status === 'empty' ? (
                      <Tooltip title="Start Session">
                        <IconButton size="small" color="success" onClick={() => startMutation.mutate(table.id)}>
                          <PlayArrow />
                        </IconButton>
                      </Tooltip>
                    ) : (
                      <Tooltip title="End Session">
                        <IconButton size="small" color="error" onClick={() => endMutation.mutate(table.id)}>
                          <Stop />
                        </IconButton>
                      </Tooltip>
                    )}
                    <Tooltip title="View Details">
                      <IconButton size="small"><Visibility /></IconButton>
                    </Tooltip>
                  </TableCell>
                </TableRow>
              ))}
            </TableBody>
          </Table>
        </TableContainer>
      </Card>

      <Dialog open={createDialogOpen} onClose={() => setCreateDialogOpen(false)} maxWidth="sm" fullWidth>
        <DialogTitle>Create New Table</DialogTitle>
        <DialogContent>
          <Box sx={{ display: 'flex', flexDirection: 'column', gap: 2, mt: 1 }}>
            <TextField label="Dealer Name" value={form.dealerName} onChange={(e) => setForm({ ...form, dealerName: e.target.value })} fullWidth required />
            <TextField label="Game Type" value={form.gameType} onChange={(e) => setForm({ ...form, gameType: e.target.value })} select fullWidth required>
              {GAME_TYPES.map(t => <MenuItem key={t} value={t}>{t}</MenuItem>)}
            </TextField>
            <TextField label="Min Bet" type="number" value={form.minBet} onChange={(e) => setForm({ ...form, minBet: Number(e.target.value) })} fullWidth />
            <TextField label="Max Bet" type="number" value={form.maxBet} onChange={(e) => setForm({ ...form, maxBet: Number(e.target.value) })} fullWidth />
          </Box>
        </DialogContent>
        <DialogActions>
          <Button onClick={() => setCreateDialogOpen(false)}>Cancel</Button>
          <Button variant="contained" onClick={handleCreate} disabled={createMutation.isPending}>Create</Button>
        </DialogActions>
      </Dialog>

      <Snackbar open={snackbar.open} autoHideDuration={4000} onClose={() => setSnackbar(s => ({ ...s, open: false }))}>
        <Alert severity={snackbar.severity} variant="filled">{snackbar.message}</Alert>
      </Snackbar>
    </Box>
  );
};

export default LiveDealer;
