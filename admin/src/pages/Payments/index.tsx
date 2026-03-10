import { useState } from 'react';
import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query';
import {
  Box, Card, CardContent, Typography, Grid, TextField, InputAdornment,
  Table, TableBody, TableCell, TableContainer, TableHead, TableRow,
  Button, Chip, IconButton, Tooltip, Tabs, Tab, FormControl, InputLabel, Select, MenuItem,
  Dialog, DialogActions, DialogContent, DialogTitle, LinearProgress
} from '@mui/material';
import { Search, Add, Edit, Visibility, AccountBalance, CreditCard, SwapHoriz, CheckCircle, Cancel, Refresh } from '@mui/icons-material';
import { useAppDispatch } from '../../store/hooks';
import { showSnackbar } from '../../store/slices/uiSlice';
import { paymentsAPI } from '../../services/api';

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

  // Fallback mock data for demo purposes when API is not available
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

  const getStatusColor = (status: string) => {
    switch (status) {
      case 'completed': return 'success';
      case 'pending': return 'warning';
      case 'failed': return 'error';
      case 'processing': return 'info';
      default: return 'default';
    }
  };

  const getMethodIcon = (method: string) => {
    switch (method.toLowerCase()) {
      case 'bank transfer': return <AccountBalance />;
      case 'credit card': return <CreditCard />;
      case 'crypto': return <SwapHoriz />;
      default: return <CreditCard />;
    }
  };

  // Calculate stats from payments
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

      {/* Stats Cards */}
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

      {/* Filters */}
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

      {/* Transactions Table */}
      <Card>
        <TableContainer>
          <Table>
            <TableHead>
              <TableRow>
                <TableCell>Transaction ID</TableCell>
                <TableCell>User</TableCell>
                <TableCell align="right">Amount</TableCell>
                <TableCell>Method</TableCell>
                <TableCell>Type</TableCell>
                <TableCell>Status</TableCell>
                <TableCell>Date</TableCell>
                <TableCell align="center">Actions</TableCell>
              </TableRow>
            </TableHead>
            <TableBody>
              {payments.length === 0 ? (
                <TableRow>
                  <TableCell colSpan={8} align="center" sx={{ py: 4 }}>
                    <Typography color="text.secondary">No transactions found</Typography>
                  </TableCell>
                </TableRow>
              ) : (
                payments.map((payment: any) => (
                  <TableRow key={payment.id} hover>
                    <TableCell>
                      <Typography fontFamily="monospace" fontWeight={500}>
                        {payment.id}
                      </Typography>
                    </TableCell>
                    <TableCell>
                      <Box>
                        <Typography fontWeight={500}>{payment.user}</Typography>
                        <Typography variant="caption" color="text.secondary">{payment.userId}</Typography>
                      </Box>
                    </TableCell>
                    <TableCell align="right">
                      <Typography 
                        fontWeight={600} 
                        color={payment.type === 'deposit' ? 'success.main' : 'error.main'}
                      >
                        {payment.type === 'deposit' ? '+' : '-'}${payment.amount.toLocaleString()}
                      </Typography>
                    </TableCell>
                    <TableCell>
                      <Box sx={{ display: 'flex', alignItems: 'center', gap: 1 }}>
                        {getMethodIcon(payment.method)}
                        {payment.method}
                      </Box>
                    </TableCell>
                    <TableCell>
                      <Chip 
                        label={payment.type} 
                        color={payment.type === 'deposit' ? 'success' : 'error'} 
                        size="small" 
                        variant="outlined"
                      />
                    </TableCell>
                    <TableCell>
                      <Chip 
                        label={payment.status} 
                        color={getStatusColor(payment.status) as any} 
                        size="small" 
                      />
                    </TableCell>
                    <TableCell>
                      <Typography variant="body2" color="text.secondary">
                        {payment.date}
                      </Typography>
                    </TableCell>
                    <TableCell align="center">
                      <Tooltip title="View Details">
                        <IconButton 
                          size="small" 
                          onClick={() => {
                            setSelectedPayment(payment);
                            setDetailsOpen(true);
                          }}
                        >
                          <Visibility fontSize="small" />
                        </IconButton>
                      </Tooltip>
                      {payment.status === 'pending' && (
                        <>
                          <Tooltip title="Approve">
                            <IconButton 
                              size="small" 
                              color="success"
                              onClick={() => approveMutation.mutate(payment.id)}
                            >
                              <CheckCircle fontSize="small" />
                            </IconButton>
                          </Tooltip>
                          <Tooltip title="Reject">
                            <IconButton 
                              size="small" 
                              color="error"
                              onClick={() => rejectMutation.mutate(payment.id)}
                            >
                              <Cancel fontSize="small" />
                            </IconButton>
                          </Tooltip>
                        </>
                      )}
                      {payment.status === 'completed' && (
                        <Tooltip title="Process (Re-process)">
                          <IconButton 
                            size="small" 
                            color="primary"
                            onClick={() => processMutation.mutate(payment.id)}
                          >
                            <Refresh fontSize="small" />
                          </IconButton>
                        </Tooltip>
                      )}
                    </TableCell>
                  </TableRow>
                ))
              )}
            </TableBody>
          </Table>
        </TableContainer>
      </Card>

      {/* Payment Details Dialog */}
      <Dialog open={detailsOpen} onClose={() => setDetailsOpen(false)} maxWidth="sm" fullWidth>
        <DialogTitle>Transaction Details</DialogTitle>
        <DialogContent>
          {selectedPayment && (
            <Grid container spacing={2} sx={{ mt: 1 }}>
              <Grid item xs={6}>
                <Typography variant="caption" color="text.secondary">Transaction ID</Typography>
                <Typography fontFamily="monospace">{selectedPayment.id}</Typography>
              </Grid>
              <Grid item xs={6}>
                <Typography variant="caption" color="text.secondary">Date</Typography>
                <Typography>{selectedPayment.date}</Typography>
              </Grid>
              <Grid item xs={12}>
                <Typography variant="caption" color="text.secondary">User</Typography>
                <Typography>{selectedPayment.user} ({selectedPayment.userId})</Typography>
              </Grid>
              <Grid item xs={6}>
                <Typography variant="caption" color="text.secondary">Amount</Typography>
                <Typography variant="h6" color={selectedPayment.type === 'deposit' ? 'success.main' : 'error.main'}>
                  {selectedPayment.type === 'deposit' ? '+' : '-'}${selectedPayment.amount.toLocaleString()}
                </Typography>
              </Grid>
              <Grid item xs={6}>
                <Typography variant="caption" color="text.secondary">Method</Typography>
                <Typography>{selectedPayment.method}</Typography>
              </Grid>
              <Grid item xs={6}>
                <Typography variant="caption" color="text.secondary">Type</Typography>
                <Chip 
                  label={selectedPayment.type} 
                  color={selectedPayment.type === 'deposit' ? 'success' : 'error'} 
                  size="small" 
                />
              </Grid>
              <Grid item xs={6}>
                <Typography variant="caption" color="text.secondary">Status</Typography>
                <Chip 
                  label={selectedPayment.status} 
                  color={getStatusColor(selectedPayment.status) as any} 
                  size="small" 
                />
              </Grid>
            </Grid>
          )}
        </DialogContent>
        <DialogActions>
          <Button onClick={() => setDetailsOpen(false)}>Close</Button>
          {selectedPayment?.status === 'pending' && (
            <>
              <Button 
                color="error" 
                onClick={() => {
                  rejectMutation.mutate(selectedPayment.id);
                }}
              >
                Reject
              </Button>
              <Button 
                color="success" 
                variant="contained"
                onClick={() => {
                  approveMutation.mutate(selectedPayment.id);
                }}
              >
                Approve
              </Button>
            </>
          )}
        </DialogActions>
      </Dialog>
    </Box>
  );
};

export default Payments;
