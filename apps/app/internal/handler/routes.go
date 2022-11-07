// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	app "fiber-g/apps/app/internal/handler/app"
	asset "fiber-g/apps/app/internal/handler/asset"
	dept "fiber-g/apps/app/internal/handler/dept"
	user "fiber-g/apps/app/internal/handler/user"
	"fiber-g/apps/app/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/health",
				Handler: app.HealthHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/create",
				Handler: dept.CreateDeptHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/v1/dept"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/create",
				Handler: user.CreateUserHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/v1/user"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/create",
				Handler: asset.CreateAssetHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/import",
				Handler: asset.ImportAssetHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/project/create",
				Handler: asset.CreateProjectHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/project/list",
				Handler: asset.ProjectListHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/v1/asset"),
	)
}
