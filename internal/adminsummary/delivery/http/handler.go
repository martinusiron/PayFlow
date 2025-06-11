package http

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/martinusiron/PayFlow/internal/adminsummary/usecase"
)

type AdminSummaryHandler struct {
	Usecase *usecase.AdminSummaryUsecase
}

func NewAdminSummaryHandler(uc *usecase.AdminSummaryUsecase) *AdminSummaryHandler {
	return &AdminSummaryHandler{Usecase: uc}
}

func (h *AdminSummaryHandler) GenerateSummary(w http.ResponseWriter, r *http.Request) {
	payrollIDStr := r.URL.Query().Get("payroll_id")
	payrollID, err := strconv.Atoi(payrollIDStr)
	if err != nil {
		http.Error(w, "invalid payroll_id", http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	summary, err := h.Usecase.GenerateSummary(ctx, payrollID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(summary)
}
