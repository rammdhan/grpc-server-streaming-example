package cache

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/redis/go-redis/v9"
)

type RedisInterface interface {
	Subscribe(ctx context.Context, channel string) (pubSub *redis.PubSub)
	Publish(ctx context.Context, channel string, message interface{}) (err error)
}

type redisClient struct {
	client *redis.Client
}

func NewRedis() (*redisClient, error) {

	redisServer := fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT_PUBLIC"))
	redisPassword := os.Getenv("REDIS_PASSWORD")
	redisDB, _ := strconv.Atoi(os.Getenv("REDIS_DATABASES"))

	client := redis.NewClient(&redis.Options{
		Addr:     redisServer,
		Password: redisPassword,
		DB:       redisDB,
	})

	return &redisClient{
		client: client,
	}, nil
}

func (r *redisClient) Subscribe(ctx context.Context, channel string) (pubSub *redis.PubSub) {
	return r.client.Subscribe(ctx, channel)
}

func (r *redisClient) Publish(ctx context.Context, channel string, message interface{}) (err error) {
	return r.client.Publish(ctx, channel, message).Err()
}
