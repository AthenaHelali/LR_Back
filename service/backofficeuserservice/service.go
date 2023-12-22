package backofficeuserservice

import "game-app/entity"



type repository interface {
	ListAllUsers() ([]entity.User, error)
	DeleteUser(UserID uint)(error)
	ListAllLaptops() ([]entity.Laptop, error)
	DeleteLaptop(LaptopID uint)(error)
	UpdateLaptop(laptopID uint,updatedLaptop entity.Laptop) (error)
	
}

type Service struct {
	repo repository
}

func New(repo repository) *Service {
	return &Service{repo: repo}
}

