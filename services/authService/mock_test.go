package authService

import (
	"context"
	"crm_go/entities"
	"time"
)

type MockUserRepository struct {
	GetUserByFunc func(field, value string) (*entities.User, error)
}

func (m *MockUserRepository) GetUserBy(field, value string) (*entities.User, error) {
	return m.GetUserByFunc(field, value)
}

type MockCache struct {
	SetFunc   func(ctx context.Context, key string, value interface{}, expiration time.Time) error
	ExistFunc func(ctx context.Context, key string) (int64, error)
	DelFunc   func(ctx context.Context, key string) error
}

func (m *MockCache) Set(ctx context.Context, key string, value interface{}, expiration time.Time) error {
	return m.SetFunc(ctx, key, value, expiration)
}
func (m *MockCache) Exist(ctx context.Context, key string) (int64, error) {
	return m.ExistFunc(ctx, key)
}
func (m *MockCache) Del(ctx context.Context, key string) error {
	return m.DelFunc(ctx, key)
}
