package http

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/martinusiron/PayFlow/internal/adminsummary/usecase"
	"github.com/martinusiron/PayFlow/internal/middleware"
)

type AdminSummaryHandler struct {
	Usecase *usecase.AdminSummaryUsecase
}

func NewAdminSummaryHandler(uc *usecase.AdminSummaryUsecase) *AdminSummaryHandler {
	return &AdminSummaryHandler{Usecase: uc}
}

// GenerateSummary godoc
// @Summary Get admin payroll summary
// @Description Generate payroll summary for a given payroll_id
// @Tags AdminSummary
// @Accept  json
// @Produce  json
// @Param payroll_id query int true "Payroll ID"
// @Success 200 {object} []domain.FullSummary
// @Failure 400 {string} string "invalid payroll_id"
// @Failure 500 {string} string "internal server error"
// @Security BearerAuth
// @Router /api/summary/admin [get]
func (h *AdminSummaryHandler) GenerateSummary(w http.ResponseWriter, r *http.Request) {
	payrollIDStr := r.URL.Query().Get("payroll_id")
	payrollID, err := strconv.Atoi(payrollIDStr)
	if err != nil {
		http.Error(w, "invalid payroll_id", http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	role := middleware.ExtractRole(ctx)
	if role != "admin" {
		http.Error(w, "forbidden", http.StatusForbidden)
		return
	}

	summary, err := h.Usecase.GenerateSummary(ctx, payrollID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(summary)
}
