package asset

import (
	"net/http"

	"fiber-g/apps/app/internal/logic/asset"
	"fiber-g/apps/app/internal/svc"
	"fiber-g/apps/app/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func CreateAssetHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Asset
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := asset.NewCreateAssetLogic(r.Context(), svcCtx)
		resp, err := l.CreateAsset(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
