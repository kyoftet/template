# TODO
## frontend
* workflowでfmtとlint
* crud実装

## TODO
* gorm ver
* no auth
* no redis

# How To Setup
```sh
cp .env.template .env
docker compose up -d --build
docker compose exec backend bash -c 'atlas migrate apply --dir "file://infrastructure/mysql/migrations" --url mysql://$MYSQL_USER:$MYSQL_PASSWORD@$MYSQL_HOST:$MYSQL_PORT/$MYSQL_DATABASE'
```

# Build For Production
```sh
cp .env.template .env
docker compose -f compose.prod.yaml up -d --build
```

# How to create model and migration file
```sh
# create table schema by go
docker compose exec backend go run -mod=mod entgo.io/ent/cmd/ent init ${Model}

### Edit ./infrastructure/mysql/ent/schema/${model}.go ###

# create ent file for database model
docker compose exec backend go generate ./ent
# create migration file
docker compose exec backend go run -mod=mod ./cmd/migration/main.go ${migration_file_name}
# migration
docker compose exec backend bash -c 'atlas migrate apply --dir "file://infrastructure/mysql/migrations" --url mysql://$MYSQL_USER:$MYSQL_PASSWORD@$MYSQL_HOST:$MYSQL_PORT/$MYSQL_DATABASE'
```
