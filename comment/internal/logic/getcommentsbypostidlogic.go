package logic

import (
	"context"

	"comment/internal/models"
	"comment/internal/svc"
	"comment/pb/comment"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCommentsByPostIDLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCommentsByPostIDLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCommentsByPostIDLogic {
	return &GetCommentsByPostIDLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCommentsByPostIDLogic) GetCommentsByPostID(in *comment.GetCommentsByPostIDRequest) (*comment.GetCommentsByPostIDResponse, error) {
	var comments []models.Comment
	if err := l.svcCtx.DB.Where("post_id = ?", in.PostId).Order("created_at desc").Find(&comments).Error; err != nil {
		return nil, err
	}

	var commentInfos []*comment.CommentInfo
	for _, c := range comments {
		commentInfos = append(commentInfos, &comment.CommentInfo{
			Id:        c.ID,
			Content:   c.Content,
			UserId:    c.UserID,
			PostId:    c.PostID,
			CreatedAt: c.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	return &comment.GetCommentsByPostIDResponse{
		Comments: commentInfos,
	}, nil
}
