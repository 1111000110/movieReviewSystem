package logic

import (
	"context"
	"movieReviewSystem/movieReviewSystemApi/shared/tool"

	"movieReviewSystem/movieReviewSystemApi/user/userRpc/generateFileV1/internal/svc"
	"movieReviewSystem/movieReviewSystemApi/user/userRpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserDeleteLogic {
	return &UserDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserDeleteLogic) UserDelete(in *__.UserDeleteReq) (*__.UserDeleteResp, error) {
	// todo: add your logic here and delete this line
	respData, err := l.svcCtx.UsersModel.FindOne(l.ctx, in.UserId)
	if err != nil {
		return nil, err
	}
	phone, err := tool.Encrypt(in.GetPhone())
	if err != nil {
		return nil, err
	}
	if respData.Phone == phone && tool.ComparePassword(respData.Password, in.GetPassword()) {
		updataServer := NewUserUpdateLogic(l.ctx, l.svcCtx)
		_, err := updataServer.UserUpdate(&__.UserUpdateReq{UserId: in.UserId, Status: __.Account_cancellation})
		if err != nil {
			return nil, err
		}
	}
	return &__.UserDeleteResp{}, nil
}
