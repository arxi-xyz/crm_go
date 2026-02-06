package authService

import "crm_go/pkg/validation"

func (s *AuthService) Login(request LoginRequest) (LoginResponse, error) {

	err := validation.V().Struct(request)

	if err != nil {
		return LoginResponse{}, err
	}

	user, err := s.UserRepository.GetUserByPhone(request.Phone)

	if err != nil {
		panic("User not found")
	}

	return LoginResponse{
		Token:        "token",
		RefreshToken: "refresh",
		User:         UserLoginResponse{FirstName: user.FirstName, LastName: user.LastName, Phone: user.Phone},
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
