package authService

import (
	"context"
	"crm_go/pkg/appError"
	"crm_go/pkg/validation"
	"time"
)

func (s *AuthService) Logout(request LogoutRequest) error {
	if err := validation.V().Struct(request); err != nil {
		return err
	}

	tokenClaims, err := s.parseRefreshClaim(request.RefreshToken)
	if err != nil {
		return appError.Unauthorized("invalid_token", "invalid_token", err)
	}

	if tokenClaims.TokenType != "refresh" {
		return appError.Unauthorized("invalid_token", "invalid_token", nil)
	}
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

	cmd := s.Cache.Exists(context.Background(), "sess:"+tokenClaims.ID)

	if err := cmd.Err(); err != nil {
		return err
	}

	if cmd.Val() == 0 {
		return appError.Unauthorized("invalid_token", "invalid_token", nil)
	}

	s.Cache.Del(context.Background(), "sess:"+tokenClaims.ID)

	return nil
}

type LogoutRequest struct {
	RefreshToken string `json:"refresh_token" validate:"required,jwt_token"`
}
