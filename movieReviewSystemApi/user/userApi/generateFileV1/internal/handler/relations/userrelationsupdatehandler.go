package relations

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"movieReviewSystem/movieReviewSystemApi/user/userApi/generateFileV1/internal/logic/relations"
	"movieReviewSystem/movieReviewSystemApi/user/userApi/generateFileV1/internal/svc"
	"movieReviewSystem/movieReviewSystemApi/user/userApi/generateFileV1/internal/types"
)

func UserRelationsUpdateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserRelationsUpdateReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := relations.NewUserRelationsUpdateLogic(r.Context(), svcCtx)
		resp, err := l.UserRelationsUpdate(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
