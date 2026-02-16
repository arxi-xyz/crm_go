package middlewares

import "crm_go/services/authService"

type authServiceInterface interface {
	ValidateAccessToken(tokenString string) (*authService.TokenClaims, error)
}
