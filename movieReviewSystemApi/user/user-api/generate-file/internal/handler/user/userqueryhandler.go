package user

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"movieReviewSystem/movieReviewSystemApi/user/user-api/generate-file/internal/logic/user"
	"movieReviewSystem/movieReviewSystemApi/user/user-api/generate-file/internal/svc"
	"movieReviewSystem/movieReviewSystemApi/user/user-api/generate-file/internal/types"
)

func UserQueryHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserQueryReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewUserQueryLogic(r.Context(), svcCtx)
		resp, err := l.UserQuery(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
