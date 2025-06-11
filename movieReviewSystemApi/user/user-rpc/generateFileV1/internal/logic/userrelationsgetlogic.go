package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"movieReviewSystem/movieReviewSystemApi/user/user-rpc/generateFileV1/internal/svc"
	"movieReviewSystem/movieReviewSystemApi/user/user-rpc/pb"
	relationsModel "movieReviewSystem/movieReviewSystemApi/user/userModel/mongo/userRelations"
)

type UserRelationsGetLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserRelationsGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRelationsGetLogic {
	return &UserRelationsGetLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserRelationsGetLogic) UserRelationsGet(in *__.UserRelationsGetReq) (*__.UserRelationsGetResp, error) {
	// todo: add your logic here and delete this line
	var data *relationsModel.UserRelations
	var err error
	data, err = l.svcCtx.UserRelationsModel.FindRelationsByUserIdAndOtherId(l.ctx, in.GetUserId(), in.GetOUserId())
	if err != nil || data == nil {
		return nil, err
	}
	UserRelations, err := __.ModelUserRelationsToUserRelations(*data)
	if err != nil {
		return nil, err
	}
	return &__.UserRelationsGetResp{
		Relations: &UserRelations,
	}, nil
}
