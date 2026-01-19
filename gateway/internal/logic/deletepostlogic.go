package logic

import (
	"context"

	"gateway/internal/svc"
	"gateway/internal/types"
	"gateway/blogclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeletePostLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeletePostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeletePostLogic {
	return &DeletePostLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeletePostLogic) DeletePost(req *types.DeletePostRequest) (resp *types.DeletePostResponse, err error) {
	// TODO: Get user ID from JWT token
	userId := uint32(1) // Mock user ID, should get from JWT

	// Call blog service
	_, err = l.svcCtx.BlogRpc.DeletePost(l.ctx, &blogclient.DeletePostRequest{
		PostId: req.PostId,
		UserId: userId,
	})
	if err != nil {
		return nil, err
	}

	return &types.DeletePostResponse{
		Message: "Post deleted successfully",
	}, nil
}