package Post

import (
	"context"
	model "movieReviewSystem/movieReviewSystemApi/post/postModel"
	"movieReviewSystem/movieReviewSystemApi/shared/tool"

	"movieReviewSystem/movieReviewSystemApi/post/postApi/generateFileV1/internal/svc"
	"movieReviewSystem/movieReviewSystemApi/post/postApi/generateFileV1/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PostCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPostCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PostCreateLogic {
	return &PostCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PostCreateLogic) PostCreate(req *types.PostCreateReq) (resp *types.PostCreateResp, err error) {
	// todo: add your logic here and delete this line
	snowflake, err := tool.NewSnowflake(1, 1)
	if err != nil {
		return nil, err
	}
	postId, err := snowflake.Generate() // 雪花唯一id
	if err != nil {
		return nil, err
	}
	modelType := model.Post{
		PostId:  postId,
		UserId:  req.UserId,
		ThemeId: req.ThemeId,
		Title:   req.Title,
		Source:  req.Source,
		Text:    req.Text,
		Image:   req.Image,
		Video:   req.Video,
	}
	err = l.svcCtx.PostModel.Insert(l.ctx, &modelType)
	if err != nil {
		return nil, err
	}
	return
}
