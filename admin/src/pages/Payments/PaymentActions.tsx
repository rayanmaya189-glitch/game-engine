import {
  Table, TableBody, TableCell, TableContainer, TableHead, TableRow,
  Typography, Chip, IconButton, Tooltip, Box
} from '@mui/material';
import { Visibility, AccountBalance, CreditCard, SwapHoriz, CheckCircle, Cancel, Refresh } from '@mui/icons-material';

interface PaymentActionsProps {
  payments: any[];
  onViewDetails: (payment: any) => void;
  onApprove: (id: string) => void;
  onReject: (id: string) => void;
  onProcess: (id: string) => void;
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

const getMethodIcon = (method: string) => {
  switch (method.toLowerCase()) {
    case 'bank transfer': return <AccountBalance />;
    case 'credit card': return <CreditCard />;
    case 'crypto': return <SwapHoriz />;
    default: return <CreditCard />;
  }
};

const PaymentActions = ({ payments, onViewDetails, onApprove, onReject, onProcess }: PaymentActionsProps) => {
  return (
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
                      onClick={() => onViewDetails(payment)}
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
                          onClick={() => onApprove(payment.id)}
                        >
                          <CheckCircle fontSize="small" />
                        </IconButton>
                      </Tooltip>
                      <Tooltip title="Reject">
                        <IconButton
                          size="small"
                          color="error"
                          onClick={() => onReject(payment.id)}
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
                        onClick={() => onProcess(payment.id)}
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
  );
};

export default PaymentActions;
