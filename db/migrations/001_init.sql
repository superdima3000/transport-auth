-- +goose Up
CREATE TABLE keys (
                      id      INTEGER PRIMARY KEY AUTOINCREMENT,
                      value   TEXT    NOT NULL,
                      name    TEXT
);

CREATE TABLE terminals (
                           id              INTEGER PRIMARY KEY AUTOINCREMENT,
                           serial_number   TEXT    NOT NULL UNIQUE,
                           address         TEXT    NOT NULL,
                           name            TEXT    NOT NULL
);

CREATE TABLE cards (
                       id          INTEGER PRIMARY KEY AUTOINCREMENT,
                       card_number TEXT    NOT NULL UNIQUE,
                       balance     REAL    NOT NULL DEFAULT 0,
                       is_blocked  INTEGER NOT NULL DEFAULT 0,  -- 0 = false, 1 = true
                       owner_name  TEXT,
                       key_id      INTEGER REFERENCES keys(id)  -- один ключ → много карт
);

CREATE TABLE transactions (
                              id          INTEGER PRIMARY KEY AUTOINCREMENT,
                              amount      REAL    NOT NULL,
                              created_at  TEXT    NOT NULL DEFAULT (datetime('now')),
                              card_id     INTEGER NOT NULL REFERENCES cards(id),
                              terminal_id INTEGER NOT NULL REFERENCES terminals(id)
);

CREATE TABLE users (
                       id           INTEGER PRIMARY KEY AUTOINCREMENT,
                       login        TEXT    NOT NULL UNIQUE,
                       password     TEXT    NOT NULL,  -- хранить хэш!
                       is_admin     INTEGER NOT NULL DEFAULT 0  -- 0 = user, 1 = admin
);

-- +goose Down
DROP TABLE IF EXISTS transactions;
DROP TABLE IF EXISTS cards;
DROP TABLE IF EXISTS terminals;
DROP TABLE IF EXISTS keys;
DROP TABLE IF EXISTS users;