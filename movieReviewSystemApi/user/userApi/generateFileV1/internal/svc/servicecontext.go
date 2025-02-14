package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"movieReviewSystem/movieReviewSystemApi/user/userApi/generateFileV1/internal/config"
	"movieReviewSystem/movieReviewSystemApi/user/userApi/generateFileV1/internal/middleware"
	"movieReviewSystem/movieReviewSystemApi/user/userRpc/generateFileV1/userservice"
)

type ServiceContext struct {
	Config         config.Config
	Middleware     rest.Middleware
	UserRpcService userservice.UserService
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:         c,
		Middleware:     middleware.NewMiddleware().Handle,
		UserRpcService: userservice.NewUserService(zrpc.MustNewClient(c.UserRpcConf)),
	}
}
