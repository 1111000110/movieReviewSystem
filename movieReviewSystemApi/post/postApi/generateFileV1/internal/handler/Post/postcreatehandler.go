package Post

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"movieReviewSystem/movieReviewSystemApi/post/postApi/generateFileV1/internal/logic/Post"
	"movieReviewSystem/movieReviewSystemApi/post/postApi/generateFileV1/internal/svc"
	"movieReviewSystem/movieReviewSystemApi/post/postApi/generateFileV1/internal/types"
)

func PostCreateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PostCreateReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := Post.NewPostCreateLogic(r.Context(), svcCtx)
		resp, err := l.PostCreate(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
