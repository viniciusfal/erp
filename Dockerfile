
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Copia os arquivos de dependência primeiro (para melhor cache)
COPY go.mod go.sum ./
RUN go mod download

# Copia o código fonte e compila
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /api ./cmd/main.go  

# Estágio final (imagem leve)
FROM alpine:latest

WORKDIR /root/

# Copia o binário compilado
COPY --from=builder /api ./

# Expõe a porta da API
EXPOSE 8000

# Comando de execução
CMD ["./api"]