package model

import "github.com/zeromicro/go-zero/core/stores/mon"

var _ UserRealtiosnModel = (*customUserRealtiosnModel)(nil)

type (
	// UserRealtiosnModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserRealtiosnModel.
	UserRealtiosnModel interface {
		userRealtiosnModel
	}

	customUserRealtiosnModel struct {
		*defaultUserRealtiosnModel
	}
)

// NewUserRealtiosnModel returns a model for the mongo.
func NewUserRealtiosnModel(url, db, collection string) UserRealtiosnModel {
	conn := mon.MustNewModel(url, db, collection)
	return &customUserRealtiosnModel{
		defaultUserRealtiosnModel: newDefaultUserRealtiosnModel(conn),
	}
}
