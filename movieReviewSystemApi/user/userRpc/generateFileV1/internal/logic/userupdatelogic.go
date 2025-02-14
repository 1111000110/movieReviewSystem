package logic

import (
	"context"
	"database/sql"
	"github.com/pkg/errors"
	"movieReviewSystem/movieReviewSystemApi/shared/tool"
	"movieReviewSystem/movieReviewSystemApi/user/userRpc/generateFileV1/internal/svc"
	"movieReviewSystem/movieReviewSystemApi/user/userRpc/pb"
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
	userResp, err := l.svcCtx.UsersModel.FindOne(l.ctx, in.GetUserId())
	if err != nil {
		return nil, err
	}
	if userResp == nil {
		return nil, errors.New("user not found")
	}
	if in.GetNickName() != "" {
		userResp.NickName = sql.NullString{in.GetEmail(), true}
	}
	if in.GetEmail() != "" {
		userResp.Email = sql.NullString{in.GetEmail(), true}
	}
	if in.GetPhone() != "" {
		phone, err := tool.Encrypt(in.GetPhone())
		if err != nil {
			return nil, err
		}
		userResp.Phone = phone
	}
	if in.GetPassword() != "" {
		password, err := tool.HashPassword(in.GetPassword())
		if err != nil {
			return nil, err
		}
		userResp.Password = password
	}
	if in.GetAvatar() != "" {
		userResp.Avatar = sql.NullString{in.GetAvatar(), true}
	}
	if in.GetGender() != "" {
		userResp.Gender = sql.NullString{in.GetGender(), true}
	}
	if in.GetBirthDate() != 0 {
		userResp.BirthDate = sql.NullInt64{in.GetBirthDate(), true}
	}
	if in.GetRole() != "" {
		userResp.Role = sql.NullString{in.GetRole(), true}
	}
	if in.GetStatus() != 0 {
		userResp.Status = sql.NullInt64{in.GetStatus(), true}
	}
	userResp.UpdataDate = sql.NullInt64{time.Now().Unix(), true}
	err = l.svcCtx.UsersModel.Update(l.ctx, userResp)
	return &__.UserUpdateResp{}, nil
}
