package authService

import (
	"context"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type TokenClaims struct {
	jwt.RegisteredClaims
	TokenType string `json:"typ"`
}

func (s *AuthService) generateTokens(uUid string) (string, string, error) {
	now := time.Now()

	accessClaims := TokenClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    s.Config.Issuer,
			Subject:   uUid,
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(s.Config.AccessTTL)),
		},
		TokenType: "access",
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	access, err := accessToken.SignedString(s.Config.JWTSecret)
	if err != nil {
		return "", "", err
	}

	jti := uuid.NewString()
	refreshClaims := TokenClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    s.Config.Issuer,
			Subject:   uUid,
			ID:        jti,
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(s.Config.RefreshTTL)),
		},
		TokenType: "refresh",
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refresh, err := refreshToken.SignedString(s.Config.JWTSecret)
	if err != nil {
		return "", "", err
	}

	redisInfo := map[string]any{
		"uid": uUid,
		"iat": now.Unix(),
		"exp": now.Add(s.Config.RefreshTTL).Unix(),
	}

	err = s.Cache.Set(context.Background(), "sess:"+jti, redisInfo, refreshClaims.ExpiresAt.Time)

	if err != nil {
		return "", "", err
	}

	return access, refresh, nil
}

func (s *AuthService) parseToken(tokenStr string) (*TokenClaims, error) {
	claims := &TokenClaims{}
	tok, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		if token.Method != jwt.SigningMethodHS256 {
			return nil, errors.New("unexpected signing method")
		}
		return s.Config.JWTSecret, nil
	})
	if err != nil {
		return nil, err
	}
	if !tok.Valid {
		return nil, errors.New("invalid token")
	}
	return claims, nil
}
