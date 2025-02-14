package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"movieReviewSystem/movieReviewSystemApi/movieInformation/movieInformationApi/generateFileV1/internal/logic"
	"movieReviewSystem/movieReviewSystemApi/movieInformation/movieInformationApi/generateFileV1/internal/svc"
	"movieReviewSystem/movieReviewSystemApi/movieInformation/movieInformationApi/generateFileV1/internal/types"
)

func getMovieInformationHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.MovielnInformationReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewGetMovieInformationLogic(r.Context(), svcCtx)
		resp, err := l.GetMovieInformation(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
