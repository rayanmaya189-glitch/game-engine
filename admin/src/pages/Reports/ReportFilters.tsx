import {
  Box, Card, CardContent, Grid, TextField, Button,
  FormControl, InputLabel, Select, MenuItem, Tabs, Tab
} from '@mui/material';
import { Download } from '@mui/icons-material';

interface ReportFiltersProps {
  activeTab: number;
  dateRange: { from: string; to: string };
  period: string;
  category: string;
  onTabChange: (tab: number) => void;
  onDateRangeChange: (range: { from: string; to: string }) => void;
  onPeriodChange: (period: string) => void;
  onCategoryChange: (category: string) => void;
  onExport: () => void;
}

const ReportFilters = ({
  activeTab,
  dateRange,
  period,
  category,
  onTabChange,
  onDateRangeChange,
  onPeriodChange,
  onCategoryChange,
  onExport,
}: ReportFiltersProps) => {
  return (
    <Card sx={{ mb: 3 }}>
      <Box sx={{ borderBottom: 1, borderColor: 'divider' }}>
        <Tabs value={activeTab} onChange={(_, v) => onTabChange(v)}>
          <Tab label="Revenue" />
          <Tab label="Users" />
          <Tab label="Games" />
        </Tabs>
      </Box>
      <CardContent>
        <Grid container spacing={2} alignItems="center">
          <Grid item xs={12} md={3}>
            <TextField
              type="date"
              label="From"
              fullWidth
              size="small"
              value={dateRange.from}
              onChange={(e) => onDateRangeChange({ ...dateRange, from: e.target.value })}
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
              onChange={(e) => onDateRangeChange({ ...dateRange, to: e.target.value })}
              InputLabelProps={{ shrink: true }}
            />
          </Grid>
          {activeTab === 1 && (
            <Grid item xs={12} md={3}>
              <FormControl fullWidth size="small">
                <InputLabel>Period</InputLabel>
                <Select value={period} label="Period" onChange={(e) => onPeriodChange(e.target.value)}>
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
                <Select value={category} label="Category" onChange={(e) => onCategoryChange(e.target.value)}>
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
      </CardContent>
    </Card>
  );
};

export default ReportFilters;
