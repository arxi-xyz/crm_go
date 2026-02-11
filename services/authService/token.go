package authService

import (
	"context"
	"crm_go/entities"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type AccessClaims struct {
	jwt.RegisteredClaims
	TokenType string `json:"typ"`
}

type RefreshClaims struct {
	jwt.RegisteredClaims
	TokenType string `json:"typ"`
}

func (s *AuthService) generateTokens(user entities.User) (string, string, error) {
	now := time.Now()

	accessClaims := AccessClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    s.Config.Issuer,
			Subject:   user.UUID,
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
	refreshClaims := RefreshClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    s.Config.Issuer,
			Subject:   user.UUID,
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
		"uid": user.UUID,
		"iat": now.Unix(),
		"exp": now.Add(s.Config.RefreshTTL).Unix(),
	}

	ctx := context.Background()
	pipe := s.Cache.TxPipeline()

	pipe.HSet(ctx, "sess:"+jti, redisInfo)
	pipe.ExpireAt(ctx, "sess:"+jti, refreshClaims.ExpiresAt.Time)
	_, err = pipe.Exec(ctx)

	if err != nil {
		return "", "", err
	}

	return access, refresh, nil
}
