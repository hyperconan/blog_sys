package logic

import (
	"context"

	"gateway/internal/svc"
	"gateway/internal/types"
	"gateway/blogclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAllPostsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAllPostsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAllPostsLogic {
	return &GetAllPostsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAllPostsLogic) GetAllPosts() (resp *types.GetPostsResponse, err error) {
	// Call blog service
	result, err := l.svcCtx.BlogRpc.GetAllPosts(l.ctx, &blogclient.GetAllPostsRequest{})
	if err != nil {
		return nil, err
	}

	// Convert response
	posts := make([]types.PostInfo, len(result.Posts))
	for i, post := range result.Posts {
		posts[i] = types.PostInfo{
			Id:       post.Id,
			Title:    post.Title,
			UserId:   post.UserId,
			CreateAt: post.CreatedAt,
			UpdateAt: post.UpdatedAt,
		}
	}

	return &types.GetPostsResponse{
		Posts: posts,
	}, nil
}