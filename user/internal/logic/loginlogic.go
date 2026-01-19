package logic

import (
	"context"
	"errors"
	"time"

	"user/internal/models"
	"user/internal/svc"
	userpb "user/pb/user"

	"github.com/golang-jwt/jwt/v4"
	"github.com/zeromicro/go-zero/core/logx"
	"golang.org/x/crypto/bcrypt"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *userpb.LoginRequest) (*userpb.LoginResponse, error) {
	var user models.User
	if err := l.svcCtx.DB.Where("username = ?", in.Username).First(&user).Error; err != nil {
		return nil, errors.New("invalid username or password")
	}

	// Verify password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(in.Password)); err != nil {
		return nil, errors.New("invalid username or password")
	}

	// Generate JWT token
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       user.ID,
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}).SignedString([]byte("hyperconan"))
	if err != nil {
		return nil, err
	}

	return &userpb.LoginResponse{
		Token: token,
	}, nil
}
