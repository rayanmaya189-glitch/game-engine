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
import { usersAPI } from '../../services/api';

const Users = () => {
  const dispatch = useAppDispatch();
  const queryClient = useQueryClient();
  const [search, setSearch] = useState('');
  const [kycFilter, setKycFilter] = useState('');
  const [statusFilter, setStatusFilter] = useState('');

  const { data, isLoading } = useQuery({
    queryKey: ['users', search, kycFilter, statusFilter],
    queryFn: () => usersAPI.getAll({ search, page: 1, limit: 50 }),
  });

  const updateStatusMutation = useMutation({
    mutationFn: ({ id, status }: { id: string; status: string }) =>
      usersAPI.updateStatus(id, { status }),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['users'] });
      dispatch(showSnackbar({ message: 'User status updated', severity: 'success' }));
    },
    onError: (error: any) => {
      dispatch(showSnackbar({ message: error.message, severity: 'error' }));
    },
  });

  const users = data?.data?.users || [
    { id: 'USR001', name: 'John Doe', email: 'john@example.com', status: 'active', kycLevel: 'verified', balance: 1250.50, joined: '2024-01-15', avatar: 'JD' },
    { id: 'USR002', name: 'Sarah Smith', email: 'sarah@example.com', status: 'active', kycLevel: 'pending', balance: 500.00, joined: '2024-01-18', avatar: 'SS' },
    { id: 'USR003', name: 'Mike Johnson', email: 'mike@example.com', status: 'suspended', kycLevel: 'verified', balance: 0, joined: '2024-01-20', avatar: 'MJ' },
    { id: 'USR004', name: 'Emily Brown', email: 'emily@example.com', status: 'active', kycLevel: 'none', balance: 250.75, joined: '2024-01-22', avatar: 'EB' },
    { id: 'USR005', name: 'David Wilson', email: 'david@example.com', status: 'active', kycLevel: 'verified', balance: 3500.00, joined: '2024-01-25', avatar: 'DW' },
  ];

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
        <Typography variant="h4" fontWeight="bold">
          Users Management
        </Typography>
        <Tooltip title="Add New User">
          <IconButton color="primary">
            <PersonAdd />
          </IconButton>
        </Tooltip>
      </Box>

      <Card sx={{ mb: 3 }}>
        <CardContent>
          <Box sx={{ display: 'flex', gap: 2, flexWrap: 'wrap' }}>
            <TextField
              placeholder="Search users by name, email, or ID..."
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
              <InputLabel>User Status</InputLabel>
              <Select
                value={statusFilter}
                label="User Status"
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
                <TableCell>User</TableCell>
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
                    Loading users...
                  </TableCell>
                </TableRow>
              ) : users.length === 0 ? (
                <TableRow>
                  <TableCell colSpan={7} align="center" sx={{ py: 4 }}>
                    No users found
                  </TableCell>
                </TableRow>
              ) : (
                users.map((user: any) => (
                  <TableRow key={user.id} hover>
                    <TableCell>
                      <Box sx={{ display: 'flex', alignItems: 'center', gap: 1.5 }}>
                        <Avatar sx={{ bgcolor: 'primary.main', width: 36, height: 36 }}>
                          {user.avatar}
                        </Avatar>
                        <Box>
                          <Typography fontWeight={500}>{user.name}</Typography>
                          <Typography variant="caption" color="text.secondary">{user.id}</Typography>
                        </Box>
                      </Box>
                    </TableCell>
                    <TableCell>{user.email}</TableCell>
                    <TableCell>
                      <Chip 
                        label={user.status} 
                        size="small" 
                        color={getStatusColor(user.status) as any}
                      />
                    </TableCell>
                    <TableCell>
                      <Chip 
                        label={user.kycLevel} 
                        size="small" 
                        variant="outlined"
                        color={getKycColor(user.kycLevel) as any}
                      />
                    </TableCell>
                    <TableCell align="right">
                      <Typography fontWeight={600} color={user.balance > 0 ? 'success.main' : 'text.secondary'}>
                        ${user.balance.toLocaleString('en-US', { minimumFractionDigits: 2 })}
                      </Typography>
                    </TableCell>
                    <TableCell>{user.joined}</TableCell>
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
                      <Tooltip title={user.status === 'active' ? 'Suspend' : 'Activate'}>
                        <IconButton 
                          size="small" 
                          color={user.status === 'active' ? 'error' : 'success'}
                          onClick={() => {
                            const newStatus = user.status === 'active' ? 'suspended' : 'active';
                            updateStatusMutation.mutate({ id: user.id, status: newStatus });
                          }}
                        >
                          {user.status === 'active' ? <Block fontSize="small" /> : <CheckCircle fontSize="small" />}
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
