import { BrowserRouter, Routes, Route, Navigate } from 'react-router-dom';
import { Snackbar, Alert } from '@mui/material';
import { useAppDispatch, useAppSelector } from './store/hooks';
import { hideSnackbar } from './store/slices/uiSlice';
import { canAccessAdmin } from './utils/permissions';
import Layout from './components/Layout';
import Login from './pages/Login';
import Dashboard from './pages/Dashboard';
import ClaimsManagement from './pages/ClaimsManagement';
import Users from './pages/Users';
import Merchants from './pages/Merchants';
import Agents from './pages/Agents';
import Games from './pages/Games';
import Tournaments from './pages/Tournaments';
import Jackpots from './pages/Jackpots';
import Bonuses from './pages/Bonuses';
import Payments from './pages/Payments';
import Reports from './pages/Reports';
import Settings from './pages/Settings';
import KYC from './pages/KYC';
import BannerManagement from './pages/BannerManagement';
import ReferralManagement from './pages/ReferralManagement';

const PrivateRoute = ({ children }: { children: React.ReactNode }) => {
  const { isAuthenticated, user } = useAppSelector((state: any) => state.auth);
  
  if (!isAuthenticated) return <Navigate to="/login" />;
  
  // Check if user has admin access
  if (user && !canAccessAdmin(user.role)) {
    return <Navigate to="/login" />;
  }
  
  return <>{children}</>;
};

// Permission-based route guard
const PermissionRoute = ({ 
  children, 
  permission 
}: { 
  children: React.ReactNode; 
  permission?: string 
}) => {
  const { user } = useAppSelector((state: any) => state.auth);
  
  if (!user) return <Navigate to="/" />;
  
  // Superadmin has access to everything
  if (user.role === 'superadmin') return <>{children}</>;
  
  // Check specific permission
  if (permission && !user.permissions?.includes(permission)) {
    return <Navigate to="/" />;
  }
  
  return <>{children}</>;
};

function App() {
  const dispatch = useAppDispatch();
  const { snackbar } = useAppSelector((state: any) => state.ui);

  return (
    <BrowserRouter>
      <Routes>
        <Route path="/login" element={<Login />} />
        <Route
          path="/"
          element={
            <PrivateRoute>
              <Layout />
            </PrivateRoute>
          }
        >
          <Route index element={<Dashboard />} />
          <Route path="claims" element={
            <PermissionRoute permission="claims:view">
              <ClaimsManagement />
            </PermissionRoute>
          } />
          <Route path="users" element={
            <PermissionRoute permission="players:view">
              <Users />
            </PermissionRoute>
          } />
          <Route path="merchants" element={
            <PermissionRoute permission="merchants:view">
              <Merchants />
            </PermissionRoute>
          } />
          <Route path="agents" element={
            <PermissionRoute permission="agents:view">
              <Agents />
            </PermissionRoute>
          } />
          <Route path="games" element={
            <PermissionRoute permission="games:view">
              <Games />
            </PermissionRoute>
          } />
          <Route path="tournaments" element={
            <PermissionRoute permission="tournaments:view">
              <Tournaments />
            </PermissionRoute>
          } />
          <Route path="jackpots" element={
            <PermissionRoute permission="jackpots:view">
              <Jackpots />
            </PermissionRoute>
          } />
          <Route path="bonuses" element={
            <PermissionRoute permission="bonuses:view">
              <Bonuses />
            </PermissionRoute>
          } />
          <Route path="payments" element={
            <PermissionRoute permission="payments:view">
              <Payments />
            </PermissionRoute>
          } />
          <Route path="reports" element={
            <PermissionRoute permission="reports:view">
              <Reports />
            </PermissionRoute>
          } />
          <Route path="settings" element={
            <PermissionRoute permission="settings:view">
              <Settings />
            </PermissionRoute>
          } />
          <Route path="kyc" element={
            <PermissionRoute permission="kyc:view">
              <KYC />
            </PermissionRoute>
          } />
          <Route path="banners" element={
            <PermissionRoute permission="banners:view">
              <BannerManagement />
            </PermissionRoute>
          } />
          <Route path="referrals" element={
            <PermissionRoute permission="referrals:view">
              <ReferralManagement />
            </PermissionRoute>
          } />
        </Route>
      </Routes>
      
      <Snackbar
        open={snackbar.open}
        autoHideDuration={6000}
        onClose={() => dispatch(hideSnackbar())}
        anchorOrigin={{ vertical: 'bottom', horizontal: 'right' }}
      >
        <Alert
          onClose={() => dispatch(hideSnackbar())}
          severity={snackbar.severity}
          variant="filled"
          sx={{ width: '100%' }}
        >
          {snackbar.message}
        </Alert>
      </Snackbar>
    </BrowserRouter>
  );
}

export default App;
