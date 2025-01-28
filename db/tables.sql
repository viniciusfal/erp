-- Tabela meta
CREATE TABLE meta (
    id VARCHAR PRIMARY KEY,
    month VARCHAR NOT NULL,
    metaValue FLOAT NOT NULL
);

-- Tabela safe
CREATE TABLE safe (
    id VARCHAR PRIMARY KEY,
    send_date DATE NOT NULL,
    send_amount INT NOT NULL,
    active BOOLEAN NOT NULL,
    code VARCHAR NOT NULL,
    resp VARCHAR,
    details TEXT
);

-- Tabela transactions
CREATE TABLE transactions (
    id VARCHAR PRIMARY KEY,
    title VARCHAR NOT NULL,
    value FLOAT NOT NULL,
    type VARCHAR NOT NULL,
    category VARCHAR NOT NULL,
    scheduling BOOLEAN NOT NULL,
    annex TEXT,
    payment_date DATE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    pay BOOLEAN,
    details TEXT,
    method VARCHAR,
    nf VARCHAR,
    account VARCHAR
);

-- Tabela users
CREATE TABLE users (
    id VARCHAR PRIMARY KEY,
    name VARCHAR NOT NULL,
    password VARCHAR NOT NULL,
    email VARCHAR NOT NULL UNIQUE,
    rope VARCHAR NOT NULL
);
