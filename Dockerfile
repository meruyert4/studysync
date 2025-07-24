# syntax=docker/dockerfile:1

FROM golang:1.23 AS builder

# Работаем внутри user-service
WORKDIR /app

# Копируем только user-service
COPY user-service/ ./user-service/

# Переходим в папку с go.mod
WORKDIR /app/user-service

# Качаем зависимости
RUN go mod download

# Сборка бинарника
RUN go build -o service ./cmd/user

# Финальный образ
FROM gcr.io/distroless/base-debian12

WORKDIR /app

COPY --from=builder /app/user-service/service .

EXPOSE 8082
EXPOSE 50051

ENTRYPOINT ["/app/service"]