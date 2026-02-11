package authService

import (
	"crm_go/entities"
	"time"

	"github.com/redis/go-redis/v9"
)

type AuthService struct {
	UserRepository userRepositoryInterface
	Config         Config
	Cache          redis.Client
	// todo: cache is completely coupled
}

func New(userRepository userRepositoryInterface, cache redis.Client, config Config) *AuthService {
	return &AuthService{
		UserRepository: userRepository,
		Config:         config,
		Cache:          cache,
	}
}

type userRepositoryInterface interface {
	GetUserByPhone(phone string) (*entities.User, error)
}
type Config struct {
	JWTSecret  []byte
	AccessTTL  time.Duration
	RefreshTTL time.Duration
	Issuer     string
}
