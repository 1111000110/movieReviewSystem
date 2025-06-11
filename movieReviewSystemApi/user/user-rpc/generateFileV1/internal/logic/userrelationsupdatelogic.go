package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"movieReviewSystem/movieReviewSystemApi/shared/tool"
	"movieReviewSystem/movieReviewSystemApi/user/user-rpc/generateFileV1/internal/svc"
	"movieReviewSystem/movieReviewSystemApi/user/user-rpc/pb"
	relationsModel "movieReviewSystem/movieReviewSystemApi/user/userModel/mongo/userRelations"
	"time"
)

type UserRelationsUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserRelationsUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRelationsUpdateLogic {
	return &UserRelationsUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserRelationsUpdateLogic) UserRelationsUpdate(in *__.UserRelationsUpdateReq) (*__.UserRelationsUpdateResp, error) {
	// todo: add your logic here and delete this line
	GetLogic := NewUserRelationsGetLogic(l.ctx, l.svcCtx)
	data, err := GetLogic.UserRelationsGet(&__.UserRelationsGetReq{
		UserId:  in.GetUserId(),
		OUserId: in.GetOUserId(),
	})
	if err != nil {
		if err == relationsModel.ErrNotFound {
			//insert
			sf, err := tool.NewSnowflake(1, 1) // 假设当前节点是数据中心1，机器1
			if err != nil || sf == nil {
				return nil, err
			}
			id, err := sf.Generate()
			if err != nil {
				panic(err)
			}
			_, err = l.svcCtx.UserRelationsModel.Upsert(l.ctx, &relationsModel.UserRelations{
				ID:               id,
				UserId:           in.GetUserId(),
				OtherId:          in.GetOUserId(),
				RelationshipType: in.GetRelationshipType(),
				UpdateAt:         time.Now().Unix(),
				CreateAt:         time.Now().Unix(),
			})
			return &__.UserRelationsUpdateResp{}, nil
		} else {
			return nil, err
		}

	}
	//关系业务逻辑，后面写
	data.Relations.RelationshipType = in.GetRelationshipType()
	modelRelations, err := __.UserRelationsToModelUserRelations(*data.Relations)
	if err != nil {
		return &__.UserRelationsUpdateResp{}, err
	}
	_, err = l.svcCtx.UserRelationsModel.Upsert(l.ctx, &modelRelations)
	return &__.UserRelationsUpdateResp{}, err
}
