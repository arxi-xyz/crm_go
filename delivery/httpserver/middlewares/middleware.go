package middlewares

import (
	"crm_go/pkg/appError"
	"crm_go/services/authService"
)

type authServiceInterface interface {
	ValidateAccessToken(tokenString string) (*authService.TokenClaims, *appError.AppError)
}

type AuthorizationServiceInterface interface {
	HasPermission(userUuid string, perm string) (bool, *appError.AppError)
}
