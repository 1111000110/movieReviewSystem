package logic

import (
	"context"

	"movieReviewSystem/movieReviewSystemApi/thumbup/thumbupRpc/generateFileV1/internal/svc"
	"movieReviewSystem/movieReviewSystemApi/thumbup/thumbupRpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ThumpUpCreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewThumpUpCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ThumpUpCreateLogic {
	return &ThumpUpCreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ThumpUpCreateLogic) ThumpUpCreate(in *__.ThumbUpCreateReq) (*__.ThumbUpCreateResp, error) {
	// todo: add your logic here and delete this line

	return &__.ThumbUpCreateResp{}, nil
}
