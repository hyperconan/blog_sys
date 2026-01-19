package middleware

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v4"
	"github.com/zeromicro/go-zero/rest/httpx"
)

type userIdKey struct{}

type JwtMiddleware struct {
	Secret string
}

func NewJwtMiddleware(secret string) *JwtMiddleware {
	return &JwtMiddleware{
		Secret: secret,
	}
}

func (m *JwtMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get token from Authorization header
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			httpx.Error(w, errors.New("authorization header is required"))
			return
		}

		// Check Bearer token format
		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			httpx.Error(w, errors.New("authorization header format must be Bearer {token}"))
			return
		}

		tokenString := parts[1]

		// Parse and validate token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("unexpected signing method")
			}
			return []byte(m.Secret), nil
		})

		if err != nil {
			httpx.Error(w, err)
			return
		}

		if !token.Valid {
			httpx.Error(w, errors.New("invalid token"))
			return
		}

		// Extract claims and add user ID to context
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			if userIdFloat, ok := claims["id"].(float64); ok {
				userId := uint32(userIdFloat)
				ctx := context.WithValue(r.Context(), userIdKey{}, userId)
				r = r.WithContext(ctx)
			}
		}

		next(w, r)
	}
}

// GetUserIdFromContext extracts user ID from context
func GetUserIdFromContext(ctx context.Context) (uint32, error) {
	userId, ok := ctx.Value(userIdKey{}).(uint32)
	if !ok {
		return 0, errors.New("user ID not found in context")
	}
	return userId, nil
}
