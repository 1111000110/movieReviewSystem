package model

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/mon"
	"go.mongodb.org/mongo-driver/bson"
)

var _ UserModel = (*customUserModel)(nil)

type (
	// UserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserModel.
	UserModel interface {
		userModel
		FindOneByPhone(ctx context.Context, phone string) (*User, error)
	}

	customUserModel struct {
		*defaultUserModel
	}
)

func (m *customUserModel) FindOneByPhone(ctx context.Context, phone string) (*User, error) {
	var data User
	err := m.conn.FindOne(ctx, &data, bson.M{"phone": phone})
	switch err {
	case nil:
		return &data, nil
	case mon.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

// NewUserModel returns a model for the mongo.
func NewUserModel(url, db, collection string) UserModel {
	conn := mon.MustNewModel(url, db, collection)
	return &customUserModel{
		defaultUserModel: newDefaultUserModel(conn),
	}
}
