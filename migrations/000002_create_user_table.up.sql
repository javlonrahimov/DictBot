CREATE TABLE IF NOT EXISTS users
(
    id         bigserial PRIMARY KEY,
    name       text    NOT NULL,
    created_at  date    NOT NULL,
    version    integer NOT NULL DEFAULT 1
);