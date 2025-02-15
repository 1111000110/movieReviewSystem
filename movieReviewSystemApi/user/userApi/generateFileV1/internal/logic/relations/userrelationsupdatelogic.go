package relations

import (
	"context"

	"movieReviewSystem/movieReviewSystemApi/user/userApi/generateFileV1/internal/svc"
	"movieReviewSystem/movieReviewSystemApi/user/userApi/generateFileV1/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRelationsUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRelationsUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRelationsUpdateLogic {
	return &UserRelationsUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRelationsUpdateLogic) UserRelationsUpdate(req *types.UserRelationsUpdateReq) (resp *types.UserRelationsUpdateResp, err error) {
	// todo: add your logic here and delete this line

	return
}
