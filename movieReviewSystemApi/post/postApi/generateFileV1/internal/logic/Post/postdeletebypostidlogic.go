package Post

import (
	"context"

	"movieReviewSystem/movieReviewSystemApi/post/postApi/generateFileV1/internal/svc"
	"movieReviewSystem/movieReviewSystemApi/post/postApi/generateFileV1/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PostDeleteByPostIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPostDeleteByPostIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PostDeleteByPostIdLogic {
	return &PostDeleteByPostIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PostDeleteByPostIdLogic) PostDeleteByPostId(req *types.PostDeleteByPostIdReq) (resp *types.PostDeleteResp, err error) {
	// todo: add your logic here and delete this line

	return
}
