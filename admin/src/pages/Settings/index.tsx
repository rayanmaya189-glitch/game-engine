import { useState } from 'react';
import {
  Box, Card, CardContent, Typography, Grid, TextField, Button,
  Tabs, Tab, Alert, CircularProgress
} from '@mui/material';
import { Save, Refresh } from '@mui/icons-material';
import { useAppDispatch } from '../../store/hooks';
import { showSnackbar } from '../../store/slices/uiSlice';
import GeneralSettings from './GeneralSettings';
import SecuritySettings from './SecuritySettings';
import NotificationSettings from './NotificationSettings';

const Settings = () => {
  const dispatch = useAppDispatch();
  const [activeTab, setActiveTab] = useState(0);
  const [saving, setSaving] = useState(false);
  const [saved, setSaved] = useState(false);

  const [generalSettings, setGeneralSettings] = useState({
    siteName: 'Casino Admin',
    supportEmail: 'support@casino.com',
    supportPhone: '+1 234 567 8900',
    timezone: 'UTC',
    maintenanceMode: false,
  });

  const [securitySettings, setSecuritySettings] = useState({
    twoFactorAuth: true,
    ipWhitelist: true,
    sessionTimeout: false,
    loginNotifications: true,
    adminPassword: '',
    confirmPassword: '',
  });

  const [apiSettings, setApiSettings] = useState({
    apiKey: 'sk_live_xxxxxxxxxxxxx',
    webhookUrl: '',
    enableApiAccess: true,
    rateLimiting: false,
  });

  const [notificationSettings, setNotificationSettings] = useState({
    emailNotifications: true,
    smsNotifications: true,
    pushNotifications: true,
    newUserRegistrations: true,
    largeTransactions: true,
    kycApprovals: true,
    systemAlerts: true,
  });

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
        <GeneralSettings
          generalSettings={generalSettings}
          onGeneralChange={setGeneralSettings}
        />
      )}

      {activeTab === 1 && (
        <SecuritySettings
          securitySettings={securitySettings}
          apiSettings={apiSettings}
          onSecurityChange={setSecuritySettings}
          onApiChange={setApiSettings}
          onCopyApiKey={handleCopyApiKey}
          onGenerateNewKey={handleGenerateNewKey}
        />
      )}

      {activeTab === 2 && (
        <NotificationSettings
          notificationSettings={notificationSettings}
          onNotificationChange={setNotificationSettings}
        />
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
