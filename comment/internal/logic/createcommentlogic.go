package logic

import (
	"context"

	"comment/internal/models"
	"comment/internal/svc"
	"comment/pb/comment"

	"github.com/zeromicro/go-zero/core/logx"
)

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
	newComment := &models.Comment{
		Content: in.Content,
		UserID:  in.UserId,
		PostID:  in.PostId,
	}

	if err := l.svcCtx.DB.Create(newComment).Error; err != nil {
		return nil, err
	}

	return &comment.CreateCommentResponse{
		CommentId: newComment.ID,
	}, nil
}
