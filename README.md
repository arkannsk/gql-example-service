# Перед началом

### Подготовка файлов, генерация ключей
```bash
make setup
```

### Установка необходимых приложений
### [dbmate](https://github.com/amacneil/dbmate#installation "dbmate")
### [gqlgen](https://github.com/99designs/gqlgen#quick-start "gqlgen")


## Основные команды

```bash
# Поднимает нужные контейнеры для проекта
docker/up
# Опустить контейнеры
docker/down
# запустить приложение
make run
# заново cгенерировать файлы на основе graphql-схемы.
make gen-graphql

# ОБЯЗАТЕЛЬНО из корня проекта
# создание миграций
dbmate new {migration_name}
# применить миграций
dbmate -e POSTGRES_URL --no-dump-schema up
# откатить миграций
dbmate -e POSTGRES_URL --no-dump-schema down

#залить немного данных
cat seed.sql | psql -h localhost -p 7232 -d postgres -U staging
```
