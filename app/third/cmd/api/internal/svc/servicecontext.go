package svc

import (
	"github.com/hibiken/asynq"
	"github.com/zeromicro/go-zero/rest"
	"jekka-api-go/app/third/cmd/api/internal/config"
	"jekka-api-go/app/third/cmd/api/internal/middleware"
)

type ServiceContext struct {
	Config          config.Config
	AuthInterceptor rest.Middleware
	AsynqClient     *asynq.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:          c,
		AuthInterceptor: middleware.NewAuthInterceptorMiddleware().Handle,
		AsynqClient:     NewAsynqClient(c.Redis.Host, c.Redis.Pass, c.Redis.DB),
	}
}
