package model

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/mon"
	"go.mongodb.org/mongo-driver/bson"
)

var _ PostModel = (*customPostModel)(nil)

type (
	// PostModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPostModel.
	PostModel interface {
		postModel
		FindAll(ctx context.Context, offset, limit int64) (*[]Post, error)
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

func (c *customPostModel) FindAll(ctx context.Context, offset, limit int64) (*[]Post, error) {
	var err error
	var data []Post

	err = c.conn.Find(ctx, &data, bson.M{
		"limit":  limit,
		"offset": offset,
	})
	switch err {
	case nil:
		return &data, nil
	case mon.ErrNotFound:
		return &data, ErrNotFound
	default:
		return &data, err
	}
}
