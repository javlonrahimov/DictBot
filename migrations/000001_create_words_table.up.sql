CREATE TABLE IF NOT EXISTS words
(
    id         bigserial PRIMARY KEY,
    word       text    NOT NULL,
    word_type  text    NOT NULL,
    definition text    NOT NULL,
    version    integer NOT NULL DEFAULT 1
);