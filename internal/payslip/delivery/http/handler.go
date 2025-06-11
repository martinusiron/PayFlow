package http

import (
	"encoding/json"
	"net/http"

	"github.com/martinusiron/PayFlow/internal/middleware"
	"github.com/martinusiron/PayFlow/internal/payslip/usecase"
)

type PayslipHandler struct {
	Usecase *usecase.PayslipUsecase
}

func NewPayslipHandler(uc *usecase.PayslipUsecase) *PayslipHandler {
	return &PayslipHandler{Usecase: uc}
}

func (h *PayslipHandler) GetMyPayslips(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID := middleware.ExtractUserID(ctx)
	if userID == 0 {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	payslips, err := h.Usecase.GenerateLatestPayslip(ctx, userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(payslips)
}
