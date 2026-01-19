package logic

import (
	"context"
	"errors"

	"user/internal/models"
	"user/internal/svc"
	userpb "user/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
	"golang.org/x/crypto/bcrypt"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *userpb.RegisterRequest) (*userpb.RegisterResponse, error) {
	// Check if username already exists
	var existingUser models.User
	if err := l.svcCtx.DB.Where("username = ?", in.Username).First(&existingUser).Error; err == nil {
		return nil, errors.New("username already exists")
	}

	// Check if email already exists
	if err := l.svcCtx.DB.Where("email = ?", in.Email).First(&existingUser).Error; err == nil {
		return nil, errors.New("email already exists")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// Create user
	newUser := &models.User{
		Username: in.Username,
		Password: string(hashedPassword),
		Email:    in.Email,
	}

	if err := l.svcCtx.DB.Create(newUser).Error; err != nil {
		return nil, err
	}

	return &userpb.RegisterResponse{
		Message: "User registered successfully",
	}, nil
}
