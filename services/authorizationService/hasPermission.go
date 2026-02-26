package authorizationService

import (
	"context"
	"crm_go/pkg/appError"
	"fmt"
)

func (s *AuthorizationService) HasPermission(userUuid string, perm string) (bool, *appError.AppError) {
	key := fmt.Sprintf("perm:%s", userUuid)

	res, err := s.Cache.IsMember(context.Background(), key, perm)

	if err != nil {
		return false, err
	}

	return res, err
}
