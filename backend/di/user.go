package di

import (
	"backend/application/service"
	"backend/infrastructure/mysql/ent"
	mysqlpersistance "backend/infrastructure/mysql/persistance"
	"backend/presentation/api/router"
)

func User(mysql *ent.Client) router.User {
	applicationService := service.NewUser(
		mysqlpersistance.NewUser(mysql),
	)
	return router.NewUser(applicationService)
}
