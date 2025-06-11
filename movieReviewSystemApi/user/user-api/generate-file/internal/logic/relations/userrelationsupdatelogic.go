package relations

import (
	"context"
	"movieReviewSystem/movieReviewSystemApi/user/user-rpc/generateFileV1/userservice"

	"movieReviewSystem/movieReviewSystemApi/user/user-api/generate-file/internal/svc"
	"movieReviewSystem/movieReviewSystemApi/user/user-api/generate-file/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRelationsUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRelationsUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRelationsUpdateLogic {
	return &UserRelationsUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRelationsUpdateLogic) UserRelationsUpdate(req *types.UserRelationsUpdateReq) (resp *types.UserRelationsUpdateResp, err error) {
	// todo: add your logic here and delete this line
	reverseUserId := false
	if req.UserId > req.OUserId {
		reverseUserId = true
	}
	if reverseUserId {
		req.UserId, req.OUserId = req.OUserId, req.UserId
	}
	data, err := l.svcCtx.UserRpcService.UserRelationsUpdate(l.ctx, &userservice.UserRelationsUpdateReq{
		UserId:           req.UserId,
		OUserId:          req.OUserId,
		RelationshipType: req.RelationshipType,
	})
	if err != nil {
		return nil, err
	}
	resp = &types.UserRelationsUpdateResp{
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
