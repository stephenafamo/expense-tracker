CREATE TABLE users (
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	email STRING NOT NULL UNIQUE,
	password STRING NULL,
	attempts INTEGER NULL DEFAULT NULL,
	last_attempt_time TIMESTAMP WITH TIME ZONE NULL DEFAULT NULL,
	locked TIMESTAMP WITH TIME ZONE NULL DEFAULT NULL,
	created_at TIMESTAMP WITH TIME ZONE NULL DEFAULT current_timestamp(),
	updated_at TIMESTAMP WITH TIME ZONE NULL,
	FAMILY "primary" (id, email, password, attempts, last_attempt_time, locked, created_at, updated_at)
);

CREATE TABLE user_tokens (
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	created_at TIMESTAMP WITH TIME ZONE NULL DEFAULT current_timestamp(),
	updated_at TIMESTAMP WITH TIME ZONE NULL,
	deleted_at TIMESTAMP WITH TIME ZONE NULL,
	user_id UUID NOT NULL REFERENCES users (id) ON DELETE CASCADE ON UPDATE CASCADE,
	token STRING NOT NULL,
	INDEX idx_user_tokens_token (token ASC),
	FAMILY "primary" (id, created_at, updated_at, deleted_at, user_id, token)
);

CREATE TABLE http_sessions (
	id INT NOT NULL DEFAULT unique_rowid(),
	key BYTES NULL,
	data BYTES NULL,
	created_on TIMESTAMP WITH TIME ZONE NULL DEFAULT current_timestamp(),
	modified_on TIMESTAMP WITH TIME ZONE NULL,
	expires_on TIMESTAMP WITH TIME ZONE NULL,
	CONSTRAINT "primary" PRIMARY KEY (id ASC),
	FAMILY "primary" (id, key, data, created_on, modified_on, expires_on)
);
