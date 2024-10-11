-- +goose Up
-- +goose StatementBegin
CREATE TABLE users
(
	user_id bigserial,
	email text NOT NULL UNIQUE,
	is_email_verified boolean NOT NULL DEFAULT false,
	is_blocked boolean NOT NULL DEFAULT false,
	block_reason text,
	password_hash text,
	is_internal_auth boolean NOT NULL,
	updated_at timestamp NOT NULL,
	created_at timestamp NOT NULL,
	PRIMARY KEY(user_id)
);

CREATE TABLE oidc_providers
(
	user_id bigint NOT NULL REFERENCES users(user_id) ON DELETE CASCADE,
	provider_tag text NOT NULL,
	provider_user_id text NOT NULL,
	provider_id_token text NOT NULL UNIQUE,
	PRIMARY KEY(user_id)
);

CREATE TABLE profiles
(
	profile_id bigserial,
	user_id bigint NOT NULL REFERENCES users(user_id) ON DELETE CASCADE,
	name text NOT NULL UNIQUE,
	display_name text NOT NULL,
	description text,
	avatar_url text NOT NULL,
	PRIMARY KEY(profile_id)
);

CREATE TABLE sessions
(
	session_id bigserial,
	user_id bigint NOT NULL REFERENCES users(user_id) ON DELETE CASCADE,
	user_agent text NOT NULL,
	session_token text NOT NULL UNIQUE,
	-- refresh_token text NOT NULL UNIQUE,
	expires_in timestamp NOT NULL,
	created_at timestamp NOT NULL,
	PRIMARY KEY(session_id)
);

-- CREATE TABLE refresh_tokens_blacklist
-- (
-- 	refresh_token_id bigserial,
-- 	user_id bigint NOT NULL REFERENCES users(user_id) ON DELETE CASCADE,
-- 	refresh_token text NOT NULL UNIQUE,
-- 	blacklisted_at timestamp NOT NULL,
-- 	PRIMARY KEY(refresh_token_id)
-- );
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- ...
-- +goose StatementEnd
