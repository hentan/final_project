# HTTP Методы

## 1. Книги

Каждый запрос, кроме GET запросов на все книги или авторов возвращает 2 сущности: Status Code и json с 2 полями:

- error - true или false, случилась ли ошибка при обработке запроса
- message - сообщение об ошибке либо сообщение об успешной записи

### POST /books

Описание: Добавляет новую книгу в коллекцию.

Тело запроса:

```json
{
  "id": 3,
  "title": "Kashtanka",
  "author": "Антон Чехов",
  "year": 1887,
  "isbn": "975-7-759-06256-6"
}
```

Ответ:
201 Created: Книга успешно добавлена.

```json
{
  "error": false,
  "message": "Книга с id 5 успешно добавлена"
}
```

400 Bad Request: Некорректные данные в запросе.
возвращается json

```json
{
  "error": true,
  "message": "json: unknown field \"birthdy\""
}
```

### GET /books

Описание: Получает список всех книг в коллекции.

Ответ:
200 OK: Возвращает массив книг.
500 Internal Server Error: Ошибка на сервере.

### GET /books/{id}

Описание: Получает информацию о книге по её уникальному идентификатору.

Ответ:
200 OK: Возвращает данные книги.
404 Not Found: Книга с указанным идентификатором не найдена.

```json
{
  "error": true,
  "message": "sql: no rows in result set"
}
```

### PUT /books/{id}

Описание: Обновляет информацию о книге по её уникальному идентификатору.

Тело запроса:

```json
{
  "title": "Kashtanka",
  "author_id": 3,
  "year": 1897,
  "isbn": "975-7-759-06256-6"
}
```

Ответ:
200 OK: Книга успешно обновлена.

```json
{
  "error": false,
  "message": "Книга c id 3 успешно обновлена"
}
```

400 Bad Request: Некорректные данные в запросе.

```json
{
  "error": true,
  "message": "json: unknown field \"author\""
}
```

404 Not Found: Книга с указанным идентификатором не найдена.

```json
{
  "error": true,
  "message": "sql: no rows in result set"
}
```

### DELETE /books/{id}

Описание: Удаляет книгу по её уникальному идентификатору.

Ответ:
200 OK: Книга успешно удалена.

```json
{
  "error": false,
  "message": "Книга успешно удалена"
}
```

404 Not Found: Книга с указанным идентификатором не найдена.

```json
{
  "error": true,
  "message": "sql: no rows in result set"
}
```

## 2. Авторы

### POST /authors

Описание: Добавляет нового автора в коллекцию.

Тело запроса:

```json
{
  "name_author": "Антон",
  "sirname_author": "Чехов",
  "biography": "Русский писатель, драматург и врач.",
  "birthday": "1860-01-29"
}
```

Ответ:
201 Created: Автор успешно добавлен.

```json
{
  "error": false,
  "message": "Автор с id 9 успешно добавлен"
}
```

400 Bad Request: Некорректные данные в запросе.

```json
{
  "error": true,
  "message": "json: unknown field \"bivography\""
}
```

### GET /authors

Описание: Получает список всех авторов в коллекции.

Ответ:
200 OK: Возвращает массив авторов.
500 Internal Server Error: Ошибка на сервере.

### GET /authors/{id}

Описание: Получает информацию об авторе по его уникальному идентификатору.

Ответ:
200 OK: Возвращает данные автора.
404 Not Found: Автор с указанным идентификатором не найден.

```json
{
  "error": true,
  "message": "sql: no rows in result set"
}
```

### PUT /authors/{id}

Описание: Обновляет информацию об авторе по его уникальному идентификатору.

Тело запроса:

```json
{
  "name_author": "Антон",
  "sirname_author": "Чехов",
  "biography": "Русский писатель, драматург и врач.",
  "birthday": "1890-01-29"
}
```

Ответ:
200 OK: Автор успешно обновлен.

```json
{
  "error": false,
  "message": "Автор успешно обновлен"
}
```

400 Bad Request: Некорректные данные в запросе.

```json
{
  "error": true,
  "message": "json: unknown field \"birthdsay\""
}
```

404 Not Found: Автор с указанным идентификатором не найден.

```json
{
  "error": true,
  "message": "sql: no rows in result set"
}
```

### DELETE /authors/{id}

Описание: Удаляет автора по его уникальному идентификатору.

Ответ:
200 OK: Автор успешно удален.

```json
{
  "error": false,
  "message": "Автор с id 7 успешно удален"
}
```

404 Not Found: Автор с указанным идентификатором не найден.
{
"error": true,
"message": "sql: no rows in result set"
}

## 3. Транзакционное обновление

### PUT /books/{book_id}/authors/{author_id}

Описание: Одновременно обновляет сведения о книге и авторе.

Тело запроса:

```json
{
  "author": {
    "name_author": "Антон",
    "sirname_author": "Чехов",
    "biography": "Русский писатель, драматург и врач.",
    "birthday": "1890-01-29"
  },
  "book": {
    "title": "Kashtanka",
    "author_id": 3,
    "year": 1887,
    "isbn": "975-7-759-06256-6"
  }
}
```

Ответ:
200 OK: Книга и автор успешно обновлены.

```json
{
  "error": false,
  "message": "Автор и книга успешно обновлены"
}
```

400 Bad Request: Некорректные данные в запросе.

```json
{
  "error": true,
  "message": "json: unknown field \"nam_author\""
}
```

404 Not Found: Книга или автор с указанными идентификаторами не найдены.

```json
{
  "error": true,
  "message": "sql: no rows in result set"
}
```
