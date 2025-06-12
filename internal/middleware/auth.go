package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/martinusiron/PayFlow/internal/auth/usecase"
)

type AuthMiddleware struct {
	AuthUC *usecase.AuthUsecase
}

func NewAuthMiddleware(uc *usecase.AuthUsecase) *AuthMiddleware {
	return &AuthMiddleware{AuthUC: uc}
}

type contextKey string

const (
	userIDKey contextKey = "user_id"
	roleKey   contextKey = "role"
)

func (m *AuthMiddleware) JWTAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		token := strings.TrimPrefix(authHeader, "Bearer ")
		userID, role, err := m.AuthUC.VerifyAccessToken(r.Context(), token)
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(r.Context(), userIDKey, userID)
		ctx = context.WithValue(ctx, roleKey, role)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// ExtractUserID helper
func ExtractUserID(ctx context.Context) (int, bool) {
	id, ok := ctx.Value(userIDKey).(int)
	return id, ok
}

func ExtractRole(ctx context.Context) string {
	role, _ := ctx.Value(roleKey).(string)
	return role
}
