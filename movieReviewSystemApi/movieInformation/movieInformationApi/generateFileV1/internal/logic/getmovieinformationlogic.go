package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"movieReviewSystem/movieReviewSystemApi/movieInformation/movieInformationApi/generateFileV1/internal/svc"
	"movieReviewSystem/movieReviewSystemApi/movieInformation/movieInformationApi/generateFileV1/internal/types"
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
	movieInfo, err := l.svcCtx.Movie.FindOne(l.ctx, req.Keywords)
	if err != nil {
		return nil, err
	}
	resp = &types.MovieInfomationReq{
		MovieInformation: make([]types.MovieInformation, 0),
	}
	resp.MovieInformation = append(resp.MovieInformation, types.MovieInformation{
		MovieInformationId: movieInfo.MovieInformationId,
		Titlt:              movieInfo.Title,
		Desc:               movieInfo.Desc,
		Author:             movieInfo.Author,
		Actors:             movieInfo.Actors,
		Language:           movieInfo.Language,
		Duration:           movieInfo.Duration,
		ReleaseDate:        movieInfo.ReleaseDate,
		Genre:              movieInfo.Genre,
		Poster:             movieInfo.Poster,
	})
	return
}
