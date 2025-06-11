package http

import (
	"context"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/martinusiron/PayFlow/pkg/utils"
)

func AuditLogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		ip := getIP(r)
		reqID := r.Header.Get("X-Request-ID")
		if reqID == "" {
			reqID = uuid.NewString()
		}

		ctx = context.WithValue(ctx, utils.ContextKeyRequestID, reqID)
		ctx = context.WithValue(ctx, utils.ContextKeyIP, ip)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getIP(r *http.Request) string {
	ip := r.Header.Get("X-Forwarded-For")
	if ip == "" {
		ip = strings.Split(r.RemoteAddr, ":")[0]
	}
	return ip
}
