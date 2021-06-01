CREATE TABLE IF NOT EXISTS users (
	id serial PRIMARY KEY,
	username VARCHAR ( 32 ) UNIQUE NOT NULL,
	name VARCHAR ( 128 ) NOT NULL,
	email VARCHAR ( 128 ) UNIQUE NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ
);