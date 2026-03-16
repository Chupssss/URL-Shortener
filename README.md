# URL Shortener Service

Простой сервис для сокращения ссылок, написанный на **Go**.\
Позволяет создавать короткие ссылки, перенаправлять пользователей и
отслеживать количество переходов.

---

# Технологии

-   Go
-   Gin
-   PostgreSQL
-   pgx
-   golang-migrate
-   Docker

---

# Архитектура проекта

    URL-Shortener
    │
    ├── cmd/
    │   └── api/                # точка входа приложения
    │
    ├── internal/
    │   ├── config/             # конфигурация приложения
    │   ├── handler/            # HTTP handlers
    │   ├── routers/            # настройка роутера
    │   ├── service/            # бизнес логика
    │   └── repos/              # работа с БД
    │
    ├── db/
    │   └── migrations/         # SQL миграции
    │
    ├── docs/
    │   └── api.md              # документация API
    │
    ├── Dockerfile              # инструкция для сборки образа (backend)
    ├── .env.example
    ├── go.mod
    └── README.md
