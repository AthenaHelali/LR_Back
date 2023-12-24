package backofficeuserservice

import "game-app/entity"



type repository interface {
	ListAllUsers() ([]entity.User, error)
	ListAllLaptops() ([]entity.Laptop, error)
	DeleteLaptop(LaptopID uint64)(error)	
}

type Service struct {
	repo repository
}

func New(repo repository) *Service {
	return &Service{repo: repo}
}

