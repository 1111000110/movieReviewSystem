package logic

import (
	"context"

	"movieReviewSystem/movieReviewSystemApi/group/groupRpc/generateFileV1/internal/svc"
	"movieReviewSystem/movieReviewSystemApi/group/groupRpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GroupCreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGroupCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupCreateLogic {
	return &GroupCreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GroupCreateLogic) GroupCreate(in *__.GroupCreateReq) (*__.GroupCreateResp, error) {
	// todo: add your logic here and delete this line

	return &__.GroupCreateResp{}, nil
}
