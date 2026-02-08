package authService

import (
	"crm_go/pkg/appError"
	"crm_go/pkg/validation"
	"errors"
)

func (s *AuthService) Login(request LoginRequest) (LoginResponse, error) {

	if err := validation.V().Struct(request); err != nil {
		return LoginResponse{}, err
	}

	user, err := s.UserRepository.GetUserByPhone(request.Phone)

	if err != nil {
		if errors.Is(err, appError.ErrUserNotFound) {
			return LoginResponse{}, appError.NotFound("user_not_found", "user not found", err)
		}

		return LoginResponse{}, appError.Internal(err)
	}

	return LoginResponse{
		Token:        "token",
		RefreshToken: "refresh",
		User:         UserLoginResponse{FirstName: user.FirstName.String, LastName: user.LastName.String, Phone: user.Phone},
	}, nil
}

type LoginRequest struct {
	Phone    string `json:"phone" validate:"required,ir_phone_number"`
	Password string `json:"password" validate:"required,min=8"`
}

type LoginResponse struct {
	Token        string            `json:"token"`
	RefreshToken string            `json:"refresh_token"`
	User         UserLoginResponse `json:"user"`
}

type UserLoginResponse struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Phone     string `json:"phone"`
}
