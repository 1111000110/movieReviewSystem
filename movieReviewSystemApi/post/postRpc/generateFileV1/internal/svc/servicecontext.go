package svc

import (
	model "movieReviewSystem/movieReviewSystemApi/post/postModel"
	"movieReviewSystem/movieReviewSystemApi/post/postRpc/generateFileV1/internal/config"
)

type ServiceContext struct {
	Config     config.Config
	UsersModel model.PostModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		UsersModel: model.NewPostModel(c.DB.Url, c.DB.Db, c.DB.Collection),
	}
}
