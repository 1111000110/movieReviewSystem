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
	if l.ctx.Value("user_id").(int64) == req.User.UserId {
		_, err = l.svcCtx.UserRpcService.UserUpdate(l.ctx, &userservice.UserUpdateReq{
			User: &userservice.User{
				UserId:    req.User.UserId,
				Phone:     req.User.Phone,
				Email:     req.User.Email,
				Password:  req.User.Password,
				Nickname:  req.User.NickName,
				Avatar:    req.User.Avatar,
				Gender:    req.User.Gender,
				BirthDate: req.User.BirthDate,
				Role:      req.User.Role,
				Status:    req.User.Status,
			},
		})
	}
	return &types.UserUpdateResp{}, err
}
