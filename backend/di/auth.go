package di

import (
	"backend/application/service"
	"backend/infrastructure/mysql/ent"
	mysqlqueryservice "backend/infrastructure/mysql/queryservice"
	redispersistance "backend/infrastructure/redis/persistance"
	"backend/presentation/api/router"

	"github.com/redis/go-redis/v9"
)

func Auth(mysql *ent.Client, redis *redis.Client) router.Auth {
	applicationService := service.NewAuth(
		redispersistance.NewAuth(redis),
		mysqlqueryservice.NewAuth(mysql),
	)
	return router.NewAuth(applicationService)
}
