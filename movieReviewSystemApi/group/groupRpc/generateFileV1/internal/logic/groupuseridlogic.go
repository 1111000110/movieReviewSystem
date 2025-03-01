package logic

import (
	"context"

	"movieReviewSystem/movieReviewSystemApi/group/groupRpc/generateFileV1/internal/svc"
	"movieReviewSystem/movieReviewSystemApi/group/groupRpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GroupUserIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGroupUserIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupUserIdLogic {
	return &GroupUserIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GroupUserIdLogic) GroupUserId(in *__.GroupUserIdReq) (*__.GroupUserIdResp, error) {
	// todo: add your logic here and delete this line

	return &__.GroupUserIdResp{}, nil
}
