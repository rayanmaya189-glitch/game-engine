CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE ws_sessions (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL,
    connection_id VARCHAR(255) NOT NULL UNIQUE,
    status VARCHAR(20) NOT NULL DEFAULT 'connected',
    ip_address INET,
    user_agent VARCHAR(500),
    connected_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    disconnected_at TIMESTAMP WITH TIME ZONE,
    last_heartbeat_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    metadata JSONB,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_ws_sessions_user_id ON ws_sessions (user_id);
CREATE INDEX idx_ws_sessions_connection_id ON ws_sessions (connection_id);
CREATE INDEX idx_ws_sessions_status ON ws_sessions (status);
CREATE INDEX idx_ws_sessions_connected_at ON ws_sessions (connected_at);
CREATE INDEX idx_ws_sessions_last_heartbeat_at ON ws_sessions (last_heartbeat_at);

CREATE TABLE ws_channels (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    channel_name VARCHAR(255) NOT NULL UNIQUE,
    channel_type VARCHAR(30) NOT NULL DEFAULT 'broadcast',
    status VARCHAR(20) NOT NULL DEFAULT 'active',
    max_subscribers INTEGER,
    current_subscribers INTEGER NOT NULL DEFAULT 0,
    metadata JSONB,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_ws_channels_channel_name ON ws_channels (channel_name);
CREATE INDEX idx_ws_channels_channel_type ON ws_channels (channel_type);
CREATE INDEX idx_ws_channels_status ON ws_channels (status);

CREATE TABLE ws_presence (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    session_id UUID NOT NULL REFERENCES ws_sessions (id),
    channel_id UUID NOT NULL REFERENCES ws_channels (id),
    user_id UUID NOT NULL,
    status VARCHAR(20) NOT NULL DEFAULT 'online',
    joined_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    last_active_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    left_at TIMESTAMP WITH TIME ZONE,
    metadata JSONB,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_ws_presence_session_id ON ws_presence (session_id);
CREATE INDEX idx_ws_presence_channel_id ON ws_presence (channel_id);
CREATE INDEX idx_ws_presence_user_id ON ws_presence (user_id);
CREATE INDEX idx_ws_presence_status ON ws_presence (status);
CREATE INDEX idx_ws_presence_joined_at ON ws_presence (joined_at);
