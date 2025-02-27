package di

import (
	"backend/application/service"
	"backend/infrastructure/mysql/ent"
	mysqlpersistance "backend/infrastructure/mysql/persistance"
	"backend/presentation/rest/controller"
)

func User(mysql *ent.Client) controller.User {
	applicationService := service.NewUser(
		mysqlpersistance.NewUser(mysql),
	)
	return controller.NewUser(applicationService)
}
