package logic

import (
	"context"

	"blog/internal/models"
	"blog/internal/svc"
	"blog/pb/blog"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreatePostLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreatePostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreatePostLogic {
	return &CreatePostLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreatePostLogic) CreatePost(in *blog.CreatePostRequest) (*blog.CreatePostResponse, error) {
	post := &models.Post{
		Title:   in.Title,
		Content: in.Content,
		UserID:  in.UserId,
	}

	if err := l.svcCtx.DB.Create(post).Error; err != nil {
		return nil, err
	}

	return &blog.CreatePostResponse{
		PostId: post.ID,
	}, nil
}
