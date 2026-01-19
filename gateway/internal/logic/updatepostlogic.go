package logic

import (
	"context"

	"gateway/blogclient"
	"gateway/internal/middleware"
	"gateway/internal/svc"
	"gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdatePostLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdatePostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatePostLogic {
	return &UpdatePostLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdatePostLogic) UpdatePost(req *types.UpdatePostRequest) (resp *types.UpdatePostResponse, err error) {
	userId, err := middleware.GetUserIdFromContext(l.ctx)
	if err != nil {
		return nil, err
	}

	_, err = l.svcCtx.BlogRpc.UpdatePost(l.ctx, &blogclient.UpdatePostRequest{
		PostId:  req.PostId,
		UserId:  userId,
		Title:   req.Title,
		Content: req.Content,
	})
	if err != nil {
		return nil, err
	}

	return &types.UpdatePostResponse{
		Message: "Post updated successfully",
	}, nil
}
