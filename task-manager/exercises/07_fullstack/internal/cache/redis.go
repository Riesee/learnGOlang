package cache

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

var Client *redis.Client
var Ctx = context.Background()

func Connect(host, port string) error {
	Client = redis.NewClient(&redis.Options{
		Addr: host + ":" + port,
	})

	_, err := Client.Ping(Ctx).Result()
	return err
}

func Set(key string, value string, expiration time.Duration) error {
	return Client.Set(Ctx, key, value, expiration).Err()
}

func Get(key string) (string, error) {
	return Client.Get(Ctx, key).Result()
}

func Delete(key string) error {
	return Client.Del(Ctx, key).Err()
}
