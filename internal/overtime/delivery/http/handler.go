package http

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/martinusiron/PayFlow/internal/overtime/usecase"
)

type OvertimeHandler struct {
	Usecase *usecase.OvertimeUsecase
}

func NewOvertimeHandler(uc *usecase.OvertimeUsecase) *OvertimeHandler {
	return &OvertimeHandler{Usecase: uc}
}

type SubmitOvertimeRequest struct {
	Date  string  `json:"date"`
	Hours float64 `json:"hours"`
}

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

	userID := 1
	ip := r.RemoteAddr
	reqID := r.Header.Get("X-Request-ID")

	err = h.Usecase.Submit(ctx, userID, date, req.Hours, userID, ip, reqID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
