-- 'user' is reserved word in postgres.
CREATE TABLE user_table (
    id      SERIAL  PRIMARY KEY,
    name    TEXT    NOT NULL,
    surname TEXT    NOT NULL
);