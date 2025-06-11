package relations

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"movieReviewSystem/movieReviewSystemApi/user/user-api/generate-file/internal/logic/relations"
	"movieReviewSystem/movieReviewSystemApi/user/user-api/generate-file/internal/svc"
	"movieReviewSystem/movieReviewSystemApi/user/user-api/generate-file/internal/types"
)

func UserRelationsGetHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserRelationsGetReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := relations.NewUserRelationsGetLogic(r.Context(), svcCtx)
		resp, err := l.UserRelationsGet(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
