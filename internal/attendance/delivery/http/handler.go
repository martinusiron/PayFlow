package http

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/martinusiron/PayFlow/internal/attendance/usecase"
	"github.com/martinusiron/PayFlow/internal/middleware"
	"github.com/martinusiron/PayFlow/internal/shared"
	mdl "github.com/martinusiron/PayFlow/pkg/middleware"
)

type AttendanceHandler struct {
	Usecase *usecase.AttendanceUsecase
	Shared  *shared.Service
}

func NewAttendanceHandlerr(
	uc *usecase.AttendanceUsecase,
	shared *shared.Service) *AttendanceHandler {
	return &AttendanceHandler{
		Usecase: uc,
		Shared:  shared,
	}
}

type SubmitAttendanceRequest struct {
	Date string `json:"date"` // e.g. "2025-06-10"
}

func (h *AttendanceHandler) SubmitAttendance(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req SubmitAttendanceRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}

	date, err := time.Parse("2006-01-02", req.Date)
	if err != nil {
		http.Error(w, "invalid date format", http.StatusBadRequest)
		return
	}

	userID := middleware.ExtractUserID(ctx)
	if userID == 0 {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}
	ip := mdl.GetIPAddress(r)
	reqID := r.Header.Get("X-Request-ID")

	err = h.Usecase.Submit(ctx, userID, date, userID, ip, reqID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
