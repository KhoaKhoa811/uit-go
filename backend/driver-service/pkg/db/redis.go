package db

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

type RedisHandler struct {
	Client *redis.Client
	Ctx    context.Context
}

func InitRedis(url string) RedisHandler {
	ctx := context.Background()

	opt, err := redis.ParseURL(url)

	if err != nil {
		panic(err)
	}

	rdb := redis.NewClient(opt)

	pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}

	log.Println("Connected to Redis successfully", pong)

	return RedisHandler{
		Client: rdb,
		Ctx:    ctx,
	}
}
