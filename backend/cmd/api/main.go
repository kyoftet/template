package main

import (
	"log"

	mysqlconf "backend/infrastructure/mysql/conf"
	redisconf "backend/infrastructure/redis/conf"

	"backend/di"
	"backend/presentation/api"
)

func main() {
	mysqlDB, err := mysqlconf.GetDB()
	if err != nil {
		log.Fatal(err)
	}
	defer mysqlDB.Close()

	redisDB, err := redisconf.GetDB()
	if err != nil {
		log.Fatal(err)
	}
	defer redisDB.Close()

	api.InitRoute(
		di.User(mysqlDB),
		di.Auth(mysqlDB, redisDB),
		di.Sample(),
	)
}
