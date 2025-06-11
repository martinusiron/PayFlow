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
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}

func (h *PayrollHandler) RunPayroll(w http.ResponseWriter, r *http.Request) {
	var req RunPayrollRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	adminID := middleware.ExtractUserID(ctx)
	if adminID == 0 {
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

	w.WriteHeader(http.StatusCreated)
}
