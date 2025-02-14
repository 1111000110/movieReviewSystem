package logic

import (
	"context"
	"errors"
	"movieReviewSystem/movieReviewSystemApi/shared/tool"
	"movieReviewSystem/movieReviewSystemApi/user/userModel"
	"movieReviewSystem/movieReviewSystemApi/user/userRpc/generateFileV1/internal/svc"
	"movieReviewSystem/movieReviewSystemApi/user/userRpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginLogic {
	return &UserLoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserLoginLogic) UserLogin(in *__.UserLoginReq) (*__.UserLoginResp, error) {
	// todo: add your logic here and delete this line
	var userResp *userModel.Users
	var err error
	switch {
	case in.GetUserId() != 0:
		userResp, err = l.svcCtx.UsersModel.FindOne(l.ctx, in.GetUserId())
	case in.GetPhone() != "":
		phone, err := tool.Encrypt(in.Phone)
		if err != nil {
			return nil, err
		}
		userResp, err = l.svcCtx.UsersModel.FindOneByPhone(l.ctx, phone)
	default:
		return nil, errors.New("UserID is nil and Phone is empty")
	}
	if err != nil || userResp == nil {
		return nil, err
	}
	status, err := userResp.Status.Value()
	if err != nil {
		return nil, err
	}
	if tool.ComparePassword(userResp.Password, in.GetPassword()) && status != __.Account_cancellation {
		return &__.UserLoginResp{
			UserId: userResp.UserId,
		}, err
	}
	return &__.UserLoginResp{
		UserId: -1,
	}, err
}
