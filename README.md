# Gin Testing

## Migrations

### Setup Migrations Go

- **To install the migrate CLI tool using scoop on Windows, you can follow these steps:**

```bash
$ scoop install migrate
```

- **To install the migrate CLI tool using scoop on Mac, you can follow these steps:**

```bash
$ brew install golang-migrate
```

- **To install the migrate CLI tool using curl on Linux, you can follow these steps:**

```bash
$ curl -L https://packagecloud.io/golang-migrate/migrate/gpgkey| apt-key add -
$ echo "deb https://packagecloud.io/golang-migrate/migrate/ubuntu/ $(lsb_release -sc) main" > /etc/apt/sources.list.d/migrate.list
$ apt-get update
$ apt-get install -y migrate
```

1. Create migration files using the following command:

```bash
$ migrate create -ext sql -dir database/migration/ -seq init_mg
```

2. Command for migrations in windows :

```bash
migrate -database "postgres://postgres:password@localhost:5432/database_name?sslmode=disable&TimeZone=Asia/Jakarta" -path migration up
```
