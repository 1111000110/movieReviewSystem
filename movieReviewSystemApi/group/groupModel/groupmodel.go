package model

import "github.com/zeromicro/go-zero/core/stores/mon"

var _ GroupModel = (*customGroupModel)(nil)

type (
	// GroupModel is an interface to be customized, add more methods here,
	// and implement the added methods in customGroupModel.
	GroupModel interface {
		groupModel
	}

	customGroupModel struct {
		*defaultGroupModel
	}
)

// NewGroupModel returns a model for the mongo.
func NewGroupModel(url, db, collection string) GroupModel {
	conn := mon.MustNewModel(url, db, collection)
	return &customGroupModel{
		defaultGroupModel: newDefaultGroupModel(conn),
	}
}
