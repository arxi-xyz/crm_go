package authorizationService

import (
	"context"
	"crm_go/entities"
	"crm_go/pkg/appError"
	"fmt"
	"time"
)

func (s *AuthorizationService) SetPermissionsToCache(User *entities.User) *appError.AppError {
	rolePermissions, appErr := s.RoleRepository.GetUserRoleWithPermissions(int(User.ID))
	if appErr != nil {
		return appErr
	}

	userPermissions, appErr := s.PermissionRepository.GetUserPermissions(int(User.ID))
	if appErr != nil {
		return appErr
	}

	allPerms := append(rolePermissions, userPermissions...)

	cacheItems := make([]interface{}, 0, len(allPerms))
	for _, p := range allPerms {
		cacheItems = append(cacheItems, p.UniqueKey)
	}

	err := s.Cache.SetSet(
		context.Background(),
		fmt.Sprintf("perm:%s", User.UUID),
		cacheItems,
		time.Now().Add(time.Hour*24*30),
	)

	if err != nil {
		return appError.Internal(err)
	}

	return nil
}

type cachePermission struct {
	UUID      string `json:"uuid"`
	UniqueKey string `json:"unique_key"`
}
