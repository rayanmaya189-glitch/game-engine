import {
  Box, Button, Card, CardContent, Typography, Grid, TextField, Switch,
  FormControlLabel, Avatar, IconButton
} from '@mui/material';
import { PhotoCamera } from '@mui/icons-material';

interface GeneralSettingsProps {
  generalSettings: {
    siteName: string;
    supportEmail: string;
    supportPhone: string;
    timezone: string;
    maintenanceMode: boolean;
  };
  onGeneralChange: (settings: GeneralSettingsProps['generalSettings']) => void;
}

const GeneralSettings = ({ generalSettings, onGeneralChange }: GeneralSettingsProps) => {
  return (
    <Grid container spacing={3}>
      <Grid item xs={12} md={6}>
        <Card>
          <CardContent>
            <Typography variant="h6" gutterBottom>Site Information</Typography>
            <TextField
              fullWidth
              label="Site Name"
              value={generalSettings.siteName}
              onChange={(e) => onGeneralChange({ ...generalSettings, siteName: e.target.value })}
              sx={{ mb: 2 }}
            />
            <TextField
              fullWidth
              label="Support Email"
              value={generalSettings.supportEmail}
              onChange={(e) => onGeneralChange({ ...generalSettings, supportEmail: e.target.value })}
              sx={{ mb: 2 }}
            />
            <TextField
              fullWidth
              label="Support Phone"
              value={generalSettings.supportPhone}
              onChange={(e) => onGeneralChange({ ...generalSettings, supportPhone: e.target.value })}
              sx={{ mb: 2 }}
            />
            <TextField
              fullWidth
              label="Timezone"
              value={generalSettings.timezone}
              onChange={(e) => onGeneralChange({ ...generalSettings, timezone: e.target.value })}
              sx={{ mb: 2 }}
            />
            <FormControlLabel
              control={
                <Switch
                  checked={generalSettings.maintenanceMode}
                  onChange={(e) => onGeneralChange({ ...generalSettings, maintenanceMode: e.target.checked })}
                />
              }
              label="Maintenance Mode"
            />
          </CardContent>
        </Card>
      </Grid>
      <Grid item xs={12} md={6}>
        <Card>
          <CardContent>
            <Typography variant="h6" gutterBottom>Logo</Typography>
            <Box sx={{ display: 'flex', alignItems: 'center', gap: 2, mb: 2 }}>
              <Avatar sx={{ width: 80, height: 80, bgcolor: 'primary.main', fontSize: 32 }}>
                {generalSettings.siteName.charAt(0)}
              </Avatar>
              <IconButton><PhotoCamera /></IconButton>
            </Box>
            <Button variant="outlined">Upload Logo</Button>
            <Typography variant="caption" display="block" color="text.secondary" sx={{ mt: 1 }}>
              Recommended size: 512x512px, PNG or JPG
            </Typography>
          </CardContent>
        </Card>
      </Grid>
    </Grid>
  );
};

export default GeneralSettings;
