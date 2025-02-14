package logic

import (
	"context"
	"encoding/json"
	"github.com/pkg/errors"

	"movieReviewSystem/movieReviewSystemApi/movieInformation/movieInformationApi/generateFileV1/internal/svc"
	"movieReviewSystem/movieReviewSystemApi/movieInformation/movieInformationApi/generateFileV1/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMovieInformationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMovieInformationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMovieInformationLogic {
	return &GetMovieInformationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMovieInformationLogic) GetMovieInformation(req *types.MovielnInformationReq) (resp *types.MovieInfomationReq, err error) {
	// todo: add your logic here and delete this line
	userIdValue := l.ctx.Value("user_id")
	if userId, ok := userIdValue.(json.Number); ok {
		// 将 json.Number 转换为 int64
		userIdInt, err := userId.Int64()
		if err != nil {
			return nil, errors.New("failed to parse user_id from context")
		}
	}
	return
}
