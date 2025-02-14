package model

import "github.com/zeromicro/go-zero/core/stores/mon"

var _ PostModel = (*customPostModel)(nil)

type (
	// PostModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPostModel.
	PostModel interface {
		postModel
	}

	customPostModel struct {
		*defaultPostModel
	}
)

// NewPostModel returns a model for the mongo.
func NewPostModel(url, db, collection string) PostModel {
	conn := mon.MustNewModel(url, db, collection)
	return &customPostModel{
		defaultPostModel: newDefaultPostModel(conn),
	}
}
