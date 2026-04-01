import { useState } from 'react';
import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query';
import {
  Box, Card, CardContent, Typography, Grid, TextField, InputAdornment,
  Table, TableBody, TableCell, TableContainer, TableHead, TableRow,
  Button, Chip, IconButton, Tooltip, Dialog, DialogTitle, DialogContent,
  DialogActions, Tabs, Tab, LinearProgress, Divider, Avatar, MenuItem
} from '@mui/material';
import { Search, Refresh, Visibility, Edit, AccountTree, People, AttachMoney, TrendingUp } from '@mui/icons-material';
import { useAppDispatch } from '../../store/hooks';
import { showSnackbar } from '../../store/slices/uiSlice';
import { referralAPI } from '../../services/api';

interface ReferralCode { id: string; code: string; ownerId: string; ownerName: string; uses: number; totalEarnings: number; status: 'active' | 'disabled'; createdAt: string; }
interface ReferralReward { id: string; name: string; type: 'percentage' | 'fixed'; value: number; level: number; minReferrals: number; status: 'active' | 'inactive'; }
interface TreeNode { id: string; name: string; referrals: number; level: number; children?: TreeNode[]; }

const ReferralManagement = () => {
  const dispatch = useAppDispatch();
  const queryClient = useQueryClient();
  const [search, setSearch] = useState('');
  const [activeTab, setActiveTab] = useState(0);
  const [statusFilter, setStatusFilter] = useState('');
  const [treeOpen, setTreeOpen] = useState(false);
  const [selectedUser, setSelectedUser] = useState<ReferralCode | null>(null);
  const [rewardOpen, setRewardOpen] = useState(false);
  const [editingReward, setEditingReward] = useState<ReferralReward | null>(null);
  const [form, setForm] = useState({ name: '', type: 'percentage' as const, value: 0, level: 1, minReferrals: 0, status: 'active' as const });

  const { data: codesData, isLoading } = useQuery({
    queryKey: ['referralCodes', statusFilter],
    queryFn: () => referralAPI.getCodes({ status: statusFilter || undefined, page: 1, limit: 50 }),
  });
  const { data: statsData } = useQuery({ queryKey: ['referralStats'], queryFn: () => referralAPI.getStats() });
  const { data: rewardsData } = useQuery({ queryKey: ['referralRewards'], queryFn: () => referralAPI.getRewards() });

  const updateRewardMut = useMutation({
    mutationFn: ({ id, data }: { id: string; data: Record<string, any> }) => referralAPI.updateReward(id, data),
    onSuccess: () => { queryClient.invalidateQueries({ queryKey: ['referralRewards'] }); dispatch(showSnackbar({ message: 'Reward updated', severity: 'success' })); setRewardOpen(false); setEditingReward(null); },
    onError: (err: any) => dispatch(showSnackbar({ message: err.message || 'Failed', severity: 'error' })),
  });

  const mockCodes: ReferralCode[] = [
    { id: 'REF001', code: 'REFER-JOHN', ownerId: 'USR001', ownerName: 'John Doe', uses: 45, totalEarnings: 2250, status: 'active', createdAt: '2024-01-10' },
    { id: 'REF002', code: 'SARAH-VIP', ownerId: 'USR002', ownerName: 'Sarah Smith', uses: 120, totalEarnings: 8400, status: 'active', createdAt: '2024-01-05' },
    { id: 'REF003', code: 'MIKE-PRO', ownerId: 'USR003', ownerName: 'Mike Johnson', uses: 12, totalEarnings: 600, status: 'active', createdAt: '2024-02-01' },
    { id: 'REF004', code: 'EMILY-PLAY', ownerId: 'USR004', ownerName: 'Emily Brown', uses: 0, totalEarnings: 0, status: 'disabled', createdAt: '2024-01-20' },
    { id: 'REF005', code: 'DAVID-DEAL', ownerId: 'USR005', ownerName: 'David Wilson', uses: 67, totalEarnings: 3350, status: 'active', createdAt: '2024-01-15' },
  ];
  const mockStats = { totalCodes: 156, activeCodes: 134, totalReferrals: 4520, totalPayouts: 125800, conversionRate: 68.5 };
  const mockRewards: ReferralReward[] = [
    { id: 'RW001', name: 'Bronze Tier', type: 'percentage', value: 5, level: 1, minReferrals: 0, status: 'active' },
    { id: 'RW002', name: 'Silver Tier', type: 'percentage', value: 8, level: 2, minReferrals: 10, status: 'active' },
    { id: 'RW003', name: 'Gold Tier', type: 'percentage', value: 12, level: 3, minReferrals: 50, status: 'active' },
    { id: 'RW004', name: 'Platinum Tier', type: 'fixed', value: 25, level: 4, minReferrals: 100, status: 'active' },
  ];
  const mockTree: TreeNode = { id: 'USR001', name: 'John Doe', referrals: 45, level: 0, children: [
    { id: 'USR010', name: 'User A', referrals: 12, level: 1, children: [{ id: 'USR020', name: 'User C', referrals: 3, level: 2 }, { id: 'USR021', name: 'User D', referrals: 5, level: 2 }] },
    { id: 'USR011', name: 'User B', referrals: 8, level: 1, children: [{ id: 'USR022', name: 'User E', referrals: 2, level: 2 }] },
  ]};
  const mockCommissions = [
    { id: 'COM001', referrer: 'John Doe', referred: 'New User 1', amount: 50, level: 1, date: '2024-01-15', status: 'paid' },
    { id: 'COM002', referrer: 'Sarah Smith', referred: 'New User 2', amount: 120, level: 2, date: '2024-01-14', status: 'pending' },
    { id: 'COM003', referrer: 'Mike Johnson', referred: 'New User 3', amount: 25, level: 1, date: '2024-01-13', status: 'paid' },
    { id: 'COM004', referrer: 'David Wilson', referred: 'New User 4', amount: 75, level: 1, date: '2024-01-12', status: 'paid' },
  ];

  const codes: ReferralCode[] = codesData?.data?.codes || mockCodes;
  const stats = statsData?.data || mockStats;
  const rewards: ReferralReward[] = rewardsData?.data?.rewards || mockRewards;
  const filteredCodes = codes.filter((c) => !search || c.code.toLowerCase().includes(search.toLowerCase()) || c.ownerName.toLowerCase().includes(search.toLowerCase()));

  const handleEditReward = (r: ReferralReward) => { setEditingReward(r); setForm({ name: r.name, type: r.type, value: r.value, level: r.level, minReferrals: r.minReferrals, status: r.status }); setRewardOpen(true); };
  const renderTree = (node: TreeNode, depth = 0) => (
    <Box key={node.id} sx={{ ml: depth * 3, mb: 1 }}>
      <Box sx={{ display: 'flex', alignItems: 'center', gap: 1, py: 0.5 }}>
        <AccountTree fontSize="small" color={depth === 0 ? 'primary' : 'action'} />
        <Typography variant="body2" fontWeight={depth === 0 ? 600 : 400}>{node.name}</Typography>
        <Chip label={`${node.referrals} referrals`} size="small" variant="outlined" />
      </Box>
      {node.children?.map((c) => renderTree(c, depth + 1))}
    </Box>
  );

  return (
    <Box>
      <Box sx={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', mb: 3 }}>
        <Typography variant="h4" fontWeight="bold">Referral Management</Typography>
        <Button variant="outlined" startIcon={<Refresh />} onClick={() => queryClient.invalidateQueries({ queryKey: ['referralCodes'] })}>Refresh</Button>
      </Box>
      {isLoading && <LinearProgress sx={{ mb: 2 }} />}

      <Grid container spacing={3} sx={{ mb: 3 }}>
        {[{ label: 'Total Codes', val: stats.totalCodes, color: '#3b82f6' }, { label: 'Active Codes', val: stats.activeCodes, color: '#22c55e' },
          { label: 'Total Referrals', val: stats.totalReferrals.toLocaleString(), color: '#8b5cf6' }, { label: 'Total Payouts', val: `$${stats.totalPayouts.toLocaleString()}`, color: '#f59e0b' },
          { label: 'Conversion Rate', val: `${stats.conversionRate}%`, color: '#06b6d4' }].map((s) => (
          <Grid item xs={12} md={2.4} key={s.label}><Card sx={{ borderLeft: `4px solid ${s.color}` }}><CardContent><Typography color="text.secondary" variant="body2">{s.label}</Typography><Typography variant="h5" fontWeight="bold">{s.val}</Typography></CardContent></Card></Grid>
        ))}
      </Grid>

      <Card sx={{ mb: 3 }}>
        <Box sx={{ borderBottom: 1, borderColor: 'divider' }}>
          <Tabs value={activeTab} onChange={(_, v) => setActiveTab(v)}>
            <Tab label="Referral Codes" icon={<People />} iconPosition="start" />
            <Tab label="Rewards Config" icon={<AttachMoney />} iconPosition="start" />
            <Tab label="Commission Payouts" icon={<TrendingUp />} iconPosition="start" />
          </Tabs>
        </Box>
        <CardContent>
          {activeTab === 0 && (<>
            <Grid container spacing={2} sx={{ mb: 2 }}>
              <Grid item xs={12} md={6}><TextField fullWidth placeholder="Search by code or owner..." value={search} onChange={(e) => setSearch(e.target.value)}
                InputProps={{ startAdornment: <InputAdornment position="start"><Search /></InputAdornment> }} size="small" /></Grid>
              <Grid item xs={12} md={3}><TextField fullWidth select label="Status" value={statusFilter} onChange={(e) => setStatusFilter(e.target.value)} size="small">
                <MenuItem value="">All</MenuItem><MenuItem value="active">Active</MenuItem><MenuItem value="disabled">Disabled</MenuItem>
              </TextField></Grid>
            </Grid>
            <TableContainer><Table size="small">
              <TableHead><TableRow><TableCell>Code</TableCell><TableCell>Owner</TableCell><TableCell>Uses</TableCell><TableCell>Earnings</TableCell><TableCell>Status</TableCell><TableCell>Created</TableCell><TableCell align="center">Actions</TableCell></TableRow></TableHead>
              <TableBody>{filteredCodes.map((code) => (
                <TableRow key={code.id} hover>
                  <TableCell><Typography fontFamily="monospace" fontWeight={600}>{code.code}</Typography></TableCell>
                  <TableCell><Box sx={{ display: 'flex', alignItems: 'center', gap: 1 }}><Avatar sx={{ width: 24, height: 24, fontSize: 12 }}>{code.ownerName[0]}</Avatar>{code.ownerName}</Box></TableCell>
                  <TableCell>{code.uses}</TableCell><TableCell>${code.totalEarnings.toLocaleString()}</TableCell>
                  <TableCell><Chip label={code.status} size="small" color={code.status === 'active' ? 'success' : 'default'} /></TableCell>
                  <TableCell>{code.createdAt}</TableCell>
                  <TableCell align="center"><Tooltip title="View Tree"><IconButton size="small" onClick={() => { setSelectedUser(code); setTreeOpen(true); }}><Visibility fontSize="small" /></IconButton></Tooltip></TableCell>
                </TableRow>
              ))}</TableBody>
            </Table></TableContainer>
          </>)}
          {activeTab === 1 && (<TableContainer><Table size="small">
            <TableHead><TableRow><TableCell>Tier Name</TableCell><TableCell>Type</TableCell><TableCell>Value</TableCell><TableCell>Level</TableCell><TableCell>Min Referrals</TableCell><TableCell>Status</TableCell><TableCell align="center">Actions</TableCell></TableRow></TableHead>
            <TableBody>{rewards.map((r) => (<TableRow key={r.id} hover>
              <TableCell fontWeight={500}>{r.name}</TableCell><TableCell><Chip label={r.type} size="small" /></TableCell>
              <TableCell>{r.type === 'percentage' ? `${r.value}%` : `$${r.value}`}</TableCell><TableCell>Level {r.level}</TableCell>
              <TableCell>{r.minReferrals}</TableCell><TableCell><Chip label={r.status} size="small" color={r.status === 'active' ? 'success' : 'default'} /></TableCell>
              <TableCell align="center"><Tooltip title="Edit"><IconButton size="small" onClick={() => handleEditReward(r)}><Edit fontSize="small" /></IconButton></Tooltip></TableCell>
            </TableRow>))}</TableBody>
          </Table></TableContainer>)}
          {activeTab === 2 && (<TableContainer><Table size="small">
            <TableHead><TableRow><TableCell>ID</TableCell><TableCell>Referrer</TableCell><TableCell>Referred</TableCell><TableCell>Amount</TableCell><TableCell>Level</TableCell><TableCell>Date</TableCell><TableCell>Status</TableCell></TableRow></TableHead>
            <TableBody>{mockCommissions.map((c) => (<TableRow key={c.id} hover>
              <TableCell>{c.id}</TableCell><TableCell>{c.referrer}</TableCell><TableCell>{c.referred}</TableCell>
              <TableCell>${c.amount}</TableCell><TableCell>Level {c.level}</TableCell><TableCell>{c.date}</TableCell>
              <TableCell><Chip label={c.status} size="small" color={c.status === 'paid' ? 'success' : 'warning'} /></TableCell>
            </TableRow>))}</TableBody>
          </Table></TableContainer>)}
        </CardContent>
      </Card>

      <Dialog open={treeOpen} onClose={() => setTreeOpen(false)} maxWidth="sm" fullWidth>
        <DialogTitle>Referral Tree - {selectedUser?.ownerName}</DialogTitle>
        <DialogContent><Divider sx={{ mb: 2 }} />{renderTree(mockTree)}</DialogContent>
        <DialogActions><Button onClick={() => setTreeOpen(false)}>Close</Button></DialogActions>
      </Dialog>

      <Dialog open={rewardOpen} onClose={() => { setRewardOpen(false); setEditingReward(null); }} maxWidth="sm" fullWidth>
        <DialogTitle>Edit Reward Tier</DialogTitle>
        <DialogContent>
          <Grid container spacing={2} sx={{ mt: 1 }}>
            <Grid item xs={12}><TextField fullWidth label="Tier Name" value={form.name} onChange={(e) => setForm({ ...form, name: e.target.value })} /></Grid>
            <Grid item xs={6}><TextField fullWidth select label="Type" value={form.type} onChange={(e) => setForm({ ...form, type: e.target.value as any })}>
              <MenuItem value="percentage">Percentage</MenuItem><MenuItem value="fixed">Fixed</MenuItem></TextField></Grid>
            <Grid item xs={6}><TextField fullWidth label="Value" type="number" value={form.value} onChange={(e) => setForm({ ...form, value: Number(e.target.value) })} /></Grid>
            <Grid item xs={6}><TextField fullWidth label="Level" type="number" value={form.level} onChange={(e) => setForm({ ...form, level: Number(e.target.value) })} /></Grid>
            <Grid item xs={6}><TextField fullWidth label="Min Referrals" type="number" value={form.minReferrals} onChange={(e) => setForm({ ...form, minReferrals: Number(e.target.value) })} /></Grid>
            <Grid item xs={12}><TextField fullWidth select label="Status" value={form.status} onChange={(e) => setForm({ ...form, status: e.target.value as any })}>
              <MenuItem value="active">Active</MenuItem><MenuItem value="inactive">Inactive</MenuItem></TextField></Grid>
          </Grid>
        </DialogContent>
        <DialogActions>
          <Button onClick={() => { setRewardOpen(false); setEditingReward(null); }}>Cancel</Button>
          <Button variant="contained" onClick={() => editingReward && updateRewardMut.mutate({ id: editingReward.id, data: form })} disabled={updateRewardMut.isPending}>Save</Button>
        </DialogActions>
      </Dialog>
    </Box>
  );
};

export default ReferralManagement;
