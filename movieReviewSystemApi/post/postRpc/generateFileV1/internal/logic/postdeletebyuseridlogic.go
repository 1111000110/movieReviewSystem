package logic

import (
	"context"

	"movieReviewSystem/movieReviewSystemApi/post/postRpc/generateFileV1/internal/svc"
	"movieReviewSystem/movieReviewSystemApi/post/postRpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type PostDeleteByUserIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPostDeleteByUserIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PostDeleteByUserIdLogic {
	return &PostDeleteByUserIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PostDeleteByUserIdLogic) PostDeleteByUserId(in *__.PostDeleteByUserIdReq) (*__.PostDeleteResp, error) {
	// todo: add your logic here and delete this line

	return &__.PostDeleteResp{}, nil
}
