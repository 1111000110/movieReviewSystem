package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"movieReviewSystem/movieReviewSystemApi/movieInformation/movieInformationApi/generateFileV1/internal/config"
	"movieReviewSystem/movieReviewSystemApi/movieInformation/movieInformationApi/generateFileV1/internal/middleware"
)

type ServiceContext struct {
	Config     config.Config
	Middleware rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		Middleware: middleware.NewMiddleware().Handle,
	}
}
