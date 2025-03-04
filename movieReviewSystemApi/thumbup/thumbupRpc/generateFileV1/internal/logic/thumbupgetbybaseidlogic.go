package logic

import (
	"context"

	"movieReviewSystem/movieReviewSystemApi/thumbup/thumbupRpc/generateFileV1/internal/svc"
	"movieReviewSystem/movieReviewSystemApi/thumbup/thumbupRpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ThumbUpGetByBaseIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewThumbUpGetByBaseIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ThumbUpGetByBaseIdLogic {
	return &ThumbUpGetByBaseIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ThumbUpGetByBaseIdLogic) ThumbUpGetByBaseId(in *__.ThumbUpGetByBaseIdReq) (*__.ThumbUpGetResp, error) {
	// todo: add your logic here and delete this line

	return &__.ThumbUpGetResp{}, nil
}
