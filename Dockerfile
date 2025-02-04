# Etapa 1: Build do aplicativo
FROM golang:1.21-alpine AS builder

# Instalar dependências necessárias para build (como git)
RUN apk add --no-cache git

# Definir o diretório de trabalho dentro do container
WORKDIR /app

# Copiar os arquivos de dependência e baixar as dependências
COPY go.mod go.sum ./
RUN go mod download

# Copiar o código-fonte para o diretório de trabalho
COPY . .

# Compilar o aplicativo
RUN go build -o main ./cmd

# Etapa 2: Imagem final para execução
FROM alpine:latest

# Instalar dependências para execução
RUN apk add --no-cache ca-certificates bash curl

# Definir o diretório de trabalho no container
WORKDIR /app

# Copiar o binário compilado da etapa de build
COPY --from=builder /app/main .

COPY wait-for-it.sh /wait-for-it.sh
RUN chmod +x /wait-for-it.sh


# Expor a porta que o aplicativo irá rodar
EXPOSE 8000

# Comando para rodar o aplicativo após o banco estar pronto
CMD ["/wait-for-it.sh", "db:5432", "--", "./main"]
