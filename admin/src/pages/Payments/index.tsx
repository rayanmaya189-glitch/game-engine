import { useState } from 'react';
import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query';
import {
  Box, Card, CardContent, Typography, Grid, TextField, InputAdornment,
  Table, TableBody, TableCell, TableContainer, TableHead, TableRow,
  Button, Chip, IconButton, Tooltip, Tabs, Tab, FormControl, InputLabel, Select, MenuItem
} from '@mui/material';
import { Search, Add, Edit, Visibility, AccountBalance, CreditCard, SwapHoriz } from '@mui/icons-material';
import { useAppDispatch } from '../../store/hooks';
import { showSnackbar } from '../../store/slices/uiSlice';
import { paymentsAPI } from '../../services/api';

const Payments = () => {
  const dispatch = useAppDispatch();
  const queryClient = useQueryClient();
  const [search, setSearch] = useState('');
  const [activeTab, setActiveTab] = useState(0);

  const { data, isLoading } = useQuery({
    queryKey: ['payments', activeTab],
    queryFn: () => paymentsAPI.getAll({ page: 1, limit: 20 }),
  });

  const transactions = data?.data?.payments || [];

  const getStatusColor = (status: string) => {
    switch (status) {
      case 'completed': return 'success';
      case 'pending': return 'warning';
      case 'failed': return 'error';
      default: return 'default';
    }
  };

  return (
    <Box>
      <Typography variant="h4" fontWeight="bold" gutterBottom>Payments Management</Typography>

      <Grid container spacing={3} sx={{ mb: 3 }}>
        <Grid item xs={12} md={3}>
          <Card><CardContent><Typography color="text.secondary">Total Deposits</Typography><Typography variant="h4" color="success.main">$125,000</Typography></CardContent></Card>
        </Grid>
        <Grid item xs={12} md={3}>
          <Card><CardContent><Typography color="text.secondary">Total Withdrawals</Typography><Typography variant="h4" color="error.main">$45,000</Typography></CardContent></Card>
        </Grid>
        <Grid item xs={12} md={3}>
          <Card><CardContent><Typography color="text.secondary">Pending</Typography><Typography variant="h4" color="warning.main">$12,500</Typography></CardContent></Card>
        </Grid>
        <Grid item xs={12} md={3}>
          <Card><CardContent><Typography color="text.secondary">Failed</Typography><Typography variant="h4" color="error.main">$1,200</Typography></CardContent></Card>
        </Grid>
      </Grid>

      <Card sx={{ mb: 3 }}>
        <Box sx={{ borderBottom: 1, borderColor: 'divider' }}>
          <Tabs value={activeTab} onChange={(_, v) => setActiveTab(v)}>
            <Tab label="All Transactions" />
            <Tab label="Deposits" />
            <Tab label="Withdrawals" />
            <Tab label="Payment Methods" />
          </Tabs>
        </Box>
        <CardContent>
          <Grid container spacing={2}>
            <Grid item xs={12} md={6}>
              <TextField
                fullWidth
                placeholder="Search transactions..."
                value={search}
                onChange={(e) => setSearch(e.target.value)}
                InputProps={{ startAdornment: <InputAdornment position="start"><Search /></InputAdornment> }}
              />
            </Grid>
            <Grid item xs={12} md={3}>
              <FormControl fullWidth>
                <InputLabel>Status</InputLabel>
                <Select label="Status" defaultValue="">
                  <MenuItem value="">All</MenuItem>
                  <MenuItem value="completed">Completed</MenuItem>
                  <MenuItem value="pending">Pending</MenuItem>
                  <MenuItem value="failed">Failed</MenuItem>
                </Select>
              </FormControl>
            </Grid>
            <Grid item xs={12} md={3}>
              <FormControl fullWidth>
                <InputLabel>Method</InputLabel>
                <Select label="Method" defaultValue="">
                  <MenuItem value="">All</MenuItem>
                  <MenuItem value="bank">Bank Transfer</MenuItem>
                  <MenuItem value="card">Credit Card</MenuItem>
                  <MenuItem value="crypto">Crypto</MenuItem>
                  <MenuItem value="ewallet">E-Wallet</MenuItem>
                </Select>
              </FormControl>
            </Grid>
          </Grid>
        </CardContent>
      </Card>

      <Card>
        <TableContainer>
          <Table>
            <TableHead>
              <TableRow>
                <TableCell>Transaction ID</TableCell>
                <TableCell>User</TableCell>
                <TableCell>Amount</TableCell>
                <TableCell>Method</TableCell>
                <TableCell>Type</TableCell>
                <TableCell>Status</TableCell>
                <TableCell>Date</TableCell>
                <TableCell align="right">Actions</TableCell>
              </TableRow>
            </TableHead>
            <TableBody>
              {transactions.map((txn) => (
                <TableRow key={txn.id} hover>
                  <TableCell fontWeight="500">{txn.id}</TableCell>
                  <TableCell>{txn.user}</TableCell>
                  <TableCell>${txn.amount.toLocaleString()}</TableCell>
                  <TableCell>{txn.method}</TableCell>
                  <TableCell><Chip label={txn.type} color={txn.type === 'deposit' ? 'success' : 'error'} size="small