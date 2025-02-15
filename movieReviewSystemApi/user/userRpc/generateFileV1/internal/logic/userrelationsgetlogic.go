package logic

import (
	"context"

	"movieReviewSystem/movieReviewSystemApi/user/userRpc/generateFileV1/internal/svc"
	"movieReviewSystem/movieReviewSystemApi/user/userRpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
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
	if in.GetUserId() < in.GetOUserId() {
		//l.svcCtx.UserRelationsModel.FindOne()
	}

	return &__.UserRelationsGetResp{}, nil
}
