package Post

import (
	"context"

	"movieReviewSystem/movieReviewSystemApi/post/postApi/generateFileV1/internal/svc"
	"movieReviewSystem/movieReviewSystemApi/post/postApi/generateFileV1/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PostFindByPostIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPostFindByPostIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PostFindByPostIdLogic {
	return &PostFindByPostIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PostFindByPostIdLogic) PostFindByPostId(req *types.PostFindByPostIdReq) (resp *types.PostFindResp, err error) {
	// todo: add your logic here and delete this line

	return
}
