package authService

import (
	"context"
	"crm_go/pkg/appError"
	"crm_go/pkg/validation"
)

func (s *AuthService) Refresh(request RefreshRequest) (RefreshResponse, error) {
	if err := validation.V().Struct(request); err != nil {
		return RefreshResponse{}, err
	}

	claims, appErr := s.ValidateRefreshToken(request.RefreshToken)
	if appErr != nil {
		return RefreshResponse{}, appErr
	}

	uUid, err := claims.GetSubject()
	if err != nil {
		return RefreshResponse{}, appError.Unauthorized(appError.UnauthorizedAccess, "invalid token", err)
	}

	err = s.Cache.Del(context.Background(), "sess:"+claims.ID)

	if err != nil {
		return RefreshResponse{}, appError.Internal(err)
	}

	token, refreshToken, appErr := s.generateTokens(uUid)
	if appErr != nil {
		return RefreshResponse{}, appErr
	}

	return RefreshResponse{
		Token:        token,
		RefreshToken: refreshToken,
	}, nil
}

type RefreshRequest struct {
	RefreshToken string `json:"refresh_token" validate:"required,jwt_token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
}

type RefreshResponse struct {
	Token        string `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
	RefreshToken string `json:"refresh_token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
}
