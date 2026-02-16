package authService

import (
	"context"
	"crm_go/pkg/validation"
)

func (s *AuthService) Logout(request LogoutRequest) error {
	if err := validation.V().Struct(request); err != nil {
		return err
	}

	claims, err := s.ValidateRefreshToken(request.RefreshToken)
	if err != nil {
		return err
	}

	s.Cache.Del(context.Background(), "sess:"+claims.ID)

	return nil
}

type LogoutRequest struct {
	RefreshToken string `json:"refresh_token" validate:"required,jwt_token"`
}
