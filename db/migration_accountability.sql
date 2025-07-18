-- Migração para adicionar novas colunas à tabela accountability
-- Este script adiciona as colunas que estão no schema tables.sql mas não existem na tabela atual

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

-- Verificar se a migração foi aplicada corretamente
-- SELECT column_name, data_type FROM information_schema.columns 
-- WHERE table_name = 'accountability' ORDER BY ordinal_position; 