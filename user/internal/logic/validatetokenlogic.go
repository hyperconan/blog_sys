package logic

import (
	"context"

	"user/internal/svc"
	userpb "user/pb/user"

	"github.com/golang-jwt/jwt/v4"
	"github.com/zeromicro/go-zero/core/logx"
)

type ValidateTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewValidateTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ValidateTokenLogic {
	return &ValidateTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ValidateTokenLogic) ValidateToken(in *userpb.ValidateTokenRequest) (*userpb.ValidateTokenResponse, error) {
	token, err := jwt.Parse(in.Token, func(token *jwt.Token) (interface{}, error) {
		return []byte("hyperconan"), nil
	})

	if err != nil {
		return &userpb.ValidateTokenResponse{
			Valid: false,
		}, nil
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID, ok := claims["id"].(float64)
		if !ok {
			return &userpb.ValidateTokenResponse{
				Valid: false,
			}, nil
		}

		userName, ok := claims["username"].(string)
		if !ok {
			return &userpb.ValidateTokenResponse{
				Valid: false,
			}, nil
		}

		return &userpb.ValidateTokenResponse{
			UserId:   uint32(userID),
			Username: userName,
			Valid:    true,
		}, nil
	}

	return &userpb.ValidateTokenResponse{
		Valid: false,
	}, nil
}