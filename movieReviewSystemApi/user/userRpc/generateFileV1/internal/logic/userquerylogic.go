package logic

import (
	"context"
	"movieReviewSystem/movieReviewSystemApi/shared/tool"

	"movieReviewSystem/movieReviewSystemApi/user/userRpc/generateFileV1/internal/svc"
	"movieReviewSystem/movieReviewSystemApi/user/userRpc/pb"

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
	queryResp, err := l.svcCtx.UsersModel.FindOne(l.ctx, in.UserId)
	if err != nil {
		return nil, err
	}
	resp := &__.UserQueryResp{}
	resp.UserId = queryResp.UserId
	phone, err := tool.Decrypt(queryResp.Phone)
	if err != nil {
		return nil, err
	}
	resp.Phone = phone
	if queryResp.Email.Valid {
		email, err := queryResp.Email.Value()
		if err != nil {
			return nil, err
		}
		resp.Email = email.(string)
	}
	if queryResp.NickName.Valid {
		nickName, err := queryResp.NickName.Value()
		if err != nil {
			return nil, err
		}
		resp.NickName = nickName.(string)
	}
	if queryResp.Avatar.Valid {
		avatar, err := queryResp.Avatar.Value()
		if err != nil {
			return nil, err
		}
		resp.Avatar = avatar.(string)
	}
	if queryResp.Gender.Valid {
		gender, err := queryResp.Gender.Value()
		if err != nil {
			return nil, err
		}
		resp.Gender = gender.(string)
	}
	if queryResp.BirthDate.Valid {
		birthDate, err := queryResp.BirthDate.Value()
		if err != nil {
			return nil, err
		}
		resp.BirthDate = birthDate.(int64)
	}
	if queryResp.Role.Valid {
		role, err := queryResp.Role.Value()
		if err != nil {
			return nil, err
		}
		resp.Role = role.(string)
	}
	if queryResp.Status.Valid {
		status, err := queryResp.Status.Value()
		if err != nil {
			return nil, err
		}
		resp.Status = status.(int64)
	}
	return resp, nil
}
