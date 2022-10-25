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

## Запуск

1. Добавить `.env` файл в проект:

``` 
    /* 
        Добавить в .env файл
    */

    POSTGRES_PASSWORD=postgres
    API_KEY=8f7a589bb238eb36173737ebbe1ec8c6
```

2. Включить Docker

3. Открыть терминал и набрать:

```
    $ make
```

По стандарту запускается цель `docker-compose`

## Тестирование

```
http://localhost:8080/api/v1/cities
```

### Get

- `/cities` Список всех городов
- `/cities/:name` Прогноз погоды на 5 дней для определенного города
- `/detail_weather/:name/:date` Детальная информация по дню для определенного города

## Makefile (список целей)

```
$ make test  - запуска тестов
$ make lint - запуска линтера
```

## Swagger (список целей)

```
http://localhost:8080/swagger/index.html#/
```
