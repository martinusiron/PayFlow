package utils

type contextKey string

const (
	ContextKeyUserID    contextKey = "user_id"
	ContextKeyRequestID contextKey = "request_id"
	ContextKeyIP        contextKey = "ip_address"
)
