package svc

import (
	"gateway/internal/config"
	"gateway/userclient"
	"gateway/blogclient"
	"gateway/commentclient"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config

	UserRpc    userclient.User
	BlogRpc    blogclient.Blog
	CommentRpc commentclient.Comment
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		UserRpc:    userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
		BlogRpc:    blogclient.NewBlog(zrpc.MustNewClient(c.BlogRpc)),
		CommentRpc: commentclient.NewComment(zrpc.MustNewClient(c.CommentRpc)),
	}
}