package logic

import (
	"context"
	"database/sql"
	"movieReviewSystem/movieReviewSystemApi/shared/tool"
	"movieReviewSystem/movieReviewSystemApi/user/userModel"
	"time"

	"movieReviewSystem/movieReviewSystemApi/user/userRpc/generateFileV1/internal/svc"
	"movieReviewSystem/movieReviewSystemApi/user/userRpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRegisterLogic {
	return &UserRegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserRegisterLogic) UserRegister(in *__.UserRegisterReq) (*__.UserRegisterResp, error) {
	// todo: add your logic here and delete this line
	if in.GetPhone() == "" || in.GetPassword() == "" {
		return &__.UserRegisterResp{UserId: -1}, nil
	}
	_phone, err := tool.Encrypt(in.GetPhone())
	if err != nil {
		return nil, err
	}
	_password, err := tool.HashPassword(in.GetPassword())
	if err != nil {
		return nil, err
	}
	insertData := &userModel.Users{
		Phone:      _phone,
		Password:   _password,
		Status:     sql.NullInt64{Int64: __.Account_normal, Valid: true},
		Role:       sql.NullString{String: in.GetRole(), Valid: true},
		CreateDate: sql.NullInt64{time.Now().Unix(), true},
		UpdataDate: sql.NullInt64{time.Now().Unix(), true},
	}

	_, err = l.svcCtx.UsersModel.Insert(l.ctx, insertData)
	logx.Info(err)
	if err != nil {
		return nil, err
	}
	data, err := NewUserLoginLogic(l.ctx, l.svcCtx).UserLogin(&__.UserLoginReq{
		Phone:    in.GetPhone(),
		Password: in.GetPassword(),
	})
	if err != nil {
		return nil, err
	}
	return &__.UserRegisterResp{UserId: data.UserId}, err
}
