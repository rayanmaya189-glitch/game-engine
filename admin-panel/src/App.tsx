import { BrowserRouter, Routes, Route, Navigate } from 'react-router-dom';
import { QueryClient, QueryClientProvider } from '@tanstack/react-query';
import { Sidebar } from './components/layout/Sidebar';
import { Dashboard } from './components/dashboard/Dashboard';
import { UsersPage } from './components/users/UsersPage';
import { GamesPage } from './components/games/GamesPage';
import { TransactionsPage } from './components/financials/TransactionsPage';
import { TournamentsPage } from './components/tournaments/TournamentsPage';
import { BonusesPage } from './components/bonuses/BonusesPage';
import { MerchantsPage } from './components/merchants/MerchantsPage';
import { AgentsPage } from './components/agents/AgentsPage';
import { LoginPage } from './components/auth/LoginPage';
import { useAuthStore } from './context/authStore';

const queryClient = new QueryClient({
  defaultOptions: {
    queries: {
      retry: 1,
      refetchOnWindowFocus: false,
    },
  },
});

function ProtectedRoute({ children }: { children: React.ReactNode }) {
  const isAuthenticated = useAuthStore((state) => state.isAuthenticated);
  return isAuthenticated ? <Sidebar>{children}</Sidebar> : <Navigate to="/login" />;
}

function AppRoutes() {
  const isAuthenticated = useAuthStore((state) => state.isAuthenticated);

  return (
    <Routes>
      <Route path="/login" element={isAuthenticated ? <Navigate to="/" /> : <LoginPage />} />
      <Route
        path="/"
        element={
          <ProtectedRoute>
            <Dashboard />
          </ProtectedRoute>
        }
      />
      <Route
        path="/users"
        element={
          <ProtectedRoute>
            <UsersPage />
          </ProtectedRoute>
        }
      />
      <Route
        path="/games"
        element={
          <ProtectedRoute>
            <GamesPage />
          </ProtectedRoute>
        }
      />
      <Route
        path="/transactions"
        element={
          <ProtectedRoute>
            <TransactionsPage />
          </ProtectedRoute>
        }
      />
      <Route
        path="/tournaments"
        element={
          <ProtectedRoute>
            <TournamentsPage />
          </ProtectedRoute>
        }
      />
      <Route
        path="/bonuses"
        element={
          <ProtectedRoute>
            <BonusesPage />
          </ProtectedRoute>
        }
      />
      <Route
        path="/merchants"
        element={
          <ProtectedRoute>
            <MerchantsPage />
          </ProtectedRoute>
        }
      />
      <Route
        path="/agents"
        element={
          <ProtectedRoute>
            <AgentsPage />
          </ProtectedRoute>
        }
      />
      <Route path="*" element={<Navigate to="/" />} />
    </Routes>
  );
}

function App() {
  return (
    <QueryClientProvider client={queryClient}>
      <BrowserRouter>
        <AppRoutes />
      </BrowserRouter>
    </QueryClientProvider>
  );
}

export default App;
