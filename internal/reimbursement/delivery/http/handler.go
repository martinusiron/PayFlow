package http

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/martinusiron/PayFlow/internal/middleware"
	"github.com/martinusiron/PayFlow/internal/reimbursement/usecase"
	mdl "github.com/martinusiron/PayFlow/pkg/middleware"
)

type ReimbursementHandler struct {
	Usecase *usecase.ReimbursementUsecase
}

func NewReimbursementHandler(uc *usecase.ReimbursementUsecase) *ReimbursementHandler {
	return &ReimbursementHandler{Usecase: uc}
}

type SubmitReimbursementRequest struct {
	Date        string  `json:"date" example:"2025-06-12"`
	Amount      float64 `json:"amount" example:"1000000"`
	Description string  `json:"description" example:"Medical reimbursement`
}

// SubmitReimbursement godoc
// @Summary Submit reimbursement
// @Description Digunakan oleh employee untuk mengajukan reimbursement berdasarkan tanggal, jumlah dan deskripsi
// @Tags Reimbursement
// @Accept json
// @Produce json
// @Param X-Request-ID header string false "Unique request ID"
// @Security BearerAuth
// @Param body body SubmitReimbursementRequest true "Data reimbursement (tanggal, jumlah, deskripsi)"
// @Success 201 {string} string "created"
// @Failure 400 {string} string "invalid body / invalid date format / other errors"
// @Failure 401 {string} string "unauthorized"
// @Router /api/reimbursement/submit [post]
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

	err = h.Usecase.Submit(ctx, userID, date, req.Amount, req.Description, userID, ip, reqID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
