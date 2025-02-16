package relations

import (
	"context"
	"movieReviewSystem/movieReviewSystemApi/user/userRpc/generateFileV1/userservice"

	"movieReviewSystem/movieReviewSystemApi/user/userApi/generateFileV1/internal/svc"
	"movieReviewSystem/movieReviewSystemApi/user/userApi/generateFileV1/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRelationsGetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRelationsGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRelationsGetLogic {
	return &UserRelationsGetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRelationsGetLogic) UserRelationsGet(req *types.UserRelationsGetReq) (resp *types.UserRelationsGetResp, err error) {
	// todo: add your logic here and delete this line
	reverseUserId := false
	if req.UserId > req.OUserId {
		reverseUserId = true
	}
	if reverseUserId {
		req.UserId, req.OUserId = req.OUserId, req.UserId
	}
	data, err := l.svcCtx.UserRpcService.UserRelationsGet(l.ctx, &userservice.UserRelationsGetReq{
		UserId:  req.UserId,
		OUserId: req.OUserId,
	})
	if err != nil {
		return nil, err
	}
	resp = &types.UserRelationsGetResp{
		Relations: types.Relations{
			RelationsId:      data.GetRelations().GetRelationsId(),
			UserId:           data.GetRelations().GetUserId(),
			OtherId:          data.GetRelations().GetOtherId(),
			RelationshipType: data.GetRelations().GetRelationshipType(),
			CreateAt:         data.GetRelations().GetCreateAt(),
			UpdateAt:         data.GetRelations().GetUpdateAt(),
		},
	}
	if reverseUserId {
		resp.Relations.UserId, resp.Relations.OtherId = resp.Relations.OtherId, resp.Relations.UserId
	}
	return resp, err
}
