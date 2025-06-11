package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/martinusiron/PayFlow/internal/auth/domain"
	"github.com/martinusiron/PayFlow/internal/auth/repository"
	"golang.org/x/crypto/bcrypt"
)

type AuthUsecase struct {
	repo      repository.AuthRepository
	jwtSecret string
}

func NewAuthUsecase(repo repository.AuthRepository, jwtSecret string) *AuthUsecase {
	return &AuthUsecase{repo: repo, jwtSecret: jwtSecret}
}

func (u *AuthUsecase) Login(ctx context.Context, creds domain.Credentials) (*domain.Token, error) {
	hash, err := u.repo.GetPasswordHash(ctx, creds.Username)
	if err != nil {
		return nil, err
	}
	if bcrypt.CompareHashAndPassword([]byte(hash), []byte(creds.Password)) != nil {
		return nil, errors.New("invalid credentials")
	}
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": creds.Username,
		"exp":      time.Now().Add(15 * time.Minute).Unix(),
	})
	accessToken, err := at.SignedString(u.jwtSecret)
	if err != nil {
		return nil, err
	}
	token := &domain.Token{AccessToken: accessToken, ExpiresAt: time.Now().Add(15 * time.Minute)}
	return token, nil
}

func (u *AuthUsecase) VerifyAccessToken(ctx context.Context, tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return u.jwtSecret, nil
	})

	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		username, ok := claims["username"].(string)
		if !ok {
			return "", errors.New("invalid claims")
		}
		return username, nil
	}

	return "", errors.New("invalid token")
}
