import { useState } from 'react';
import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query';
import {
  Box, Card, CardContent, Typography, Grid, TextField, InputAdornment,
  Table, TableBody, TableCell, TableContainer, TableHead, TableRow,
  Button, Chip, IconButton, Tooltip, Dialog, DialogTitle, DialogContent,
  DialogActions, FormControl, InputLabel, Select, MenuItem, LinearProgress
} from '@mui/material';
import {
  Search, Add, Edit, Delete, Image as ImageIcon, Refresh
} from '@mui/icons-material';
import { useAppDispatch } from '../../store/hooks';
import { showSnackbar } from '../../store/slices/uiSlice';
import { bannerAPI } from '../../services/api';

interface Banner {
  id: string;
  title: string;
  type: 'HERO' | 'SIDEBAR' | 'POPUP';
  image_url: string;
  click_url: string;
  status: 'active' | 'inactive' | 'scheduled';
  priority: number;
  start_date: string;
  end_date: string;
}

const emptyForm: Partial<Banner> = {
  title: '', type: 'HERO', image_url: '', click_url: '', status: 'active', priority: 0, start_date: '', end_date: '',
};

const BannerManagement = () => {
  const dispatch = useAppDispatch();
  const queryClient = useQueryClient();
  const [search, setSearch] = useState('');
  const [typeFilter, setTypeFilter] = useState('');
  const [statusFilter, setStatusFilter] = useState('');
  const [dialogOpen, setDialogOpen] = useState(false);
  const [editingBanner, setEditingBanner] = useState<Banner | null>(null);
  const [form, setForm] = useState<Partial<Banner>>(emptyForm);
  const [deleteConfirm, setDeleteConfirm] = useState<string | null>(null);

  const { data, isLoading, refetch } = useQuery({
    queryKey: ['banners', typeFilter, statusFilter],
    queryFn: () => bannerAPI.getAll({ type: typeFilter || undefined, status: statusFilter || undefined, page: 1, limit: 50 }),
  });

  const createMutation = useMutation({
    mutationFn: (bannerData: Partial<Banner>) => bannerAPI.create(bannerData),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['banners'] });
      dispatch(showSnackbar({ message: 'Banner created successfully', severity: 'success' }));
      handleCloseDialog();
    },
    onError: (error: any) => dispatch(showSnackbar({ message: error.message || 'Failed to create banner', severity: 'error' })),
  });

  const updateMutation = useMutation({
    mutationFn: ({ id, data }: { id: string; data: Partial<Banner> }) => bannerAPI.update(id, data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['banners'] });
      dispatch(showSnackbar({ message: 'Banner updated successfully', severity: 'success' }));
      handleCloseDialog();
    },
    onError: (error: any) => dispatch(showSnackbar({ message: error.message || 'Failed to update banner', severity: 'error' })),
  });

  const deleteMutation = useMutation({
    mutationFn: (id: string) => bannerAPI.delete(id),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['banners'] });
      dispatch(showSnackbar({ message: 'Banner deleted', severity: 'success' }));
      setDeleteConfirm(null);
    },
    onError: (error: any) => dispatch(showSnackbar({ message: error.message || 'Failed to delete banner', severity: 'error' })),
  });

  const statusMutation = useMutation({
    mutationFn: ({ id, status }: { id: string; status: string }) => bannerAPI.updateStatus(id, { status }),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['banners'] });
      dispatch(showSnackbar({ message: 'Banner status updated', severity: 'success' }));
    },
    onError: (error: any) => dispatch(showSnackbar({ message: error.message || 'Failed to update status', severity: 'error' })),
  });

  const mockBanners: Banner[] = [
    { id: 'BNR001', title: 'Summer Sale', type: 'HERO', image_url: '/banners/summer.jpg', click_url: '/promotions/summer', status: 'active', priority: 1, start_date: '2024-06-01', end_date: '2024-08-31' },
    { id: 'BNR002', title: 'New Game Launch', type: 'SIDEBAR', image_url: '/banners/newgame.jpg', click_url: '/games/new', status: 'active', priority: 2, start_date: '2024-01-01', end_date: '2024-12-31' },
    { id: 'BNR003', title: 'Welcome Bonus', type: 'POPUP', image_url: '/banners/welcome.jpg', click_url: '/bonuses/welcome', status: 'active', priority: 1, start_date: '2024-01-01', end_date: '2024-12-31' },
    { id: 'BNR004', title: 'Winter Tournament', type: 'HERO', image_url: '/banners/winter.jpg', click_url: '/tournaments/winter', status: 'scheduled', priority: 3, start_date: '2024-12-01', end_date: '2025-02-28' },
    { id: 'BNR005', title: 'VIP Promotion', type: 'SIDEBAR', image_url: '/banners/vip.jpg', click_url: '/vip', status: 'inactive', priority: 5, start_date: '2024-03-01', end_date: '2024-05-31' },
  ];

  const banners: Banner[] = data?.data?.banners || mockBanners;

  const filteredBanners = banners.filter((b) => {
    if (search && !b.title.toLowerCase().includes(search.toLowerCase())) return false;
    return true;
  });

  const handleCloseDialog = () => {
    setDialogOpen(false);
    setEditingBanner(null);
    setForm(emptyForm);
  };

  const handleOpenEdit = (banner: Banner) => {
    setEditingBanner(banner);
    setForm(banner);
    setDialogOpen(true);
  };

  const handleSubmit = () => {
    if (editingBanner) {
      updateMutation.mutate({ id: editingBanner.id, data: form });
    } else {
      createMutation.mutate(form);
    }
  };

  const getTypeColor = (type: string) => {
    switch (type) {
      case 'HERO': return 'primary';
      case 'SIDEBAR': return 'info';
      case 'POPUP': return 'secondary';
      default: return 'default';
    }
  };

  const getStatusColor = (status: string) => {
    switch (status) {
      case 'active': return 'success';
      case 'inactive': return 'default';
      case 'scheduled': return 'warning';
      default: return 'default';
    }
  };

  return (
    <Box>
      <Box sx={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', mb: 3 }}>
        <Typography variant="h4" fontWeight="bold">Banner Management</Typography>
        <Box sx={{ display: 'flex', gap: 1 }}>
          <Button variant="outlined" startIcon={<Refresh />} onClick={() => refetch()}>Refresh</Button>
          <Button variant="contained" startIcon={<Add />} onClick={() => setDialogOpen(true)}>Create Banner</Button>
        </Box>
      </Box>

      {isLoading && <LinearProgress sx={{ mb: 2 }} />}

      <Grid container spacing={3} sx={{ mb: 3 }}>
        <Grid item xs={12} md={4}>
          <Card sx={{ borderLeft: '4px solid #22c55e' }}>
            <CardContent><Typography color="text.secondary" variant="body2">Active</Typography><Typography variant="h4" fontWeight="bold" color="success.main">{banners.filter(b => b.status === 'active').length}</Typography></CardContent>
          </Card>
        </Grid>
        <Grid item xs={12} md={4}>
          <Card sx={{ borderLeft: '4px solid #f59e0b' }}>
            <CardContent><Typography color="text.secondary" variant="body2">Scheduled</Typography><Typography variant="h4" fontWeight="bold" color="warning.main">{banners.filter(b => b.status === 'scheduled').length}</Typography></CardContent>
          </Card>
        </Grid>
        <Grid item xs={12} md={4}>
          <Card sx={{ borderLeft: '4px solid #6b7280' }}>
            <CardContent><Typography color="text.secondary" variant="body2">Inactive</Typography><Typography variant="h4" fontWeight="bold">{banners.filter(b => b.status === 'inactive').length}</Typography></CardContent>
          </Card>
        </Grid>
      </Grid>

      <Card sx={{ mb: 3 }}>
        <CardContent>
          <Grid container spacing={2}>
            <Grid item xs={12} md={4}>
              <TextField fullWidth placeholder="Search banners..." value={search} onChange={(e) => setSearch(e.target.value)}
                InputProps={{ startAdornment: <InputAdornment position="start"><Search /></InputAdornment> }} size="small" />
            </Grid>
            <Grid item xs={12} md={3}>
              <FormControl fullWidth size="small">
                <InputLabel>Type</InputLabel>
                <Select value={typeFilter} label="Type" onChange={(e) => setTypeFilter(e.target.value)}>
                  <MenuItem value="">All</MenuItem>
                  <MenuItem value="HERO">Hero</MenuItem>
                  <MenuItem value="SIDEBAR">Sidebar</MenuItem>
                  <MenuItem value="POPUP">Popup</MenuItem>
                </Select>
              </FormControl>
            </Grid>
            <Grid item xs={12} md={3}>
              <FormControl fullWidth size="small">
                <InputLabel>Status</InputLabel>
                <Select value={statusFilter} label="Status" onChange={(e) => setStatusFilter(e.target.value)}>
                  <MenuItem value="">All</MenuItem>
                  <MenuItem value="active">Active</MenuItem>
                  <MenuItem value="inactive">Inactive</MenuItem>
                  <MenuItem value="scheduled">Scheduled</MenuItem>
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
                <TableCell>Title</TableCell>
                <TableCell>Type</TableCell>
                <TableCell>Priority</TableCell>
                <TableCell>Start Date</TableCell>
                <TableCell>End Date</TableCell>
                <TableCell>Status</TableCell>
                <TableCell align="center">Actions</TableCell>
              </TableRow>
            </TableHead>
            <TableBody>
              {filteredBanners.length === 0 ? (
                <TableRow><TableCell colSpan={7} align="center">No banners found</TableCell></TableRow>
              ) : (
                filteredBanners.map((banner) => (
                  <TableRow key={banner.id} hover>
                    <TableCell>
                      <Box sx={{ display: 'flex', alignItems: 'center', gap: 1 }}>
                        <ImageIcon color="primary" fontSize="small" />
                        {banner.title}
                      </Box>
                    </TableCell>
                    <TableCell><Chip label={banner.type} size="small" color={getTypeColor(banner.type) as any} /></TableCell>
                    <TableCell>{banner.priority}</TableCell>
                    <TableCell>{banner.start_date}</TableCell>
                    <TableCell>{banner.end_date}</TableCell>
                    <TableCell><Chip label={banner.status} size="small" color={getStatusColor(banner.status) as any} /></TableCell>
                    <TableCell align="center">
                      <Tooltip title="Edit"><IconButton size="small" onClick={() => handleOpenEdit(banner)}><Edit fontSize="small" /></IconButton></Tooltip>
                      <Tooltip title="Delete"><IconButton size="small" color="error" onClick={() => setDeleteConfirm(banner.id)}><Delete fontSize="small" /></IconButton></Tooltip>
                    </TableCell>
                  </TableRow>
                ))
              )}
            </TableBody>
          </Table>
        </TableContainer>
      </Card>

      <Dialog open={dialogOpen} onClose={handleCloseDialog} maxWidth="md" fullWidth>
        <DialogTitle>{editingBanner ? 'Edit Banner' : 'Create Banner'}</DialogTitle>
        <DialogContent>
          <Grid container spacing={2} sx={{ mt: 1 }}>
            <Grid item xs={12}><TextField fullWidth label="Title" value={form.title} onChange={(e) => setForm({ ...form, title: e.target.value })} required /></Grid>
            <Grid item xs={12} md={6}>
              <FormControl fullWidth><InputLabel>Type</InputLabel>
                <Select value={form.type} label="Type" onChange={(e) => setForm({ ...form, type: e.target.value as Banner['type'] })}>
                  <MenuItem value="HERO">Hero</MenuItem><MenuItem value="SIDEBAR">Sidebar</MenuItem><MenuItem value="POPUP">Popup</MenuItem>
                </Select>
              </FormControl>
            </Grid>
            <Grid item xs={12} md={6}><TextField fullWidth label="Priority" type="number" value={form.priority} onChange={(e) => setForm({ ...form, priority: Number(e.target.value) })} /></Grid>
            <Grid item xs={12}><TextField fullWidth label="Image URL" value={form.image_url} onChange={(e) => setForm({ ...form, image_url: e.target.value })} required /></Grid>
            <Grid item xs={12}><TextField fullWidth label="Click URL" value={form.click_url} onChange={(e) => setForm({ ...form, click_url: e.target.value })} /></Grid>
            <Grid item xs={12} md={6}><TextField fullWidth label="Start Date" type="date" value={form.start_date} onChange={(e) => setForm({ ...form, start_date: e.target.value })} InputLabelProps={{ shrink: true }} /></Grid>
            <Grid item xs={12} md={6}><TextField fullWidth label="End Date" type="date" value={form.end_date} onChange={(e) => setForm({ ...form, end_date: e.target.value })} InputLabelProps={{ shrink: true }} /></Grid>
            <Grid item xs={12} md={6}>
              <FormControl fullWidth><InputLabel>Status</InputLabel>
                <Select value={form.status} label="Status" onChange={(e) => setForm({ ...form, status: e.target.value as Banner['status'] })}>
                  <MenuItem value="active">Active</MenuItem><MenuItem value="inactive">Inactive</MenuItem><MenuItem value="scheduled">Scheduled</MenuItem>
                </Select>
              </FormControl>
            </Grid>
          </Grid>
        </DialogContent>
        <DialogActions>
          <Button onClick={handleCloseDialog}>Cancel</Button>
          <Button variant="contained" onClick={handleSubmit}
            disabled={createMutation.isPending || updateMutation.isPending || !form.title || !form.image_url}>
            {editingBanner ? 'Update' : 'Create'}
          </Button>
        </DialogActions>
      </Dialog>

      <Dialog open={!!deleteConfirm} onClose={() => setDeleteConfirm(null)}>
        <DialogTitle>Delete Banner</DialogTitle>
        <DialogContent><Typography>Are you sure you want to delete this banner?</Typography></DialogContent>
        <DialogActions>
          <Button onClick={() => setDeleteConfirm(null)}>Cancel</Button>
          <Button variant="contained" color="error" onClick={() => deleteConfirm && deleteMutation.mutate(deleteConfirm)}
            disabled={deleteMutation.isPending}>Delete</Button>
        </DialogActions>
      </Dialog>
    </Box>
  );
};

export default BannerManagement;
