FROM golang:1.21-alpine

LABEL authors="esnas"

# Устанавливаем зависимости для компиляции
RUN apk add --no-cache gcc musl-dev git

WORKDIR /app

# Копируем файлы модулей
COPY go.mod go.sum ./
RUN go mod download

# Копируем исходный код
COPY . .

# Собираем приложение
RUN go build -o main ./cmd/app

EXPOSE 8080

CMD ["./main"]