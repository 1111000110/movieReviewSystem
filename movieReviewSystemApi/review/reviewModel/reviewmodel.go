package model

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/mon"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

var _ ReviewModel = (*customReviewModel)(nil)

type (
	// ReviewModel is an interface to be customized, add more methods here,
	// and implement the added methods in customReviewModel.
	ReviewModel interface {
		reviewModel
		FindByRootId(ctx context.Context, rootId int64) (*[]Review, error)
		FindByBaseId(ctx context.Context, baseId int64) (*[]Review, error)
		InsertMany(ctx context.Context, review *[]Review) (int64, error)
		UpdataRootcommitCountByReviewId(ctx context.Context, reviewId int64, addData int64) (*Review, error)
		DeleteByIds(ctx context.Context, reviewIds []int64) (int64, error)
		DeleteByBaseIds(ctx context.Context, baseIds []int64) (int64, error)
		DeleteByRootIds(ctx context.Context, rootIds []int64) (int64, error)
	}

	customReviewModel struct {
		*defaultReviewModel
	}
)

func (cm customReviewModel) DeleteByRootIds(ctx context.Context, rootIds []int64) (int64, error) {
	count, err := cm.conn.DeleteMany(ctx, bson.M{"rootId": bson.M{"$in": rootIds}}, nil)
	return count, err
}
func (cm customReviewModel) DeleteByBaseIds(ctx context.Context, baseIds []int64) (int64, error) {
	count, err := cm.conn.DeleteMany(ctx, bson.M{"baseId": bson.M{"$in": baseIds}}, nil)
	return count, err
}
func (cm customReviewModel) DeleteByIds(ctx context.Context, reviewIds []int64) (int64, error) {
	count, err := cm.conn.DeleteMany(ctx, bson.M{"_id": bson.M{"$in": reviewIds}}, nil)
	return count, err
}
func (cm customReviewModel) FindByRootId(ctx context.Context, rootId int64) (*[]Review, error) {
	data := make([]Review, 0)
	err := cm.conn.Find(ctx, &data, bson.M{"rootId": rootId})
	if len(data) == 0 {
		return nil, ErrNotFound
	}
	switch err {
	case nil:
		return &data, nil
	case mon.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
func (cm customReviewModel) UpdataRootcommitCountByReviewId(ctx context.Context, reviewId int64, count int64) (*Review, error) {
	data := &Review{}
	_, err := cm.conn.UpdateOne(ctx, bson.M{"_id": reviewId}, bson.M{"$inc": bson.M{"rootCommentCount": count}})
	if err != nil {
		return nil, err
	}
	err = cm.conn.FindOne(ctx, data, bson.M{"_id": reviewId})
	if err != nil {
		return nil, err
	}
	return data, nil

}
func (cm customReviewModel) FindByBaseId(ctx context.Context, baseId int64) (*[]Review, error) {
	data := make([]Review, 0)
	err := cm.conn.Find(ctx, &data, bson.M{"baseId": baseId})
	if len(data) == 0 {
		return nil, ErrNotFound
	}
	switch err {
	case nil:
		return &data, nil
	case mon.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
func (cm customReviewModel) InsertMany(ctx context.Context, review *[]Review) (int64, error) {
	data := make([]any, len(*review))
	for i := 0; i < len(*review); i++ {
		(*review)[i].UpdateAt = time.Now().Unix()
		(*review)[i].CreateAt = time.Now().Unix()
		data[i] = (*review)[i]
	}
	_, err := cm.conn.InsertMany(ctx, data)
	if err != nil {
		return 0, err
	}
	return int64(len(*review)), nil
}

// NewReviewModel returns a model for the mongo.
func NewReviewModel(url, db, collection string) ReviewModel {
	conn := mon.MustNewModel(url, db, collection)
	return &customReviewModel{
		defaultReviewModel: newDefaultReviewModel(conn),
	}
}
