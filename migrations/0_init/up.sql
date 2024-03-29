CREATE TABLE counter (
    id SERIAL PRIMARY KEY,
    url TEXT NOT NULL,
    code VARCHAR(50) UNIQUE NOT NULL,
    name TEXT UNIQUE NOT NULL
);

CREATE TABLE redirect (
    id SERIAL PRIMARY KEY,
    date TIMESTAMPTZ NOT NULL,
    counter_id INT REFERENCES counter(id)
);