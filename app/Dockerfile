# ===============================
# Etapa 1 - Build da aplicação
# ===============================
FROM golang:1.26-alpine AS builder

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o http-server-projeto-korp .

# ===============================
# Etapa 2 - Imagem final
# ===============================
FROM alpine:3.22

WORKDIR /app

COPY --from=builder /app/http-server-projeto-korp .

RUN adduser -D appuser

USER appuser

EXPOSE 8080

CMD ["./http-server-projeto-korp"]