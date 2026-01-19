package handler

import (
	"net/http"

	"gateway/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/user/register",
				Handler: UserRegisterHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/login",
				Handler: UserLoginHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/blog",
				Handler: CreatePostHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/blog/all",
				Handler: GetAllPostsHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/blog/:post_id",
				Handler: UpdatePostHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/blog/:post_id",
				Handler: DeletePostHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/blog/comment",
				Handler: CreateCommentHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/blog/comment/:post_id",
				Handler: GetCommentsByPostIDHandler(serverCtx),
			},
		},
	)
}