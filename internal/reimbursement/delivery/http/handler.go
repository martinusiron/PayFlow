package http

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/martinusiron/PayFlow/internal/reimbursement/usecase"
)

type ReimbursementHandler struct {
	Usecase *usecase.ReimbursementUsecase
}

func NewReimbursementHandler(uc *usecase.ReimbursementUsecase) *ReimbursementHandler {
	return &ReimbursementHandler{Usecase: uc}
}

type SubmitReimbursementRequest struct {
	Date        string  `json:"date"`
	Amount      float64 `json:"amount"`
	Description string  `json:"description"`
}

func (h *ReimbursementHandler) SubmitReimbursement(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req SubmitReimbursementRequest
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

	err = h.Usecase.Submit(ctx, userID, date, req.Amount, req.Description, userID, ip, reqID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
