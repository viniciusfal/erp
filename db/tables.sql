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

-- Tabela accountability
CREATE TABLE IF NOT EXISTS accountability (
    id UUID PRIMARY KEY,
    send_date DATE NOT NULL,
    resp_id VARCHAR NOT NULL,
    deb DECIMAL NOT NULL,
    cred DECIMAL NOT NULL,
    pix DECIMAL NOT NULL,
    coin DECIMAL NOT NULL,
    total_of_day DECIMAL NOT NULL,
    total_atlas DECIMAL,
    guiche VARCHAR,
    vias INT,
    ter_vias INT,
    vias_atlas INT,
    total_sec_vias DECIMAL,
    total_ter_vias DECIMAL,
    details TEXT,
    desconto DECIMAL,
    annex TEXT[],
    updated_at TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    resp_name VARCHAR
);

-- Tabela accountability_change_request
CREATE TABLE IF NOT EXISTS accountability_change_request (
    id UUID PRIMARY KEY,
    original_accountability_id UUID NOT NULL,
    requested_by VARCHAR NOT NULL,
    send_date DATE NOT NULL,
    new_deb DECIMAL,
    new_cred DECIMAL,
    new_pix DECIMAL,
    new_coin DECIMAL,
    new_total_of_day DECIMAL,
    new_vias INT,
    new_guiche VARCHAR,
    status VARCHAR NOT NULL,
    request_reason TEXT,
    rejection_reason TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    reviewed_at TIMESTAMP,
    old_accountability JSONB,
    new_total_sec_vias DECIMAL,
    new_ter_vias INT,
    new_total_ter_vias DECIMAL,
    new_desconto DECIMAL,
    new_annex TEXT[]
);

-- Tabela sale_point
CREATE TABLE IF NOT EXISTS sale_point (
    id UUID PRIMARY KEY,
    guiche_name VARCHAR NOT NULL
);

-- Corrigir tabela users (role)
DROP TABLE IF EXISTS users CASCADE;
CREATE TABLE users (
    id UUID PRIMARY KEY,
    name VARCHAR NOT NULL,
    password VARCHAR NOT NULL,
    email VARCHAR NOT NULL UNIQUE,
    role VARCHAR NOT NULL
);

-- Corrigir tabela partners (padronizar campo)
DROP TABLE IF EXISTS partners CASCADE;
CREATE TABLE partners (
    id UUID PRIMARY KEY,
    name VARCHAR NOT NULL,
    taxa_parceiro DECIMAL(5,2) NOT NULL,
    guiche_name VARCHAR
);