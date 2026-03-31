import {
  Card, CardContent, Typography, Grid, Switch, FormControlLabel
} from '@mui/material';

interface NotificationSettingsProps {
  notificationSettings: {
    emailNotifications: boolean;
    smsNotifications: boolean;
    pushNotifications: boolean;
    newUserRegistrations: boolean;
    largeTransactions: boolean;
    kycApprovals: boolean;
    systemAlerts: boolean;
  };
  onNotificationChange: (settings: NotificationSettingsProps['notificationSettings']) => void;
}

const NotificationSettings = ({ notificationSettings, onNotificationChange }: NotificationSettingsProps) => {
  return (
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
                  onChange={(e) => onNotificationChange({ ...notificationSettings, emailNotifications: e.target.checked })}
                />
              }
              label="Email Notifications"
              sx={{ display: 'block', mb: 1 }}
            />
            <FormControlLabel
              control={
                <Switch
                  checked={notificationSettings.smsNotifications}
                  onChange={(e) => onNotificationChange({ ...notificationSettings, smsNotifications: e.target.checked })}
                />
              }
              label="SMS Notifications"
              sx={{ display: 'block', mb: 1 }}
            />
            <FormControlLabel
              control={
                <Switch
                  checked={notificationSettings.pushNotifications}
                  onChange={(e) => onNotificationChange({ ...notificationSettings, pushNotifications: e.target.checked })}
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
                  onChange={(e) => onNotificationChange({ ...notificationSettings, newUserRegistrations: e.target.checked })}
                />
              }
              label="New user registrations"
              sx={{ display: 'block', mb: 1 }}
            />
            <FormControlLabel
              control={
                <Switch
                  checked={notificationSettings.largeTransactions}
                  onChange={(e) => onNotificationChange({ ...notificationSettings, largeTransactions: e.target.checked })}
                />
              }
              label="Large transactions"
              sx={{ display: 'block', mb: 1 }}
            />
            <FormControlLabel
              control={
                <Switch
                  checked={notificationSettings.kycApprovals}
                  onChange={(e) => onNotificationChange({ ...notificationSettings, kycApprovals: e.target.checked })}
                />
              }
              label="KYC approvals"
              sx={{ display: 'block', mb: 1 }}
            />
            <FormControlLabel
              control={
                <Switch
                  checked={notificationSettings.systemAlerts}
                  onChange={(e) => onNotificationChange({ ...notificationSettings, systemAlerts: e.target.checked })}
                />
              }
              label="System alerts"
              sx={{ display: 'block', mb: 1 }}
            />
          </Grid>
        </Grid>
      </CardContent>
    </Card>
  );
};

export default NotificationSettings;
