package logic

import (
	"context"
	"sync"
	"time"

	"blog/internal/svc"
	"blog/pb/blog"

	"github.com/zeromicro/go-zero/core/logx"
)

// Simple in-memory post storage
var (
	posts   = make(map[uint32]*Post)
	postsMu sync.RWMutex
	postID  uint32 = 1
)

type Post struct {
	ID        uint32
	Title     string
	Content   string
	UserID    uint32
	CreatedAt time.Time
	UpdatedAt time.Time
}

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
	postsMu.Lock()
	defer postsMu.Unlock()

	now := time.Now()
	post := &Post{
		ID:        postID,
		Title:     in.Title,
		Content:   in.Content,
		UserID:    in.UserId,
		CreatedAt: now,
		UpdatedAt: now,
	}
	posts[postID] = post
	postID++

	return &blog.CreatePostResponse{
		PostId: postID - 1,
	}, nil
}
