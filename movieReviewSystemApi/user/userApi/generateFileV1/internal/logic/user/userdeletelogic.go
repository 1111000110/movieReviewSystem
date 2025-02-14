package user

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"movieReviewSystem/movieReviewSystemApi/user/userApi/generateFileV1/internal/svc"
	"movieReviewSystem/movieReviewSystemApi/user/userApi/generateFileV1/internal/types"
	"movieReviewSystem/movieReviewSystemApi/user/userRpc/generateFileV1/userservice"
)

type UserDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserDeleteLogic {
	return &UserDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}
func (l *UserDeleteLogic) UserDelete(req *types.UserDeleteReq) (resp *types.UserDeleteResp, err error) {
	// 获取并断言 user_id 为 json.Number
	userIdValue := l.ctx.Value("user_id")
	if userId, ok := userIdValue.(json.Number); ok {
		// 将 json.Number 转换为 int64
		userIdInt, err := userId.Int64()
		if err != nil {
			return nil, errors.New("failed to parse user_id from context")
		}

		// 判断是否是当前用户，进行删除
		if userIdInt == req.UserId {
			_, err := l.svcCtx.UserRpcService.UserDelete(l.ctx, &userservice.UserDeleteReq{
				Phone:    req.Phone,
				UserId:   req.UserId,
				Password: req.Password,
			})
			if err != nil {
				return nil, fmt.Errorf("failed to delete user: %w", err)
			}

			// 返回删除成功的响应
			return &types.UserDeleteResp{}, nil
		}
	}
	return nil, errors.New("user_id does not match")
}
