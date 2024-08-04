# Running

If running locally and not using docker-compose, gcc will be needed to be accesible from your PATH to compile the SQLite driver\

The run command is to be run from the "src" directory

```sh
go run .
```

Currently, failure to connect to the analytics service is treated as a fatal error, so the analytics server should be run first.

# Environment Variables

The runtime can be customized with environment variables. If they are not present, the default setting will be used

- HTTP_PORT
    - The port for the HTTP API to listen on
- MIGRATIONS_DIR
    - Directory for migration files
- SQLITE_FILE
    - The SQLite database file
- ANALYTICS_HOST
    - Analytics service hostname
- ANALYTICS_PORT
    - Analytics service port
- CLEAR_DB_ON_STARTUP
    - Deletes the DB file on startup. Used because the Analytics service does not use a database yet

# API

The API for this application is an HTTP API with two routes

## GET /:linkId

Redirects to the link referenced by linkId

## POST /links/

Creates a shortened link

Example request body:
```json
{
    "link": "http://google.com"
}
```

# Testing

To run the tests, you can use this command from the "src" directory

```sh
go test -v ./...
```

Mocks for testing generated with [mockgen](https://github.com/golang/mock)

Mocks go in src/mocks, and belong to the "mocks" package.

# Migrations

Migrations are done via [golang-migrate](https://github.com/golang-migrate/migrate)
