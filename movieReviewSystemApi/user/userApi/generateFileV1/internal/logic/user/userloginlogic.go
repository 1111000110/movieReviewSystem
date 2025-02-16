package user

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"movieReviewSystem/movieReviewSystemApi/user/userApi/generateFileV1/internal/svc"
	"movieReviewSystem/movieReviewSystemApi/user/userApi/generateFileV1/internal/types"
	"movieReviewSystem/movieReviewSystemApi/user/userRpc/generateFileV1/userservice"
)

type UserLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginLogic {
	return &UserLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLoginLogic) UserLogin(req *types.UserLoginReq) (resp *types.UserLoginResp, err error) {
	// todo: add your logic here and delete this line
	userResp, err := l.svcCtx.UserRpcService.UserLogin(l.ctx, &userservice.UserLoginReq{
		Phone:    req.Phone,
		Password: req.Password,
		UserId:   req.UserId,
	})
	if err != nil {
		return nil, err
	}
	resp.Token = userResp.Token
	return resp, nil
}
