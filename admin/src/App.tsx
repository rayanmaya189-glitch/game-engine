import { BrowserRouter, Routes, Route, Navigate } from 'react-router-dom';
import { Snackbar, Alert } from '@mui/material';
import { useAppDispatch, useAppSelector } from './store/hooks';
import { hideSnackbar } from './store/slices/uiSlice';
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

const PrivateRoute = ({ children }: { children: React.ReactNode }) => {
  const { isAuthenticated } = useAppSelector((state: any) => state.auth);
  return isAuthenticated ? <>{children}</> : <Navigate to="/login" />;
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
          <Route path="claims" element={<ClaimsManagement />} />
          <Route path="users" element={<Users />} />
          <Route path="merchants" element={<Merchants />} />
          <Route path="agents" element={<Agents />} />
          <Route path="games" element={<Games />} />
          <Route path="tournaments" element={<Tournaments />} />
          <Route path="jackpots" element={<Jackpots />} />
          <Route path="bonuses" element={<Bonuses />} />
          <Route path="payments" element={<Payments />} />
          <Route path="reports" element={<Reports />} />
          <Route path="settings" element={<Settings />} />
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
