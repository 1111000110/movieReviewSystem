package svc

import (
	model "movieReviewSystem/movieReviewSystemApi/group/groupModel"
	"movieReviewSystem/movieReviewSystemApi/group/groupRpc/generateFileV1/internal/config"
)

type ServiceContext struct {
	Config     config.Config
	GroupModel model.GroupModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		GroupModel: model.NewGroupModel(c.DB.Group.Url, c.DB.Group.Db, c.DB.Group.Collection),
	}
}
