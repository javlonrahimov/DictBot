CREATE TABLE IF NOT EXISTS users (
    id bigserial PRIMARY KEY,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    first_name text,
    last_name text,
    username text,
    version integer NOT NULL DEFAULT 1
);
