package logic

import (
	"context"
	"errors"
	"time"

	"blog/internal/svc"
	"blog/pb/blog"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdatePostLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdatePostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatePostLogic {
	return &UpdatePostLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdatePostLogic) UpdatePost(in *blog.UpdatePostRequest) (*blog.UpdatePostResponse, error) {
	postsMu.Lock()
	defer postsMu.Unlock()

	post, exists := posts[in.PostId]
	if !exists {
		return nil, errors.New("post not found")
	}

	if post.UserID != in.UserId {
		return nil, errors.New("unauthorized")
	}

	post.Title = in.Title
	post.Content = in.Content
	post.UpdatedAt = time.Now()

	return &blog.UpdatePostResponse{
		Message: "Post updated successfully",
	}, nil
}