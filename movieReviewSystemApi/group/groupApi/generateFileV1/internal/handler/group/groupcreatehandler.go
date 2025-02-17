package group

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"movieReviewSystem/movieReviewSystemApi/group/groupApi/generateFileV1/internal/logic/group"
	"movieReviewSystem/movieReviewSystemApi/group/groupApi/generateFileV1/internal/svc"
	"movieReviewSystem/movieReviewSystemApi/group/groupApi/generateFileV1/internal/types"
)

func GroupCreateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GroupCreateReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := group.NewGroupCreateLogic(r.Context(), svcCtx)
		resp, err := l.GroupCreate(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
