package logic

import (
	"context"
	"errors"

	"blog/internal/svc"
	"blog/pb/blog"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeletePostLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeletePostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeletePostLogic {
	return &DeletePostLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeletePostLogic) DeletePost(in *blog.DeletePostRequest) (*blog.DeletePostResponse, error) {
	postsMu.Lock()
	defer postsMu.Unlock()

	post, exists := posts[in.PostId]
	if !exists {
		return nil, errors.New("post not found")
	}

	if post.UserID != in.UserId {
		return nil, errors.New("unauthorized")
	}

	delete(posts, in.PostId)

	return &blog.DeletePostResponse{
		Message: "Post deleted successfully",
	}, nil
}