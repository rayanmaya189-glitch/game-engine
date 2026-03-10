import { createSlice, PayloadAction } from '@reduxjs/toolkit';

interface UIState {
  sidebarOpen: boolean;
  notifications: Notification[];
  snackbar: {
    open: boolean;
    message: string;
    severity: 'success' | 'error' | 'warning' | 'info';
  };
}

interface Notification {
  id: string;
  type: 'success' | 'error' | 'warning' | 'info';
  title: string;
  message: string;
  timestamp: string;
  read: boolean;
}

const initialState: UIState = {
  sidebarOpen: true,
  notifications: [],
  snackbar: {
    open: false,
    message: '',
    severity: 'info',
  },
};

const uiSlice = createSlice({
  name: 'ui',
  initialState,
  reducers: {
    toggleSidebar: (state) => {
      state.sidebarOpen = !state.sidebarOpen;
    },
    setSidebarOpen: (state, action: PayloadAction<boolean>) => {
      state.sidebarOpen = action.payload;
    },
    addNotification: (state, action: PayloadAction<Omit<Notification, 'id' | 'timestamp' | 'read'>>) => {
      state.notifications.unshift({
        ...action.payload,
        id: Date.now().toString(),
        timestamp: new Date().toISOString(),
        read: false,
      });
    },
    markNotificationRead: (state, action: PayloadAction<string>) => {
      const notification = state.notifications.find(n => n.id === action.payload);
      if (notification) {
        notification.read = true;
      }
    },
    clearNotifications: (state) => {
      state.notifications = [];
    },
    showSnackbar: (state, action: PayloadAction<{ message: string; severity: 'success' | 'error' | 'warning' | 'info' }>) => {
      state.snackbar = {
        open: true,
        message: action.payload.message,
        severity: action.payload.severity,
      };
    },
    hideSnackbar: (state) => {
      state.snackbar.open = false;
    },
  },
});

export const { 
  toggleSidebar, 
  setSidebarOpen,
  addNotification, 
  markNotificationRead,
  clearNotifications,
  showSnackbar,
  hideSnackbar 
} = uiSlice.actions;

export default uiSlice.reducer;
