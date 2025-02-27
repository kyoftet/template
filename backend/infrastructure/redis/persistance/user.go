package redispersistance

import (
	"backend/domain/entity"
	"backend/domain/repository"
	"context"

	"github.com/redis/go-redis/v9"
)

type infra struct {
	client *redis.Client
}

func NewAuth(client *redis.Client) repository.Auth {
	return &infra{client: client}
}

func (i *infra) Set(auth *entity.Auth) error {
	ctx := context.Background()
	return i.client.Set(
		ctx,
		auth.SessionID.Value(),
		auth.UserID,
		auth.DeadLine,
	).Err()
}

func (i *infra) Delete(sessionId string) error {
	ctx := context.Background()
	return i.client.Del(ctx, sessionId).Err()
}

func (i *infra) Get(sessionId string) (string, error) {
	ctx := context.Background()
	return i.client.Get(ctx, sessionId).Result()
}
