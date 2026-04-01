import { useState } from 'react';
import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query';
import {
  Box, Card, CardContent, Typography, Grid, TextField, InputAdornment,
  Table, TableBody, TableCell, TableContainer, TableHead, TableRow,
  Button, Chip, IconButton, Tooltip, Dialog, DialogTitle, DialogContent,
  DialogActions, FormControl, InputLabel, Select, MenuItem, Avatar, Tabs, Tab, LinearProgress
} from '@mui/material';
import { Search, Visibility, CheckCircle, Cancel, Refresh, Description, Person } from '@mui/icons-material';
import { useAppDispatch } from '../../store/hooks';
import { showSnackbar } from '../../store/slices/uiSlice';
import { kycAPI } from '../../services/api';

interface KYCRequest {
  id: string; userId: string; userName: string; email: string;
  documentType: string; documentUrl: string;
  status: 'pending' | 'approved' | 'rejected'; level: string;
  submittedAt: string; reviewedAt?: string; rejectionReason?: string;
}

const statusColor = (s: string) => ({ approved: 'success', rejected: 'error', pending: 'warning' }[s] || 'default') as any;

const KYC = () => {
  const dispatch = useAppDispatch();
  const queryClient = useQueryClient();
  const [search, setSearch] = useState('');
  const [statusFilter, setStatusFilter] = useState('');
  const [activeTab, setActiveTab] = useState(0);
  const [selected, setSelected] = useState<KYCRequest | null>(null);
  const [detailsOpen, setDetailsOpen] = useState(false);
  const [confirmAction, setConfirmAction] = useState<{ type: 'approve' | 'reject'; id: string } | null>(null);
  const [rejectReason, setRejectReason] = useState('');
  const [adminNote, setAdminNote] = useState('');

  const { data, isLoading, refetch } = useQuery({
    queryKey: ['kyc', statusFilter, activeTab],
    queryFn: () => kycAPI.getAll({
      status: statusFilter || undefined,
      level: activeTab === 1 ? 'basic' : activeTab === 2 ? 'advanced' : undefined,
      page: 1, limit: 50,
    }),
  });

  const approveMut = useMutation({
    mutationFn: (id: string) => kycAPI.approve(id, { adminNote: adminNote || undefined }),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['kyc'] });
      dispatch(showSnackbar({ message: 'KYC approved successfully', severity: 'success' }));
      setConfirmAction(null); setAdminNote(''); setDetailsOpen(false);
    },
    onError: (err: any) => dispatch(showSnackbar({ message: err.message || 'Failed to approve KYC', severity: 'error' })),
  });

  const rejectMut = useMutation({
    mutationFn: ({ id, reason }: { id: string; reason: string }) => kycAPI.reject(id, { reason, adminNote: adminNote || undefined }),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['kyc'] });
      dispatch(showSnackbar({ message: 'KYC rejected', severity: 'success' }));
      setConfirmAction(null); setRejectReason(''); setAdminNote(''); setDetailsOpen(false);
    },
    onError: (err: any) => dispatch(showSnackbar({ message: err.message || 'Failed to reject KYC', severity: 'error' })),
  });

  const mockRequests: KYCRequest[] = [
    { id: 'KYC001', userId: 'USR001', userName: 'John Doe', email: 'john@example.com', documentType: 'Passport', documentUrl: '/docs/passport_001.jpg', status: 'pending', level: 'basic', submittedAt: '2024-01-15 10:30' },
    { id: 'KYC002', userId: 'USR002', userName: 'Sarah Smith', email: 'sarah@example.com', documentType: 'ID Card', documentUrl: '/docs/idcard_002.jpg', status: 'pending', level: 'advanced', submittedAt: '2024-01-15 11:45' },
    { id: 'KYC003', userId: 'USR003', userName: 'Mike Johnson', email: 'mike@example.com', documentType: 'Driver License', documentUrl: '/docs/license_003.jpg', status: 'approved', level: 'basic', submittedAt: '2024-01-14 09:00', reviewedAt: '2024-01-14 14:30' },
    { id: 'KYC004', userId: 'USR004', userName: 'Emily Brown', email: 'emily@example.com', documentType: 'Passport', documentUrl: '/docs/passport_004.jpg', status: 'rejected', level: 'basic', submittedAt: '2024-01-13 16:00', reviewedAt: '2024-01-14 10:00', rejectionReason: 'Document expired' },
    { id: 'KYC005', userId: 'USR005', userName: 'David Wilson', email: 'david@example.com', documentType: 'ID Card', documentUrl: '/docs/idcard_005.jpg', status: 'pending', level: 'advanced', submittedAt: '2024-01-16 08:15' },
    { id: 'KYC006', userId: 'USR006', userName: 'Lisa Anderson', email: 'lisa@example.com', documentType: 'Passport', documentUrl: '/docs/passport_006.jpg', status: 'approved', level: 'advanced', submittedAt: '2024-01-12 11:00', reviewedAt: '2024-01-12 16:00' },
  ];

  const requests: KYCRequest[] = data?.data?.requests || mockRequests;
  const filtered = requests.filter((r) => {
    if (statusFilter && r.status !== statusFilter) return false;
    if (activeTab === 1 && r.level !== 'basic') return false;
    if (activeTab === 2 && r.level !== 'advanced') return false;
    if (search && !r.userName.toLowerCase().includes(search.toLowerCase()) && !r.id.toLowerCase().includes(search.toLowerCase()) && !r.email.toLowerCase().includes(search.toLowerCase())) return false;
    return true;
  });
  const pendingCount = requests.filter((r) => r.status === 'pending').length;
  const approvedCount = requests.filter((r) => r.status === 'approved').length;
  const rejectedCount = requests.filter((r) => r.status === 'rejected').length;

  return (
    <Box>
      <Box sx={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', mb: 3 }}>
        <Typography variant="h4" fontWeight="bold">KYC Verification</Typography>
        <Button variant="outlined" startIcon={<Refresh />} onClick={() => refetch()}>Refresh</Button>
      </Box>
      {isLoading && <LinearProgress sx={{ mb: 2 }} />}

      <Grid container spacing={3} sx={{ mb: 3 }}>
        <Grid item xs={12} md={4}><Card sx={{ borderLeft: '4px solid #f59e0b' }}><CardContent><Typography color="text.secondary" variant="body2">Pending</Typography><Typography variant="h4" fontWeight="bold" color="warning.main">{pendingCount}</Typography></CardContent></Card></Grid>
        <Grid item xs={12} md={4}><Card sx={{ borderLeft: '4px solid #22c55e' }}><CardContent><Typography color="text.secondary" variant="body2">Approved</Typography><Typography variant="h4" fontWeight="bold" color="success.main">{approvedCount}</Typography></CardContent></Card></Grid>
        <Grid item xs={12} md={4}><Card sx={{ borderLeft: '4px solid #ef4444' }}><CardContent><Typography color="text.secondary" variant="body2">Rejected</Typography><Typography variant="h4" fontWeight="bold" color="error.main">{rejectedCount}</Typography></CardContent></Card></Grid>
      </Grid>

      <Card sx={{ mb: 3 }}>
        <Box sx={{ borderBottom: 1, borderColor: 'divider' }}>
          <Tabs value={activeTab} onChange={(_, v) => setActiveTab(v)}>
            <Tab label="All Levels" /><Tab label="Basic" /><Tab label="Advanced" />
          </Tabs>
        </Box>
        <CardContent>
          <Grid container spacing={2}>
            <Grid item xs={12} md={6}>
              <TextField fullWidth placeholder="Search by name, email, or ID..." value={search} onChange={(e) => setSearch(e.target.value)}
                InputProps={{ startAdornment: <InputAdornment position="start"><Search /></InputAdornment> }} size="small" />
            </Grid>
            <Grid item xs={12} md={3}>
              <FormControl fullWidth size="small">
                <InputLabel>Status</InputLabel>
                <Select value={statusFilter} label="Status" onChange={(e) => setStatusFilter(e.target.value)}>
                  <MenuItem value="">All</MenuItem><MenuItem value="pending">Pending</MenuItem>
                  <MenuItem value="approved">Approved</MenuItem><MenuItem value="rejected">Rejected</MenuItem>
                </Select>
              </FormControl>
            </Grid>
          </Grid>
        </CardContent>
      </Card>

      <Card>
        <TableContainer>
          <Table>
            <TableHead><TableRow><TableCell>User</TableCell><TableCell>Email</TableCell><TableCell>Document</TableCell><TableCell>Level</TableCell><TableCell>Submitted</TableCell><TableCell>Status</TableCell><TableCell align="center">Actions</TableCell></TableRow></TableHead>
            <TableBody>
              {filtered.length === 0 ? <TableRow><TableCell colSpan={7} align="center">No KYC requests found</TableCell></TableRow> :
                filtered.map((req) => (
                  <TableRow key={req.id} hover>
                    <TableCell><Box sx={{ display: 'flex', alignItems: 'center', gap: 1 }}><Avatar sx={{ width: 32, height: 32 }}><Person fontSize="small" /></Avatar>{req.userName}</Box></TableCell>
                    <TableCell>{req.email}</TableCell>
                    <TableCell><Chip icon={<Description />} label={req.documentType} size="small" variant="outlined" /></TableCell>
                    <TableCell><Chip label={req.level} size="small" /></TableCell>
                    <TableCell>{req.submittedAt}</TableCell>
                    <TableCell><Chip label={req.status} size="small" color={statusColor(req.status)} /></TableCell>
                    <TableCell align="center">
                      <Tooltip title="View Details"><IconButton size="small" onClick={() => { setSelected(req); setDetailsOpen(true); }}><Visibility fontSize="small" /></IconButton></Tooltip>
                      {req.status === 'pending' && <>
                        <Tooltip title="Approve"><IconButton size="small" color="success" onClick={() => setConfirmAction({ type: 'approve', id: req.id })}><CheckCircle fontSize="small" /></IconButton></Tooltip>
                        <Tooltip title="Reject"><IconButton size="small" color="error" onClick={() => setConfirmAction({ type: 'reject', id: req.id })}><Cancel fontSize="small" /></IconButton></Tooltip>
                      </>}
                    </TableCell>
                  </TableRow>
                ))}
            </TableBody>
          </Table>
        </TableContainer>
      </Card>

      <Dialog open={detailsOpen} onClose={() => setDetailsOpen(false)} maxWidth="md" fullWidth>
        <DialogTitle>KYC Request Details</DialogTitle>
        <DialogContent>
          {selected && <Grid container spacing={2} sx={{ mt: 1 }}>
            <Grid item xs={6}><Typography variant="body2" color="text.secondary">Request ID</Typography><Typography>{selected.id}</Typography></Grid>
            <Grid item xs={6}><Typography variant="body2" color="text.secondary">User</Typography><Typography>{selected.userName}</Typography></Grid>
            <Grid item xs={6}><Typography variant="body2" color="text.secondary">Email</Typography><Typography>{selected.email}</Typography></Grid>
            <Grid item xs={6}><Typography variant="body2" color="text.secondary">Document Type</Typography><Typography>{selected.documentType}</Typography></Grid>
            <Grid item xs={6}><Typography variant="body2" color="text.secondary">Level</Typography><Typography>{selected.level}</Typography></Grid>
            <Grid item xs={6}><Typography variant="body2" color="text.secondary">Status</Typography><Chip label={selected.status} size="small" color={statusColor(selected.status)} /></Grid>
            <Grid item xs={12}><Typography variant="body2" color="text.secondary">Document Preview</Typography>
              <Box sx={{ mt: 1, p: 2, bgcolor: 'grey.100', borderRadius: 1, textAlign: 'center' }}>
                <Description sx={{ fontSize: 48, color: 'grey.500' }} />
                <Typography variant="body2" color="text.secondary">{selected.documentUrl}</Typography>
              </Box>
            </Grid>
            {selected.rejectionReason && <Grid item xs={12}><Typography variant="body2" color="text.secondary">Rejection Reason</Typography><Typography color="error">{selected.rejectionReason}</Typography></Grid>}
          </Grid>}
        </DialogContent>
        <DialogActions>
          {selected?.status === 'pending' && <>
            <Button color="success" variant="contained" onClick={() => setConfirmAction({ type: 'approve', id: selected.id })}>Approve</Button>
            <Button color="error" variant="outlined" onClick={() => setConfirmAction({ type: 'reject', id: selected.id })}>Reject</Button>
          </>}
          <Button onClick={() => setDetailsOpen(false)}>Close</Button>
        </DialogActions>
      </Dialog>

      <Dialog open={!!confirmAction} onClose={() => { setConfirmAction(null); setRejectReason(''); setAdminNote(''); }} maxWidth="sm" fullWidth>
        <DialogTitle>{confirmAction?.type === 'approve' ? 'Approve KYC' : 'Reject KYC'}</DialogTitle>
        <DialogContent>
          <Typography sx={{ mb: 2 }}>{confirmAction?.type === 'approve' ? 'Are you sure you want to approve this KYC verification?' : 'Please provide a reason for rejection.'}</Typography>
          {confirmAction?.type === 'reject' && <TextField fullWidth label="Rejection Reason" value={rejectReason} onChange={(e) => setRejectReason(e.target.value)} sx={{ mb: 2 }} required />}
          <TextField fullWidth label="Admin Note (optional)" value={adminNote} onChange={(e) => setAdminNote(e.target.value)} multiline rows={2} />
        </DialogContent>
        <DialogActions>
          <Button onClick={() => { setConfirmAction(null); setRejectReason(''); setAdminNote(''); }}>Cancel</Button>
          {confirmAction?.type === 'approve' ? (
            <Button variant="contained" color="success" onClick={() => approveMut.mutate(confirmAction.id)} disabled={approveMut.isPending}>Approve</Button>
          ) : (
            <Button variant="contained" color="error" onClick={() => confirmAction && rejectMut.mutate({ id: confirmAction.id, reason: rejectReason })} disabled={!rejectReason || rejectMut.isPending}>Reject</Button>
          )}
        </DialogActions>
      </Dialog>
    </Box>
  );
};

export default KYC;
