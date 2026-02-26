package authorizationService

import (
	"context"
	"crm_go/entities"
	"crm_go/pkg/appError"
	"time"
)

type AuthorizationService struct {
	Cache                redisClientInterface
	RoleRepository       roleRepositoryInterface
	PermissionRepository permissionRepositoryInterface
}

func New(cache redisClientInterface, roleRepository roleRepositoryInterface, permissionRepository permissionRepositoryInterface) *AuthorizationService {
	return &AuthorizationService{
		Cache:                cache,
		RoleRepository:       roleRepository,
		PermissionRepository: permissionRepository,
	}
}

type roleRepositoryInterface interface {
	GetUserRoleWithPermissions(userId int) ([]entities.Permission, *appError.AppError)
}

type permissionRepositoryInterface interface {
	GetUserPermissions(userId int) ([]entities.Permission, *appError.AppError)
}

type redisClientInterface interface {
	SetSet(ctx context.Context, key string, members []interface{}, expireAt time.Time) *appError.AppError
}
