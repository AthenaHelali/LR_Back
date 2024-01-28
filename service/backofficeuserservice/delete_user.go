package backofficeuserservice

import (
	"fmt"
	"game-app/param"
)

func (s *Service) DeleteUser(req param.DeleteUserRequest) (param.DeleteUserResponse, error) {
	err := s.repo.DeleteUser(req.UserID)
	if err != nil {
		return param.DeleteUserResponse{}, fmt.Errorf("unexpected error: %w", err)
	}

	return param.DeleteUserResponse{Success: true}, nil
}
