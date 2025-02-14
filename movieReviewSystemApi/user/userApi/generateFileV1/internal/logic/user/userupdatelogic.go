package user

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"movieReviewSystem/movieReviewSystemApi/user/userApi/generateFileV1/internal/svc"
	"movieReviewSystem/movieReviewSystemApi/user/userApi/generateFileV1/internal/types"
	"movieReviewSystem/movieReviewSystemApi/user/userRpc/generateFileV1/userservice"
)

type UserUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserUpdateLogic {
	return &UserUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserUpdateLogic) UserUpdate(req *types.UserUpdateReq) (resp *types.UserUpdateResp, err error) {
	// todo: add your logic here and delete this line
	if l.ctx.Value("user_id").(int64) == req.UserId {
		_, err = l.svcCtx.UserRpcService.UserUpdate(l.ctx, &userservice.UserUpdateReq{
			UserId:    req.UserId,
			Phone:     req.Phone,
			Email:     req.Email,
			Password:  req.Password,
			NickName:  req.NickName,
			Avatar:    req.Avatar,
			Gender:    req.Gender,
			BirthDate: req.BirthDate,
			Role:      req.Role,
			Status:    req.Status,
		})
	}
	return &types.UserUpdateResp{}, err
}
