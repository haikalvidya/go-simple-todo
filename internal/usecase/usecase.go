package usecase

import (
	"github.com/haikalvidya/go-simple-todo/config"
	"github.com/haikalvidya/go-simple-todo/internal/middlewares"
	"github.com/haikalvidya/go-simple-todo/internal/repository"

	"github.com/go-redis/redis"
)

type Usecase struct {
	Activity IActivityUsecase
	Todo     ITodoUsecase
}

type usecaseType struct {
	Repo        *repository.Repository
	Middleware  *middlewares.CustomMiddleware
	RedisClient *redis.Client
	ServerInfo  *config.ServerConfig
}

func NewUsecase(repo *repository.Repository, mid *middlewares.CustomMiddleware, redis *redis.Client, serverInfo *config.ServerConfig) *Usecase {
	usc := &usecaseType{Repo: repo, Middleware: mid, RedisClient: redis, ServerInfo: serverInfo}

	return &Usecase{
		Activity: (*activityUsecase)(usc),
		Todo:     (*todoUsecase)(usc),
	}
}
