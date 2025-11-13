FROM golang:1.24

LABEL authors="esnas"

# Устанавливаем зависимости
RUN apk add --no-cache gcc musl-dev git

WORKDIR /app

# Копируем зависимости и скачиваем их
COPY go.mod go.sum ./
RUN go mod download

# Копируем исходный код
COPY . .

# Собираем приложение
RUN go build -o main ./cmd/app

EXPOSE 8080

# Запускаем приложение
CMD ["./main"]