// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"lebron/apps/order/admin/internal/svc"
	"net/http"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/from/:name",
				Handler: AdminHandler(serverCtx),
			},
		},
	)
}
