package logic

import (
	"context"

	"movieReviewSystem/movieReviewSystemApi/post/postRpc/generateFileV1/internal/svc"
	"movieReviewSystem/movieReviewSystemApi/post/postRpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type PostFindByPostIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPostFindByPostIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PostFindByPostIdLogic {
	return &PostFindByPostIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PostFindByPostIdLogic) PostFindByPostId(in *__.PostFindByPostIdReq) (*__.PostFindResp, error) {
	// todo: add your logic here and delete this line

	return &__.PostFindResp{}, nil
}
