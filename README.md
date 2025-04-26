# football-tables
Tables for football tournaments: groups, play-off, and games.

Structure of .env:
```bash
DB_HOST=localhost
DB_PORT=5432
DB_USER=football_user
DB_PASSWORD=football_tables_2025
DB_NAME=football_db

POSTGRES_SUPERUSER=postgres
POSTGRES_SUPERUSER_PASSWORD=AlexHSE2005

CORS_ALLOWED_ORIGINS=*
CORS_DEBUG=true

ENV=development

JWT_SECRET=
```

Скрипт, чтобы дать права пользователю на создание таблиц в БД:
```bash
\c postgres

GRANT ALL PRIVILEGES ON DATABASE football_db TO football_user;
ALTER USER football_user CREATEDB;

\c football_db

GRANT ALL PRIVILEGES ON SCHEMA public TO football_user;
GRANT USAGE ON SCHEMA public TO football_user;
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO football_user;
```

Удаление таблиц в БД:
```bash
psql -U football_user -d football_db -c "DROP TABLE IF EXISTS matches, teams, schema_migrations CASCADE;"
```

Миграция БД:
```bash
migrate -path ./backend/migrations -database 'postgres://football_user:football_tables_2025@localhost:5432/football_db?sslmode=disable' up
```

Инициализация Swagger-документации:
```bash
 swag init -g backend/cmd/server/main.go
```