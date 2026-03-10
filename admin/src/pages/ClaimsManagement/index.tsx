import { useState } from 'react';
import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query';
import {
  Box, Card, CardContent, Typography, Tabs, Tab, Table, TableBody,
  TableCell, TableContainer, TableHead, TableRow, Chip, Button,
  Dialog, DialogTitle, DialogContent, DialogActions, TextField,
  FormControl, InputLabel, Select, MenuItem, Grid, LinearProgress,
  IconButton, Tooltip, Pagination
} from '@mui/material';
import {
  CheckCircle, Cancel, Visibility, Refresh
} from '@mui/icons-material';
import { claimsAPI } from '../../services/api';
import { useAppDispatch } from '../../store/hooks';
import { showSnackbar } from '../../store/slices/uiSlice';

type ClaimType = 'commission' | 'rebet' | 'insurance';

interface Claim {
  id: number;
  userId: number;
  affiliateId?: number;
  claimType?: string;
  amount: string;
  status: string;
  claimReason?: string;
  requestedAt: string;
  processedAt?: string;
  adminNote?: string;
}

const ClaimsManagement = () => {
  const dispatch = useAppDispatch();
  const queryClient = useQueryClient();
  const [activeTab, setActiveTab] = useState<ClaimType>('commission');
  const [selectedClaim, setSelectedClaim] = useState<Claim | null>(null);
  const [detailsOpen, setDetailsOpen] = useState(false);
  const [actionDialog, setActionDialog] = useState<{ open: boolean; type: 'approve' | 'reject' | 'pay'; claim: Claim | null }>({
    open: false,
    type: 'approve',
    claim: null,
  });
  const [adminNote, setAdminNote] = useState('');
  const [page, setPage] = useState(1);

  const { data: commissionData, isLoading: commissionLoading } = useQuery({
    queryKey: ['commission-claims', activeTab === 'commission' ? page : 1],
    queryFn: () => claimsAPI.getAllCommissionClaims({ page, limit: 10 }),
    enabled: activeTab === 'commission',
  });

  const { data: rebetData, isLoading: rebetLoading } = useQuery({
    queryKey: ['rebet-claims', activeTab === 'rebet' ? page : 1],
    queryFn: () => claimsAPI.getAllRebetClaims({ page, limit: 10 }),
    enabled: activeTab === 'rebet',
  });

  const { data: insuranceData, isLoading: insuranceLoading } = useQuery({
    queryKey: ['insurance-claims', activeTab === 'insurance' ? page : 1],
    queryFn: () => claimsAPI.getAllInsuranceClaims({ page, limit: 10 }),
    enabled: activeTab === 'insurance',
  });

  const approveMutation = useMutation({
    mutationFn: ({ id, type, note }: { id: string; type: ClaimType; note: string }) => {
      if (type === 'commission') return claimsAPI.approveCommissionClaim(id, { adminNote: note });
      if (type === 'rebet') return claimsAPI.approveRebetClaim(id, { adminNote: note });
      return claimsAPI.approveInsuranceClaim(id, { adminNote: note });
    },
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: [`${activeTab}-claims`] });
      dispatch(showSnackbar({ message: 'Claim approved successfully', severity: 'success' }));
      setActionDialog({ open: false, type: 'approve', claim: null });
      setAdminNote('');
    },
    onError: (error: any) => {
      dispatch(showSnackbar({ message: error.message || 'Failed to approve claim', severity: 'error' }));
    },
  });

  const rejectMutation = useMutation({
    mutationFn: ({ id, type, note }: { id: string; type: ClaimType; note: string }) => {
      if (type === 'commission') return claimsAPI.rejectCommissionClaim(id, { adminNote: note });
      if (type === 'rebet') return claimsAPI.rejectRebetClaim(id, { adminNote: note });
      return claimsAPI.rejectInsuranceClaim(id, { adminNote: note });
    },
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: [`${activeTab}-claims`] });
      dispatch(showSnackbar({ message: 'Claim rejected', severity: 'success' }));
      setActionDialog({ open: false, type: 'reject', claim: null });
      setAdminNote('');
    },
    onError: (error: any) => {
      dispatch(showSnackbar({ message: error.message || 'Failed to reject claim', severity: 'error' }));
    },
  });

  const payMutation = useMutation({
    mutationFn: ({ id, type }: { id: string; type: ClaimType }) => {
      if (type === 'commission') return claimsAPI.payCommissionClaim(id);
      if (type === 'rebet') return claimsAPI.approveRebetClaim(id, { adminNote: 'Auto-approved for payment' });
      return claimsAPI.payInsuranceClaim(id);
    },
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: [`${activeTab}-claims`] });
      dispatch(showSnackbar({ message: 'Payment processed successfully', severity: 'success' }));
      setActionDialog({ open: false, type: 'pay', claim: null });
    },
    onError: (error: any) => {
      dispatch(showSnackbar({ message: error.message || 'Failed to process payment', severity: 'error' }));
    },
  });

  const getStatusColor = (status: string) => {
    const colors: Record<string, 'default' | 'primary' | 'secondary' | 'error' | 'info' | 'success' | 'warning'> = {
      PENDING: 'warning',
      APPROVED: 'info',
      REJECTED: 'error',
      PAID: 'success',
      CLAIMABLE: 'success',
      CLAIMED: 'success',
      IN_PROGRESS: 'info',
    };
    return colors[status] || 'default';
  };

  const handleAction = () => {
    if (!actionDialog.claim) return;
    const id = actionDialog.claim.id.toString();
    
    if (actionDialog.type === 'approve') {
      approveMutation.mutate({ id, type: activeTab, note: adminNote });
    } else if (actionDialog.type === 'reject') {
      rejectMutation.mutate({ id, type: activeTab, note: adminNote });
    } else {
      payMutation.mutate({ id, type: activeTab });
    }
  };

  const getClaims = () => {
    if (activeTab === 'commission') return commissionData?.data?.claims || [];
    if (activeTab === 'rebet') return rebetData?.data?.claims || [];
    return insuranceData?.data?.claims || [];
  };

  const isLoading = activeTab === 'commission' ? commissionLoading : activeTab === 'rebet' ? rebetLoading : insuranceLoading;

  return (
    <Box>
      <Typography variant="h4" fontWeight="bold" gutterBottom>
        Claims Management
      </Typography>

      <Card sx={{ mb: 3 }}>
        <CardContent>
          <Grid container spacing={3}>
            <Grid item xs={12} md={3}>
              <Box sx={{ textAlign: 'center', p: 2, bgcolor: 'warning.light', borderRadius: 2 }}>
                <Typography variant="h3">{commissionData?.data?.totalPending || 0}</Typography>
                <Typography color="warning.dark">Pending</Typography>
              </Box>
            </Grid>
            <Grid item xs={12} md={3}>
              <Box sx={{ textAlign: 'center', p: 2, bgcolor: 'info.light', borderRadius: 2 }}>
                <Typography variant="h3">{rebetData?.data?.totalInProgress || 0}</Typography>
                <Typography color="info.dark">In Progress</Typography>
              </Box>
            </Grid>
            <Grid item xs={12} md={3}>
              <Box sx={{ textAlign: 'center', p: 2, bgcolor: 'success.light', borderRadius: 2 }}>
                <Typography variant="h3">{commissionData?.data?.totalPaid || 0}</Typography>
                <Typography color="success.dark">Paid</Typography>
              </Box>
            </Grid>
            <Grid item xs={12} md={3}>
              <Box sx={{ textAlign: 'center', p: 2, bgcolor: 'error.light', borderRadius: 2 }}>
                <Typography variant="h3">{insuranceData?.data?.totalRejected || 0}</Typography>
                <Typography color="error.dark">Rejected</Typography>
              </Box>
            </Grid>
          </Grid>
        </CardContent>
      </Card>

      <Card>
        <Box sx={{ borderBottom: 1, borderColor: 'divider' }}>
          <Tabs value={activeTab} onChange={(_, v) => { setActiveTab(v); setPage(1); }}>
            <Tab value="commission" label="Commission Claims" />
            <Tab value="rebet" label="Rebet Claims" />
            <Tab value="insurance" label="Insurance Claims" />
          </Tabs>
        </Box>

        <CardContent>
          {isLoading && <LinearProgress sx={{ mb: 2 }} />}
          
          <TableContainer>
            <Table>
              <TableHead>
                <TableRow>
                  <TableCell>ID</TableCell>
                  <TableCell>User ID</TableCell>
                  <TableCell>Type</TableCell>
                  <TableCell>Amount</TableCell>
                  <TableCell>Status</TableCell>
                  <TableCell>Reason</TableCell>
                  <TableCell>Date</TableCell>
                  <TableCell align="right">Actions</TableCell>
                </TableRow>
              </TableHead>
              <TableBody>
                {getClaims().length === 0 ? (
                  <TableRow>
                    <TableCell colSpan={8} align="center" sx={{ py: 4 }}>
                      No claims found
                    </TableCell>
                  </TableRow>
                ) : (
                  getClaims().map((claim: Claim) => (
                    <TableRow key={claim.id} hover>
                      <TableCell>{claim.id}</TableCell>
                      <TableCell>{claim.userId}</TableCell>
                      <TableCell>{claim.claimType || activeTab}</TableCell>
                      <TableCell>${claim.amount}</TableCell>
                      <TableCell>
                        <Chip 
                          label={claim.status} 
                          color={getStatusColor(claim.status)} 
                          size="small" 
                        />
                      </TableCell>
                      <TableCell>{claim.claimReason || '-'}</TableCell>
                      <TableCell>{new Date(claim.requestedAt).toLocaleDateString()}</TableCell>
                      <TableCell align="right">
                        <Tooltip title="View Details">
                          <IconButton 
                            size="small" 
                            onClick={() => { setSelectedClaim(claim); setDetailsOpen(true); }}
                          >
                            <Visibility />
                          </IconButton>
                        </Tooltip>
                        {claim.status === 'PENDING' && (
                          <>
                            <Tooltip title="Approve">
                              <IconButton 
                                size="small" 
                                color="success"
                                onClick={() => setActionDialog({ open: true, type: 'approve', claim })}
                              >
                                <CheckCircle />
                              </IconButton>
                            </Tooltip>
                            <Tooltip title="Reject">
                              <IconButton 
                                size="small" 
                                color="error"
                                onClick={() => setActionDialog({ open: true, type: 'reject', claim })}
                              >
                                <Cancel />
                              </IconButton>
                            </Tooltip>
                          </>
                        )}
                        {claim.status === 'APPROVED' && (
                          <Tooltip title="Pay">
                            <IconButton 
                              size="small" 
                              color="primary"
                              onClick={() => setActionDialog({ open: true, type: 'pay', claim })}
                            >
                              <Refresh />
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

          <Box sx={{ mt: 2, display: 'flex', justifyContent: 'center' }}>
            <Pagination 
              count={10} 
              page={page} 
              onChange={(_, v) => setPage(v)} 
              color="primary" 
            />
          </Box>
        </CardContent>
      </Card>

      {/* Details Dialog */}
      <Dialog open={detailsOpen} onClose={() => setDetailsOpen(false)} maxWidth="sm" fullWidth>
        <DialogTitle>Claim Details</DialogTitle>
        <DialogContent>
          {selectedClaim && (
            <Grid container spacing={2} sx={{ mt: 1 }}>
              <Grid item xs={6}>
                <Typography variant="subtitle2" color="text.secondary">ID</Typography>
                <Typography>{selectedClaim.id}</Typography>
              </Grid>
              <Grid item xs={6}>
                <Typography variant="subtitle2" color="text.secondary">User ID</Typography>
                <Typography>{selectedClaim.userId}</Typography>
              </Grid>
              <Grid item xs={6}>
                <Typography variant="subtitle2" color="text.secondary">Amount</Typography>
                <Typography>${selectedClaim.amount}</Typography>
              </Grid>
              <Grid item xs={6}>
                <Typography variant="subtitle2" color="text.secondary">Status</Typography>
                <Chip label={selectedClaim.status} color={getStatusColor(selectedClaim.status)} size="small" />
              </Grid>
              <Grid item xs={12}>
                <Typography variant="subtitle2" color="text.secondary">Reason</Typography>
                <Typography>{selectedClaim.claimReason || 'N/A'}</Typography>
              </Grid>
              {selectedClaim.adminNote && (
                <Grid item xs={12}>
                  <Typography variant="subtitle2" color="text.secondary">Admin Note</Typography>
                  <Typography>{selectedClaim.adminNote}</Typography>
                </Grid>
              )}
            </Grid>
          )}
        </DialogContent>
        <DialogActions>
          <Button onClick={() => setDetailsOpen(false)}>Close</Button>
        </DialogActions>
      </Dialog>

      {/* Action Dialog */}
      <Dialog open={actionDialog.open} onClose={() => setActionDialog({ open: false, type: 'approve', claim: null })} maxWidth="sm" fullWidth>
        <DialogTitle>
          {actionDialog.type === 'approve' ? 'Approve' : actionDialog.type === 'reject' ? 'Reject' : 'Process Payment'}
          {' '}Claim
        </DialogTitle>
        <DialogContent>
          {actionDialog.type !== 'pay' && (
            <TextField
              fullWidth
              label="Admin Note"
              multiline
              rows={3}
              value={adminNote}
              onChange={(e) => setAdminNote(e.target.value)}
              sx={{ mt: 2 }}
              placeholder="Enter reason for this action..."
            />
          )}
        </DialogContent>
        <DialogActions>
          <Button onClick={() => setActionDialog({ open: false, type: 'approve', claim: null })}>
            Cancel
          </Button>
          <Button 
            variant="contained" 
            color={actionDialog.type === 'reject' ? 'error' : 'success'}
            onClick={handleAction}
            disabled={approveMutation.isPending || rejectMutation.isPending || payMutation.isPending}
          >
            {actionDialog.type === 'approve' ? 'Approve' : actionDialog.type === 'reject' ? 'Reject' : 'Pay'}
          </Button>
        </DialogActions>
      </Dialog>
    </Box>
  );
};

export default ClaimsManagement;
