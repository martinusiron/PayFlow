package http

import (
	"encoding/json"
	"net/http"

	"github.com/martinusiron/PayFlow/internal/auth/domain"
	"github.com/martinusiron/PayFlow/internal/auth/usecase"
)

type AuthHandler struct{ UC *usecase.AuthUsecase }

func NewAuthHandler(uc *usecase.AuthUsecase) *AuthHandler {
	return &AuthHandler{UC: uc}
}

// Login godoc
// @Summary Login untuk mendapatkan token JWT
// @Description Autentikasi dengan username dan password
// @Tags Auth
// @Accept json
// @Produce json
// @Param credentials body domain.Credentials true "Login credentials"
// @Success 200 {object} domain.Token
// @Failure 400 {string} string "invalid body"
// @Failure 401 {string} string "unauthorized"
// @Router /api/auth/login [post]
func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var creds domain.Credentials
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}
	token, err := h.UC.Login(r.Context(), creds)
	if err != nil {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	json.NewEncoder(w).Encode(token)
}
