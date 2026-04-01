import { useState } from 'react';
import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query';
import { 
  Box, Typography, Card, CardContent, TextField, InputAdornment, 
  Table, TableBody, TableCell, TableContainer, TableHead, TableRow, 
  Chip, IconButton, Tooltip, Avatar, MenuItem, Select, FormControl, InputLabel
} from '@mui/material';
import { Search, Visibility, Edit, Block, CheckCircle, PersonAdd } from '@mui/icons-material';
import { useAppDispatch } from '../../store/hooks';
import { showSnackbar } from '../../store/slices/uiSlice';
import { playersAPI } from '../../services/api';

const Users = () => {
  const dispatch = useAppDispatch();
  const queryClient = useQueryClient();
  const [search, setSearch] = useState('');
  const [kycFilter, setKycFilter] = useState('');
  const [statusFilter, setStatusFilter] = useState('');

  const { data, isLoading } = useQuery({
    queryKey: ['players', search, kycFilter, statusFilter],
    queryFn: () => playersAPI.getAll({ search, page: 1, limit: 50, status: statusFilter }),
  });

  const updateStatusMutation = useMutation({
    mutationFn: ({ id, status }: { id: string; status: string }) =>
      playersAPI.updateStatus(id, { status }),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['players'] });
      dispatch(showSnackbar({ message: 'Player status updated', severity: 'success' }));
    },
    onError: (error: any) => {
      dispatch(showSnackbar({ message: error?.response?.data?.error || error.message, severity: 'error' }));
    },
  });

  const players = data?.data?.data?.players || data?.data?.players || [];
  const total = data?.data?.data?.total || data?.data?.total || 0;

  const getKycColor = (kyc: string) => {
    switch (kyc) {
      case 'verified': return 'success';
      case 'pending': return 'warning';
      case 'none': return 'default';
      default: return 'default';
    }
  };

  const getStatusColor = (status: string) => {
    switch (status) {
      case 'active': return 'success';
      case 'suspended': return 'error';
      case 'banned': return 'error';
      default: return 'default';
    }
  };

  return (
    <Box>
      <Box sx={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', mb: 3 }}>
        <Box>
          <Typography variant="h4" fontWeight="bold">
            Players Management
          </Typography>
          <Typography variant="body2" color="text.secondary">
            {total} total players
          </Typography>
        </Box>
        <Tooltip title="Add New Player">
          <IconButton color="primary">
            <PersonAdd />
          </IconButton>
        </Tooltip>
      </Box>

      <Card sx={{ mb: 3 }}>
        <CardContent>
          <Box sx={{ display: 'flex', gap: 2, flexWrap: 'wrap' }}>
            <TextField
              placeholder="Search players by name, email, or ID..."
              value={search}
              onChange={(e) => setSearch(e.target.value)}
              sx={{ flex: 1, minWidth: 250 }}
              InputProps={{
                startAdornment: (
                  <InputAdornment position="start">
                    <Search />
                  </InputAdornment>
                ),
              }}
            />
            <FormControl sx={{ minWidth: 150 }}>
              <InputLabel>KYC Status</InputLabel>
              <Select
                value={kycFilter}
                label="KYC Status"
                onChange={(e) => setKycFilter(e.target.value)}
              >
                <MenuItem value="">All</MenuItem>
                <MenuItem value="verified">Verified</MenuItem>
                <MenuItem value="pending">Pending</MenuItem>
                <MenuItem value="none">None</MenuItem>
              </Select>
            </FormControl>
            <FormControl sx={{ minWidth: 150 }}>
              <InputLabel>Status</InputLabel>
              <Select
                value={statusFilter}
                label="Status"
                onChange={(e) => setStatusFilter(e.target.value)}
              >
                <MenuItem value="">All</MenuItem>
                <MenuItem value="active">Active</MenuItem>
                <MenuItem value="suspended">Suspended</MenuItem>
                <MenuItem value="banned">Banned</MenuItem>
              </Select>
            </FormControl>
          </Box>
        </CardContent>
      </Card>

      <Card>
        <TableContainer>
          <Table>
            <TableHead>
              <TableRow>
                <TableCell>Player</TableCell>
                <TableCell>Email</TableCell>
                <TableCell>Status</TableCell>
                <TableCell>KYC</TableCell>
                <TableCell align="right">Balance</TableCell>
                <TableCell>Joined</TableCell>
                <TableCell align="center">Actions</TableCell>
              </TableRow>
            </TableHead>
            <TableBody>
              {isLoading ? (
                <TableRow>
                  <TableCell colSpan={7} align="center" sx={{ py: 4 }}>
                    Loading players...
                  </TableCell>
                </TableRow>
              ) : players.length === 0 ? (
                <TableRow>
                  <TableCell colSpan={7} align="center" sx={{ py: 4 }}>
                    No players found
                  </TableCell>
                </TableRow>
              ) : (
                players.map((player: any) => (
                  <TableRow key={player.id || player.user_id} hover>
                    <TableCell>
                      <Box sx={{ display: 'flex', alignItems: 'center', gap: 1.5 }}>
                        <Avatar sx={{ bgcolor: 'primary.main', width: 36, height: 36 }}>
                          {(player.name || player.username || 'P').slice(0, 2).toUpperCase()}
                        </Avatar>
                        <Box>
                          <Typography fontWeight={500}>{player.name || player.username}</Typography>
                          <Typography variant="caption" color="text.secondary">{player.id || player.user_id}</Typography>
                        </Box>
                      </Box>
                    </TableCell>
                    <TableCell>{player.email}</TableCell>
                    <TableCell>
                      <Chip 
                        label={player.status} 
                        size="small" 
                        color={getStatusColor(player.status) as any}
                      />
                    </TableCell>
                    <TableCell>
                      <Chip 
                        label={player.kyc_level || player.kyc_status || 'none'} 
                        size="small" 
                        variant="outlined"
                        color={getKycColor(player.kyc_level || player.kyc_status) as any}
                      />
                    </TableCell>
                    <TableCell align="right">
                      <Typography fontWeight={600} color={player.balance > 0 ? 'success.main' : 'text.secondary'}>
                        ${(player.balance || 0).toLocaleString('en-US', { minimumFractionDigits: 2 })}
                      </Typography>
                    </TableCell>
                    <TableCell>{player.created_at || player.joined}</TableCell>
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
                      <Tooltip title={player.status === 'active' ? 'Suspend' : 'Activate'}>
                        <IconButton 
                          size="small" 
                          color={player.status === 'active' ? 'error' : 'success'}
                          onClick={() => {
                            const newStatus = player.status === 'active' ? 'suspended' : 'active';
                            updateStatusMutation.mutate({ id: player.id || player.user_id, status: newStatus });
                          }}
                        >
                          {player.status === 'active' ? <Block fontSize="small" /> : <CheckCircle fontSize="small" />}
                        </IconButton>
                      </Tooltip>
                    </TableCell>
                  </TableRow>
                ))
              )}
            </TableBody>
          </Table>
        </TableContainer>
      </Card>
    </Box>
  );
};

export default Users;
