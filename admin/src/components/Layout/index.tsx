import { useState } from 'react';
import { Outlet, useNavigate, useLocation } from 'react-router-dom';
import { 
  Box, AppBar, Toolbar, Typography, Drawer, List, ListItem, 
  ListItemButton, ListItemIcon, ListItemText, IconButton,
  Avatar, Menu, MenuItem, Badge, Divider
} from '@mui/material';
import {
  Menu as MenuIcon,
  Dashboard, Assignment, People, Games, Assessment,
  Notifications, Logout, Person, Settings, EmojiEvents,
  CardGiftcard, AccountBalance, Business, SupervisorAccount
} from '@mui/icons-material';
import { useAppDispatch, useAppSelector } from '../../store/hooks';
import { toggleSidebar } from '../../store/slices/uiSlice';
import { logout } from '../../store/slices/authSlice';

const drawerWidth = 260;

const menuItems = [
  { text: 'Dashboard', icon: <Dashboard />, path: '/' },
  { text: 'Claims Management', icon: <Assignment />, path: '/claims' },
  { text: 'Users', icon: <People />, path: '/users' },
  { text: 'Merchants', icon: <Business />, path: '/merchants' },
  { text: 'Agents', icon: <SupervisorAccount />, path: '/agents' },
  { text: 'Games', icon: <Games />, path: '/games' },
  { text: 'Tournaments', icon: <EmojiEvents />, path: '/tournaments' },
  { text: 'Jackpots', icon: <EmojiEvents />, path: '/jackpots' },
  { text: 'Bonuses', icon: <CardGiftcard />, path: '/bonuses' },
  { text: 'Payments', icon: <AccountBalance />, path: '/payments' },
  { text: 'Reports', icon: <Assessment />, path: '/reports' },
  { text: 'Settings', icon: <Settings />, path: '/settings' },
];

const Layout = () => {
  const dispatch = useAppDispatch();
  const navigate = useNavigate();
  const location = useLocation();
  const { sidebarOpen } = useAppSelector((state: any) => state.ui);
  const { user } = useAppSelector((state: any) => state.auth);
  const [anchorEl, setAnchorEl] = useState<null | HTMLElement>(null);

  const handleMenuOpen = (event: React.MouseEvent<HTMLElement>) => {
    setAnchorEl(event.currentTarget);
  };

  const handleMenuClose = () => {
    setAnchorEl(null);
  };

  const handleLogout = () => {
    dispatch(logout());
    navigate('/login');
  };

  const drawer = (
    <Box sx={{ height: '100%', display: 'flex', flexDirection: 'column' }}>
      <Box sx={{ p: 2, display: 'flex', alignItems: 'center', gap: 1 }}>
        <Games sx={{ color: 'primary.main', fontSize: 32 }} />
        <Typography variant="h6" fontWeight="bold">
          Casino Admin
        </Typography>
      </Box>
      <Divider />
      <List sx={{ flex: 1, px: 1, overflow: 'auto' }}>
        {menuItems.map((item) => (
          <ListItem key={item.text} disablePadding>
            <ListItemButton
              selected={location.pathname === item.path}
              onClick={() => navigate(item.path)}
              sx={{
                '&.Mui-selected': {
                  bgcolor: 'primary.main',
                  color: 'white',
                  '&:hover': { bgcolor: 'primary.dark' },
                  '& .MuiListItemIcon-root': { color: 'white' },
                },
              }}
            >
              <ListItemIcon sx={{ minWidth: 40 }}>{item.icon}</ListItemIcon>
              <ListItemText primary={item.text} />
            </ListItemButton>
          </ListItem>
        ))}
      </List>
      <Divider />
      <Box sx={{ p: 2 }}>
        <Typography variant="body2" color="text.secondary">v1.0.0</Typography>
      </Box>
    </Box>
  );

  return (
    <Box sx={{ display: 'flex' }}>
      <AppBar
        position="fixed"
        sx={{
          zIndex: (theme: any) => theme.zIndex.drawer + 1,
          bgcolor: 'background.paper',
          color: 'text.primary',
          boxShadow: 'none',
          borderBottom: '1px solid',
          borderColor: 'divider',
        }}
      >
        <Toolbar>
          <IconButton color="inherit" edge="start" onClick={() => dispatch(toggleSidebar())} sx={{ mr: 2 }}>
            <MenuIcon />
          </IconButton>
          <Typography variant="h6" sx={{ flexGrow: 1 }}>
            Casino Admin Panel
          </Typography>
          <IconButton color="inherit">
            <Badge badgeContent={4} color="error">
              <Notifications />
            </Badge>
          </IconButton>
          <IconButton onClick={handleMenuOpen} sx={{ ml: 1 }}>
            <Avatar sx={{ bgcolor: 'primary.main', width: 36, height: 36 }}>
              {user?.name?.[0] || 'A'}
            </Avatar>
          </IconButton>
          <Menu
            anchorEl={anchorEl}
            open={Boolean(anchorEl)}
            onClose={handleMenuClose}
            transformOrigin={{ horizontal: 'right', vertical: 'top' }}
            anchorOrigin={{ horizontal: 'right', vertical: 'bottom' }}
          >
            <MenuItem onClick={handleMenuClose}><ListItemIcon><Person fontSize="small" /></ListItemIcon>Profile</MenuItem>
            <MenuItem onClick={handleMenuClose}><ListItemIcon><Settings fontSize="small" /></ListItemIcon>Settings</MenuItem>
            <Divider />
            <MenuItem onClick={handleLogout}><ListItemIcon><Logout fontSize="small" /></ListItemIcon>Logout</MenuItem>
          </Menu>
        </Toolbar>
      </AppBar>

      <Drawer
        variant="permanent"
        open={sidebarOpen}
        sx={{
          width: sidebarOpen ? drawerWidth : 72,
          flexShrink: 0,
          '& .MuiDrawer-paper': {
            width: sidebarOpen ? drawerWidth : 72,
            boxSizing: 'border-box',
            transition: 'width 0.2s',
            overflowX: 'hidden',
          },
        }}
      >
        <Toolbar />
        {drawer}
      </Drawer>

      <Box component="main" sx={{ flexGrow: 1, p: 3, bgcolor: 'background.default', minHeight: '100vh' }}>
        <Toolbar />
        <Outlet />
      </Box>
    </Box>
  );
};

export default Layout;
