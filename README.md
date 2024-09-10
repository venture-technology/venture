# venture

### migrations

- Please, install golang-migration

> linux
```bash
$ curl -L https://packagecloud.io/golang-migrate/migrate/gpgkey | apt-key add -
$ echo "deb https://packagecloud.io/golang-migrate/migrate/ubuntu/ $(lsb_release -sc) main" > /etc/apt/sources.list.d/migrate.list
$ apt-get update
$ apt-get install -y migrate
```

> Windows
```bash
$ scoop install migrate
```

- To create a new migration
```bash
migrate create -ext sql -dir database/migrations description-of-migration
```
> f.e: migrate create -ext sql -dir database/migrations add_profile_image_to_drivers

- To run your migration
```bash
migrate -path=database/migrations -database "postgres://user:password@localhost:5432/mydb?sslmode=disable" up
```

- If you need run rollback
```bash
migrate -path=database/migrations -database "postgres://user:password@localhost:5432/mydb?sslmode=disable" down
```

- Case, for some reason, need run migration from the beginning
> The same command to run migrations
