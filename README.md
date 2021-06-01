# Golang API Example

## Project Dependencies

- [Golang](https://golang.org/)
- [Docker](https://www.docker.com/get-started) (this is for the DB)
- [golang-migrate](https://github.com/golang-migrate/migrate) - migration runner of choice
- [go-chi](https://github.com/go-chi/chi) - REST routing library
- [gorm](https://gorm.io/docs/) - Database integration (DOESN'T CONTROL THE SCHEMA)

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

### Common migration commands

All of the migrations for this project will have to be written by hand and be unlinked from the models. This has it's
positives and negatives, but so long as it's known about it shouldn't make too much of a difference.

Migrations are stored in the `migrations` directory and each will require a `*.down.sql` and a `*.up.sql` to allow for
reversions.

Because all of the migrations will need to be created by hand, we should follow a naming convention. I've chosen
`YYYYMMDD_000N.*.sql` for this example, but we can change that up if we wish to do so.

As the above example illustrates how to run all the migrations, we can also revert a single migration

```
docker run -v "$(PWD)"/migrations:/migrations --network host migrate/migrate -path=/migrations/ -database postgres://go_user:password@localhost:5432/go_api\?sslmode=disable down 1
```
