package mysqlpersistance

import (
	"backend/domain/entity"
	"backend/domain/repository"
	"backend/infrastructure/mysql/ent"
	entUser "backend/infrastructure/mysql/ent/user"
	"context"
)

type infra struct {
	client *ent.Client
}

func NewUser(client *ent.Client) repository.User {
	return &infra{client: client}
}

func (i *infra) Insert(user *entity.User) error {
	ctx := context.Background()
	_, err := i.client.User.Create().
		SetUserID(user.ID).
		SetName(user.Name).
		SetPassword(user.Password.Hash()).
		SetMailAddress(user.MailAddress).
		SetIntroduction(user.Introduction).
		Save(ctx)
	return err
}

func (i *infra) Delete(userId string) error {
	ctx := context.Background()
	user, err := i.client.User.Query().Where(entUser.UserID(userId)).Only(ctx)
	if err != nil {
		return err
	}
	err = i.client.User.DeleteOneID(user.ID).Exec(ctx)
	return err
}

func (i *infra) Update(user *entity.User) error {
	ctx := context.Background()
	storedUser, err := i.client.User.Query().Where(entUser.UserID(user.ID)).Only(ctx)
	if err != nil {
		return err
	}
	_, err = i.client.User.UpdateOneID(storedUser.ID).
		SetName(user.Name).
		SetMailAddress(user.MailAddress).
		SetIntroduction(user.Introduction).
		Save(ctx)
	return err
}
