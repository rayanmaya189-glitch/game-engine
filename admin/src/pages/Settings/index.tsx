import { useState } from 'react';
import {
  Box, Card, CardContent, Typography, Grid, TextField, Button, Switch, 
  FormControlLabel, Divider, Avatar, IconButton, Tabs, Tab, Alert, CircularProgress
} from '@mui/material';
import { Save, PhotoCamera, ContentCopy, Refresh } from '@mui/icons-material';
import { useAppDispatch } from '../../store/hooks';
import { showSnackbar } from '../../store/slices/uiSlice';

const Settings = () => {
  const dispatch = useAppDispatch();
  const [activeTab, setActiveTab] = useState(0);
  const [saving, setSaving] = useState(false);
  const [saved, setSaved] = useState(false);

  // General Settings State
  const [generalSettings, setGeneralSettings] = useState({
    siteName: 'Casino Admin',
    supportEmail: 'support@casino.com',
    supportPhone: '+1 234 567 8900',
    timezone: 'UTC',
    maintenanceMode: false,
  });

  // Security Settings State
  const [securitySettings, setSecuritySettings] = useState({
    twoFactorAuth: true,
    ipWhitelist: true,
    sessionTimeout: false,
    loginNotifications: true,
    adminPassword: '',
    confirmPassword: '',
  });

  // API Security State
  const [apiSettings, setApiSettings] = useState({
    apiKey: 'sk_live_xxxxxxxxxxxxx',
    webhookUrl: '',
    enableApiAccess: true,
    rateLimiting: false,
  });

  // Notification Settings State
  const [notificationSettings, setNotificationSettings] = useState({
    emailNotifications: true,
    smsNotifications: true,
    pushNotifications: true,
    newUserRegistrations: true,
    largeTransactions: true,
    kycApprovals: true,
    systemAlerts: true,
  });

  // System Settings State
  const [systemSettings, setSystemSettings] = useState({
    dbHost: 'localhost',
    dbPort: '5432',
    dbName: 'casino_db',
    redisHost: 'localhost',
    redisPort: '6379',
    natsServer: 'nats://localhost:4222',
  });

  const handleSave = async () => {
    setSaving(true);
    setSaved(false);
    
    // Simulate API call
    await new Promise(resolve => setTimeout(resolve, 1500));
    
    setSaving(false);
    setSaved(true);
    dispatch(showSnackbar({ message: 'Settings saved successfully', severity: 'success' }));
    
    setTimeout(() => setSaved(false), 3000);
  };

  const handleCopyApiKey = () => {
    navigator.clipboard.writeText(apiSettings.apiKey);
    dispatch(showSnackbar({ message: 'API Key copied to clipboard', severity: 'success' }));
  };

  const handleGenerateNewKey = () => {
    const newKey = 'sk_live_' + Math.random().toString(36).substring(2, 15);
    setApiSettings({ ...apiSettings, apiKey: newKey });
    dispatch(showSnackbar({ message: 'New API key generated', severity: 'success' }));
  };

  return (
    <Box>
      <Typography variant="h4" fontWeight="bold" gutterBottom>Settings</Typography>

      {saved && (
        <Alert severity="success" sx={{ mb: 2 }}>
          All settings have been saved successfully!
        </Alert>
      )}

      <Box sx={{ borderBottom: 1, borderColor: 'divider', mb: 3 }}>
        <Tabs value={activeTab} onChange={(_, v) => setActiveTab(v)}>
          <Tab label="General" />
          <Tab label="Security" />
          <Tab label="Notifications" />
          <Tab label="API Keys" />
          <Tab label="System" />
        </Tabs>
      </Box>

      {activeTab === 0 && (
        <Grid container spacing={3}>
          <Grid item xs={12} md={6}>
            <Card>
              <CardContent>
                <Typography variant="h6" gutterBottom>Site Information</Typography>
                <TextField 
                  fullWidth 
                  label="Site Name" 
                  value={generalSettings.siteName}
                  onChange={(e) => setGeneralSettings({ ...generalSettings, siteName: e.target.value })}
                  sx={{ mb: 2 }} 
                />
                <TextField 
                  fullWidth 
                  label="Support Email" 
                  value={generalSettings.supportEmail}
                  onChange={(e) => setGeneralSettings({ ...generalSettings, supportEmail: e.target.value })}
                  sx={{ mb: 2 }} 
                />
                <TextField 
                  fullWidth 
                  label="Support Phone" 
                  value={generalSettings.supportPhone}
                  onChange={(e) => setGeneralSettings({ ...generalSettings, supportPhone: e.target.value })}
                  sx={{ mb: 2 }} 
                />
                <TextField 
                  fullWidth 
                  label="Timezone" 
                  value={generalSettings.timezone}
                  onChange={(e) => setGeneralSettings({ ...generalSettings, timezone: e.target.value })}
                  sx={{ mb: 2 }} 
                />
                <FormControlLabel 
                  control={
                    <Switch 
                      checked={generalSettings.maintenanceMode}
                      onChange={(e) => setGeneralSettings({ ...generalSettings, maintenanceMode: e.target.checked })}
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
      )}

      {activeTab === 1 && (
        <Grid container spacing={3}>
          <Grid item xs={12} md={6}>
            <Card>
              <CardContent>
                <Typography variant="h6" gutterBottom>Security Settings</Typography>
                <FormControlLabel 
                  control={
                    <Switch 
                      checked={securitySettings.twoFactorAuth}
                      onChange={(e) => setSecuritySettings({ ...securitySettings, twoFactorAuth: e.target.checked })}
                    />
                  } 
                  label="Two-Factor Authentication" 
                  sx={{ display: 'block', mb: 1 }} 
                />
                <FormControlLabel 
                  control={
                    <Switch 
                      checked={securitySettings.ipWhitelist}
                      onChange={(e) => setSecuritySettings({ ...securitySettings, ipWhitelist: e.target.checked })}
                    />
                  } 
                  label="IP Whitelist" 
                  sx={{ display: 'block', mb: 1 }} 
                />
                <FormControlLabel 
                  control={
                    <Switch 
                      checked={securitySettings.sessionTimeout}
                      onChange={(e) => setSecuritySettings({ ...securitySettings, sessionTimeout: e.target.checked })}
                    />
                  } 
                  label="Session Timeout (30 min)" 
                  sx={{ display: 'block', mb: 1 }} 
                />
                <FormControlLabel 
                  control={
                    <Switch 
                      checked={securitySettings.loginNotifications}
                      onChange={(e) => setSecuritySettings({ ...securitySettings, loginNotifications: e.target.checked })}
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
                  onChange={(e) => setSecuritySettings({ ...securitySettings, adminPassword: e.target.value })}
                  sx={{ mb: 2 }} 
                />
                <TextField 
                  fullWidth 
                  type="password" 
                  label="Confirm Password" 
                  value={securitySettings.confirmPassword}
                  onChange={(e) => setSecuritySettings({ ...securitySettings, confirmPassword: e.target.value })}
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
                  <IconButton onClick={handleCopyApiKey} title="Copy API Key">
                    <ContentCopy />
                  </IconButton>
                </Box>
                <TextField 
                  fullWidth 
                  label="Webhook URL" 
                  placeholder="https://..."
                  value={apiSettings.webhookUrl}
                  onChange={(e) => setApiSettings({ ...apiSettings, webhookUrl: e.target.value })}
                  sx={{ mb: 2 }} 
                />
                <FormControlLabel 
                  control={
                    <Switch 
                      checked={apiSettings.enableApiAccess}
                      onChange={(e) => setApiSettings({ ...apiSettings, enableApiAccess: e.target.checked })}
                    />
                  } 
                  label="Enable API Access" 
                  sx={{ display: 'block', mb: 1 }} 
                />
                <FormControlLabel 
                  control={
                    <Switch 
                      checked={apiSettings.rateLimiting}
                      onChange={(e) => setApiSettings({ ...apiSettings, rateLimiting: e.target.checked })}
                    />
                  } 
                  label="Rate Limiting" 
                  sx={{ display: 'block', mb: 1 }} 
                />
                <Button 
                  variant="outlined" 
                  color="warning" 
                  startIcon={<Refresh />}
                  onClick={handleGenerateNewKey}
                  sx={{ mt: 1 }}
                >
                  Generate New Key
                </Button>
              </CardContent>
            </Card>
          </Grid>
        </Grid>
      )}

      {activeTab === 2 && (
        <Card>
          <CardContent>
            <Typography variant="h6" gutterBottom>Notification Preferences</Typography>
            <Grid container spacing={2}>
              <Grid item xs={12} md={6}>
                <Typography variant="subtitle2" gutterBottom>Delivery Methods</Typography>
                <FormControlLabel 
                  control={
                    <Switch 
                      checked={notificationSettings.emailNotifications}
                      onChange={(e) => setNotificationSettings({ ...notificationSettings, emailNotifications: e.target.checked })}
                    />
                  } 
                  label="Email Notifications" 
                  sx={{ display: 'block', mb: 1 }} 
                />
                <FormControlLabel 
                  control={
                    <Switch 
                      checked={notificationSettings.smsNotifications}
                      onChange={(e) => setNotificationSettings({ ...notificationSettings, smsNotifications: e.target.checked })}
                    />
                  } 
                  label="SMS Notifications" 
                  sx={{ display: 'block', mb: 1 }} 
                />
                <FormControlLabel 
                  control={
                    <Switch 
                      checked={notificationSettings.pushNotifications}
                      onChange={(e) => setNotificationSettings({ ...notificationSettings, pushNotifications: e.target.checked })}
                    />
                  } 
                  label="Push Notifications" 
                  sx={{ display: 'block', mb: 1 }} 
                />
              </Grid>
              <Grid item xs={12} md={6}>
                <Typography variant="subtitle2" gutterBottom>Notify me about:</Typography>
                <FormControlLabel 
                  control={
                    <Switch 
                      checked={notificationSettings.newUserRegistrations}
                      onChange={(e) => setNotificationSettings({ ...notificationSettings, newUserRegistrations: e.target.checked })}
                    />
                  } 
                  label="New user registrations" 
                  sx={{ display: 'block', mb: 1 }} 
                />
                <FormControlLabel 
                  control={
                    <Switch 
                      checked={notificationSettings.largeTransactions}
                      onChange={(e) => setNotificationSettings({ ...notificationSettings, largeTransactions: e.target.checked })}
                    />
                  } 
                  label="Large transactions" 
                  sx={{ display: 'block', mb: 1 }} 
                />
                <FormControlLabel 
                  control={
                    <Switch 
                      checked={notificationSettings.kycApprovals}
                      onChange={(e) => setNotificationSettings({ ...notificationSettings, kycApprovals: e.target.checked })}
                    />
                  } 
                  label="KYC approvals" 
                  sx={{ display: 'block', mb: 1 }} 
                />
                <FormControlLabel 
                  control={
                    <Switch 
                      checked={notificationSettings.systemAlerts}
                      onChange={(e) => setNotificationSettings({ ...notificationSettings, systemAlerts: e.target.checked })}
                    />
                  } 
                  label="System alerts" 
                  sx={{ display: 'block', mb: 1 }} 
                />
              </Grid>
            </Grid>
          </CardContent>
        </Card>
      )}

      {activeTab === 3 && (
        <Card>
          <CardContent>
            <Typography variant="h6" gutterBottom>API Keys Management</Typography>
            <Grid container spacing={2}>
              <Grid item xs={12}>
                <Box sx={{ p: 2, bgcolor: 'background.paper', borderRadius: 1, display: 'flex', justifyContent: 'space-between', alignItems: 'center', border: '1px solid', borderColor: 'divider' }}>
                  <Box>
                    <Typography fontWeight="500">Production Key</Typography>
                    <Typography variant="body2" color="text.secondary" sx={{ fontFamily: 'monospace' }}>
                      {apiSettings.apiKey}
                    </Typography>
                  </Box>
                  <Box>
                    <Button variant="outlined" color="error" size="small">Revoke</Button>
                  </Box>
                </Box>
              </Grid>
              <Grid item xs={12}>
                <Box sx={{ p: 2, bgcolor: 'background.paper', borderRadius: 1, display: 'flex', justifyContent: 'space-between', alignItems: 'center', border: '1px solid', borderColor: 'divider' }}>
                  <Box>
                    <Typography fontWeight="500">Test Key</Typography>
                    <Typography variant="body2" color="text.secondary" sx={{ fontFamily: 'monospace' }}>
                      sk_test_xxxxxxxxxxxxx
                    </Typography>
                  </Box>
                  <Box>
                    <Button variant="outlined" color="error" size="small">Revoke</Button>
                  </Box>
                </Box>
              </Grid>
            </Grid>
            <Button variant="contained" startIcon={<Refresh />} onClick={handleGenerateNewKey} sx={{ mt: 2 }}>
              Generate New Key
            </Button>
          </CardContent>
        </Card>
      )}

      {activeTab === 4 && (
        <Grid container spacing={3}>
          <Grid item xs={12} md={6}>
            <Card>
              <CardContent>
                <Typography variant="h6" gutterBottom>Database</Typography>
                <TextField 
                  fullWidth 
                  label="Host" 
                  value={systemSettings.dbHost}
                  onChange={(e) => setSystemSettings({ ...systemSettings, dbHost: e.target.value })}
                  sx={{ mb: 2 }} 
                />
                <TextField 
                  fullWidth 
                  label="Port" 
                  value={systemSettings.dbPort}
                  onChange={(e) => setSystemSettings({ ...systemSettings, dbPort: e.target.value })}
                  sx={{ mb: 2 }} 
                />
                <TextField 
                  fullWidth 
                  label="Database Name" 
                  value={systemSettings.dbName}
                  onChange={(e) => setSystemSettings({ ...systemSettings, dbName: e.target.value })}
                  sx={{ mb: 2 }} 
                />
              </CardContent>
            </Card>
          </Grid>
          <Grid item xs={12} md={6}>
            <Card>
              <CardContent>
                <Typography variant="h6" gutterBottom>Cache & Queue</Typography>
                <TextField 
                  fullWidth 
                  label="Redis Host" 
                  value={systemSettings.redisHost}
                  onChange={(e) => setSystemSettings({ ...systemSettings, redisHost: e.target.value })}
                  sx={{ mb: 2 }} 
                />
                <TextField 
                  fullWidth 
                  label="Redis Port" 
                  value={systemSettings.redisPort}
                  onChange={(e) => setSystemSettings({ ...systemSettings, redisPort: e.target.value })}
                  sx={{ mb: 2 }} 
                />
                <TextField 
                  fullWidth 
                  label="NATS Server" 
                  value={systemSettings.natsServer}
                  onChange={(e) => setSystemSettings({ ...systemSettings, natsServer: e.target.value })}
                  sx={{ mb: 2 }} 
                />
              </CardContent>
            </Card>
          </Grid>
        </Grid>
      )}

      <Box sx={{ mt: 3, display: 'flex', justifyContent: 'flex-end', gap: 2 }}>
        <Button variant="outlined">Cancel</Button>
        <Button 
          variant="contained" 
          startIcon={saving ? <CircularProgress size={20} color="inherit" /> : <Save />} 
          onClick={handleSave}
          disabled={saving}
        >
          {saving ? 'Saving...' : 'Save Changes'}
        </Button>
      </Box>
    </Box>
  );
};

export default Settings;
