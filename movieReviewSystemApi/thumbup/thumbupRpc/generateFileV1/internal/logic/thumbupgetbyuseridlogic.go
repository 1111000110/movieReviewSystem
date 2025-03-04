package logic

import (
	"context"

	"movieReviewSystem/movieReviewSystemApi/thumbup/thumbupRpc/generateFileV1/internal/svc"
	"movieReviewSystem/movieReviewSystemApi/thumbup/thumbupRpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ThumbUpGetByUserIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewThumbUpGetByUserIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ThumbUpGetByUserIdLogic {
	return &ThumbUpGetByUserIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ThumbUpGetByUserIdLogic) ThumbUpGetByUserId(in *__.ThumbUpGetByUserIdReq) (*__.ThumbUpGetResp, error) {
	// todo: add your logic here and delete this line

	return &__.ThumbUpGetResp{}, nil
}
