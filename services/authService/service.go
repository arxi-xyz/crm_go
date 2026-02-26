package authService

import (
	"context"
	"crm_go/entities"
	"crm_go/pkg/appError"
	"time"
)

const (
	REFRESH = "refresh"
	ACCESS  = "access"
)

type AuthService struct {
	UserRepository       userRepositoryInterface
	AuthorizationService authorizationServiceInterface
	Config               Config
	Cache                redisClientInterface
}

func New(userRepository userRepositoryInterface, authorizationService authorizationServiceInterface, cache redisClientInterface, config Config) *AuthService {
	return &AuthService{
		UserRepository:       userRepository,
		Config:               config,
		Cache:                cache,
		AuthorizationService: authorizationService,
	}
}

type authorizationServiceInterface interface {
	SetPermissionsToCache(User *entities.User) *appError.AppError
}

type userRepositoryInterface interface {
	GetUserBy(field, value string) (*entities.User, error)
}

type redisClientInterface interface {
	Set(ctx context.Context, key string, value interface{}, expiration time.Time) error
	Exist(ctx context.Context, key string) (int64, error)
	Del(ctx context.Context, key string) error
}
type Config struct {
	JWTSecret  []byte
	AccessTTL  time.Duration
	RefreshTTL time.Duration
	Issuer     string
}
