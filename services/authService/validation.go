package authService

import (
	"context"
	"crm_go/pkg/appError"
	"time"
)

func (s *AuthService) ValidateRefreshToken(tokenString string) (*TokenClaims, error) {
	claims, err := s.parseToken(tokenString)
	if err != nil {
		return nil, appError.Unauthorized("invalid_token", "invalid_token", err)
	}

	if claims.TokenType != "refresh" {
		return nil, appError.Unauthorized("invalid_token", "invalid_token_type", nil)
	}

	if err := s.validateExpireTime(claims); err != nil {
		return nil, err
	}

	if err := s.validateSessionExists(claims); err != nil {
		return nil, err
	}

	return claims, nil
}

func (s *AuthService) ValidateAccessToken(tokenString string) (*TokenClaims, error) {
	claims, err := s.parseToken(tokenString)
	if err != nil {
		return nil, appError.Unauthorized("invalid_token", "invalid_token", err)
	}

	if claims.TokenType != "access" {
		return nil, appError.Unauthorized("invalid_token", "invalid_token_type", nil)
	}

	if err := s.validateExpireTime(claims); err != nil {
		return nil, err
	}

	return claims, nil
}

func (s *AuthService) validateExpireTime(claims *TokenClaims) error {
	expirationTime, err := claims.GetExpirationTime()
	if err != nil {
		return appError.Unauthorized("invalid_token", "invalid_token", err)
	}

	if expirationTime == nil {
		return appError.Unauthorized("invalid_token", "missing_exp", nil)
	}

	if expirationTime.Before(time.Now()) {
		return appError.Unauthorized("invalid_token", "token_expired", nil)
	}

	return nil
}

func (s *AuthService) validateSessionExists(claims *TokenClaims) error {
	cmd := s.Cache.Exists(context.Background(), "sess:"+claims.ID)

	if err := cmd.Err(); err != nil {
		return appError.Internal(err)
	}

	if cmd.Val() == 0 {
		return appError.Unauthorized("invalid_token", "invalid_token", nil)
	}

	return nil
}
