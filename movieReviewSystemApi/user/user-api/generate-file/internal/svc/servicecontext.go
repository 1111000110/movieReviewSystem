package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"movieReviewSystem/movieReviewSystemApi/user/user-api/generate-file/internal/config"
	"movieReviewSystem/movieReviewSystemApi/user/user-api/generate-file/internal/middleware"
	"movieReviewSystem/movieReviewSystemApi/user/user-rpc/generateFileV1/userservice"
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
