import { Grid, TextField, Button, FormControl, InputLabel, Select, MenuItem } from '@mui/material';
import { Download } from '@mui/icons-material';

interface ReportFiltersProps {
  activeTab: number;
  dateRange: { from: string; to: string };
  setDateRange: (range: { from: string; to: string }) => void;
  period: string;
  setPeriod: (period: string) => void;
  category: string;
  setCategory: (category: string) => void;
  onExport: () => void;
}

const ReportFilters = ({
  activeTab,
  dateRange,
  setDateRange,
  period,
  setPeriod,
  category,
  setCategory,
  onExport,
}: ReportFiltersProps) => (
  <Grid container spacing={2} alignItems="center">
    <Grid item xs={12} md={3}>
      <TextField
        type="date"
        label="From"
        fullWidth
        size="small"
        value={dateRange.from}
        onChange={(e) => setDateRange({ ...dateRange, from: e.target.value })}
        InputLabelProps={{ shrink: true }}
      />
    </Grid>
    <Grid item xs={12} md={3}>
      <TextField
        type="date"
        label="To"
        fullWidth
        size="small"
        value={dateRange.to}
        onChange={(e) => setDateRange({ ...dateRange, to: e.target.value })}
        InputLabelProps={{ shrink: true }}
      />
    </Grid>
    {activeTab === 1 && (
      <Grid item xs={12} md={3}>
        <FormControl fullWidth size="small">
          <InputLabel>Period</InputLabel>
          <Select value={period} label="Period" onChange={(e) => setPeriod(e.target.value)}>
            <MenuItem value="7">Last 7 days</MenuItem>
            <MenuItem value="30">Last 30 days</MenuItem>
            <MenuItem value="90">Last 90 days</MenuItem>
            <MenuItem value="365">Last year</MenuItem>
          </Select>
        </FormControl>
      </Grid>
    )}
    {activeTab === 2 && (
      <Grid item xs={12} md={3}>
        <FormControl fullWidth size="small">
          <InputLabel>Category</InputLabel>
          <Select value={category} label="Category" onChange={(e) => setCategory(e.target.value)}>
            <MenuItem value="">All Games</MenuItem>
            <MenuItem value="slot">Slots</MenuItem>
            <MenuItem value="live_casino">Live Casino</MenuItem>
            <MenuItem value="table_games">Table Games</MenuItem>
            <MenuItem value="dice">Dice</MenuItem>
          </Select>
        </FormControl>
      </Grid>
    )}
    <Grid item xs={12} md={activeTab === 0 ? 6 : 3}>
      <Button variant="contained" startIcon={<Download />} onClick={onExport}>
        Export
      </Button>
    </Grid>
  </Grid>
);

export default ReportFilters;
