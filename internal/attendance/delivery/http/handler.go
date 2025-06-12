package http

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/martinusiron/PayFlow/internal/attendance/usecase"
	"github.com/martinusiron/PayFlow/internal/middleware"
	mdl "github.com/martinusiron/PayFlow/pkg/middleware"
)

type AttendanceHandler struct {
	Usecase *usecase.AttendanceUsecase
}

func NewAttendanceHandlerr(
	uc *usecase.AttendanceUsecase) *AttendanceHandler {
	return &AttendanceHandler{
		Usecase: uc,
	}
}

type SubmitAttendanceRequest struct {
	Date string `json:"date" example:"2025-06-10"`
}

// SubmitAttendance godoc
// @Summary Submit kehadiran harian
// @Description Digunakan oleh employee untuk submit kehadiran berdasarkan tanggal
// @Tags Attendance
// @Accept json
// @Produce json
// @Param X-Request-ID header string false "Unique request ID"
// @Security BearerAuth
// @Param body body SubmitAttendanceRequest true "Tanggal kehadiran (format YYYY-MM-DD)"
// @Success 201 {string} string "created"
// @Failure 400 {string} string "invalid body / invalid date format / other errors"
// @Failure 401 {string} string "unauthorized"
// @Router /api/attendance/submit [post]
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

	role := middleware.ExtractRole(ctx)
	if role != "employee" {
		http.Error(w, "forbidden", http.StatusForbidden)
		return
	}

	userID, ok := middleware.ExtractUserID(ctx)
	if !ok || userID == 0 {
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

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
