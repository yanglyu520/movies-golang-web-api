CREATE TABLE IF NOT EXISTS movies
(
    id         BIGSERIAL PRIMARY KEY,
    created_at INTEGER                     NOT NULL,
    title      TEXT                        NOT NULL,
    year       INTEGER                     NOT NULL,
    runtime    INTEGER                     NOT NULL,
    genres     TEXT[]                      NOT NULL,
    version    INTEGER                     NOT NULL DEFAULT 1
    );