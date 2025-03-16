# Используем официальный образ Go с поддержкой cgo
FROM golang:1.21 AS builder

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем файлы модулей и загружаем зависимости
COPY go.mod go.sum ./
RUN go mod download

# Копируем исходный код
COPY . .

# Собираем приложение
RUN go build -o sre-app .

# Используем образ с поддержкой cgo
FROM debian:bookworm-slim

# Устанавливаем рабочую директорию
WORKDIR /app

# Устанавливаем необходимые зависимости для SQLite
RUN apt-get update && apt-get install -y sqlite3

# Копируем собранное приложение из builder
COPY --from=builder /app/sre-app .

# Создаем non-root пользователя
RUN useradd -m appuser
USER appuser

# Открываем порт, на котором работает приложение
EXPOSE 8080

# Добавляем health-check
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
  CMD curl -f http://localhost:8080/health || exit 1

# Запускаем приложение
CMD ["./sre-app"]
