package logic

import (
	"context"
	model "movieReviewSystem/movieReviewSystemApi/post/postModel"

	"movieReviewSystem/movieReviewSystemApi/post/postRpc/generateFileV1/internal/svc"
	"movieReviewSystem/movieReviewSystemApi/post/postRpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type PostCreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPostCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PostCreateLogic {
	return &PostCreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PostCreateLogic) PostCreate(in *__.PostCreateReq) (*__.PostCreateResp, error) {
	// todo: add your logic here and delete this line
	l.svcCtx.UsersModel.Insert(l.ctx, &model.Post{})
	return &__.PostCreateResp{}, nil
}
