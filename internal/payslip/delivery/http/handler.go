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

// GetMyPayslips godoc
// @Summary Ambil slip gaji terakhir milik user
// @Description Mengambil slip gaji terbaru berdasarkan user yang sedang login
// @Tags Payslip
// @Produce json
// @Security BearerAuth
// @Success 200 {object} []domain.Payslip "daftar payslip terbaru"
// @Failure 401 {string} string "unauthorized"
// @Failure 500 {string} string "internal server error"
// @Router /api/payslip/get [get]
func (h *PayslipHandler) GetMyPayslips(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID, ok := middleware.ExtractUserID(ctx)
	if !ok || userID == 0 {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	role := middleware.ExtractRole(ctx)
	if role != "employee" {
		http.Error(w, "forbidden", http.StatusForbidden)
		return
	}

	payslips, err := h.Usecase.GenerateLatestPayslip(ctx, 1)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(payslips)
}
