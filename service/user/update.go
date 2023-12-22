package user

import (
	"fmt"
	"game-app/entity"
	"game-app/param"

	"golang.org/x/crypto/bcrypt"
)

func (s Service) UpdateUser(req param.UpdateUserRequest) (param.UpdateUserResponse, error) {
	pass := []byte(req.Password)
	hashedPass, err := bcrypt.GenerateFromPassword(pass, bcrypt.DefaultCost)
	
	if err != nil {
		return param.UpdateUserResponse{},fmt.Errorf("unexpected error: %w", err)
	}

	user := &entity.User{
		ID: req.ID,
		PhoneNumber: req.PhoneNumber,
		Name: req.Name,
		Password: string(hashedPass),
	}
	err = s.repo.UpdateUser(*user)
	if err != nil {
		return param.UpdateUserResponse{},fmt.Errorf("unexpected error: %w", err)
	}
	return param.UpdateUserResponse{Success: true},err
}
