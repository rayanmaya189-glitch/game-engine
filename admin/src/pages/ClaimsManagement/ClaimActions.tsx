import { useState } from 'react';
import { useMutation, useQueryClient } from '@tanstack/react-query';
import {
  Box, Typography, Chip, Button, IconButton, Tooltip,
  Dialog, DialogTitle, DialogContent, DialogActions, TextField
} from '@mui/material';
import { CheckCircle, Cancel, Refresh } from '@mui/icons-material';
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

interface ClaimActionsProps {
  claim: Claim;
  activeTab: ClaimType;
}

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

const ClaimActions = ({ claim, activeTab }: ClaimActionsProps) => {
  const dispatch = useAppDispatch();
  const queryClient = useQueryClient();
  const [actionDialog, setActionDialog] = useState<{ open: boolean; type: 'approve' | 'reject' | 'pay' }>({
    open: false,
    type: 'approve',
  });
  const [adminNote, setAdminNote] = useState('');

  const approveMutation = useMutation({
    mutationFn: ({ id, type, note }: { id: string; type: ClaimType; note: string }) => {
      if (type === 'commission') return claimsAPI.approveCommissionClaim(id, { adminNote: note });
      if (type === 'rebet') return claimsAPI.approveRebetClaim(id, { adminNote: note });
      return claimsAPI.approveInsuranceClaim(id, { adminNote: note });
    },
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: [`${activeTab}-claims`] });
      dispatch(showSnackbar({ message: 'Claim approved successfully', severity: 'success' }));
      setActionDialog({ open: false, type: 'approve' });
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
      setActionDialog({ open: false, type: 'reject' });
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
      setActionDialog({ open: false, type: 'pay' });
    },
    onError: (error: any) => {
      dispatch(showSnackbar({ message: error.message || 'Failed to process payment', severity: 'error' }));
    },
  });

  const handleAction = () => {
    const id = claim.id.toString();
    if (actionDialog.type === 'approve') {
      approveMutation.mutate({ id, type: activeTab, note: adminNote });
    } else if (actionDialog.type === 'reject') {
      rejectMutation.mutate({ id, type: activeTab, note: adminNote });
    } else {
      payMutation.mutate({ id, type: activeTab });
    }
  };

  return (
    <>
      <Box sx={{ display: 'flex', alignItems: 'center', gap: 0.5 }}>
        <Chip label={claim.status} color={getStatusColor(claim.status)} size="small" />
        {claim.status === 'PENDING' && (
          <>
            <Tooltip title="Approve">
              <IconButton
                size="small"
                color="success"
                onClick={() => setActionDialog({ open: true, type: 'approve' })}
              >
                <CheckCircle fontSize="small" />
              </IconButton>
            </Tooltip>
            <Tooltip title="Reject">
              <IconButton
                size="small"
                color="error"
                onClick={() => setActionDialog({ open: true, type: 'reject' })}
              >
                <Cancel fontSize="small" />
              </IconButton>
            </Tooltip>
          </>
        )}
        {claim.status === 'APPROVED' && (
          <Tooltip title="Pay">
            <IconButton
              size="small"
              color="primary"
              onClick={() => setActionDialog({ open: true, type: 'pay' })}
            >
              <Refresh fontSize="small" />
            </IconButton>
          </Tooltip>
        )}
      </Box>

      <Dialog open={actionDialog.open} onClose={() => setActionDialog({ open: false, type: 'approve' })} maxWidth="sm" fullWidth>
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
          <Button onClick={() => setActionDialog({ open: false, type: 'approve' })}>
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
    </>
  );
};

export { getStatusColor };
export default ClaimActions;
