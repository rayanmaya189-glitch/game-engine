import {
  Box, Card, CardContent, Typography, Grid, TextField, Switch,
  FormControlLabel, Divider, IconButton, Button
} from '@mui/material';
import { ContentCopy, Refresh } from '@mui/icons-material';

interface SecuritySettingsProps {
  securitySettings: {
    twoFactorAuth: boolean;
    ipWhitelist: boolean;
    sessionTimeout: boolean;
    loginNotifications: boolean;
    adminPassword: string;
    confirmPassword: string;
  };
  apiSettings: {
    apiKey: string;
    webhookUrl: string;
    enableApiAccess: boolean;
    rateLimiting: boolean;
  };
  onSecurityChange: (settings: SecuritySettingsProps['securitySettings']) => void;
  onApiChange: (settings: SecuritySettingsProps['apiSettings']) => void;
  onCopyApiKey: () => void;
  onGenerateNewKey: () => void;
}

const SecuritySettings = ({
  securitySettings,
  apiSettings,
  onSecurityChange,
  onApiChange,
  onCopyApiKey,
  onGenerateNewKey,
}: SecuritySettingsProps) => {
  return (
    <Grid container spacing={3}>
      <Grid item xs={12} md={6}>
        <Card>
          <CardContent>
            <Typography variant="h6" gutterBottom>Security Settings</Typography>
            <FormControlLabel
              control={
                <Switch
                  checked={securitySettings.twoFactorAuth}
                  onChange={(e) => onSecurityChange({ ...securitySettings, twoFactorAuth: e.target.checked })}
                />
              }
              label="Two-Factor Authentication"
              sx={{ display: 'block', mb: 1 }}
            />
            <FormControlLabel
              control={
                <Switch
                  checked={securitySettings.ipWhitelist}
                  onChange={(e) => onSecurityChange({ ...securitySettings, ipWhitelist: e.target.checked })}
                />
              }
              label="IP Whitelist"
              sx={{ display: 'block', mb: 1 }}
            />
            <FormControlLabel
              control={
                <Switch
                  checked={securitySettings.sessionTimeout}
                  onChange={(e) => onSecurityChange({ ...securitySettings, sessionTimeout: e.target.checked })}
                />
              }
              label="Session Timeout (30 min)"
              sx={{ display: 'block', mb: 1 }}
            />
            <FormControlLabel
              control={
                <Switch
                  checked={securitySettings.loginNotifications}
                  onChange={(e) => onSecurityChange({ ...securitySettings, loginNotifications: e.target.checked })}
                />
              }
              label="Login Notifications"
              sx={{ display: 'block', mb: 1 }}
            />
            <Divider sx={{ my: 2 }} />
            <Typography variant="subtitle2" gutterBottom>Change Admin Password</Typography>
            <TextField
              fullWidth
              type="password"
              label="Admin Password"
              value={securitySettings.adminPassword}
              onChange={(e) => onSecurityChange({ ...securitySettings, adminPassword: e.target.value })}
              sx={{ mb: 2 }}
            />
            <TextField
              fullWidth
              type="password"
              label="Confirm Password"
              value={securitySettings.confirmPassword}
              onChange={(e) => onSecurityChange({ ...securitySettings, confirmPassword: e.target.value })}
              sx={{ mb: 2 }}
            />
          </CardContent>
        </Card>
      </Grid>
      <Grid item xs={12} md={6}>
        <Card>
          <CardContent>
            <Typography variant="h6" gutterBottom>API Security</Typography>
            <Box sx={{ display: 'flex', alignItems: 'center', gap: 1, mb: 2 }}>
              <TextField
                fullWidth
                label="API Key"
                value={apiSettings.apiKey}
                disabled
                size="small"
              />
              <IconButton onClick={onCopyApiKey} title="Copy API Key">
                <ContentCopy />
              </IconButton>
            </Box>
            <TextField
              fullWidth
              label="Webhook URL"
              placeholder="https://..."
              value={apiSettings.webhookUrl}
              onChange={(e) => onApiChange({ ...apiSettings, webhookUrl: e.target.value })}
              sx={{ mb: 2 }}
            />
            <FormControlLabel
              control={
                <Switch
                  checked={apiSettings.enableApiAccess}
                  onChange={(e) => onApiChange({ ...apiSettings, enableApiAccess: e.target.checked })}
                />
              }
              label="Enable API Access"
              sx={{ display: 'block', mb: 1 }}
            />
            <FormControlLabel
              control={
                <Switch
                  checked={apiSettings.rateLimiting}
                  onChange={(e) => onApiChange({ ...apiSettings, rateLimiting: e.target.checked })}
                />
              }
              label="Rate Limiting"
              sx={{ display: 'block', mb: 1 }}
            />
            <Button
              variant="outlined"
              color="warning"
              startIcon={<Refresh />}
              onClick={onGenerateNewKey}
              sx={{ mt: 1 }}
            >
              Generate New Key
            </Button>
          </CardContent>
        </Card>
      </Grid>
    </Grid>
  );
};

export default SecuritySettings;
