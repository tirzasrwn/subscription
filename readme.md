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

```sh
# Start the web app.
make start
# Stop the web app.
make stop
```

### Run Test

```sh
go test -coverprofile=coverage.out
go tool cover -html=coverage.out -o coverage.html
# Open coverage.html file with your browser to preview.
open coverage.html
```

## Ports

- Go Web App: 8080
- Mailhog: 8025
- Redis: 6379
- Postgres: 5432
