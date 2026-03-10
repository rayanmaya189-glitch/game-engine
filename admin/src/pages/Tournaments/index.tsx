import { useState } from 'react';
import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query';
import {
  Box, Card, CardContent, Typography, Grid, TextField, InputAdornment,
  Table, TableBody, TableCell, TableContainer, TableHead, TableRow,
  Button, Chip, IconButton, Tooltip, Avatar, AvatarGroup
} from '@mui/material';
import { Search, Add, Edit, Visibility, EmojiEvents, Timer } from '@mui/icons-material';
import { useAppDispatch } from '../../store/hooks';
import { showSnackbar } from '../../store/slices/uiSlice';
import { tournamentsAPI } from '../../services/api';

const Tournaments = () => {
  const dispatch = useAppDispatch();
  const queryClient = useQueryClient();
  const [search, setSearch] = useState('');

  const { data, isLoading } = useQuery({
    queryKey: ['tournaments', search],
    queryFn: () => tournamentsAPI.getAll({ search, page: 1, limit: 20 }),
  });

  const tournaments = data?.data?.tournaments || [];

  const getStatusColor = (status: string) => {
    switch (status) {
      case 'active': return 'success';
      case 'scheduled': return 'info';
      case 'completed': return 'default';
      case 'cancelled': return 'error';
      default: return 'default';
    }
  };

  return (
    <Box>
      <Box sx={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', mb: 3 }}>
        <Typography variant="h4" fontWeight="bold">Tournaments Management</Typography>
        <Button variant="contained" startIcon={<Add />}>Create Tournament</Button>
      </Box>

      <Grid container spacing={3} sx={{ mb: 3 }}>
        <Grid item xs={12} md={3}>
          <Card><CardContent><Typography color="text.secondary">Active Tournaments</Typography><Typography variant="h4">3</Typography></CardContent></Card>
        </Grid>
        <Grid item xs={12} md={3}>
          <Card><CardContent><Typography color="text.secondary">Total Prize Pool</Typography><Typography variant="h4">$67,500</Typography></CardContent></Card>
        </Grid>
        <Grid item xs={12} md={3}>
          <Card><CardContent><Typography color="text.secondary">Total Participants</Typography><Typography variant="h4">290</Typography></CardContent></Card>
        </Grid>
        <Grid item xs={12} md={3}>
          <Card><CardContent><Typography color="text.secondary">Completed</Typography><Typography variant="h4">12</Typography></CardContent></Card>
        </Grid>
      </Grid>

      <Card sx={{ mb: 3 }}>
        <CardContent>
          <TextField
            fullWidth
            placeholder="Search tournaments..."
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
                <TableCell>Tournament</TableCell>
                <TableCell>Game</TableCell>
                <TableCell>Prize Pool</TableCell>
                <TableCell>Players</TableCell>
                <TableCell>Status</TableCell>
                <TableCell>Start Date</TableCell>
                <TableCell>End Date</TableCell>
                <TableCell align="right">Actions</TableCell>
              </TableRow>
            </TableHead>
            <TableBody>
              {tournaments.map((tournament) => (
                <TableRow key={tournament.id} hover>
                  <TableCell fontWeight="500">{tournament.name}</TableCell>
                  <TableCell>{tournament.game}</TableCell>
                  <TableCell><Box sx={{ display: 'flex', alignItems: 'center', gap: 0.5 }}><EmojiEvents color="warning" /> ${tournament.prizePool.toLocaleString()}</Box></TableCell>
                  <TableCell>{tournament.players}</TableCell>
                  <TableCell><Chip label={tournament.status} color={getStatusColor(tournament.status) as any} size="small" /></TableCell>
                  <TableCell>{tournament.startDate}</TableCell>
                  <TableCell>{tournament.endDate}</TableCell>
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

export default Tournaments;
