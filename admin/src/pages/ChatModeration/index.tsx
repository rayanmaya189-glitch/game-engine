import { useState } from 'react';
import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query';
import {
  Box, Card, CardContent, Typography, Grid, Table, TableBody, TableCell,
  TableContainer, TableHead, TableRow, Button, Chip, IconButton, Tooltip,
  TextField, InputAdornment, Dialog, DialogTitle, DialogContent, DialogActions,
  MenuItem, Tabs, Tab, Alert, Snackbar
} from '@mui/material';
import { Search, Delete, Block, Visibility, Chat, VolumeOff } from '@mui/icons-material';
import { chatModerationAPI } from '../../services/api';

interface ChatRoom {
  id: string;
  name: string;
  activeMessages: number;
  participants: number;
}

interface ChatMessage {
  id: string;
  userId: string;
  username: string;
  content: string;
  roomName: string;
  createdAt: string;
}

interface MutedUser {
  id: string;
  username: string;
  mutedUntil: string;
  reason: string;
}

const ChatModeration = () => {
  const queryClient = useQueryClient();
  const [tab, setTab] = useState(0);
  const [search, setSearch] = useState('');
  const [deleteDialog, setDeleteDialog] = useState<{ open: boolean; id: string | null }>({ open: false, id: null });
  const [deleteReason, setDeleteReason] = useState('');
  const [banDialog, setBanDialog] = useState<{ open: boolean; id: string | null }>({ open: false, id: null });
  const [banReason, setBanReason] = useState('');
  const [snackbar, setSnackbar] = useState({ open: false, message: '', severity: 'success' as 'success' | 'error' });

  const { data: roomsData } = useQuery({
    queryKey: ['chat-rooms'],
    queryFn: () => chatModerationAPI.getRooms(),
  });

  const { data: messagesData } = useQuery({
    queryKey: ['chat-messages', search],
    queryFn: () => chatModerationAPI.getMessages({ search, limit: 50 }),
  });

  const { data: mutedData } = useQuery({
    queryKey: ['muted-users'],
    queryFn: () => chatModerationAPI.getMutedUsers(),
  });

  const rooms: ChatRoom[] = roomsData?.data?.rooms || [];
  const messages: ChatMessage[] = messagesData?.data?.messages || [];
  const mutedUsers: MutedUser[] = mutedData?.data?.users || [];

  const deleteMutation = useMutation({
    mutationFn: (id: string) => chatModerationAPI.deleteMessage(id, { reason: deleteReason }),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['chat-messages'] });
      setDeleteDialog({ open: false, id: null });
      setDeleteReason('');
      setSnackbar({ open: true, message: 'Message deleted', severity: 'success' });
    },
    onError: () => setSnackbar({ open: true, message: 'Failed to delete message', severity: 'error' }),
  });

  const banMutation = useMutation({
    mutationFn: (id: string) => chatModerationAPI.banUser(id, { reason: banReason }),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['chat-rooms'] });
      queryClient.invalidateQueries({ queryKey: ['muted-users'] });
      setBanDialog({ open: false, id: null });
      setBanReason('');
      setSnackbar({ open: true, message: 'User banned from chat', severity: 'success' });
    },
    onError: () => setSnackbar({ open: true, message: 'Failed to ban user', severity: 'error' }),
  });

  return (
    <Box>
      <Typography variant="h4" fontWeight="bold" sx={{ mb: 3 }}>Chat Moderation</Typography>

      <Grid container spacing={3} sx={{ mb: 3 }}>
        <Grid item xs={12} md={4}>
          <Card><CardContent>
            <Typography color="text.secondary">Active Rooms</Typography>
            <Typography variant="h4">{rooms.length}</Typography>
          </CardContent></Card>
        </Grid>
        <Grid item xs={12} md={4}>
          <Card><CardContent>
            <Typography color="text.secondary">Total Messages</Typography>
            <Typography variant="h4">{rooms.reduce((s, r) => s + r.activeMessages, 0)}</Typography>
          </CardContent></Card>
        </Grid>
        <Grid item xs={12} md={4}>
          <Card><CardContent>
            <Typography color="text.secondary">Muted Users</Typography>
            <Typography variant="h4" color="warning.main">{mutedUsers.length}</Typography>
          </CardContent></Card>
        </Grid>
      </Grid>

      <Tabs value={tab} onChange={(_, v) => setTab(v)} sx={{ mb: 2 }}>
        <Tab label="Rooms" icon={<Chat />} iconPosition="start" />
        <Tab label="Messages" icon={<Visibility />} iconPosition="start" />
        <Tab label="Muted Users" icon={<VolumeOff />} iconPosition="start" />
      </Tabs>

      {tab === 0 && (
        <Card>
          <TableContainer>
            <Table>
              <TableHead>
                <TableRow>
                  <TableCell>Room Name</TableCell>
                  <TableCell>Active Messages</TableCell>
                  <TableCell>Participants</TableCell>
                  <TableCell align="right">Actions</TableCell>
                </TableRow>
              </TableHead>
              <TableBody>
                {rooms.map((room) => (
                  <TableRow key={room.id} hover>
                    <TableCell><Box sx={{ display: 'flex', alignItems: 'center', gap: 1 }}><Chat color="primary" fontSize="small" />{room.name}</Box></TableCell>
                    <TableCell><Chip label={room.activeMessages} size="small" color="primary" /></TableCell>
                    <TableCell>{room.participants}</TableCell>
                    <TableCell align="right">
                      <Tooltip title="View Messages"><IconButton size="small"><Visibility /></IconButton></Tooltip>
                    </TableCell>
                  </TableRow>
                ))}
              </TableBody>
            </Table>
          </TableContainer>
        </Card>
      )}

      {tab === 1 && (
        <>
          <Card sx={{ mb: 2 }}>
            <CardContent>
              <TextField
                fullWidth
                placeholder="Search messages by user or keyword..."
                value={search}
                onChange={(e) => setSearch(e.target.value)}
                InputProps={{ startAdornment: <InputAdornment position="start"><Search /></InputAdornment> }}
              />
            </CardContent>
          </Card>
          <Card>
            <TableContainer>
              <Table>
                <TableHead>
                  <TableRow>
                    <TableCell>User</TableCell>
                    <TableCell>Message</TableCell>
                    <TableCell>Room</TableCell>
                    <TableCell>Time</TableCell>
                    <TableCell align="right">Actions</TableCell>
                  </TableRow>
                </TableHead>
                <TableBody>
                  {messages.map((msg) => (
                    <TableRow key={msg.id} hover>
                      <TableCell>{msg.username}</TableCell>
                      <TableCell sx={{ maxWidth: 300, overflow: 'hidden', textOverflow: 'ellipsis', whiteSpace: 'nowrap' }}>{msg.content}</TableCell>
                      <TableCell><Chip label={msg.roomName} size="small" /></TableCell>
                      <TableCell>{new Date(msg.createdAt).toLocaleString()}</TableCell>
                      <TableCell align="right">
                        <Tooltip title="Delete Message">
                          <IconButton size="small" color="error" onClick={() => setDeleteDialog({ open: true, id: msg.id })}>
                            <Delete />
                          </IconButton>
                        </Tooltip>
                        <Tooltip title="Ban User">
                          <IconButton size="small" color="warning" onClick={() => setBanDialog({ open: true, id: msg.userId })}>
                            <Block />
                          </IconButton>
                        </Tooltip>
                      </TableCell>
                    </TableRow>
                  ))}
                </TableBody>
              </Table>
            </TableContainer>
          </Card>
        </>
      )}

      {tab === 2 && (
        <Card>
          <TableContainer>
            <Table>
              <TableHead>
                <TableRow>
                  <TableCell>Username</TableCell>
                  <TableCell>Muted Until</TableCell>
                  <TableCell>Reason</TableCell>
                </TableRow>
              </TableHead>
              <TableBody>
                {mutedUsers.map((user) => (
                  <TableRow key={user.id} hover>
                    <TableCell><Box sx={{ display: 'flex', alignItems: 'center', gap: 1 }}><VolumeOff color="warning" fontSize="small" />{user.username}</Box></TableCell>
                    <TableCell>{new Date(user.mutedUntil).toLocaleString()}</TableCell>
                    <TableCell>{user.reason}</TableCell>
                  </TableRow>
                ))}
              </TableBody>
            </Table>
          </TableContainer>
        </Card>
      )}

      <Dialog open={deleteDialog.open} onClose={() => setDeleteDialog({ open: false, id: null })} maxWidth="sm" fullWidth>
        <DialogTitle>Delete Message</DialogTitle>
        <DialogContent>
          <TextField label="Reason" value={deleteReason} onChange={(e) => setDeleteReason(e.target.value)} fullWidth sx={{ mt: 1 }} required />
        </DialogContent>
        <DialogActions>
          <Button onClick={() => setDeleteDialog({ open: false, id: null })}>Cancel</Button>
          <Button variant="contained" color="error" onClick={() => deleteDialog.id && deleteMutation.mutate(deleteDialog.id)} disabled={!deleteReason || deleteMutation.isPending}>Delete</Button>
        </DialogActions>
      </Dialog>

      <Dialog open={banDialog.open} onClose={() => setBanDialog({ open: false, id: null })} maxWidth="sm" fullWidth>
        <DialogTitle>Ban User from Chat</DialogTitle>
        <DialogContent>
          <TextField label="Reason" value={banReason} onChange={(e) => setBanReason(e.target.value)} fullWidth sx={{ mt: 1 }} required />
        </DialogContent>
        <DialogActions>
          <Button onClick={() => setBanDialog({ open: false, id: null })}>Cancel</Button>
          <Button variant="contained" color="warning" onClick={() => banDialog.id && banMutation.mutate(banDialog.id)} disabled={!banReason || banMutation.isPending}>Ban</Button>
        </DialogActions>
      </Dialog>

      <Snackbar open={snackbar.open} autoHideDuration={4000} onClose={() => setSnackbar(s => ({ ...s, open: false }))}>
        <Alert severity={snackbar.severity} variant="filled">{snackbar.message}</Alert>
      </Snackbar>
    </Box>
  );
};

export default ChatModeration;
