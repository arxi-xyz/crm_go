package authService

import (
	"context"
	"crm_go/pkg/appError"
	"crm_go/pkg/validation"
	"time"
)

func (s *AuthService) Refresh(request RefreshRequest) (RefreshResponse, error) {
	if err := validation.V().Struct(request); err != nil {
		return RefreshResponse{}, err
	}

	tokenClaims, err := s.parseRefreshClaim(request.RefreshToken)
	if err != nil {
		return RefreshResponse{}, appError.Unauthorized("invalid_token", "invalid_token", err)
	}

	if tokenClaims.TokenType != "refresh" {
		return RefreshResponse{}, appError.Unauthorized("invalid_token", "invalid_token", nil)
	}
	expirationTime, err := tokenClaims.GetExpirationTime()
	if err != nil {
		return RefreshResponse{}, appError.Unauthorized("invalid_token", "invalid_token", err)
	}

	if expirationTime == nil {
		return RefreshResponse{}, appError.Unauthorized("invalid_token", "missing_exp", nil)
	}

	if expirationTime.Before(time.Now()) {
		return RefreshResponse{}, appError.Unauthorized("invalid_token", "token_expired", nil)
	}

	uUid, err := tokenClaims.GetSubject()

	if err != nil {
		return RefreshResponse{}, appError.Unauthorized("invalid_token", "invalid_token", err)
	}

	cmd := s.Cache.Exists(context.Background(), "sess:"+tokenClaims.ID)

	if err := cmd.Err(); err != nil {
		return RefreshResponse{}, err
	}

	if cmd.Val() == 0 {
		return RefreshResponse{}, appError.Unauthorized("invalid_token", "invalid_token", nil)
	}

	s.Cache.Del(context.Background(), "sess:"+tokenClaims.ID)

	token, refreshToken, err := s.generateTokens(uUid)

	return RefreshResponse{
		Token:        token,
		RefreshToken: refreshToken,
	}, nil
}

type RefreshRequest struct {
	RefreshToken string `json:"refresh_token" validate:"required,jwt_token"`
}

type RefreshResponse struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}
