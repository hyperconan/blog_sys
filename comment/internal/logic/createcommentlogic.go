package logic

import (
	"context"
	"sync"
	"time"

	"comment/internal/svc"
	"comment/pb/comment"

	"github.com/zeromicro/go-zero/core/logx"
)

// Simple in-memory comment storage
var (
	comments   = make(map[uint32]*Comment)
	commentsMu sync.RWMutex
	commentID  uint32 = 1
)

type Comment struct {
	ID        uint32
	Content   string
	UserID    uint32
	PostID    uint32
	CreatedAt time.Time
}

type CreateCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateCommentLogic {
	return &CreateCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateCommentLogic) CreateComment(in *comment.CreateCommentRequest) (*comment.CreateCommentResponse, error) {
	commentsMu.Lock()
	defer commentsMu.Unlock()

	newComment := &Comment{
		ID:        commentID,
		Content:   in.Content,
		UserID:    in.UserId,
		PostID:    in.PostId,
		CreatedAt: time.Now(),
	}
	comments[commentID] = newComment
	commentID++

	return &comment.CreateCommentResponse{
		CommentId: commentID - 1,
	}, nil
}
