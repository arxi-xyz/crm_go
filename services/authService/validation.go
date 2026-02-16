package authService

import (
	"context"
	"crm_go/pkg/appError"
	"time"
)

func (s *AuthService) ValidateRefreshToken(tokenClaims *RefreshClaims) error {
	if tokenClaims.TokenType != "refresh" {
		return appError.Unauthorized("invalid_token", "invalid_token", nil)
	}

	err := s.validateExpireTime(tokenClaims)

	if err != nil {
		return err
	}

	err = s.validateRefreshTokenExistence(tokenClaims)

	if err != nil {
		return err
	}

	return nil
}

func (s *AuthService) ValidateToken(tokenClaims *RefreshClaims) error {
	if tokenClaims.TokenType != "token" {
		return appError.Unauthorized("invalid_token", "invalid_token", nil)
	}

	err := s.validateExpireTime(tokenClaims)

	if err != nil {
		return err
	}

	err = s.validateRefreshTokenExistence(tokenClaims)

	if err != nil {
		return err
	}

	return nil
}

func (s *AuthService) validateExpireTime(tokenClaims *RefreshClaims) error {
	expirationTime, err := tokenClaims.GetExpirationTime()
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

func (s *AuthService) validateRefreshTokenExistence(tokenClaims *RefreshClaims) error {
	cmd := s.Cache.Exists(context.Background(), "sess:"+tokenClaims.ID)

	if err := cmd.Err(); err != nil {
		return appError.Internal(err)
	}

	if cmd.Val() == 0 {
		return appError.Unauthorized("invalid_token", "invalid_token", nil)
	}

	return nil
}
