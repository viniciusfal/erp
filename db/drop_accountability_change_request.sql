-- Script para remover a tabela accountability_change_request
-- ATENÇÃO: Este comando irá remover TODOS os dados da tabela permanentemente

-- Verificar se a tabela existe antes de tentar removê-la
DO $$
BEGIN
    IF EXISTS (SELECT FROM information_schema.tables WHERE table_name = 'accountability_change_request') THEN
        RAISE NOTICE 'Tabela accountability_change_request encontrada. Removendo...';
        DROP TABLE accountability_change_request CASCADE;
        RAISE NOTICE 'Tabela accountability_change_request removida com sucesso.';
    ELSE
        RAISE NOTICE 'Tabela accountability_change_request não existe no banco de dados.';
    END IF;
END $$;

-- Verificar se a tabela foi removida
SELECT 
    CASE 
        WHEN EXISTS (SELECT FROM information_schema.tables WHERE table_name = 'accountability_change_request') 
        THEN 'ERRO: Tabela accountability_change_request ainda existe' 
        ELSE 'SUCESSO: Tabela accountability_change_request foi removida' 
    END as status; 