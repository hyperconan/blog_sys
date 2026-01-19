package logic

import (
	"context"

	"blog/internal/models"
	"blog/internal/svc"
	"blog/pb/blog"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAllPostsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAllPostsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAllPostsLogic {
	return &GetAllPostsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetAllPostsLogic) GetAllPosts(in *blog.GetAllPostsRequest) (*blog.GetAllPostsResponse, error) {
	var posts []models.Post
	if err := l.svcCtx.DB.Order("created_at desc").Find(&posts).Error; err != nil {
		return nil, err
	}

	var postInfos []*blog.PostInfo
	for _, post := range posts {
		postInfos = append(postInfos, &blog.PostInfo{
			Id:        post.ID,
			Title:     post.Title,
			UserId:    post.UserID,
			CreatedAt: post.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: post.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	return &blog.GetAllPostsResponse{
		Posts: postInfos,
	}, nil
}
