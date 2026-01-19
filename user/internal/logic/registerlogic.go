package logic

import (
	"context"
	"errors"
	"sync"

	"user/internal/svc"
	userpb "user/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
	"golang.org/x/crypto/bcrypt"
)

// Simple in-memory user storage
var (
	users   = make(map[string]*User)
	usersMu sync.RWMutex
	userID  uint32 = 1
)

type User struct {
	ID       uint32
	Username string
	Password string
	Email    string
}

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
	usersMu.Lock()
	defer usersMu.Unlock()

	// Check if username already exists
	if _, exists := users[in.Username]; exists {
		return nil, errors.New("username already exists")
	}

	// Check if email already exists
	for _, u := range users {
		if u.Email == in.Email {
			return nil, errors.New("email already exists")
		}
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// Create user
	newUser := &User{
		ID:       userID,
		Username: in.Username,
		Password: string(hashedPassword),
		Email:    in.Email,
	}
	users[in.Username] = newUser
	userID++

	return &userpb.RegisterResponse{
		Message: "User registered successfully",
	}, nil
}