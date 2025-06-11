package domain

import "time"

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Token struct {
	AccessToken  string    `json:"accessToken"`
	RefreshToken string    `json:"refreshToken"`
	ExpiresAt    time.Time `json:"expiresAt"`
}
