package authorizationService

import "crm_go/pkg/appError"

func (s *AuthorizationService) HasPermission(userUuid string, field string) (bool, *appError.AppError) {
	/* Check from redis */
	return false, nil
}
