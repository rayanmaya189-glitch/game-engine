import { useState } from 'react';
import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query';
import {
  Box, Card, CardContent, Typography, Grid, TextField, InputAdornment,
  Table, TableBody, TableCell, TableContainer, TableHead, TableRow,
  Button, Chip, IconButton, Tooltip, Dialog, DialogTitle, DialogContent,
  DialogActions, FormControl, InputLabel, Select, MenuItem, Switch, FormControlLabel
} from '@mui/material';
import { Search, Add, Edit, Block, CheckCircle } from '@mui/icons-material';
import { useAppDispatch } from '../../store/hooks';
import { showSnackbar } from '../../store/slices/uiSlice';
import { merchantsAPI } from '../../services/api';

const Merchants = () => {
  const dispatch = useAppDispatch();
  const queryClient = useQueryClient();
  const [search, setSearch] = useState('');
  const [dialogOpen, setDialogOpen] = useState(false);
  const [selectedMerchant, setSelectedMerchant] = useState<any>(null);

  const { data, isLoading } = useQuery({
    queryKey: ['merchants', search],
    queryFn: () => merchantsAPI.getAll({ search, page: 1, limit: 20 }),
  });

  const updateStatusMutation = useMutation({
    mutationFn: ({ id, status }: { id: string; status: string }) =>
      merchantsAPI.updateStatus(id, { status }),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['merchants'] });
      dispatch(showSnackbar({ message: 'Merchant status updated', severity: 'success' }));
    },
    onError: (error: any) => {
      dispatch(showSnackbar({ message: error.message, severity: 'error' }));
    },
  });

  const merchants = data?.data?.merchants || [];

  return (
    <Box>
      <Box sx={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', mb: 3 }}>
        <Typography variant="h4" fontWeight="bold">Merchants Management</Typography>
        <Button variant="contained" startIcon={<Add />} onClick={() => setDialogOpen(true)}>
          Add Merchant
        </Button>
      </Box>

      <Card sx={{ mb: 3 }}>
        <CardContent>
          <Grid container spacing={2} alignItems="center">
            <Grid item xs={12} md={6}>
              <TextField
                fullWidth
                placeholder="Search merchants by name, email..."
                value={search}
                onChange={(e) => setSearch(e.target.value)}
                InputProps={{
                  startAdornment: <InputAdornment position="start"><Search /></InputAdornment>,
                }}
              />
            </Grid>
            <Grid item xs={12} md={3}>
              <FormControl fullWidth>
                <InputLabel>Status</InputLabel>
                <Select label="Status" defaultValue="">
                  <MenuItem value="">All</MenuItem>
                  <MenuItem value="active">Active</MenuItem>
                  <MenuItem value="inactive">Inactive</MenuItem>
                  <MenuItem value="suspended">Suspended</MenuItem>
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
                <TableCell>ID</TableCell>
                <TableCell>Name</TableCell>
                <TableCell>Email</TableCell>
                <TableCell>Commission %</TableCell>
                <TableCell>Players</TableCell>
                <TableCell>Revenue</TableCell>
                <TableCell>Status</TableCell>
                <TableCell>Created</TableCell>
                <TableCell align="right">Actions</TableCell>
              </TableRow>
            </TableHead>
            <TableBody>
              {isLoading ? (
                <TableRow>
                  <TableCell colSpan={9} align="center">Loading...</TableCell>
                </TableRow>
              ) : merchants.length === 0 ? (
                <TableRow>
                  <TableCell colSpan={9} align="center">No merchants found</TableCell>
                </TableRow>
              ) : (
                merchants.map((merchant: any) => (
                  <TableRow key={merchant.id} hover>
                    <TableCell>{merchant.id}</TableCell>
                    <TableCell variant='head'>{merchant.name}</TableCell>
                    <TableCell>{merchant.email}</TableCell>
                    <TableCell>{merchant.commissionRate}%</TableCell>
                    <TableCell>{merchant.playerCount || 0}</TableCell>
                    <TableCell>${merchant.revenue || 0}</TableCell>
                    <TableCell>
                      <Chip
                        label={merchant.status}
                        color={merchant.status === 'active' ? 'success' : merchant.status === 'suspended' ? 'error' : 'default'}
                        size="small"
                      />
                    </TableCell>
                    <TableCell>{new Date(merchant.createdAt).toLocaleDateString()}</TableCell>
                    <TableCell align="right">
                      <Tooltip title="Edit">
                        <IconButton size="small" onClick={() => { setSelectedMerchant(merchant); setDialogOpen(true); }}>
                          <Edit />
                        </IconButton>
                      </Tooltip>
                      <Tooltip title={merchant.status === 'active' ? 'Suspend' : 'Activate'}>
                        <IconButton size="small" color={merchant.status === 'active' ? 'error' : 'success'}>
                          {merchant.status === 'active' ? <Block /> : <CheckCircle />}
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

      <Dialog open={dialogOpen} onClose={() => { setDialogOpen(false); setSelectedMerchant(null); }} maxWidth="md" fullWidth>
        <DialogTitle>{selectedMerchant ? 'Edit Merchant' : 'Add New Merchant'}</DialogTitle>
        <DialogContent>
          <Grid container spacing={2} sx={{ mt: 1 }}>
            <Grid item xs={12} md={6}>
              <TextField fullWidth label="Merchant Name" defaultValue={selectedMerchant?.name} />
            </Grid>
            <Grid item xs={12} md={6}>
              <TextField fullWidth label="Email" defaultValue={selectedMerchant?.email} />
            </Grid>
            <Grid item xs={12} md={6}>
              <TextField fullWidth label="Contact Phone" defaultValue={selectedMerchant?.phone} />
            </Grid>
            <Grid item xs={12} md={6}>
              <TextField fullWidth label="Commission Rate (%)" type="number" defaultValue={selectedMerchant?.commissionRate || 10} />
            </Grid>
            <Grid item xs={12} md={6}>
              <FormControl fullWidth>
                <InputLabel>Status</InputLabel>
                <Select label="Status" defaultValue={selectedMerchant?.status || 'active'}>
                  <MenuItem value="active">Active</MenuItem>
                  <MenuItem value="inactive">Inactive</MenuItem>
                  <MenuItem value="suspended">Suspended</MenuItem>
                </Select>
              </FormControl>
            </Grid>
            <Grid item xs={12}>
              <FormControlLabel
                control={<Switch defaultChecked={selectedMerchant?.isActive ?? true} />}
                label="Active"
              />
            </Grid>
          </Grid>
        </DialogContent>
        <DialogActions>
          <Button onClick={() => { setDialogOpen(false); setSelectedMerchant(null); }}>Cancel</Button>
          <Button variant="contained" onClick={() => { setDialogOpen(false); setSelectedMerchant(null); }}>
            {selectedMerchant ? 'Update' : 'Create'}
          </Button>
        </DialogActions>
      </Dialog>
    </Box>
  );
};

export default Merchants;
