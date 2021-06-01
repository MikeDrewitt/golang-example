# Golang API Example

## Project Dependencies

- [Golang](https://golang.org/)
- [Docker](https://www.docker.com/get-started) (this is for the DB)

Once you've got these installed, we can build our container and run it

```
docker run \
  --name go-postgres \
  -p 5432:5432 \
  -e POSTGRES_PASSWORD=password \
  -e POSTGRES_DB=go_api \
  -e POSTGRES_USER=go_user \
  -d postgres
```

Run your migrations

```
docker run -v "$(PWD)"/migrations:/migrations --network host migrate/migrate -path=/migrations/ -database postgres://go_user:password@localhost:5432/go_api\?sslmode=disable up
```

## Running the project

This project uses a `Makefile` to build and run the project.

We can run a local copy of the project with

```
make run
```
