package redisconf

import "github.com/redis/go-redis/v9"

func GetDB() (*redis.Client, error) {
	dsn := "redis://user:password@localhost:6379/0?protocol=3"
	opts, err := redis.ParseURL(dsn)
	if err != nil {
		return nil, err
	}
	return redis.NewClient(opts), nil
}
