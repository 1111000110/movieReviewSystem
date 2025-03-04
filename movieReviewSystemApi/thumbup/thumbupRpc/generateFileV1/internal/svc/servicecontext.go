package svc

import (
	"movieReviewSystem/movieReviewSystemApi/thumbup/thumbupModel/mysql"
	"movieReviewSystem/movieReviewSystemApi/thumbup/thumbupRpc/generateFileV1/internal/config"
)

type ServiceContext struct {
	Config       config.Config
	ThumbUpModel []mysql.ThumbUpModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
