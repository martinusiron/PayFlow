package http

import (
	"context"
	"net/http"
	"strings"

	"github.com/martinusiron/PayFlow/internal/auth/usecase"
)

type JWTMiddleware struct{ UC *usecase.AuthUsecase }

func NewJWTMiddleware(uc *usecase.AuthUsecase) *JWTMiddleware {
	return &JWTMiddleware{UC: uc}
}

func (m *JWTMiddleware) Verify(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		if !strings.HasPrefix(auth, "Bearer ") {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}
		tokenStr := strings.TrimPrefix(auth, "Bearer ")
		username, err := m.UC.VerifyAccessToken(r.Context(), tokenStr)
		if err != nil {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}

		// Attach username to context
		ctx := context.WithValue(r.Context(), "username", username)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
