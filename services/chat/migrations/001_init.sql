-- Chat Service Database Schema

-- Chat Rooms table: stores chat room configurations
CREATE TABLE IF NOT EXISTS chat_rooms (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    room_id VARCHAR(255) UNIQUE NOT NULL,
    name VARCHAR(255) NOT NULL,
    room_type VARCHAR(50) NOT NULL, -- GLOBAL, GAME, PRIVATE, SUPPORT, TOURNAMENT
    reference_id VARCHAR(255), -- Game room, tournament, or user ID
    status VARCHAR(50) DEFAULT 'ACTIVE', -- ACTIVE, ARCHIVED, DISABLED
    max_participants INTEGER DEFAULT 100,
    current_participants INTEGER DEFAULT 0,
    is_moderated BOOLEAN DEFAULT FALSE,
    is_persistent BOOLEAN DEFAULT TRUE,
    config JSONB DEFAULT '{}',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    archived_at TIMESTAMP WITH TIME ZONE
);

CREATE INDEX idx_chat_rooms_room_id ON chat_rooms(room_id);
CREATE INDEX idx_chat_rooms_room_type ON chat_rooms(room_type);
CREATE INDEX idx_chat_rooms_reference_id ON chat_rooms(reference_id);
CREATE INDEX idx_chat_rooms_status ON chat_rooms(status);
CREATE INDEX idx_chat_rooms_created_at ON chat_rooms(created_at);

-- Chat Messages table: stores all chat messages
CREATE TABLE IF NOT EXISTS chat_messages (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    message_id VARCHAR(255) UNIQUE NOT NULL,
    room_id VARCHAR(255) NOT NULL REFERENCES chat_rooms(room_id) ON DELETE CASCADE,
    sender_id VARCHAR(255) NOT NULL,
    sender_username VARCHAR(255),
    message_type VARCHAR(50) DEFAULT 'TEXT', -- TEXT, IMAGE, SYSTEM, EMOJI, STICKER
    content TEXT NOT NULL,
    reply_to_id VARCHAR(255),
    is_edited BOOLEAN DEFAULT FALSE,
    is_deleted BOOLEAN DEFAULT FALSE,
    is_pinned BOOLEAN DEFAULT FALSE,
    metadata JSONB DEFAULT '{}',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    deleted_at TIMESTAMP WITH TIME ZONE
);

CREATE INDEX idx_chat_messages_message_id ON chat_messages(message_id);
CREATE INDEX idx_chat_messages_room_id ON chat_messages(room_id);
CREATE INDEX idx_chat_messages_sender_id ON chat_messages(sender_id);
CREATE INDEX idx_chat_messages_reply_to_id ON chat_messages(reply_to_id);
CREATE INDEX idx_chat_messages_created_at ON chat_messages(created_at);
CREATE INDEX idx_chat_messages_is_pinned ON chat_messages(is_pinned) WHERE is_pinned = TRUE;

-- Chat Moderation table: stores moderation actions and rules
CREATE TABLE IF NOT EXISTS chat_moderation (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    moderation_id VARCHAR(255) UNIQUE NOT NULL,
    target_user_id VARCHAR(255) NOT NULL,
    room_id VARCHAR(255) REFERENCES chat_rooms(room_id),
    action VARCHAR(50) NOT NULL, -- WARN, MUTE, KICK, BAN, DELETE_MESSAGE
    reason TEXT,
    message_id VARCHAR(255),
    duration_minutes INTEGER, -- NULL for permanent
    moderator_id VARCHAR(255),
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    expires_at TIMESTAMP WITH TIME ZONE
);

CREATE INDEX idx_chat_moderation_moderation_id ON chat_moderation(moderation_id);
CREATE INDEX idx_chat_moderation_target_user_id ON chat_moderation(target_user_id);
CREATE INDEX idx_chat_moderation_room_id ON chat_moderation(room_id);
CREATE INDEX idx_chat_moderation_action ON chat_moderation(action);
CREATE INDEX idx_chat_moderation_is_active ON chat_moderation(is_active);
CREATE INDEX idx_chat_moderation_expires_at ON chat_moderation(expires_at);

-- Trigger to update updated_at on chat_rooms changes
CREATE TRIGGER update_chat_rooms_updated_at BEFORE UPDATE ON chat_rooms
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

-- Trigger to update updated_at on chat_messages changes
CREATE TRIGGER update_chat_messages_updated_at BEFORE UPDATE ON chat_messages
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

-- Trigger to update updated_at on chat_moderation changes
CREATE TRIGGER update_chat_moderation_updated_at BEFORE UPDATE ON chat_moderation
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
