package http

import (
	"encoding/json"
	"net/http"

	"github.com/martinusiron/PayFlow/internal/middleware"
	"github.com/martinusiron/PayFlow/internal/payroll/usecase"
	mdl "github.com/martinusiron/PayFlow/pkg/middleware"
)

type PayrollHandler struct {
	Usecase *usecase.PayrollUsecase
}

func NewPayrollHandler(uc *usecase.PayrollUsecase) *PayrollHandler {
	return &PayrollHandler{Usecase: uc}
}

type RunPayrollRequest struct {
	StartDate string `json:"start_date" example:"2025-06-01"`
	EndDate   string `json:"end_date" example:"2025-06-30"`
}

// RunPayroll godoc
// @Summary Jalankan proses payroll untuk semua karyawan
// @Description Hanya dapat dijalankan oleh admin untuk memproses payroll berdasarkan rentang tanggal
// @Tags Payroll
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param body body RunPayrollRequest true "Tanggal mulai dan akhir payroll"
// @Success 201 {string} string "created"
// @Failure 400 {string} string "invalid request body / processing error"
// @Failure 401 {string} string "unauthorized"
// @Router /api/payroll/run [post]
func (h *PayrollHandler) RunPayroll(w http.ResponseWriter, r *http.Request) {
	var req RunPayrollRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	role := middleware.ExtractRole(ctx)
	if role != "admin" {
		http.Error(w, "forbidden", http.StatusForbidden)
		return
	}
	adminID, ok := middleware.ExtractUserID(ctx)
	if !ok || adminID == 0 {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}
	ip := mdl.GetIPAddress(r)
	reqID := r.Header.Get("X-Request-ID")

	err := h.Usecase.RunPayroll(ctx, req.StartDate, req.EndDate, adminID, ip, reqID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
