package logic

import (
	"context"
	"errors"

	"blog/internal/models"
	"blog/internal/svc"
	"blog/pb/blog"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
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
	var post models.Post
	if err := l.svcCtx.DB.First(&post, in.PostId).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("post not found")
		}
		return nil, err
	}

	if post.UserID != in.UserId {
		return nil, errors.New("unauthorized")
	}

	updates := map[string]interface{}{
		"title":   in.Title,
		"content": in.Content,
	}

	if err := l.svcCtx.DB.Model(&post).Updates(updates).Error; err != nil {
		return nil, err
	}

	return &blog.UpdatePostResponse{
		Message: "Post updated successfully",
	}, nil
}
