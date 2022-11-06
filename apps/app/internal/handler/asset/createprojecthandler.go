package asset

import (
	"net/http"

	"fiber-g/apps/app/internal/logic/asset"
	"fiber-g/apps/app/internal/svc"
	"fiber-g/apps/app/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func CreateProjectHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ProjectReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := asset.NewCreateProjectLogic(r.Context(), svcCtx)
		resp, err := l.CreateProject(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
