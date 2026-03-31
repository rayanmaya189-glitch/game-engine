import { useState } from 'react';
import { useQuery } from '@tanstack/react-query';
import {
  Box, Card, CardContent, Typography, Tabs, Tab, Table, TableBody,
  TableCell, TableContainer, TableHead, TableRow, Chip,
  Dialog, DialogTitle, DialogContent, DialogActions, Button,
  Grid, LinearProgress, IconButton, Tooltip, Pagination
} from '@mui/material';
import { Visibility } from '@mui/icons-material';
import { claimsAPI } from '../../services/api';
import ClaimActions, { getStatusColor } from './ClaimActions';

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
  const [activeTab, setActiveTab] = useState<ClaimType>('commission');
  const [selectedClaim, setSelectedClaim] = useState<Claim | null>(null);
  const [detailsOpen, setDetailsOpen] = useState(false);
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
                        <ClaimActions claim={claim} activeTab={activeTab} />
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
    </Box>
  );
};

export default ClaimsManagement;
