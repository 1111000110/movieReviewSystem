package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"movieReviewSystem/movieReviewSystemApi/movieInformation/model"
	"movieReviewSystem/movieReviewSystemApi/movieInformation/movieInformationApi/generateFileV1/internal/config"
	"movieReviewSystem/movieReviewSystemApi/movieInformation/movieInformationApi/generateFileV1/internal/middleware"
)

type ServiceContext struct {
	Config     config.Config
	Middleware rest.Middleware
	Movie      model.MovieModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		Middleware: middleware.NewMiddleware().Handle,
		Movie:      model.NewMovieModel(c.DB.Url, c.DB.Db, c.DB.Collection),
	}
}
