package model

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/mon"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

var _ UserRelationsModel = (*customUserRelationsModel)(nil)

type (
	// UserRelationsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserRelationsModel.
	UserRelationsModel interface {
		userRelationsModel
		FindRelationsByUserIdAndOtherId(ctx context.Context, UserId, OtherId int64) (*UserRelations, error)
		Upsert(ctx context.Context, data *UserRelations) (*mongo.UpdateResult, error)
	}

	customUserRelationsModel struct {
		*defaultUserRelationsModel
	}
)

func (m *customUserRelationsModel) FindRelationsByUserIdAndOtherId(ctx context.Context, UserId, OtherId int64) (*UserRelations, error) {
	var data UserRelations
	var err error
	err = m.conn.FindOne(ctx, &data, bson.M{"userId": UserId, "otherId": OtherId})
	switch err {
	case nil:
		return &data, nil
	case mon.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
func (m *customUserRelationsModel) Upsert(ctx context.Context, data *UserRelations) (*mongo.UpdateResult, error) {
	data.UpdateAt = time.Now().Unix()
	opts := options.Update().SetUpsert(true) // 启用 upsert
	res, err := m.conn.UpdateOne(ctx, bson.M{"userId": data.UserId, "otherId": data.OtherId}, bson.M{"$set": data}, opts)
	return res, err
}

// NewUserRelationsModel returns a model for the mongo.
func NewUserRelationsModel(url, db, collection string) UserRelationsModel {
	conn := mon.MustNewModel(url, db, collection)
	return &customUserRelationsModel{
		defaultUserRelationsModel: newDefaultUserRelationsModel(conn),
	}
}
