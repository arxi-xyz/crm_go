package authService

func (s *AuthService) Login(request LoginRequest) (LoginResponse, error) {

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
	Phone    string `json:"phone"`
	Password string `json:"password"`
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
