package logic

import (
	"context"

	"movieReviewSystem/movieReviewSystemApi/group/groupRpc/generateFileV1/internal/svc"
	"movieReviewSystem/movieReviewSystemApi/group/groupRpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GroupDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGroupDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupDeleteLogic {
	return &GroupDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GroupDeleteLogic) GroupDelete(in *__.GroupDeleteReq) (*__.GroupDeleteResp, error) {
	// todo: add your logic here and delete this line

	return &__.GroupDeleteResp{}, nil
}
