import { useState } from 'react';
import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query';
import {
  Box, Card, CardContent, Typography, Grid, Table, TableBody, TableCell,
  TableContainer, TableHead, TableRow, Button, Chip, IconButton, Tooltip,
  TextField, MenuItem, Dialog, DialogTitle, DialogContent, DialogActions,
  Tabs, Tab, Alert, Snackbar
} from '@mui/material';
import { Add, Send, Edit, Notifications as NotificationsIcon, BarChart, Description } from '@mui/icons-material';
import { notificationsAPI } from '../../services/api';
interface Notification {
  id: string;
  title: string;
  message: string;
  type: 'push' | 'email' | 'sms';
  target: string;
  status: 'sent' | 'pending' | 'failed';
  createdAt: string;
}

interface Template {
  id: string;
  name: string;
  subject: string;
  body: string;
  type: string;
}

interface NotificationStats {
  delivered: number;
  opened: number;
  clicked: number;
}

const statusColor = (s: string) => {
  switch (s) {
    case 'sent': return 'success';
    case 'pending': return 'warning';
    case 'failed': return 'error';
    default: return 'default';
  }
};

const Notifications = () => {
  const queryClient = useQueryClient();
  const [tab, setTab] = useState(0);
  const [createDialog, setCreateDialog] = useState(false);
  const [templateDialog, setTemplateDialog] = useState(false);
  const [snackbar, setSnackbar] = useState({ open: false, message: '', severity: 'success' as 'success' | 'error' });
  const [form, setForm] = useState({ title: '', message: '', type: 'push' as string, target: 'all' });
  const [tplForm, setTplForm] = useState({ name: '', subject: '', body: '', type: 'email' });

  const { data: notifData } = useQuery({
    queryKey: ['notifications'],
    queryFn: () => notificationsAPI.getAll({ limit: 50 }),
  });

  const { data: tplData } = useQuery({
    queryKey: ['notification-templates'],
    queryFn: () => notificationsAPI.getTemplates(),
  });

  const { data: statsData } = useQuery({
    queryKey: ['notification-stats'],
    queryFn: () => notificationsAPI.getStats(),
  });

  const notifications: Notification[] = notifData?.data?.notifications || [];
  const templates: Template[] = tplData?.data?.templates || [];
  const stats: NotificationStats = statsData?.data || { delivered: 0, opened: 0, clicked: 0 };

  const createMutation = useMutation({
    mutationFn: notificationsAPI.create,
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['notifications'] });
      setCreateDialog(false);
      setForm({ title: '', message: '', type: 'push', target: 'all' });
      setSnackbar({ open: true, message: 'Notification created', severity: 'success' });
    },
    onError: () => setSnackbar({ open: true, message: 'Failed to create notification', severity: 'error' }),
  });

  const sendMutation = useMutation({
    mutationFn: notificationsAPI.send,
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['notifications'] });
      queryClient.invalidateQueries({ queryKey: ['notification-stats'] });
      setSnackbar({ open: true, message: 'Notification sent', severity: 'success' });
    },
    onError: () => setSnackbar({ open: true, message: 'Failed to send notification', severity: 'error' }),
  });

  const tplMutation = useMutation({
    mutationFn: notificationsAPI.createTemplate,
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['notification-templates'] });
      setTemplateDialog(false);
      setTplForm({ name: '', subject: '', body: '', type: 'email' });
      setSnackbar({ open: true, message: 'Template created', severity: 'success' });
    },
    onError: () => setSnackbar({ open: true, message: 'Failed to create template', severity: 'error' }),
  });

  const openRate = stats.delivered > 0 ? ((stats.opened / stats.delivered) * 100).toFixed(1) : '0.0';
  const clickRate = stats.opened > 0 ? ((stats.clicked / stats.opened) * 100).toFixed(1) : '0.0';

  return (
    <Box>
      <Box sx={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', mb: 3 }}>
        <Typography variant="h4" fontWeight="bold">Notifications</Typography>
        <Box sx={{ display: 'flex', gap: 1 }}>
          <Button variant="outlined" startIcon={<Description />} onClick={() => setTemplateDialog(true)}>New Template</Button>
          <Button variant="contained" startIcon={<Add />} onClick={() => setCreateDialog(true)}>Create Notification</Button>
        </Box>
      </Box>

      <Grid container spacing={3} sx={{ mb: 3 }}>
        <Grid item xs={12} md={3}>
          <Card><CardContent>
            <Typography color="text.secondary">Delivered</Typography>
            <Typography variant="h4" color="success.main">{stats.delivered.toLocaleString()}</Typography>
          </CardContent></Card>
        </Grid>
        <Grid item xs={12} md={3}>
          <Card><CardContent>
            <Typography color="text.secondary">Opened</Typography>
            <Typography variant="h4" color="info.main">{stats.opened.toLocaleString()}</Typography>
            <Typography variant="caption" color="text.secondary">{openRate}% open rate</Typography>
          </CardContent></Card>
        </Grid>
        <Grid item xs={12} md={3}>
          <Card><CardContent>
            <Typography color="text.secondary">Clicked</Typography>
            <Typography variant="h4" color="primary.main">{stats.clicked.toLocaleString()}</Typography>
            <Typography variant="caption" color="text.secondary">{clickRate}% click rate</Typography>
          </CardContent></Card>
        </Grid>
        <Grid item xs={12} md={3}>
          <Card><CardContent>
            <Typography color="text.secondary">Templates</Typography>
            <Typography variant="h4">{templates.length}</Typography>
          </CardContent></Card>
        </Grid>
      </Grid>

      <Tabs value={tab} onChange={(_, v) => setTab(v)} sx={{ mb: 2 }}>
        <Tab label="History" icon={<NotificationsIcon />} iconPosition="start" />
        <Tab label="Templates" icon={<Description />} iconPosition="start" />
        <Tab label="Statistics" icon={<BarChart />} iconPosition="start" />
      </Tabs>

      {tab === 0 && (
        <Card>
          <TableContainer>
            <Table>
              <TableHead>
                <TableRow>
                  <TableCell>Title</TableCell>
                  <TableCell>Type</TableCell>
                  <TableCell>Target</TableCell>
                  <TableCell>Status</TableCell>
                  <TableCell>Created</TableCell>
                  <TableCell align="right">Actions</TableCell>
                </TableRow>
              </TableHead>
              <TableBody>
                {notifications.length === 0 ? (
                  <TableRow><TableCell colSpan={6} align="center">No notifications</TableCell></TableRow>
                ) : notifications.map((n) => (
                  <TableRow key={n.id} hover>
                    <TableCell>
                      <Typography fontWeight={500}>{n.title}</Typography>
                      <Typography variant="caption" color="text.secondary" noWrap sx={{ maxWidth: 250, display: 'block' }}>{n.message}</Typography>
                    </TableCell>
                    <TableCell><Chip label={n.type.toUpperCase()} size="small" /></TableCell>
                    <TableCell>{n.target}</TableCell>
                    <TableCell><Chip label={n.status} color={statusColor(n.status)} size="small" /></TableCell>
                    <TableCell>{new Date(n.createdAt).toLocaleString()}</TableCell>
                    <TableCell align="right">
                      {n.status === 'pending' && (
                        <Tooltip title="Send Now">
                          <IconButton size="small" color="primary" onClick={() => sendMutation.mutate(n.id)}>
                            <Send />
                          </IconButton>
                        </Tooltip>
                      )}
                    </TableCell>
                  </TableRow>
                ))}
              </TableBody>
            </Table>
          </TableContainer>
        </Card>
      )}

      {tab === 1 && (
        <Card>
          <TableContainer>
            <Table>
              <TableHead>
                <TableRow>
                  <TableCell>Name</TableCell>
                  <TableCell>Subject</TableCell>
                  <TableCell>Type</TableCell>
                  <TableCell align="right">Actions</TableCell>
                </TableRow>
              </TableHead>
              <TableBody>
                {templates.length === 0 ? (
                  <TableRow><TableCell colSpan={4} align="center">No templates</TableCell></TableRow>
                ) : templates.map((t) => (
                  <TableRow key={t.id} hover>
                    <TableCell>{t.name}</TableCell>
                    <TableCell>{t.subject}</TableCell>
                    <TableCell><Chip label={t.type} size="small" /></TableCell>
                    <TableCell align="right">
                      <Tooltip title="Edit"><IconButton size="small"><Edit /></IconButton></Tooltip>
                    </TableCell>
                  </TableRow>
                ))}
              </TableBody>
            </Table>
          </TableContainer>
        </Card>
      )}

      {tab === 2 && (
        <Grid container spacing={3}>
          <Grid item xs={12} md={4}>
            <Card><CardContent sx={{ textAlign: 'center' }}>
              <Typography color="text.secondary">Total Delivered</Typography>
              <Typography variant="h3" color="success.main">{stats.delivered.toLocaleString()}</Typography>
            </CardContent></Card>
          </Grid>
          <Grid item xs={12} md={4}>
            <Card><CardContent sx={{ textAlign: 'center' }}>
              <Typography color="text.secondary">Open Rate</Typography>
              <Typography variant="h3" color="info.main">{openRate}%</Typography>
              <Typography variant="body2" color="text.secondary">{stats.opened.toLocaleString()} opened</Typography>
            </CardContent></Card>
          </Grid>
          <Grid item xs={12} md={4}>
            <Card><CardContent sx={{ textAlign: 'center' }}>
              <Typography color="text.secondary">Click Rate</Typography>
              <Typography variant="h3" color="primary.main">{clickRate}%</Typography>
              <Typography variant="body2" color="text.secondary">{stats.clicked.toLocaleString()} clicked</Typography>
            </CardContent></Card>
          </Grid>
        </Grid>
      )}

      <Dialog open={createDialog} onClose={() => setCreateDialog(false)} maxWidth="sm" fullWidth>
        <DialogTitle>Create Notification</DialogTitle>
        <DialogContent>
          <Box sx={{ display: 'flex', flexDirection: 'column', gap: 2, mt: 1 }}>
            <TextField label="Title" value={form.title} onChange={(e) => setForm({ ...form, title: e.target.value })} fullWidth required />
            <TextField label="Message" value={form.message} onChange={(e) => setForm({ ...form, message: e.target.value })} multiline rows={3} fullWidth required />
            <TextField label="Type" value={form.type} onChange={(e) => setForm({ ...form, type: e.target.value })} select fullWidth>
              <MenuItem value="push">Push</MenuItem>
              <MenuItem value="email">Email</MenuItem>
              <MenuItem value="sms">SMS</MenuItem>
            </TextField>
            <TextField label="Target" value={form.target} onChange={(e) => setForm({ ...form, target: e.target.value })} select fullWidth>
              <MenuItem value="all">All Users</MenuItem>
              <MenuItem value="segment">Segment</MenuItem>
              <MenuItem value="user">Specific User</MenuItem>
            </TextField>
          </Box>
        </DialogContent>
        <DialogActions>
          <Button onClick={() => setCreateDialog(false)}>Cancel</Button>
          <Button variant="contained" onClick={() => createMutation.mutate(form)} disabled={!form.title || !form.message || createMutation.isPending}>Create</Button>
        </DialogActions>
      </Dialog>

      <Dialog open={templateDialog} onClose={() => setTemplateDialog(false)} maxWidth="sm" fullWidth>
        <DialogTitle>Create Template</DialogTitle>
        <DialogContent>
          <Box sx={{ display: 'flex', flexDirection: 'column', gap: 2, mt: 1 }}>
            <TextField label="Name" value={tplForm.name} onChange={(e) => setTplForm({ ...tplForm, name: e.target.value })} fullWidth required />
            <TextField label="Subject" value={tplForm.subject} onChange={(e) => setTplForm({ ...tplForm, subject: e.target.value })} fullWidth required />
            <TextField label="Body" value={tplForm.body} onChange={(e) => setTplForm({ ...tplForm, body: e.target.value })} multiline rows={4} fullWidth required />
            <TextField label="Type" value={tplForm.type} onChange={(e) => setTplForm({ ...tplForm, type: e.target.value })} select fullWidth>
              <MenuItem value="email">Email</MenuItem>
              <MenuItem value="push">Push</MenuItem>
              <MenuItem value="sms">SMS</MenuItem>
            </TextField>
          </Box>
        </DialogContent>
        <DialogActions>
          <Button onClick={() => setTemplateDialog(false)}>Cancel</Button>
          <Button variant="contained" onClick={() => tplMutation.mutate(tplForm)} disabled={!tplForm.name || !tplForm.body || tplMutation.isPending}>Create</Button>
        </DialogActions>
      </Dialog>

      <Snackbar open={snackbar.open} autoHideDuration={4000} onClose={() => setSnackbar(s => ({ ...s, open: false }))}><Alert severity={snackbar.severity} variant="filled">{snackbar.message}</Alert></Snackbar>
    </Box>
  );
};
export default Notifications;
