package dept

import (
	"net/http"

	"fiber-g/apps/app/internal/logic/dept"
	"fiber-g/apps/app/internal/svc"
	"fiber-g/apps/app/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func CreateDeptHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateDeptReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := dept.NewCreateDeptLogic(r.Context(), svcCtx)
		resp, err := l.CreateDept(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
