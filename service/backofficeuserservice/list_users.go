package backofficeuserservice

import (
	"game-app/entity"
	"game-app/param"
)

func (s Service) ListAllUsers() (param.ListUsersResponse, error) {
	//TODO - implement me
	list := make([]entity.User, 0)
	list = append(list, entity.User{
		ID:          0,
		PhoneNumber: "fake",
		Name:        "fake",
		Password:    "123",
		Role:        entity.AdminRole,
	})

	return param.ListUsersResponse{Users: list}, nil
}
