import { useState } from 'react';
import { Outlet, useNavigate, useLocation } from 'react-router-dom';
import { 
  Box, AppBar, Toolbar, Typography, Drawer, List, ListItem, 
  ListItemButton, ListItemIcon, ListItemText, IconButton,
  Avatar, Menu, MenuItem, Badge, Divider, Chip
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
import { getFilteredMenuItems, isSuperAdmin } from '../../utils/permissions';

const drawerWidth = 260;

const iconMap: Record<string, React.ReactNode> = {
  Dashboard: <Dashboard />,
  Assignment: <Assignment />,
  People: <People />,
  Business: <Business />,
  SupervisorAccount: <SupervisorAccount />,
  Games: <Games />,
  EmojiEvents: <EmojiEvents />,
  CardGiftcard: <CardGiftcard />,
  AccountBalance: <AccountBalance />,
  Assessment: <Assessment />,
  Settings: <Settings />,
};

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

  // Filter menu items based on user permissions
  const userRole = user?.role || '';
  const userPermissions: string[] = user?.permissions || [];
  const filteredMenuItems = getFilteredMenuItems(userRole, userPermissions);

  const getRoleBadgeColor = (role: string) => {
    switch (role) {
      case 'superadmin': return 'error';
      case 'admin': return 'primary';
      case 'finance': return 'success';
      case 'support': return 'info';
      default: return 'default';
    }
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
        {filteredMenuItems.map((item) => (
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
              <ListItemIcon sx={{ minWidth: 40 }}>
                {iconMap[item.icon] || <Dashboard />}
              </ListItemIcon>
              <ListItemText primary={item.text} />
            </ListItemButton>
          </ListItem>
        ))}
      </List>
      <Divider />
      <Box sx={{ p: 2 }}>
        <Chip
          label={userRole}
          size="small"
          color={getRoleBadgeColor(userRole) as any}
          sx={{ mb: 1 }}
        />
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
          {isSuperAdmin(userRole) && (
            <Chip label="SUPERADMIN" color="error" size="small" sx={{ mr: 2 }} />
          )}
          <IconButton color="inherit">
            <Badge badgeContent={4} color="error">
              <Notifications />
            </Badge>
          </IconButton>
          <IconButton onClick={handleMenuOpen} sx={{ ml: 1 }}>
            <Avatar sx={{ bgcolor: 'primary.main', width: 36, height: 36 }}>
              {user?.name?.[0] || user?.username?.[0] || 'A'}
            </Avatar>
          </IconButton>
          <Menu
            anchorEl={anchorEl}
            open={Boolean(anchorEl)}
            onClose={handleMenuClose}
            transformOrigin={{ horizontal: 'right', vertical: 'top' }}
            anchorOrigin={{ horizontal: 'right', vertical: 'bottom' }}
          >
            <MenuItem disabled>
              <Typography variant="body2" color="text.secondary">
                {user?.name || user?.username} ({userRole})
              </Typography>
            </MenuItem>
            <Divider />
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
