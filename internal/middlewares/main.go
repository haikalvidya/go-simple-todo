package middlewares

import (
	"github.com/haikalvidya/go-simple-todo/config"
	"github.com/haikalvidya/go-simple-todo/pkg/logger"
)

type CustomMiddleware struct {
	Logger logger.Logger
}

func New(cfg *config.Config) *CustomMiddleware {

	logger := logger.NewApiLogger(cfg)

	return &CustomMiddleware{
		Logger: logger,
	}
}
