import { useState } from 'react';
import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query';
import {
  Box, Card, CardContent, Typography, Grid, TextField, InputAdornment,
  Table, TableBody, TableCell, TableContainer, TableHead, TableRow,
  Button, Chip, IconButton, Tooltip, Avatar, LinearProgress
} from '@mui/material';
import { Search, Add, Edit, Visibility, TrendingUp } from '@mui/icons-material';
import { useAppDispatch } from '../../store/hooks';
import { showSnackbar } from '../../store/slices/uiSlice';
import { agentsAPI } from '../../services/api';

const Agents = () => {
  const dispatch = useAppDispatch();
  const queryClient = useQueryClient();
  const [search, setSearch] = useState('');

  const { data, isLoading } = useQuery({
    queryKey: ['agents', search],
    queryFn: () => agentsAPI.getAll({ search, page: 1, limit: 20 }),
  });

  const agents = data?.data?.agents || [];

  return (
    <Box>
      <Box sx={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', mb: 3 }}>
        <Typography variant="h4" fontWeight="bold">Agents Management</Typography>
        <Button variant="contained" startIcon={<Add />}>Add Agent</Button>
      </Box>

      <Grid container spacing={3} sx={{ mb: 3 }}>
        <Grid item xs={12} md={3}>
          <Card><CardContent><Typography color="text.secondary">Total Agents</Typography><Typography variant="h4">24</Typography></CardContent></Card>
        </Grid>
        <Grid item xs={12} md={3}>
          <Card><CardContent><Typography color="text.secondary">Active Agents</Typography><Typography variant="h4">20</Typography></CardContent></Card>
        </Grid>
        <Grid item xs={12} md={3}>
          <Card><CardContent><Typography color="text.secondary">Total Players</Typography><Typography variant="h4">1,540</Typography></CardContent></Card>
        </Grid>
        <Grid item xs={12} md={3}>
          <Card><CardContent><Typography color="text.secondary">Total Revenue</Typography><Typography variant="h4">$245,000</Typography></CardContent></Card>
        </Grid>
      </Grid>

      <Card sx={{ mb: 3 }}>
        <CardContent>
          <TextField
            fullWidth
            placeholder="Search agents..."
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
                <TableCell>Agent</TableCell>
                <TableCell>Email</TableCell>
                <TableCell>Tier</TableCell>
                <TableCell>Players</TableCell>
                <TableCell>Revenue</TableCell>
                <TableCell>Commission</TableCell>
                <TableCell>Status</TableCell>
                <TableCell align="right">Actions</TableCell>
              </TableRow>
            </TableHead>
            <TableBody>
              {agents.map((agent) => (
                <TableRow key={agent.id} hover>
                  <TableCell>
                    <Box sx={{ display: 'flex', alignItems: 'center', gap: 1 }}>
                      <Avatar sx={{ width: 32, height: 32 }}>{agent.name[0]}</Avatar>
                      {agent.name}
                    </Box>
                  </TableCell>
                  <TableCell>{agent.email}</TableCell>
                  <TableCell><Chip label={agent.tier} size="small" color={agent.tier === 'Platinum' ? 'secondary' : agent.tier === 'Gold' ? 'warning' : 'default'} /></TableCell>
                  <TableCell>{agent.players}</TableCell>
                  <TableCell>${agent.revenue.toLocaleString()}</TableCell>
                  <TableCell>${agent.commission.toLocaleString()}</TableCell>
                  <TableCell><Chip label={agent.status} color="success" size="small" /></TableCell>
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

export default Agents;
