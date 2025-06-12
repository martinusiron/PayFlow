package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

type contextKey string

const (
	ContextKeyUserID    contextKey = "user_id"
	ContextKeyRequestID contextKey = "request_id"
	ContextKeyIP        contextKey = "ip_address"
)

func HashPassword(pw string) string {
	hash := sha256.Sum256([]byte(pw))
	return hex.EncodeToString(hash[:])
}
