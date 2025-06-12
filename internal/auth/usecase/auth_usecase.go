package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/martinusiron/PayFlow/internal/auth/domain"
	"github.com/martinusiron/PayFlow/internal/auth/repository"
	ur "github.com/martinusiron/PayFlow/internal/user/repository"
	"github.com/martinusiron/PayFlow/pkg/utils"
)

type AuthUsecase struct {
	authRepo  repository.AuthRepository
	userRepo  ur.UserRepository
	jwtSecret string
}

func NewAuthUsecase(
	authRepo repository.AuthRepository,
	userRepo ur.UserRepository,
	jwtSecret string) *AuthUsecase {
	return &AuthUsecase{
		authRepo:  authRepo,
		userRepo:  userRepo,
		jwtSecret: jwtSecret,
	}
}

func (u *AuthUsecase) Login(ctx context.Context, creds domain.Credentials) (*domain.Token, error) {
	hash, err := u.authRepo.GetPasswordHash(ctx, creds.Username)
	if err != nil {
		return nil, err
	}

	inputHash := utils.HashPassword(creds.Password)
	if inputHash != hash {
		return nil, errors.New("invalid credentials")
	}

	user, err := u.userRepo.GetByUsername(ctx, creds.Username)
	if err != nil {
		return nil, errors.New("user not found")
	}

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": creds.Username,
		"user_id":  user.ID,
		"role":     user.Role,
		"exp":      time.Now().Add(15 * time.Minute).Unix(),
	})

	accessToken, err := at.SignedString([]byte(u.jwtSecret))
	if err != nil {
		return nil, err
	}

	token := &domain.Token{
		AccessToken: accessToken,
		ExpiresAt:   time.Now().Add(60 * time.Minute),
	}
	return token, nil
}

func (u *AuthUsecase) VerifyAccessToken(ctx context.Context, tokenString string) (int, string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(u.jwtSecret), nil
	})

	if err != nil {
		return 0, "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userIDFloat, ok := claims["user_id"].(float64)
		if !ok {
			return 0, "", errors.New("invalid user_id claim")
		}

		role, ok := claims["role"].(string)
		if !ok {
			return 0, "", errors.New("invalid role claim")
		}

		return int(userIDFloat), role, nil
	}

	return 0, "", errors.New("invalid token")
}
