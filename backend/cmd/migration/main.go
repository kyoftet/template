package main

import (
	"backend/infrastructure/mysql/ent/migrate"
	"backend/util"
	"context"
	"fmt"
	"log"
	"os"

	atlas "ariga.io/atlas/sql/migrate"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql/schema"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dataSourceName := fmt.Sprintf("mysql://%s:%s@%s:%s/%s?charset=utf8&parseTime=True",
		util.GetEnv("MYSQL_USER", "user"),
		util.GetEnv("MYSQL_PASSWORD", "password"),
		util.GetEnv("MYSQL_HOST", "localhost"),
		util.GetEnv("MYSQL_PORT", "3306"),
		util.GetEnv("MYSQL_NAME", "template"),
	)
	ctx := context.Background()

	dir, err := atlas.NewLocalDir("../../infrastructure/mysql/ent/migrate/migrations")
	if err != nil {
		log.Fatalf("failed creating atlas migration directory: %v", err)
	}

	opts := []schema.MigrateOption{
		schema.WithDir(dir),
		schema.WithMigrationMode(schema.ModeInspect),
		schema.WithDialect(dialect.MySQL),
		schema.WithFormatter(atlas.DefaultFormatter),
	}

	err = migrate.NamedDiff(ctx, dataSourceName, os.Args[1], opts...)
	if err != nil {
		log.Fatalf("failed generating migration file: %v", err)
	}
}
