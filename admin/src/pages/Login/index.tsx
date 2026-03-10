import { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { useForm } from 'react-hook-form';
import { zodResolver } from '@hookform/resolvers/zod';
import { z } from 'zod';
import {
  Box, Card, CardContent, TextField, Button, Typography,
  Alert, CircularProgress, InputAdornment, IconButton
} from '@mui/material';
import { Visibility, VisibilityOff, Games } from '@mui/icons-material';
import { useAppDispatch } from '../../store/hooks';
import { loginSuccess, loginStart, loginFailure } from '../../store/slices/authSlice';
import { authAPI } from '../../services/api';

const loginSchema = z.object({
  username: z.string().min(1, 'Username is required'),
  password: z.string().min(1, 'Password is required'),
});

type LoginForm = z.infer<typeof loginSchema>;

const Login = () => {
  const navigate = useNavigate();
  const dispatch = useAppDispatch();
  const [showPassword, setShowPassword] = useState(false);
  const [error, setError] = useState('');

  const {
    register,
    handleSubmit,
    formState: { errors, isSubmitting },
  } = useForm<LoginForm>({
    resolver: zodResolver(loginSchema),
  });

  const onSubmit = async (data: LoginForm) => {
    dispatch(loginStart());
    setError('');
    
    try {
      const response = await authAPI.login(data);
      const { user, token } = response.data;
      
      sessionStorage.setItem('adminToken', token);
      dispatch(loginSuccess({ user, token }));
      navigate('/');
    } catch (err: any) {
      const message = err.response?.data?.message || 'Login failed. Please try again.';
      setError(message);
      dispatch(loginFailure(message));
    }
  };

  return (
    <Box
      sx={{
        minHeight: '100vh',
        display: 'flex',
        alignItems: 'center',
        justifyContent: 'center',
        bgcolor: 'background.default',
        p: 2,
      }}
    >
      <Card sx={{ maxWidth: 440, width: '100%', boxShadow: 3 }}>
        <CardContent sx={{ p: 4 }}>
          <Box sx={{ textAlign: 'center', mb: 4 }}>
            <Games sx={{ fontSize: 56, color: 'primary.main', mb: 1 }} />
            <Typography variant="h4" fontWeight="bold">
              Casino Admin
            </Typography>
            <Typography color="text.secondary">
              Sign in to your account
            </Typography>
          </Box>

          {error && (
            <Alert severity="error" sx={{ mb: 3 }}>
              {error}
            </Alert>
          )}

          <form onSubmit={handleSubmit(onSubmit)}>
            <TextField
              fullWidth
              label="Username"
              {...register('username')}
              error={!!errors.username}
              helperText={errors.username?.message}
              sx={{ mb: 2 }}
            />
            
            <TextField
              fullWidth
              label="Password"
              type={showPassword ? 'text' : 'password'}
              {...register('password')}
              error={!!errors.password}
              helperText={errors.password?.message}
              sx={{ mb: 3 }}
              InputProps={{
                endAdornment: (
                  <InputAdornment position="end">
                    <IconButton
                      onClick={() => setShowPassword(!showPassword)}
                      edge="end"
                    >
                      {showPassword ? <VisibilityOff /> : <Visibility />}
                    </IconButton>
                  </InputAdornment>
                ),
              }}
            />

            <Button
              fullWidth
              type="submit"
              variant="contained"
              size="large"
              disabled={isSubmitting}
              sx={{ py: 1.5 }}
            >
              {isSubmitting ? <CircularProgress size={24} color="inherit" /> : 'Sign In'}
            </Button>
          </form>

          <Box sx={{ mt: 3, textAlign: 'center' }}>
            <Typography variant="body2" color="text.secondary">
              Forgot your password?{' '}
              <Button variant="text" size="small">
                Reset Password
              </Button>
            </Typography>
          </Box>
        </CardContent>
      </Card>
    </Box>
  );
};

export default Login;
