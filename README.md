# graphql-articles

Articles GraphQL Service

Requierements:

- libvips

## Install PostgreSQL

```bash
$ docker run -p 5432:5432 --name postgres -e POSTGRES_DB=articles -e POSTGRES_HOST_AUTH_METHOD=trust -d postgres:latest
```

## Enter in PostgreSQL

```bash
$ docker exec -it postgres psql -U postgres articles
```

## How to build

```shell
make build
```

## How to apply DB migrations

```bash
$ migration update postgres://postgres:@localhost:5432/articles\?sslmode=disable
```

## How to run app

```bash
$ app postgres://postgres:@postgres:5432/articles\?sslmode=disable
```
