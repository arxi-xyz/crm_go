package authService

import (
	"context"
	"crm_go/pkg/appError"
	"time"
)

func (s *AuthService) ValidateRefreshToken(tokenString string) (*TokenClaims, *appError.AppError) {
	claims, err := s.parseToken(tokenString)
	if err != nil {
		return nil, err
	}

	if claims.TokenType != "refresh" {
		return nil, appError.Unauthorized(appError.UnauthorizedAccess, "invalid token", nil)
	}

	if err = s.validateExpireTime(claims); err != nil {
		return nil, err
	}

	if err = s.validateSessionExists(claims); err != nil {
		return nil, err
	}

	return claims, nil
}

func (s *AuthService) ValidateAccessToken(tokenString string) (*TokenClaims, *appError.AppError) {
	claims, err := s.parseToken(tokenString)
	if err != nil {
		return nil, appError.Unauthorized(appError.UnauthorizedAccess, "invalid token", err)
	}

	if claims.TokenType != "access" {
		return nil, appError.Unauthorized(appError.UnauthorizedAccess, "invalid token", nil)
	}

	if err := s.validateExpireTime(claims); err != nil {
		return nil, err
	}

	return claims, nil
}

func (s *AuthService) validateExpireTime(claims *TokenClaims) *appError.AppError {
	expirationTime, err := claims.GetExpirationTime()
	if err != nil {
		return appError.Unauthorized(appError.UnauthorizedAccess, "invalid token", nil)
	}

	if expirationTime == nil {
		return appError.Unauthorized(appError.UnauthorizedAccess, "invalid token", nil)
	}

	if expirationTime.Before(time.Now()) {
		return appError.Unauthorized(appError.UnauthorizedAccess, "invalid token", nil)
	}

	return nil
}

func (s *AuthService) validateSessionExists(claims *TokenClaims) *appError.AppError {
	exist, err := s.Cache.Exist(context.Background(), "sess:"+claims.ID)
	if err != nil {
		return appError.Internal(err)
	}

	if exist == 0 {
		return appError.Unauthorized(appError.UnauthorizedAccess, "invalid token", nil)
	}

	return nil
}
