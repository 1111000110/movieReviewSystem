package Post

import (
	"context"
	"movieReviewSystem/movieReviewSystemApi/shared/tool"

	"movieReviewSystem/movieReviewSystemApi/post/postApi/generateFileV1/internal/svc"
	"movieReviewSystem/movieReviewSystemApi/post/postApi/generateFileV1/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PostGetListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPostGetListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PostGetListLogic {
	return &PostGetListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PostGetListLogic) PostGetList(req *types.PostGetListRep) (resp *types.PostGetListResp, err error) {
	// todo: add your logic here and delete this line
	data, err := l.svcCtx.PostModel.FindAll(l.ctx, req.Offset, req.Limit) // 根据offset和limit获取帖子信息
	if err != nil {
		return resp, err
	}
	postInfos := make([]types.Post, 0)
	for _, modelPost := range *data {
		postInfo := &types.Post{}
		err = tool.StructConverter(&modelPost, postInfo)
		if err != nil {
			return &types.PostGetListResp{}, err
		}
		postInfos = append(postInfos, *postInfo)
	}
	resp = &types.PostGetListResp{
		PostInfos: postInfos,
	}
	return
}
