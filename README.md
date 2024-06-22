# Bookstore API

Это REST API для управления коллекцией книг и авторами, которые их пишут, созданное на языке Go и использующее PostgreSQL для хранения данных. Приложение контейнеризировано с помощью Docker.

## Структура проекта

- `cmd/main.go`: Главный файл приложения.
- `interal/config`: чтение файла конфигурации.
- `interal/handlers/handlers.go`: обработчики HTTP-запросов.
- `interal/handlers/routes.go`: маршрутизация.
- `interal/handlers/json_handlers`: вспомогательные функции для обработки запроса/ответа/возвращения ошибок
- `internal/models`: Модели для работы с json
- `internal/repository`: Подключение и работа с базой данных
- `db/migrations`: SQL файлы для миграции базы данных.
- `Dockerfile`: Инструкции для сборки Docker-образа.
- `docker-compose.yml`: Конфигурация Docker Compose.

## Требования

- Go 1.20
- Docker
- Docker Compose

## Инструкции по сборке и запуску

### Сборка и запуск

Используйте `Makefile` для автоматизации процесса сборки и запуска.

1. Сборка приложения и запуск без логов:

```sh
   make all
```

2. Сборка приложения:

```sh
   make build
```

3. Запуск приложения:

```sh
   make up
```

4. Остановка и удаление контейнеров:

```sh
   make down
```

5. Просмотр логов сборки:

```sh
   make logs
```

### Проверка состояния контейнеров

Для проверки состояния контейнеров используйте команду:

```sh
   make ps
```

Использование
API предоставляет следующие эндпоинты:

Книги
POST /books: Добавить новую книгу
GET /books: Получить все книги
GET /books/{id}: Получить книгу по ее идентификатору
PUT /books/{id}: Обновить книгу по ее идентификатору
DELETE /books/{id}: Удалить книгу по ее идентификатору

Авторы
POST /authors: Добавить нового автора
GET /authors: Получить всех авторов
GET /authors/{id}: Получить автора по его идентификатору
PUT /authors/{id}: Обновить автора по его идентификатору
DELETE /authors/{id}: Удалить автора по его идентификатору

Транзакционное обновление
PUT /books/{book_id}/authors/{author_id}: Одновременно обновить сведения о книге и авторе

Более подробное описание методов в файле [Документация API методов](Документация_API_методов.md)
