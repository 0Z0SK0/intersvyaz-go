# Intersvyaz Golang Test Task
RESTful API приложение для сбора метрики

> [!NOTE]
> Перед запуском убедитесь что выставлены данные подключения к PostgreSQL (их можно сменить в файле .env).

```bash
# Build and Run
cd intersvyaz-go-test
make
docker-compose build
docker-compose up

# Default API Endpoint : http://127.0.0.1:8000
```

## Structure
```
├── cmd
│   └── app
│       └── main.go          // Основное приложение
├── models
│   └── track.go             // Модель таблицы track
├── server
│   └── app.go               // Файл сервера + бд, здесь весь код роутинга, инициализации бд и тд тп
├── tests
│   └── main_test.go         // юнит-тесты
└── track
    ├── repository.go        // Здесь описан интерфейс репозитория трека
    ├── usecase.go           // Суть та же что и в файле выше
    ├── http
    │   └── delivery
    │       ├── handler.go   // Непосредственно контроллер для запросов к методу track
    │       └── register.go  // Простая обертка для создания роут групп к методу track
    ├── repository
    │   └── track.go         // Слой-обёртка для работы с репозиторием (бд)
    └── usecase
        └── usecase.go
```

## API

#### /track
* `POST` : Create a new track
