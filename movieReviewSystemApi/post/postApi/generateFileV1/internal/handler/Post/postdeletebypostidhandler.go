package Post

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"movieReviewSystem/movieReviewSystemApi/post/postApi/generateFileV1/internal/logic/Post"
	"movieReviewSystem/movieReviewSystemApi/post/postApi/generateFileV1/internal/svc"
	"movieReviewSystem/movieReviewSystemApi/post/postApi/generateFileV1/internal/types"
)

func PostDeleteByPostIdHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PostDeleteByPostIdReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := Post.NewPostDeleteByPostIdLogic(r.Context(), svcCtx)
		resp, err := l.PostDeleteByPostId(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
