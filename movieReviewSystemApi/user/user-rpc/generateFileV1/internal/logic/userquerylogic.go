package logic

import (
	"context"
	"movieReviewSystem/movieReviewSystemApi/user/user-rpc/generateFileV1/internal/svc"
	"movieReviewSystem/movieReviewSystemApi/user/user-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserQueryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserQueryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserQueryLogic {
	return &UserQueryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserQueryLogic) UserQuery(in *__.UserQueryReq) (*__.UserQueryResp, error) {
	// todo: add your logic here and delete this line
	queryResp, err := l.svcCtx.UsersModel.FindOne(l.ctx, in.GetUserId())
	if err != nil || queryResp == nil {
		return nil, err
	}
	User, err := __.ModelUserToUser(*queryResp)
	if err != nil {
		return nil, err
	}
	resp := &__.UserQueryResp{
		User: &User,
	}
	return resp, nil
}
