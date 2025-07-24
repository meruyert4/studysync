CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    email TEXT NOT NULL UNIQUE,
    username TEXT NOT NULL,
    password TEXT NOT NULL,
    avatar_url TEXT,
    xp INTEGER DEFAULT 0,
    created_at TIMESTAMP NOT NULL DEFAULT now()
);
