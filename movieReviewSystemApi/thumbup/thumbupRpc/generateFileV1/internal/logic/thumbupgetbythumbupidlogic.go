package logic

import (
	"context"

	"movieReviewSystem/movieReviewSystemApi/thumbup/thumbupRpc/generateFileV1/internal/svc"
	"movieReviewSystem/movieReviewSystemApi/thumbup/thumbupRpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ThumbUpGetByThumbUpIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewThumbUpGetByThumbUpIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ThumbUpGetByThumbUpIdLogic {
	return &ThumbUpGetByThumbUpIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ThumbUpGetByThumbUpIdLogic) ThumbUpGetByThumbUpId(in *__.ThumbUpGetByThumbUpIdReq) (*__.ThumbUpGetResp, error) {
	// todo: add your logic here and delete this line

	return &__.ThumbUpGetResp{}, nil
}
