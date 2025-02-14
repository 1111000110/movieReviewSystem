package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"movieReviewSystem/movieReviewSystemApi/user/userModel"
	"movieReviewSystem/movieReviewSystemApi/user/userRpc/generateFileV1/internal/config"
)

type ServiceContext struct {
	Config     config.Config
	UsersModel userModel.UsersModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		UsersModel: userModel.NewUsersModel(sqlx.NewMysql(c.DB.DataSource)),
	}
}
