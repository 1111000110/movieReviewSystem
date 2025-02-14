package user

import (
	"context"
	"movieReviewSystem/movieReviewSystemApi/shared/tool"
	"movieReviewSystem/movieReviewSystemApi/user/userApi/generateFileV1/internal/svc"
	"movieReviewSystem/movieReviewSystemApi/user/userApi/generateFileV1/internal/types"
	"movieReviewSystem/movieReviewSystemApi/user/userRpc/generateFileV1/userservice"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
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
	if err != nil || userResp.GetUserId() == 0 {
		return nil, err
	}
	tokenString, err := tool.CreateToken(userResp.GetUserId(), time.Duration(l.svcCtx.Config.Auth.AccessExpire)*time.Second, l.svcCtx.Config.Auth.AccessSecret)
	resp = &types.UserLoginResp{
		Token: tokenString,
	}
	return resp, nil
}
