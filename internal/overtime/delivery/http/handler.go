package http

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/martinusiron/PayFlow/internal/middleware"
	"github.com/martinusiron/PayFlow/internal/overtime/usecase"
	mdl "github.com/martinusiron/PayFlow/pkg/middleware"
)

type OvertimeHandler struct {
	Usecase *usecase.OvertimeUsecase
}

func NewOvertimeHandler(uc *usecase.OvertimeUsecase) *OvertimeHandler {
	return &OvertimeHandler{Usecase: uc}
}

type SubmitOvertimeRequest struct {
	Date  string  `json:"date" example:"2025-06-11"`
	Hours float64 `json:"hours" example:"2.5"`
}

// SubmitOvertime godoc
// @Summary Submit lembur
// @Description Digunakan oleh employee untuk mencatat lembur berdasarkan tanggal dan jumlah jam
// @Tags Overtime
// @Accept json
// @Produce json
// @Param X-Request-ID header string false "Unique request ID"
// @Security BearerAuth
// @Param body body SubmitOvertimeRequest true "Tanggal lembur dan jumlah jam lembur"
// @Success 201 {string} string "created"
// @Failure 400 {string} string "invalid body / invalid date format / other errors"
// @Failure 401 {string} string "unauthorized"
// @Router /api/overtime/submit [post]
func (h *OvertimeHandler) SubmitOvertime(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req SubmitOvertimeRequest
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

	err = h.Usecase.Submit(ctx, userID, date, req.Hours, userID, ip, reqID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
