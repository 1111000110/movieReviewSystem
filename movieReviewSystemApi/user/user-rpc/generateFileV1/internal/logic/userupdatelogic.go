package logic

import (
	"context"
	"github.com/pkg/errors"
	"movieReviewSystem/movieReviewSystemApi/shared/tool"
	"movieReviewSystem/movieReviewSystemApi/user/user-rpc/generateFileV1/internal/svc"
	"movieReviewSystem/movieReviewSystemApi/user/user-rpc/pb"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserUpdateLogic {
	return &UserUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserUpdateLogic) UserUpdate(in *__.UserUpdateReq) (*__.UserUpdateResp, error) {
	userResp, err := l.svcCtx.UsersModel.FindOne(l.ctx, in.GetUser().GetUserId())
	if err != nil {
		return nil, err
	}
	if userResp == nil {
		return nil, errors.New("user not found")
	}
	if in.GetUser().GetPhone() != "" {
		userResp.Phone, err = tool.Encrypt(in.GetUser().GetPhone())
		if err != nil {
			return nil, err
		}
	}
	if in.GetUser().GetEmail() != "" {
		userResp.Email = in.GetUser().GetEmail()
	}
	if in.GetUser().GetPassword() != "" {
		userResp.Password, err = tool.HashPassword(in.GetUser().GetPassword())
		if err != nil {
			return nil, err
		}
	}
	if in.GetUser().GetNickname() != "" {
		userResp.Avatar = in.GetUser().GetAvatar()
	}
	if in.GetUser().GetAvatar() != "" {
		userResp.Avatar = in.GetUser().GetAvatar()
	}
	if in.GetUser().GetGender() != 0 {
		userResp.Gender = in.GetUser().GetGender()
	}
	if in.GetUser().GetBirthDate() != 0 {
		userResp.BirthDate = in.GetUser().GetBirthDate()
	}
	if in.GetUser().GetRole() != "" {
		userResp.Role = in.GetUser().GetRole()
	}
	if in.GetUser().GetStatus() != 0 {
		userResp.Status = in.GetUser().GetStatus()
	}
	if in.GetUser().GetCreateAt() != 0 {
		userResp.CreateAt = in.GetUser().GetCreateAt()
	}
	userResp.UpdateAt = time.Now().Unix()
	_, err = l.svcCtx.UsersModel.Update(l.ctx, userResp)
	return &__.UserUpdateResp{}, nil
}
