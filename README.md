![poster](resourcer/poster.png)

# Weather_API

## Описание

Сервис, который предоставляет API для работы с информацией о погоде

## Стек

- `Go`
- `Postgres`
- Фреймворк [Gin](https://github.com/gin-gonic/gin)
- `Docker`
- Конфигурация приложения [cleanenv](https://github.com/ilyakaznacheev/cleanenv)
- Работа с БД [sqlx](https://github.com/jmoiron/sqlx)
- Логер [logrus](https://github.com/sirupsen/logrus)

# Чеклист работы над проектом

- [x] One-tap deployment with [docker-compose](#deployment)
- [x] Swagger API documentation
- [x] PostgreSQL
- [x] Auth
- [ ] Tests (very soon)

## Запуск

1. Добавить `.env` файл в проект:

``` 
POSTGRES_PASSWORD=postgres
API_KEY=8f7a589bb238eb36173737ebbe1ec8c6
```

2. Включить Docker

3. Открыть терминал и набрать:

```
make
```

По стандарту запускается цель `docker-compose`

## Тестирование

### Стандартный порт `8080`

- `/` - REST API
- `/swagger` - Swagger API

### Aутентификации

### Post

- `/sing-up` Регистрация

- `/sing-in` Вход

- `/api/v2/lists/` Создание списка

- `/api/v2/lists/:id/items/:city` Добавить определенный город в список

### Get

- `/api/v2/lists` Список всех списков пользователя

- `/api/v2/lists/:id` Информация о списке пользователя по id списка

- `/api/v2/lists/:id/items/` Информация о городах в списке клиента

### PUT

- `/api/v2/lists/:id` Обновление названия списка

### DELETE

- `/api/v2/lists/:id` Удаления списка по id

- `/api/v2/lists/:id/items/:item_id` Удаления города из списка по id

### Открытые методы

### Get

- `api/v2/cities` Список всех городов
- `api/v2/cities/:name` Прогноз погоды на 5 дней для определенного города
- `api/v2/detail_weather/:name/:date` Детальная информация по дню для определенного города

## Makefile (список целей)

```
make test  - запуска тестов
make lint - запуска линтера
```

## Swagger (список целей)

```
http://localhost:8080/swagger/index.html#/
```
