CREATE TYPE user_role AS ENUM ('guest', 'student', 'teacher');
CREATE TYPE user_status AS ENUM ('active', 'created');

CREATE TABLE IF NOT EXISTS users (
    id serial PRIMARY KEY,
    username TEXT UNIQUE NOT NULL,
    password TEXT NOT NULL,
    role   user_role   NOT NULL,
    status user_status NOT NULL,
    created TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS groups (
    id serial PRIMARY KEY,
    name TEXT UNIQUE NOT NULL,
    created TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS task (
    ...
);
