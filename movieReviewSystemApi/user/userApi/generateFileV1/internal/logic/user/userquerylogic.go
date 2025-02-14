package user

import (
	"context"
	"movieReviewSystem/movieReviewSystemApi/user/userRpc/generateFileV1/userservice"

	"movieReviewSystem/movieReviewSystemApi/user/userApi/generateFileV1/internal/svc"
	"movieReviewSystem/movieReviewSystemApi/user/userApi/generateFileV1/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserQueryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserQueryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserQueryLogic {
	return &UserQueryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserQueryLogic) UserQuery(req *types.UserQueryReq) (resp *types.UserQueryResp, err error) {
	// todo: add your logic here and delete this line
	respData, err := l.svcCtx.UserRpcService.UserQuery(l.ctx, &userservice.UserQueryReq{
		UserId: req.UserId,
	})
	if err != nil || respData == nil {
		return nil, err
	}
	resp = &types.UserQueryResp{
		UserId:    respData.UserId,
		Phone:     respData.Phone,
		Email:     respData.Email,
		NickName:  respData.NickName,
		Avatar:    respData.Avatar,
		Gender:    respData.Gender,
		BirthDate: respData.BirthDate,
		Role:      respData.Role,
		Status:    respData.Status,
	}
	return
}
