package authService

import (
	"crm_go/pkg/appError"
	"crm_go/pkg/validation"

	"golang.org/x/crypto/bcrypt"
)

func (s *AuthService) Login(request LoginRequest) (LoginResponse, error) {

	if err := validation.V().Struct(request); err != nil {
		return LoginResponse{}, err
	}

	user, err := s.UserRepository.GetUserByPhone(request.Phone)

	if user == nil {
		return LoginResponse{}, appError.Unauthorized("invalid_credential", "invalid credential", err)
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(request.Password),
	)

	if err != nil {
		return LoginResponse{}, appError.Unauthorized(
			"invalid_credential",
			"invalid_credential",
			err,
		)
	}

	token, refreshToken, err := s.generateTokens(*user)

	if err != nil {
		return LoginResponse{}, appError.Internal(err)
	}

	return LoginResponse{
		Token:        token,
		RefreshToken: refreshToken,
		User: UserLoginResponse{
			Uuid:      user.UUID,
			FirstName: user.FirstName.String,
			LastName:  user.LastName.String,
			Phone:     user.Phone,
		},
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
	Uuid      string `json:"uuid"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Phone     string `json:"phone"`
}
