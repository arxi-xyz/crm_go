package userService

import (
	"crm_go/pkg/appError"
	"database/sql"
	"errors"
)

func (s *UserService) GetMe(uuid string) (MeResponse, error) {
	user, err := s.UserRepository.GetUserBy("uuid", uuid)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return MeResponse{}, appError.NotFound("user_not_found", "user not found", err)
		}
		return MeResponse{}, appError.Internal(err)
	}

	return MeResponse{
		UUID:      user.UUID,
		Phone:     user.Phone,
		FirstName: user.FirstName.String,
		LastName:  user.LastName.String,
		CreatedAt: user.CreatedAt.Format("2006-01-02T15:04:05Z"),
	}, nil
}

type MeResponse struct {
	UUID      string `json:"uuid" example:"550e8400-e29b-41d4-a716-446655440000"`
	Phone     string `json:"phone" example:"09130108631"`
	FirstName string `json:"first_name" example:"Sharif"`
	LastName  string `json:"last_name" example:"Mohammadi"`
	CreatedAt string `json:"created_at" example:"2025-01-15T10:30:00Z"`
}
