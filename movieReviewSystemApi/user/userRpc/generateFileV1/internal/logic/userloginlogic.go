package logic

import (
	"context"
	"errors"
	"movieReviewSystem/movieReviewSystemApi/shared/tool"
	"movieReviewSystem/movieReviewSystemApi/user/userModel/mongo/user"
	"movieReviewSystem/movieReviewSystemApi/user/userRpc/generateFileV1/internal/svc"
	"movieReviewSystem/movieReviewSystemApi/user/userRpc/pb"
	"time"

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
	var userResp *model.User
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
	tokenString, err := tool.CreateToken(userResp.ID, time.Duration(l.svcCtx.Config.AuthConf.AccessExpire)*time.Second, l.svcCtx.Config.AuthConf.AccessSecret)
	if tool.ComparePassword(userResp.Password, in.GetPassword()) && userResp.Status != __.Account_cancellation {
		return &__.UserLoginResp{
			Token: tokenString,
		}, err
	}
	return nil, errors.New("password or phone err")
}
