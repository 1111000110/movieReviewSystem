package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"movieReviewSystem/movieReviewSystemApi/post/postApi/generateFileV1/internal/config"
	"movieReviewSystem/movieReviewSystemApi/post/postApi/generateFileV1/internal/middleware"
	model "movieReviewSystem/movieReviewSystemApi/post/postModel"
)

type ServiceContext struct {
	Config     config.Config
	Middleware rest.Middleware
	PostModel  model.PostModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		Middleware: middleware.NewMiddleware().Handle,
		PostModel:  model.NewPostModel(c.DB.Url, c.DB.Db, c.DB.Collection),
	}
}
