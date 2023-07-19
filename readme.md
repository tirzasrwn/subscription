# subscription

## About

Implementing concurrency in real world example with go.

## Account Info

## Web App

```
user: admin@example.com
password: verysecret
```

### Postgres

```
USER: postgres
PASSWORD: password
DB: concurrency
```

## Instalation

### Docker

```sh
# Start docker compose for postgres, redis, and mailhog.
docker compose up -d
# Stop docker compose.
docker compose down
```

### Running Web App

```
# Start the web app.
make start
# Stop the web app.
make stop
```

## Ports

- Go Web App: 8080
- Mailhog: 8025
- Redis: 6379
- Postgres: 5432
