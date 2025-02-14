package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"movieReviewSystem/movieReviewSystemApi/chat/generateFileV1/internal/config"
	"movieReviewSystem/movieReviewSystemApi/chat/generateFileV1/internal/hub"
	"movieReviewSystem/movieReviewSystemApi/chat/generateFileV1/internal/middleware"
)

type ServiceContext struct {
	Config     config.Config
	Middleware rest.Middleware
	Hub        hub.Hub
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		Middleware: middleware.NewMiddleware().Handle,
		Hub:        hub.NewHub(),
	}
}
