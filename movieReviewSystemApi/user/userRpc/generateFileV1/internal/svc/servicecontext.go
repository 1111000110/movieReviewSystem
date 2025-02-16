package svc

import (
	"movieReviewSystem/movieReviewSystemApi/user/userModel/mongo/user"
	relationsModel "movieReviewSystem/movieReviewSystemApi/user/userModel/mongo/userRelations"
	"movieReviewSystem/movieReviewSystemApi/user/userRpc/generateFileV1/internal/config"
)

type ServiceContext struct {
	Config             config.Config
	UsersModel         model.UserModel
	UserRelationsModel relationsModel.UserRelationsModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:             c,
		UsersModel:         model.NewUserModel(c.DB.User.Url, c.DB.User.Db, c.DB.User.Collection),
		UserRelationsModel: relationsModel.NewUserRelationsModel(c.DB.UserRelations.Url, c.DB.UserRelations.Db, c.DB.UserRelations.Collection),
	}
}
