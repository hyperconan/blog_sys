package logic

import (
	"context"
	"errors"

	"blog/internal/models"
	"blog/internal/svc"
	"blog/pb/blog"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPostDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPostDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPostDetailLogic {
	return &GetPostDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetPostDetailLogic) GetPostDetail(in *blog.GetPostDetailRequest) (*blog.GetPostDetailResponse, error) {
	var post models.Post
	if err := l.svcCtx.DB.First(&post, in.PostId).Error; err != nil {
		return nil, errors.New("post not found")
	}

	return &blog.GetPostDetailResponse{
		Id:        post.ID,
		Title:     post.Title,
		Content:   post.Content,
		UserId:    post.UserID,
		CreatedAt: post.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: post.UpdatedAt.Format("2006-01-02 15:04:05"),
	}, nil
}
