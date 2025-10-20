-- Trego Database Schema
-- PostgreSQL database setup for sports game management application

-- Create database (run this separately if needed)
-- CREATE DATABASE trego;

-- Enable UUID extension (for generating string IDs)
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Users table
CREATE TABLE users (
    user_id TEXT PRIMARY KEY DEFAULT uuid_generate_v4()::text,
    name TEXT NOT NULL,
    email TEXT UNIQUE NOT NULL,
    picture_url TEXT,
    phone_number TEXT,
    location TEXT,
    reputation INTEGER DEFAULT 0 CHECK (reputation >= 0),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- Sports table
CREATE TABLE sports (
    sport_name TEXT PRIMARY KEY,
    icon_url TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- User sports junction table (many-to-many relationship)
CREATE TABLE user_sports (
    user_id TEXT NOT NULL,
    sport_name TEXT NOT NULL,
    position TEXT CHECK (position IN ('front', 'back')),
    skill_level TEXT NOT NULL CHECK (skill_level IN ('beginner', 'intermediate', 'advanced')),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    PRIMARY KEY (user_id, sport_name),
    FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE,
    FOREIGN KEY (sport_name) REFERENCES sports(sport_name) ON DELETE CASCADE
);

-- Games table
CREATE TABLE games (
    game_id TEXT PRIMARY KEY DEFAULT uuid_generate_v4()::text,
    host_id TEXT NOT NULL,
    sport_name TEXT NOT NULL,
    title TEXT NOT NULL,
    description TEXT,
    start_time TIMESTAMP WITH TIME ZONE NOT NULL,
    end_time TIMESTAMP WITH TIME ZONE NOT NULL,
    location TEXT NOT NULL,
    skill_range TEXT,
    capacity INTEGER NOT NULL CHECK (capacity > 0),
    skill_level TEXT CHECK (skill_level IN ('beginner', 'intermediate', 'advanced')),
    visibility TEXT NOT NULL DEFAULT 'public' CHECK (visibility IN ('public', 'invite-only')),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    FOREIGN KEY (host_id) REFERENCES users(user_id) ON DELETE CASCADE,
    FOREIGN KEY (sport_name) REFERENCES sports(sport_name) ON DELETE RESTRICT,
    CONSTRAINT valid_time_range CHECK (end_time > start_time)
);

-- Game players junction table (many-to-many relationship)
CREATE TABLE game_players (
    user_id TEXT NOT NULL,
    game_id TEXT NOT NULL,
    attendance TEXT DEFAULT 'none' CHECK (attendance IN ('true', 'false', 'none')),
    joined_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    PRIMARY KEY (user_id, game_id),
    FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE,
    FOREIGN KEY (game_id) REFERENCES games(game_id) ON DELETE CASCADE
);

-- Indexes for better performance
CREATE INDEX idx_users_email ON users(email);
CREATE INDEX idx_users_location ON users(location);
CREATE INDEX idx_user_sports_user_id ON user_sports(user_id);
CREATE INDEX idx_user_sports_sport_name ON user_sports(sport_name);
CREATE INDEX idx_games_host_id ON games(host_id);
CREATE INDEX idx_games_sport_name ON games(sport_name);
CREATE INDEX idx_games_start_time ON games(start_time);
CREATE INDEX idx_games_location ON games(location);
CREATE INDEX idx_games_visibility ON games(visibility);
CREATE INDEX idx_game_players_game_id ON game_players(game_id);
CREATE INDEX idx_game_players_user_id ON game_players(user_id);
CREATE INDEX idx_game_players_attendance ON game_players(attendance);

-- Trigger function to update updated_at timestamp
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ language 'plpgsql';

-- Triggers for updated_at
CREATE TRIGGER update_users_updated_at BEFORE UPDATE ON users
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_games_updated_at BEFORE UPDATE ON games
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

-- Insert some default sports
INSERT INTO sports (sport_name, icon_url) VALUES
('Basketball', '/icons/basketball.svg'),
('Football', '/icons/football.svg'),
('Soccer', '/icons/soccer.svg'),
('Tennis', '/icons/tennis.svg'),
('Volleyball', '/icons/volleyball.svg'),
('Baseball', '/icons/baseball.svg'),
('Hockey', '/icons/hockey.svg'),
('Badminton', '/icons/badminton.svg'),
('Table Tennis', '/icons/table-tennis.svg'),
('Swimming', '/icons/swimming.svg')
ON CONFLICT (sport_name) DO NOTHING;
