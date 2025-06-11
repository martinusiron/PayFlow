package middleware

import (
	"net"
	"net/http"
	"strings"
)

func GetIPAddress(r *http.Request) string {
	ip := r.Header.Get("X-Forwarded-For")
	if ip == "" {
		ip, _, _ = net.SplitHostPort(r.RemoteAddr)
	} else {
		ip = strings.Split(ip, ",")[0]
	}
	return strings.TrimSpace(ip)
}
