package logic

import (
	"context"

	"gateway/commentclient"
	"gateway/internal/middleware"
	"gateway/internal/svc"
	"gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateCommentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateCommentLogic {
	return &CreateCommentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateCommentLogic) CreateComment(req *types.CreateCommentRequest) (resp *types.CreateCommentResponse, err error) {
	userId, err := middleware.GetUserIdFromContext(l.ctx)
	if err != nil {
		return nil, err
	}

	result, err := l.svcCtx.CommentRpc.CreateComment(l.ctx, &commentclient.CreateCommentRequest{
		Content: req.Content,
		UserId:  userId,
		PostId:  req.PostId,
	})
	if err != nil {
		return nil, err
	}

	return &types.CreateCommentResponse{
		CommentId: result.CommentId,
	}, nil
}
