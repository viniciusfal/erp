-- Script para verificar tabelas que precisam ser criadas
-- Execute este script para ver quais tabelas do schema tables.sql ainda não existem

-- Verificar se a tabela meta existe
SELECT 
    CASE 
        WHEN EXISTS (SELECT FROM information_schema.tables WHERE table_name = 'meta') 
        THEN 'Tabela meta: EXISTE' 
        ELSE 'Tabela meta: NÃO EXISTE - PRECISA CRIAR' 
    END as status;

-- Verificar se a tabela safe existe
SELECT 
    CASE 
        WHEN EXISTS (SELECT FROM information_schema.tables WHERE table_name = 'safe') 
        THEN 'Tabela safe: EXISTE' 
        ELSE 'Tabela safe: NÃO EXISTE - PRECISA CRIAR' 
    END as status;

-- Verificar se a tabela transactions existe
SELECT 
    CASE 
        WHEN EXISTS (SELECT FROM information_schema.tables WHERE table_name = 'transactions') 
        THEN 'Tabela transactions: EXISTE' 
        ELSE 'Tabela transactions: NÃO EXISTE - PRECISA CRIAR' 
    END as status;

-- Tabela accountability_change_request removida conforme solicitado

-- Verificar se a tabela sale_point existe
SELECT 
    CASE 
        WHEN EXISTS (SELECT FROM information_schema.tables WHERE table_name = 'sale_point') 
        THEN 'Tabela sale_point: EXISTE' 
        ELSE 'Tabela sale_point: NÃO EXISTE - PRECISA CRIAR' 
    END as status;

-- Verificar se a tabela users existe
SELECT 
    CASE 
        WHEN EXISTS (SELECT FROM information_schema.tables WHERE table_name = 'users') 
        THEN 'Tabela users: EXISTE' 
        ELSE 'Tabela users: NÃO EXISTE - PRECISA CRIAR' 
    END as status;

-- Verificar se a tabela partners existe
SELECT 
    CASE 
        WHEN EXISTS (SELECT FROM information_schema.tables WHERE table_name = 'partners') 
        THEN 'Tabela partners: EXISTE' 
        ELSE 'Tabela partners: NÃO EXISTE - PRECISA CRIAR' 
    END as status;

-- Verificar se a tabela configs existe
SELECT 
    CASE 
        WHEN EXISTS (SELECT FROM information_schema.tables WHERE table_name = 'configs') 
        THEN 'Tabela configs: EXISTE' 
        ELSE 'Tabela configs: NÃO EXISTE - PRECISA CRIAR' 
    END as status; 