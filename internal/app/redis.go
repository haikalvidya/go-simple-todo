package app

import (
	"github.com/haikalvidya/go-simple-todo/config"

	"github.com/go-redis/redis"
)

func InitRedis(cfg *config.Config) (*redis.Client, error) {
	redisConfig := &redis.Options{
		Addr:     cfg.Redis.Host + ":" + cfg.Redis.Port,
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
	}

	client := redis.NewClient(redisConfig)

	_, err := client.Ping().Result()
	if err != nil {
		return nil, err
	}

	return client, nil
}
