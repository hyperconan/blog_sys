package logic

import (
	"context"

	"gateway/blogclient"
	"gateway/internal/middleware"
	"gateway/internal/svc"
	"gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreatePostLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreatePostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreatePostLogic {
	return &CreatePostLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreatePostLogic) CreatePost(req *types.CreatePostRequest) (resp *types.CreatePostResponse, err error) {
	userId, err := middleware.GetUserIdFromContext(l.ctx)
	if err != nil {
		return nil, err
	}

	result, err := l.svcCtx.BlogRpc.CreatePost(l.ctx, &blogclient.CreatePostRequest{
		Title:   req.Title,
		Content: req.Content,
		UserId:  userId,
	})
	if err != nil {
		return nil, err
	}

	return &types.CreatePostResponse{
		PostId: result.PostId,
	}, nil
}
