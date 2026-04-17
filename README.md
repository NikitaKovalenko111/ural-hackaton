# Ural Hackaton

Полноценный fullstack-проект для управления хабами, событиями, наставниками, запросами и бронированиями.

Проект состоит из:
- frontend на React + TypeScript + Vite
- backend на Go + Fiber
- PostgreSQL как основной БД

## Что умеет проект

- Просмотр списка хабов и страницы конкретного хаба
- Просмотр событий и наставников
- Регистрация/авторизация через magic link
- Работа с сущностями: users, mentors, admins, hubs, events, requests, bookings
- Swagger-страница для API

## Технологический стек

### Frontend

- React 19
- TypeScript
- Vite
- Redux Toolkit + Redux Saga
- React Router
- Axios
- SCSS

### Backend

- Go 1.25+
- Fiber v2
- PostgreSQL (driver: github.com/lib/pq)
- cleanenv + godotenv для конфигов

## Структура репозитория

```text
ural-hackaton/
├─ client/                  # Frontend (Vite + React + TS)
├─ server/                  # Backend (Go + Fiber)
├─ docs/                    # SQL и документация
├─ markup/                  # Статичная разметка/дизайн-референсы
└─ README.md
```

## Требования

- Node.js 20+
- npm 10+
- Go 1.25+
- PostgreSQL 14+

## Быстрый старт (локально)

### 1) Клонирование

```bash
git clone <repo-url>
cd ural-hackaton
```

### 2) Запуск PostgreSQL

Создайте базу:

```sql
CREATE DATABASE "ural-hackaton";
```

### 3) Настройка backend

Перейдите в папку backend:

```bash
cd server
```

Создайте файл `.env`:

```env
CONFIG_PATH=./config/local.yaml
```

Проверьте параметры в `server/config/local.yaml`:
- `storage.db_host`
- `storage.db_port`
- `storage.db_user`
- `storage.db_pass`
- `storage.db_name`

Установите зависимости и запустите сервер:

```bash
go mod download
go run ./cmd/ural_hackaton/main.go
```

Backend стартует на `http://localhost:3000`.

### 4) Настройка frontend

В новом терминале:

```bash
cd client
npm install
npm run dev
```

Frontend стартует на `http://localhost:5173`.

## Важное про запуск из корня

Если запускать backend из корня репозитория, путь будет другим:

```bash
go run ./server/cmd/ural_hackaton/main.go
```

Ошибка вида `stat ./cmd/ural_hackaton/main.go: no such file or directory` означает, что команда была выполнена не из директории `server`.

## Переменные окружения

### Backend

- `CONFIG_PATH` — путь к yaml-конфигу (обязательно)

### Frontend

- `VITE_API_BASE_URL` (опционально)

Если `VITE_API_BASE_URL` не задан, frontend использует runtime URL:

```text
{protocol}//{hostname}:3000
```

Например при открытом frontend на `http://localhost:5173` API будет `http://localhost:3000`.

## База данных

DDL-скрипт находится в `docs/database-create-query.md`.

При старте backend выполняет подготовку схемы (`storage.Prepare()`), поэтому таблицы создаются/мигрируются автоматически.

### Быстрое добавление тестовых данных через psql

Ниже пример подключения:

```bash
PGPASSWORD='<db_password>' psql -h localhost -p 5432 -U postgres -d 'ural-hackaton'
```

Далее вставляйте тестовые сущности (`hubs`, `users`, `mentors`, `admins`, `events`) SQL-командами.

## API и Swagger

- Swagger UI: `http://localhost:3000/swagger`
- OpenAPI YAML: `http://localhost:3000/swagger/openapi.yaml`

Ключевые группы эндпоинтов:
- `/auth/*`
- `/users/*`
- `/admins/*`
- `/mentors/*`
- `/hubs/*`
- `/events/*`
- `/requests/*`
- `/bookings/*`

## Frontend команды

```bash
cd client
npm run dev
npm run build
npm run preview
npm run lint
```

## Backend команды

```bash
cd server
go run ./cmd/ural_hackaton/main.go
go test ./...
```

## SCSS и предупреждения Sass

В проекте используются `@import` и `darken(...)`, поэтому Sass показывает deprecation warnings.

Это не ломает сборку, но в перспективе стоит перейти на:
- `@use`/`@forward`
- `color.adjust` или `color.scale`

## Частые проблемы

### 1) Хабы/события/менторы не видны на сайте

Проверка:

```bash
curl -i http://127.0.0.1:3000/hubs/
curl -i http://127.0.0.1:3000/events/
curl -i http://127.0.0.1:3000/mentors/
```

Если API возвращает 200 и JSON, проблема обычно в:
- не перезапущен frontend
- не перезапущен backend
- кэш браузера (нужен hard refresh)

### 2) Ошибка подключения к БД

Проверьте:
- что PostgreSQL запущен
- параметры `storage.*` в `server/config/local.yaml`
- что БД `ural-hackaton` существует

### 3) Ошибка CORS

Список разрешенных origin задается в backend (Fiber CORS middleware) и в `http_server.cors_origins`.

## Безопасность

- Не храните реальные секреты в репозитории.
- Для production используйте отдельные конфиги и секреты из CI/CD или secret manager.
- SMTP и auth-ключи должны быть заменены на ваши значения.

## Roadmap (рекомендуемо)

- Перевести SCSS с `@import` на `@use`
- Добавить миграции БД отдельным инструментом (например, goose)
- Добавить docker-compose для PostgreSQL + backend + frontend
- Расширить e2e/интеграционные тесты

## Лицензия

Добавьте лицензию проекта в отдельный файл `LICENSE`.
