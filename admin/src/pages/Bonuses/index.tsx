import { useState } from 'react';
import { useQuery } from '@tanstack/react-query';
import {
  Box, Card, CardContent, Typography, Grid, TextField, InputAdornment,
  Table, TableBody, TableCell, TableContainer, TableHead, TableRow,
  Button, Chip, IconButton, Tooltip
} from '@mui/material';
import { Search, Add, Edit, Visibility, CardGiftcard } from '@mui/icons-material';
import { bonusesAPI } from '../../services/api';

interface Bonus {
  id: string;
  name: string;
  type: string;
  amount: number;
  maxBonus: number;
  minDeposit: number;
  wagerReq: number;
  uses: number;
  status: boolean;
}

const Bonuses = () => {
  const [search, setSearch] = useState('');

  const { data } = useQuery({
    queryKey: ['bonuses', search],
    queryFn: () => bonusesAPI.getAll({ search, page: 1, limit: 20 }),
  });

  const bonuses = data?.data?.bonuses || [];

  return (
    <Box>
      <Box sx={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', mb: 3 }}>
        <Typography variant="h4" fontWeight="bold">Bonuses & Promotions</Typography>
        <Button variant="contained" startIcon={<Add />}>Create Bonus</Button>
      </Box>

      <Grid container spacing={3} sx={{ mb: 3 }}>
        <Grid item xs={12} md={3}>
          <Card><CardContent><Typography color="text.secondary">Active Bonuses</Typography><Typography variant="h4">8</Typography></CardContent></Card>
        </Grid>
        <Grid item xs={12} md={3}>
          <Card><CardContent><Typography color="text.secondary">Total Bonus Given</Typography><Typography variant="h4">$125,000</Typography></CardContent></Card>
        </Grid>
        <Grid item xs={12} md={3}>
          <Card><CardContent><Typography color="text.secondary">Total Uses</Typography><Typography variant="h4">4,696</Typography></CardContent></Card>
        </Grid>
        <Grid item xs={12} md={3}>
          <Card><CardContent><Typography color="text.secondary">Pending Claims</Typography><Typography variant="h4">34</Typography></CardContent></Card>
        </Grid>
      </Grid>

      <Card sx={{ mb: 3 }}>
        <CardContent>
          <TextField
            fullWidth
            placeholder="Search bonuses..."
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
                <TableCell>Bonus Name</TableCell>
                <TableCell>Type</TableCell>
                <TableCell>Amount</TableCell>
                <TableCell>Max Bonus</TableCell>
                <TableCell>Min Deposit</TableCell>
                <TableCell>Wager Req</TableCell>
                <TableCell>Uses</TableCell>
                <TableCell>Active</TableCell>
                <TableCell align="right">Actions</TableCell>
              </TableRow>
            </TableHead>
            <TableBody>
              {bonuses.map((bonus: Bonus) => (
                <TableRow key={bonus.id} hover>
                  <TableCell><Box sx={{ display: 'flex', alignItems: 'center', gap: 1 }}><CardGiftcard color="primary" />{bonus.name}</Box></TableCell>
                  <TableCell><Chip label={bonus.type} size="small" /></TableCell>
                  <TableCell>{bonus.amount}%</TableCell>
                  <TableCell>${bonus.maxBonus}</TableCell>
                  <TableCell>${bonus.minDeposit}</TableCell>
                  <TableCell>{bonus.wagerReq}x</TableCell>
                  <TableCell>{bonus.uses}</TableCell>
                  <TableCell><Chip label={bonus.status ? 'Active' : 'Inactive'} color={bonus.status ? 'success' : 'default'} size="small" /></TableCell>
                  <TableCell align="right">
                    <Tooltip title="View"><IconButton size="small"><Visibility /></IconButton></Tooltip>
                    <Tooltip title="Edit"><IconButton size="small"><Edit /></IconButton></Tooltip>
                  </TableCell>
                </TableRow>
              ))}
            </TableBody>
          </Table>
        </TableContainer>
      </Card>
    </Box>
  );
};

export default Bonuses;
