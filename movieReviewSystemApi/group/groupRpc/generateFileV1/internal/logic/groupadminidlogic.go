package logic

import (
	"context"

	"movieReviewSystem/movieReviewSystemApi/group/groupRpc/generateFileV1/internal/svc"
	"movieReviewSystem/movieReviewSystemApi/group/groupRpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GroupAdminIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGroupAdminIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupAdminIdLogic {
	return &GroupAdminIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GroupAdminIdLogic) GroupAdminId(in *__.GroupAdminIdReq) (*__.GroupAdminIdResp, error) {
	// todo: add your logic here and delete this line

	return &__.GroupAdminIdResp{}, nil
}
