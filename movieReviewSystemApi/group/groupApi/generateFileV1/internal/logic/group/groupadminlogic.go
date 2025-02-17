package group

import (
	"context"

	"movieReviewSystem/movieReviewSystemApi/group/groupApi/generateFileV1/internal/svc"
	"movieReviewSystem/movieReviewSystemApi/group/groupApi/generateFileV1/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GroupAdminLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGroupAdminLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupAdminLogic {
	return &GroupAdminLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GroupAdminLogic) GroupAdmin(req *types.GroupAdminIdReq) (resp *types.GroupAdminIdResp, err error) {
	// todo: add your logic here and delete this line

	return
}
