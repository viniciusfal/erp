# ERP - Documentação Técnica

## Sumário
- [Visão Geral](#visão-geral)
- [Arquitetura](#arquitetura)
- [Funcionalidades Principais](#funcionalidades-principais)
- [Endpoints da API](#endpoints-da-api)
- [Anexos de Arquivos](#anexos-de-arquivos)
- [Uso de Go Routines](#uso-de-go-routines)
- [Setup e Execução](#setup-e-execução)
- [Testes](#testes)
- [Deploy](#deploy)

---

## Visão Geral
Este ERP é um sistema modular para gestão financeira, com controle de accountability, transações, parceiros, fornecedores, usuários, metas, cofres e anexos de arquivos. O sistema foi projetado para ser seguro, escalável e fácil de manter.

## Arquitetura
- **Backend:** Go (Gin, SQL)
- **Banco de Dados:** PostgreSQL
- **Autenticação:** JWT + RBAC (não alterada por regra de negócio)
- **Upload/Download de Arquivos:** Servidor local (pasta `uploads/`)
- **Testes:** Go test, Testify
- **Serviços em background:** Limpeza de arquivos órfãos

## Funcionalidades Principais
- Gestão de accountability e solicitações de alteração
- Controle de transações financeiras (parcelas, importação de Excel, anexos)
- Cadastro e consulta de usuários, parceiros, fornecedores, cofres, metas
- Upload e download seguro de arquivos anexos
- Serviço automático de limpeza de arquivos órfãos

## Endpoints da API
### Autenticação
- `POST /api/login` - Login de usuário
- `POST /api/refresh` - Refresh de sessão

### Accountability
- `POST /api/accountability` - Criar accountability
- `GET /api/accountability/:start_date/:end_date` - Listar por data
- `PUT /api/accountability/:id` - Atualizar accountability
- `POST /api/accountability/change-request` - Solicitar alteração
- ... (ver rotas em `http/routes/accountability-routes.go`)

### Transações
- `POST /api/transaction` - Criar transação
- `GET /api/transaction/:id` - Buscar por ID
- `POST /api/transaction/import-excel` - Importar transações via Excel
- ...

### Anexos de Arquivos
- `POST /api/upload` - Upload de arquivo (multipart/form-data, campo `file`)
- `GET /api/download/:filename` - Download seguro de arquivo

### Outros módulos
- Usuários, parceiros, fornecedores, cofres, metas, etc. (ver rotas em `http/routes/`)

## Anexos de Arquivos
- Arquivos são enviados via endpoint `/api/upload` e salvos em `uploads/`
- O campo `annex` nas entidades armazena um array de caminhos/URLs dos arquivos
- Download seguro via `/api/download/:filename` (com RBAC)
- Limpeza automática de arquivos órfãos (não referenciados no banco)
- Tipos permitidos: PDF, JPG, PNG, DOC, XLS, etc. (até 10MB)

#### Exemplo de Upload de Arquivo

```bash
curl -X POST \
  http://localhost:8000/api/upload \
  -H 'Authorization: Bearer <SEU_TOKEN_JWT>' \
  -F 'file=@/caminho/para/seuarquivo.pdf'
```

**Resposta de sucesso:**
```json
{
  "file_path": "uploads/20240607_153012_1234abcd.pdf",
  "file_name": "20240607_153012_1234abcd.pdf",
  "file_size": 123456
}
```

#### Exemplo de Download de Arquivo

```bash
curl -X GET \
  http://localhost:8000/api/download/20240607_153012_1234abcd.pdf \
  -H 'Authorization: Bearer <SEU_TOKEN_JWT>' \
  -o arquivo_baixado.pdf
```

**Resposta:**
- Download do arquivo solicitado, com headers apropriados para download seguro.
- Se o arquivo não existir ou o nome for inválido, retorna erro 404 ou 400.

## Uso de Go Routines
- **Upload/Download:** Testes unitários simulam uploads/downloads concorrentes
- **Processamento de anexos:** Validação e auditoria de anexos em background
- **Serviço de limpeza:** Limpeza de arquivos órfãos roda em background, removendo arquivos não referenciados
- **Justificativa:** Melhora performance, evita bloqueios e mantém o sistema responsivo

## Setup e Execução
1. **Pré-requisitos:** Go 1.18+, Docker, PostgreSQL
2. **Configuração:**
   - Copie `.env.example` para `.env` e configure variáveis
   - Execute `docker-compose up -d` para subir o banco
   - Rode as migrações SQL em `db/tables.sql`
3. **Build e Run:**
   - `go mod tidy`
   - `go run cmd/main.go`
4. **Acesso:**
   - API: `http://localhost:8000/api/`
   - Uploads: `http://localhost:8000/uploads/`

## Testes
- Testes unitários em `tests/unit/`
- Execute: `go test ./tests/unit/...`
- Testes de concorrência para upload/download de arquivos

## Deploy
- Use Docker para buildar e rodar em produção
- Certifique-se de configurar variáveis de ambiente e RBAC corretamente
- Recomenda-se uso de HTTPS para upload/download

---

## Observações Importantes
- **Nunca altere a lógica de autenticação/JWT/RBAC sem autorização**
- **Toda alteração relevante deve ser justificada e documentada**
- **Arquivos de upload não são versionados (ver `.gitignore`)

---

Para detalhes de tabelas e regras de negócio, consulte `docs/tables.md` e `docs/regras.md`. 