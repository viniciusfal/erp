-- Script para criar a tabela configs
-- Baseado na estrutura encontrada no ConfigRepository e docs/tables.md
-- Corrigido para relação adequada com sale_point

-- Criar tabela configs se não existir
CREATE TABLE IF NOT EXISTS configs (
    id UUID PRIMARY KEY,
    value_second_via DECIMAL NOT NULL,
    value_tera_via DECIMAL NOT NULL,
    sale_points UUID[],  -- Array de IDs dos sale_points
    partners UUID[],     -- Array de IDs dos partners
    taxa_card_deb DECIMAL(5,2) NOT NULL,
    taxa_card_cred DECIMAL(5,2) NOT NULL
);

-- Inserir configuração padrão se a tabela estiver vazia
INSERT INTO configs (id, value_second_via, value_tera_via, sale_points, partners, taxa_card_deb, taxa_card_cred)
SELECT 
    gen_random_uuid(),
    10.00,  -- valor padrão para segunda via
    15.00,  -- valor padrão para terceira via
    ARRAY[]::UUID[],  -- array vazio para sale_points (IDs)
    ARRAY[]::UUID[],  -- array vazio para partners (IDs)
    2.50,   -- taxa padrão para cartão de débito (2.5%)
    3.50    -- taxa padrão para cartão de crédito (3.5%)
WHERE NOT EXISTS (SELECT 1 FROM configs);

-- Verificar se a tabela foi criada corretamente
SELECT 'Tabela configs criada com sucesso!' as status;

-- Mostrar a estrutura da tabela
SELECT column_name, data_type, is_nullable 
FROM information_schema.columns 
WHERE table_name = 'configs' 
ORDER BY ordinal_position;

-- Mostrar os dados da configuração
SELECT * FROM configs;

-- Exemplo de como relacionar com sale_point
-- SELECT c.*, sp.guiche_name 
-- FROM configs c
-- LEFT JOIN LATERAL unnest(c.sale_points) AS sale_point_id(id) ON true
-- LEFT JOIN sale_point sp ON sp.id = sale_point_id.id; 