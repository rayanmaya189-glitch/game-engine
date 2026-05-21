# Admin Portal - Gaming Platform

A comprehensive admin dashboard for managing the gaming platform built with React, TypeScript, and TailwindCSS.

## Features

### Dashboard
- Real-time revenue statistics and charts
- Active users monitoring
- Game performance metrics
- Quick action shortcuts

### User Management
- View and search all users
- Filter by status and role
- Suspend/activate user accounts
- Delete users
- Role-based access control

### Game Management
- Browse all games with detailed information
- Filter by type, status, and provider
- Activate/deactivate games
- View game statistics (RTP, bet ranges, total plays)
- Configure game settings

### Transaction Management
- View all transactions (deposits, withdrawals, bets, wins)
- Filter by type and status
- Approve/reject pending withdrawals
- Transaction history and details

## Tech Stack

- **Frontend Framework**: React 18 with TypeScript
- **Build Tool**: Vite
- **Styling**: TailwindCSS
- **State Management**: Zustand
- **Data Fetching**: TanStack Query (React Query)
- **Routing**: React Router v6
- **Charts**: Recharts
- **HTTP Client**: Axios
- **Utilities**: date-fns, clsx, tailwind-merge

## Project Structure

```
admin-panel/
├── src/
│   ├── components/
│   │   ├── auth/          # Login page
│   │   ├── dashboard/     # Dashboard components
│   │   ├── financials/    # Transaction management
│   │   ├── games/         # Game management
│   │   ├── layout/        # Sidebar and layout
│   │   └── users/         # User management
│   ├── context/           # Auth store (Zustand)
│   ├── hooks/             # Custom hooks and utilities
│   ├── services/          # API service layer
│   ├── App.tsx            # Main app component
│   ├── main.tsx           # Entry point
│   └── index.css          # Global styles
├── package.json
├── tsconfig.json
├── vite.config.ts
├── tailwind.config.js
└── postcss.config.js
```

## Getting Started

### Prerequisites
- Node.js 18+ 
- npm or yarn

### Installation

1. Install dependencies:
```bash
npm install
```

2. Create environment file:
```bash
cp .env.example .env
```

3. Update `.env` with your API URL:
```
VITE_API_URL=http://localhost:8080/api
```

4. Start development server:
```bash
npm run dev
```

The app will be available at `http://localhost:3000`

### Build for Production

```bash
npm run build
```

## API Integration

The admin panel integrates with the backend through a centralized API client (`src/services/api.ts`). All API calls are authenticated using JWT tokens stored in localStorage.

### Available Services

- **authService**: Login, logout, token refresh
- **userService**: CRUD operations for users
- **gameService**: Game management and statistics
- **financialService**: Transactions and wallet operations

### Authentication Flow

1. User logs in with credentials
2. JWT token is stored in localStorage and Zustand store
3. Token is automatically included in all API requests
4. On 401 errors, user is redirected to login page

## Default Credentials (Development)

For testing purposes, use:
- Username: `admin`
- Password: `admin123`

(Note: These need to be configured in the backend)

## License

Proprietary - Gaming Platform
