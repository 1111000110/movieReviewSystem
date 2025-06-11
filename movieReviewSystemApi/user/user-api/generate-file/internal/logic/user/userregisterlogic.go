package user

import (
	"context"
	__ "movieReviewSystem/movieReviewSystemApi/user/user-rpc/pb"

	"movieReviewSystem/movieReviewSystemApi/user/user-api/generate-file/internal/svc"
	"movieReviewSystem/movieReviewSystemApi/user/user-api/generate-file/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRegisterLogic {
	return &UserRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRegisterLogic) UserRegister(req *types.UserRegisterReq) (resp *types.UserRegisterResp, err error) {
	// todo: add your logic here and delete this line
	registerResp, err := l.svcCtx.UserRpcService.UserRegister(l.ctx, &__.UserRegisterReq{Phone: req.Phone, Password: req.Password, Role: req.Role})
	if err != nil {
		return nil, err
	}
	resp = &types.UserRegisterResp{
		User: types.User{UserId: registerResp.User.GetUserId()},
	}
	return
}
