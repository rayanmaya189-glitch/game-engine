import { useState } from 'react';
import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query';
import {
  Box, Card, CardContent, Typography, Grid, TextField, InputAdornment,
  Button, Tabs, Tab, FormControl, InputLabel, Select, MenuItem,
  LinearProgress
} from '@mui/material';
import { Search, Refresh } from '@mui/icons-material';
import { useAppDispatch } from '../../store/hooks';
import { showSnackbar } from '../../store/slices/uiSlice';
import { paymentsAPI } from '../../services/api';
import PaymentDetails from './PaymentDetails';
import PaymentActions from './PaymentActions';

const Payments = () => {
  const dispatch = useAppDispatch();
  const queryClient = useQueryClient();
  const [search, setSearch] = useState('');
  const [statusFilter, setStatusFilter] = useState('');
  const [methodFilter, setMethodFilter] = useState('');
  const [activeTab, setActiveTab] = useState(0);
  const [selectedPayment, setSelectedPayment] = useState<any>(null);
  const [detailsOpen, setDetailsOpen] = useState(false);

  const { data, isLoading, refetch } = useQuery({
    queryKey: ['payments', activeTab, statusFilter, methodFilter],
    queryFn: () => paymentsAPI.getAll({
      status: statusFilter || undefined,
      type: activeTab === 1 ? 'deposit' : activeTab === 2 ? 'withdrawal' : undefined,
      page: 1,
      limit: 50
    }),
  });

  const approveMutation = useMutation({
    mutationFn: (id: string) => paymentsAPI.approve(id),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['payments'] });
      dispatch(showSnackbar({ message: 'Payment approved successfully', severity: 'success' }));
      setDetailsOpen(false);
    },
    onError: (error: any) => {
      dispatch(showSnackbar({ message: error.message || 'Failed to approve payment', severity: 'error' }));
    },
  });

  const rejectMutation = useMutation({
    mutationFn: (id: string) => paymentsAPI.reject(id),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['payments'] });
      dispatch(showSnackbar({ message: 'Payment rejected', severity: 'success' }));
      setDetailsOpen(false);
    },
    onError: (error: any) => {
      dispatch(showSnackbar({ message: error.message || 'Failed to reject payment', severity: 'error' }));
    },
  });

  const processMutation = useMutation({
    mutationFn: (id: string) => paymentsAPI.process(id),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['payments'] });
      dispatch(showSnackbar({ message: 'Payment processed successfully', severity: 'success' }));
      setDetailsOpen(false);
    },
    onError: (error: any) => {
      dispatch(showSnackbar({ message: error.message || 'Failed to process payment', severity: 'error' }));
    },
  });

  const mockPayments = [
    { id: 'TXN001', userId: 'USR001', user: 'John Doe', amount: 500, method: 'Bank Transfer', type: 'deposit', status: 'completed', date: '2024-01-15 10:30' },
    { id: 'TXN002', userId: 'USR002', user: 'Sarah Smith', amount: 250, method: 'Credit Card', type: 'deposit', status: 'pending', date: '2024-01-15 11:45' },
    { id: 'TXN003', userId: 'USR003', user: 'Mike Johnson', amount: 1000, method: 'Crypto', type: 'withdrawal', status: 'completed', date: '2024-01-15 12:00' },
    { id: 'TXN004', userId: 'USR004', user: 'Emily Brown', amount: 75, method: 'E-Wallet', type: 'deposit', status: 'failed', date: '2024-01-15 12:30' },
    { id: 'TXN005', userId: 'USR005', user: 'David Wilson', amount: 2000, method: 'Bank Transfer', type: 'withdrawal', status: 'pending', date: '2024-01-15 13:00' },
    { id: 'TXN006', userId: 'USR006', user: 'Lisa Anderson', amount: 150, method: 'Credit Card', type: 'deposit', status: 'completed', date: '2024-01-15 14:15' },
    { id: 'TXN007', userId: 'USR007', user: 'James Taylor', amount: 5000, method: 'Crypto', type: 'withdrawal', status: 'completed', date: '2024-01-15 15:30' },
    { id: 'TXN008', userId: 'USR008', user: 'Jennifer Martinez', amount: 300, method: 'E-Wallet', type: 'deposit', status: 'pending', date: '2024-01-15 16:00' },
  ];

  const payments = data?.data?.payments || mockPayments.filter(payment => {
    if (statusFilter && payment.status !== statusFilter) return false;
    if (methodFilter && payment.method.toLowerCase() !== methodFilter.toLowerCase()) return false;
    if (activeTab === 1 && payment.type !== 'deposit') return false;
    if (activeTab === 2 && payment.type !== 'withdrawal') return false;
    if (search && !payment.user.toLowerCase().includes(search.toLowerCase()) &&
        !payment.id.toLowerCase().includes(search.toLowerCase())) return false;
    return true;
  });

  const totalDeposits = mockPayments.filter(p => p.type === 'deposit' && p.status === 'completed').reduce((sum, p) => sum + p.amount, 0);
  const totalWithdrawals = mockPayments.filter(p => p.type === 'withdrawal' && p.status === 'completed').reduce((sum, p) => sum + p.amount, 0);
  const pendingPayments = mockPayments.filter(p => p.status === 'pending').reduce((sum, p) => sum + p.amount, 0);
  const failedPayments = mockPayments.filter(p => p.status === 'failed').reduce((sum, p) => sum + p.amount, 0);

  return (
    <Box>
      <Box sx={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', mb: 3 }}>
        <Typography variant="h4" fontWeight="bold">Payments Management</Typography>
        <Button variant="outlined" startIcon={<Refresh />} onClick={() => refetch()}>
          Refresh
        </Button>
      </Box>

      {isLoading && <LinearProgress sx={{ mb: 2 }} />}

      <Grid container spacing={3} sx={{ mb: 3 }}>
        <Grid item xs={12} md={3}>
          <Card sx={{ borderLeft: '4px solid #22c55e' }}>
            <CardContent>
              <Typography color="text.secondary" variant="body2">Total Deposits</Typography>
              <Typography variant="h4" fontWeight="bold" color="success.main">
                ${totalDeposits.toLocaleString()}
              </Typography>
            </CardContent>
          </Card>
        </Grid>
        <Grid item xs={12} md={3}>
          <Card sx={{ borderLeft: '4px solid #ef4444' }}>
            <CardContent>
              <Typography color="text.secondary" variant="body2">Total Withdrawals</Typography>
              <Typography variant="h4" fontWeight="bold" color="error.main">
                ${totalWithdrawals.toLocaleString()}
              </Typography>
            </CardContent>
          </Card>
        </Grid>
        <Grid item xs={12} md={3}>
          <Card sx={{ borderLeft: '4px solid #f59e0b' }}>
            <CardContent>
              <Typography color="text.secondary" variant="body2">Pending</Typography>
              <Typography variant="h4" fontWeight="bold" color="warning.main">
                ${pendingPayments.toLocaleString()}
              </Typography>
            </CardContent>
          </Card>
        </Grid>
        <Grid item xs={12} md={3}>
          <Card sx={{ borderLeft: '4px solid #dc2626' }}>
            <CardContent>
              <Typography color="text.secondary" variant="body2">Failed</Typography>
              <Typography variant="h4" fontWeight="bold" color="error.dark">
                ${failedPayments.toLocaleString()}
              </Typography>
            </CardContent>
          </Card>
        </Grid>
      </Grid>

      <Card sx={{ mb: 3 }}>
        <Box sx={{ borderBottom: 1, borderColor: 'divider' }}>
          <Tabs value={activeTab} onChange={(_, v) => setActiveTab(v)}>
            <Tab label="All Transactions" />
            <Tab label="Deposits" />
            <Tab label="Withdrawals" />
          </Tabs>
        </Box>
        <CardContent>
          <Grid container spacing={2}>
            <Grid item xs={12} md={4}>
              <TextField
                fullWidth
                placeholder="Search by user or transaction ID..."
                value={search}
                onChange={(e) => setSearch(e.target.value)}
                InputProps={{
                  startAdornment: <InputAdornment position="start"><Search /></InputAdornment>
                }}
                size="small"
              />
            </Grid>
            <Grid item xs={12} md={4}>
              <FormControl fullWidth size="small">
                <InputLabel>Status</InputLabel>
                <Select
                  value={statusFilter}
                  label="Status"
                  onChange={(e) => setStatusFilter(e.target.value)}
                >
                  <MenuItem value="">All</MenuItem>
                  <MenuItem value="completed">Completed</MenuItem>
                  <MenuItem value="pending">Pending</MenuItem>
                  <MenuItem value="processing">Processing</MenuItem>
                  <MenuItem value="failed">Failed</MenuItem>
                </Select>
              </FormControl>
            </Grid>
            <Grid item xs={12} md={4}>
              <FormControl fullWidth size="small">
                <InputLabel>Method</InputLabel>
                <Select
                  value={methodFilter}
                  label="Method"
                  onChange={(e) => setMethodFilter(e.target.value)}
                >
                  <MenuItem value="">All</MenuItem>
                  <MenuItem value="bank transfer">Bank Transfer</MenuItem>
                  <MenuItem value="credit card">Credit Card</MenuItem>
                  <MenuItem value="crypto">Crypto</MenuItem>
                  <MenuItem value="e-wallet">E-Wallet</MenuItem>
                </Select>
              </FormControl>
            </Grid>
          </Grid>
        </CardContent>
      </Card>

      <Card>
        <PaymentActions
          payments={payments}
          onViewDetails={(payment) => { setSelectedPayment(payment); setDetailsOpen(true); }}
          onApprove={(id) => approveMutation.mutate(id)}
          onReject={(id) => rejectMutation.mutate(id)}
          onProcess={(id) => processMutation.mutate(id)}
        />
      </Card>

      <PaymentDetails
        open={detailsOpen}
        payment={selectedPayment}
        onClose={() => setDetailsOpen(false)}
        onApprove={(id) => approveMutation.mutate(id)}
        onReject={(id) => rejectMutation.mutate(id)}
      />
    </Box>
  );
};

export default Payments;
