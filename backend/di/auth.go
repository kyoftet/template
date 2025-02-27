package di

import (
	"backend/application/service"
	"backend/infrastructure/mysql/ent"
	mysqlqueryservice "backend/infrastructure/mysql/queryservice"
	redispersistance "backend/infrastructure/redis/persistance"
	"backend/presentation/rest/controller"

	"github.com/redis/go-redis/v9"
)

func Auth(mysql *ent.Client, redis *redis.Client) controller.Auth {
	applicationService := service.NewAuth(
		redispersistance.NewAuth(redis),
		mysqlqueryservice.NewAuth(mysql),
	)
	return controller.NewAuth(applicationService)
}
