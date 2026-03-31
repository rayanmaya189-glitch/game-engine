import {
  Box, Typography, Grid, Chip,
  Dialog, DialogActions, DialogContent, DialogTitle, Button
} from '@mui/material';

interface PaymentDetailsProps {
  open: boolean;
  payment: any;
  onClose: () => void;
  onApprove: (id: string) => void;
  onReject: (id: string) => void;
}

const getStatusColor = (status: string) => {
  switch (status) {
    case 'completed': return 'success';
    case 'pending': return 'warning';
    case 'failed': return 'error';
    case 'processing': return 'info';
    default: return 'default';
  }
};

const PaymentDetails = ({ open, payment, onClose, onApprove, onReject }: PaymentDetailsProps) => {
  return (
    <Dialog open={open} onClose={onClose} maxWidth="sm" fullWidth>
      <DialogTitle>Transaction Details</DialogTitle>
      <DialogContent>
        {payment && (
          <Grid container spacing={2} sx={{ mt: 1 }}>
            <Grid item xs={6}>
              <Typography variant="caption" color="text.secondary">Transaction ID</Typography>
              <Typography fontFamily="monospace">{payment.id}</Typography>
            </Grid>
            <Grid item xs={6}>
              <Typography variant="caption" color="text.secondary">Date</Typography>
              <Typography>{payment.date}</Typography>
            </Grid>
            <Grid item xs={12}>
              <Typography variant="caption" color="text.secondary">User</Typography>
              <Typography>{payment.user} ({payment.userId})</Typography>
            </Grid>
            <Grid item xs={6}>
              <Typography variant="caption" color="text.secondary">Amount</Typography>
              <Typography variant="h6" color={payment.type === 'deposit' ? 'success.main' : 'error.main'}>
                {payment.type === 'deposit' ? '+' : '-'}${payment.amount.toLocaleString()}
              </Typography>
            </Grid>
            <Grid item xs={6}>
              <Typography variant="caption" color="text.secondary">Method</Typography>
              <Typography>{payment.method}</Typography>
            </Grid>
            <Grid item xs={6}>
              <Typography variant="caption" color="text.secondary">Type</Typography>
              <Chip
                label={payment.type}
                color={payment.type === 'deposit' ? 'success' : 'error'}
                size="small"
              />
            </Grid>
            <Grid item xs={6}>
              <Typography variant="caption" color="text.secondary">Status</Typography>
              <Chip
                label={payment.status}
                color={getStatusColor(payment.status) as any}
                size="small"
              />
            </Grid>
          </Grid>
        )}
      </DialogContent>
      <DialogActions>
        <Button onClick={onClose}>Close</Button>
        {payment?.status === 'pending' && (
          <>
            <Button
              color="error"
              onClick={() => onReject(payment.id)}
            >
              Reject
            </Button>
            <Button
              color="success"
              variant="contained"
              onClick={() => onApprove(payment.id)}
            >
              Approve
            </Button>
          </>
        )}
      </DialogActions>
    </Dialog>
  );
};

export default PaymentDetails;
