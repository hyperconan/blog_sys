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

	if err := l.svcCtx.DB.Delete(&post).Error; err != nil {
		return nil, err
	}

	return &blog.DeletePostResponse{
		Message: "Post deleted successfully",
	}, nil
}
