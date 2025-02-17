package group

import (
	"context"

	"movieReviewSystem/movieReviewSystemApi/group/groupApi/generateFileV1/internal/svc"
	"movieReviewSystem/movieReviewSystemApi/group/groupApi/generateFileV1/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GroupUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGroupUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupUserLogic {
	return &GroupUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GroupUserLogic) GroupUser(req *types.GroupUserIdReq) (resp *types.GroupUserIdResp, err error) {
	// todo: add your logic here and delete this line

	return
}
