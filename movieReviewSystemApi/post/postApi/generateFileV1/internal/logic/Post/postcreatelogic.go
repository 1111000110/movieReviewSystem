package Post

import (
	"context"

	"movieReviewSystem/movieReviewSystemApi/post/postApi/generateFileV1/internal/svc"
	"movieReviewSystem/movieReviewSystemApi/post/postApi/generateFileV1/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PostCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPostCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PostCreateLogic {
	return &PostCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PostCreateLogic) PostCreate(req *types.PostCreateReq) (resp *types.PostCreateResp, err error) {
	// todo: add your logic here and delete this line

	return
}
