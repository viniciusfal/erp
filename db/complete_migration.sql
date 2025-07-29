-- Script completo de migração para atualizar o banco de dados conforme tables.sql
-- Este script é seguro e não perde dados existentes

-- =====================================================
-- 1. ADICIONAR NOVAS COLUNAS À TABELA ACCOUNTABILITY
-- =====================================================

-- Adicionar coluna total_atlas (DECIMAL)
ALTER TABLE accountability ADD COLUMN IF NOT EXISTS total_atlas DECIMAL;

-- Adicionar coluna ter_vias (INT)
ALTER TABLE accountability ADD COLUMN IF NOT EXISTS ter_vias INT;

-- Adicionar coluna vias_atlas (INT)
ALTER TABLE accountability ADD COLUMN IF NOT EXISTS vias_atlas INT;

-- Adicionar coluna total_sec_vias (DECIMAL)
ALTER TABLE accountability ADD COLUMN IF NOT EXISTS total_sec_vias DECIMAL;

-- Adicionar coluna total_ter_vias (DECIMAL)
ALTER TABLE accountability ADD COLUMN IF NOT EXISTS total_ter_vias DECIMAL;

-- Adicionar coluna details (TEXT)
ALTER TABLE accountability ADD COLUMN IF NOT EXISTS details TEXT;

-- Adicionar coluna desconto (DECIMAL)
ALTER TABLE accountability ADD COLUMN IF NOT EXISTS desconto DECIMAL;

-- Adicionar coluna annex (TEXT[])
ALTER TABLE accountability ADD COLUMN IF NOT EXISTS annex TEXT[];

-- Adicionar coluna updated_at (TIMESTAMP)
ALTER TABLE accountability ADD COLUMN IF NOT EXISTS updated_at TIMESTAMP;

-- Atualizar a coluna updated_at com o valor de created_at para registros existentes
UPDATE accountability SET updated_at = created_at WHERE updated_at IS NULL;

-- =====================================================
-- 2. CRIAR TABELAS QUE PODEM NÃO EXISTIR
-- =====================================================

-- Criar tabela meta se não existir
CREATE TABLE IF NOT EXISTS meta (
    id VARCHAR PRIMARY KEY,
    month VARCHAR NOT NULL,
    metaValue FLOAT NOT NULL
);

-- Criar tabela safe se não existir
CREATE TABLE IF NOT EXISTS safe (
    id VARCHAR PRIMARY KEY,
    send_date DATE NOT NULL,
    send_amount INT NOT NULL,
    active BOOLEAN NOT NULL,
    code VARCHAR NOT NULL,
    resp VARCHAR,
    details TEXT
);

-- Criar tabela transactions se não existir
CREATE TABLE IF NOT EXISTS transactions (
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

-- Tabela accountability_change_request removida conforme solicitado

-- Criar tabela sale_point se não existir
CREATE TABLE IF NOT EXISTS sale_point (
    id UUID PRIMARY KEY,
    guiche_name VARCHAR NOT NULL
);

-- Criar tabela users se não existir (sem DROP para não perder dados)
CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY,
    name VARCHAR NOT NULL,
    password VARCHAR NOT NULL,
    email VARCHAR NOT NULL UNIQUE,
    role VARCHAR NOT NULL
);

-- Criar tabela partners se não existir (sem DROP para não perder dados)
CREATE TABLE IF NOT EXISTS partners (
    id UUID PRIMARY KEY,
    name VARCHAR NOT NULL,
    taxa_parceiro DECIMAL(5,2) NOT NULL,
    guiche_name VARCHAR
);

-- Criar tabela configs se não existir
CREATE TABLE IF NOT EXISTS configs (
    id UUID PRIMARY KEY,
    value_second_via DECIMAL NOT NULL,
    value_tera_via DECIMAL NOT NULL,
    sale_points TEXT[],
    partners TEXT[],
    taxa_card_deb DECIMAL(5,2) NOT NULL,
    taxa_card_cred DECIMAL(5,2) NOT NULL
);

-- Inserir configuração padrão se a tabela configs estiver vazia
INSERT INTO configs (id, value_second_via, value_tera_via, sale_points, partners, taxa_card_deb, taxa_card_cred)
SELECT 
    gen_random_uuid(),
    10.00,  -- valor padrão para segunda via
    15.00,  -- valor padrão para terceira via
    ARRAY[]::TEXT[],  -- array vazio para sale_points
    ARRAY[]::TEXT[],  -- array vazio para partners
    2.50,   -- taxa padrão para cartão de débito (2.5%)
    3.50    -- taxa padrão para cartão de crédito (3.5%)
WHERE NOT EXISTS (SELECT 1 FROM configs);

-- =====================================================
-- 3. VERIFICAÇÃO FINAL
-- =====================================================

-- Mostrar as colunas da tabela accountability após a migração
SELECT 'Verificação da tabela accountability:' as info;
SELECT column_name, data_type, is_nullable 
FROM information_schema.columns 
WHERE table_name = 'accountability' 
ORDER BY ordinal_position;

-- Mostrar o número de registros na tabela accountability
SELECT 'Total de registros na tabela accountability:' as info, COUNT(*) as total_registros FROM accountability; 