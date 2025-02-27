package mysqlqueryservice

import (
	"backend/application/queryservice"
	"backend/infrastructure/mysql/ent"
	entUser "backend/infrastructure/mysql/ent/user"
	"context"
)

type infra struct {
	client *ent.Client
}

func NewAuth(client *ent.Client) queryservice.Auth {
	return &infra{client: client}
}

func (i *infra) GetPasswordByUserID(userID string) (string, error) {
	ctx := context.Background()
	user, err := i.client.User.Query().Where(entUser.UserID(userID)).Only(ctx)
	if err != nil {
		return "", nil
	}
	return user.Password, nil
}
