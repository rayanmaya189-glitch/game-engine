-- Create roles table
CREATE TABLE IF NOT EXISTS roles (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(50) UNIQUE NOT NULL,
    description TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- Create permissions table
CREATE TABLE IF NOT EXISTS permissions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(100) UNIQUE NOT NULL,
    description TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- Create role_permissions join table
CREATE TABLE IF NOT EXISTS role_permissions (
    role_id UUID REFERENCES roles(id) ON DELETE CASCADE,
    permission_id UUID REFERENCES permissions(id) ON DELETE CASCADE,
    PRIMARY KEY (role_id, permission_id)
);

-- Create user_roles join table
CREATE TABLE IF NOT EXISTS user_roles (
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    role_id UUID REFERENCES roles(id) ON DELETE CASCADE,
    PRIMARY KEY (user_id, role_id)
);

-- Seed basic roles
INSERT INTO roles (name, description) VALUES 
('superadmin', 'System Super Administrator with full access'),
('admin', 'Administrator with management access'),
('support', 'Support staff with read and limited write access'),
('player', 'Regular player')
ON CONFLICT (name) DO NOTHING;

-- Seed basic permissions
INSERT INTO permissions (name, description) VALUES 
('all', 'Global permission for all actions'),
('user.read', 'Read user profiles'),
('user.write', 'Modify user profiles'),
('user.delete', 'Delete users'),
('game.manage', 'Manage games and settings'),
('finance.manage', 'Manage payments and wallets')
ON CONFLICT (name) DO NOTHING;

-- Assign 'all' permission to 'superadmin'
INSERT INTO role_permissions (role_id, permission_id)
SELECT r.id, p.id FROM roles r, permissions p 
WHERE r.name = 'superadmin' AND p.name = 'all'
ON CONFLICT DO NOTHING;

-- Create superadmin user
INSERT INTO users (
    email, 
    password_hash, 
    status, 
    email_verified, 
    marketing_consent, 
    accept_terms,
    created_at,
    updated_at
) VALUES (
    'admin@game-engine.com', 
    '$2a$12$A6Ae/bdusqjinL0zx/8CCOR50/aMbfEf6uLGU2sJDYD2TvrLgo1ga', 
    'active', 
    true, 
    true, 
    true,
    NOW(),
    NOW()
) ON CONFLICT (email) DO NOTHING;

-- Assign superadmin role to superadmin user
INSERT INTO user_roles (user_id, role_id)
SELECT u.id, r.id FROM users u, roles r
WHERE u.email = 'admin@game-engine.com' AND r.name = 'superadmin'
ON CONFLICT DO NOTHING;
