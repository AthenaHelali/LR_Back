package backofficeuserservice

import (
	"game-app/param"
	"game-app/pkg/richerror"
)

func (s Service) ListAllUsers() (param.ListUsersResponse, error) {
	const op = "backofficeuserservice.ListAllUsers"

	l, err := s.repo.ListAllUsers()
	if err != nil {
		return param.ListUsersResponse{}, richerror.New(op).WithError(err)
	}
	
	return param.ListUsersResponse{Users: l}, nil
}
