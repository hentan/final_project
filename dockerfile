# Используем конкретную версию официального изображения Go в качестве базового
FROM golang:1.22

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app

# Копируем go.mod и go.sum файлы в контейнер
COPY go.mod go.sum ./

# Загружаем все зависимости. Кэшируется, если go.mod и go.sum не изменяются
RUN go mod download

# Копируем исходный код в контейнер
COPY . .

# Копируем файлы миграций в контейнер
COPY ./db ./db

#копируем docs
COPY ./docs ./docs

# Устанавливаем утилиту для миграций
RUN go install -tags "postgres" github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# Сборка Go-приложения
RUN go build -o main ./cmd

# Выполняем миграции и запускаем приложение
CMD ["sh", "-c", "migrate -path ./db/migration -database $DATABASE_URL up && ./main"]
