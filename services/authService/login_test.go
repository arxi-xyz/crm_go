package authService

import (
	"context"
	"crm_go/entities"
	"crm_go/pkg/validation"
	"database/sql"
	"fmt"
	"testing"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func TestLogin(t *testing.T) {
	validation.Init()

	mockConfig := Config{
		JWTSecret:  []byte("test_secret"),
		AccessTTL:  1 * time.Minute,
		RefreshTTL: 10 * time.Minute,
		Issuer:     "crm",
	}

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte("mypassword"),
		bcrypt.DefaultCost,
	)

	if err != nil {
		t.Fatal(err)
	}

	mockedUser := entities.User{
		ID:        0,
		UUID:      uuid.NewString(),
		Phone:     "09130108631",
		Password:  string(hashedPassword),
		FirstName: sql.NullString{String: "ah", Valid: true},
		LastName:  sql.NullString{String: "Sharif", Valid: true},
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}

	mockRepo := &MockUserRepository{
		GetUserByPhoneFunc: func(phone string) (*entities.User, error) {
			if phone == mockedUser.Phone {
				return &mockedUser, nil
			}
			return nil, nil
		},
	}

	cacheMock := &MockCache{
		SetFunc: func(ctx context.Context, key string, value interface{}, expiration time.Time) error {
			return nil
		},
		ExistFunc: func(ctx context.Context, key string) (int64, error) {
			return 1, nil
		},
		DelFunc: func(ctx context.Context, key string) error {
			return nil
		},
	}

	service := AuthService{
		UserRepository: mockRepo,
		Config:         mockConfig,
		Cache:          cacheMock,
	}

	t.Run("test validation", func(t *testing.T) {
		tests := []struct {
			name    string
			request LoginRequest
		}{
			{"short phone", LoginRequest{"123", "password"}},
			{"empty password", LoginRequest{"09130108631", ""}},
			{"empty phone", LoginRequest{"", "password"}},
			{"long phone", LoginRequest{"0913010863112", "password"}},
			{"invalid phone format", LoginRequest{"123", "password"}},
		}

		for _, tc := range tests {
			t.Run(tc.name, func(t *testing.T) {
				_, err := service.Login(tc.request)
				if err == nil {
					t.Errorf("expected validation error for %q", tc.name)
				}
			})
		}
	})

	t.Run("test invalid credential", func(t *testing.T) {
		test := LoginRequest{
			"09130108631",
			"yourPassword",
		}

		_, err := service.Login(test)
		if err == nil {
			t.Errorf("Expected error to be returned for invalid credential")
		}
	})

	t.Run("valid tokens returned", func(t *testing.T) {
		resp, err := service.Login(LoginRequest{"09130108631", "mypassword"})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		cases := []struct {
			name     string
			token    string
			wantType string
			wantTTL  time.Duration
		}{
			{"access", resp.Token, ACCESS, mockConfig.AccessTTL},
			{"refresh", resp.RefreshToken, REFRESH, mockConfig.RefreshTTL},
		}

		for _, tc := range cases {
			t.Run(tc.name, func(t *testing.T) {
				claims, err := service.parseToken(tc.token)
				if err != nil {
					t.Fatalf("failed to parse token: %v", err)
				}

				now := time.Now()

				expectedMin := now.Add(tc.wantTTL - time.Minute)
				expectedMax := now.Add(tc.wantTTL + time.Minute)

				if claims.ExpiresAt.Before(expectedMin) || claims.ExpiresAt.After(expectedMax) {
					t.Errorf("unexpected expiration time: %v", claims.ExpiresAt)
				}

				if claims.TokenType != tc.wantType {
					t.Errorf("unexpected token type: %v", claims.TokenType)
				}

				if claims.Issuer != service.Config.Issuer {
					t.Errorf("unexpected Issuer for token: %v", claims.TokenType)
				}
			})
		}
	})

	t.Run("valid user info", func(t *testing.T) {
		resp, err := service.Login(LoginRequest{"09130108631", "mypassword"})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		cases := []struct {
			wanted string
			given  string
			field  string
		}{
			{wanted: mockedUser.UUID, given: resp.User.Uuid, field: "uuid"},
			{wanted: mockedUser.LastName.String, given: resp.User.LastName, field: "last name"},
			{wanted: mockedUser.FirstName.String, given: resp.User.FirstName, field: "first name"},
			{wanted: mockedUser.Phone, given: resp.User.Phone, field: "phone"},
		}

		for _, tc := range cases {
			t.Run(fmt.Sprintf("run test for field %s", tc.field), func(t *testing.T) {
				if tc.given != tc.wanted {
					t.Errorf("wanted: %v, got: %v", tc.wanted, tc.given)
				}
			})
		}
	})

	t.Run("invalid User login", func(t *testing.T) {
		_, err := service.Login(LoginRequest{"09004636353", "invalidUserPassword"})

		if err == nil {
			t.Fatalf("User doesnt exists and login doesnt return error: %v", err)
		}
	})

}
