package logic

import (
	"context"

	"gateway/commentclient"
	"gateway/internal/svc"
	"gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCommentsByPostIDLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCommentsByPostIDLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCommentsByPostIDLogic {
	return &GetCommentsByPostIDLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCommentsByPostIDLogic) GetCommentsByPostID(req *types.GetCommentsRequest) (resp *types.GetCommentsResponse, err error) {
	result, err := l.svcCtx.CommentRpc.GetCommentsByPostID(l.ctx, &commentclient.GetCommentsByPostIDRequest{
		PostId: req.PostId,
	})
	if err != nil {
		return nil, err
	}

	comments := make([]types.CommentInfo, len(result.Comments))
	for i, comment := range result.Comments {
		comments[i] = types.CommentInfo{
			Id:       comment.Id,
			Content:  comment.Content,
			UserId:   comment.UserId,
			PostId:   comment.PostId,
			CreateAt: comment.CreatedAt,
		}
	}

	return &types.GetCommentsResponse{
		Comments: comments,
	}, nil
}
