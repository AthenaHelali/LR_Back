package user

import (
	"game-app/entity"
	"game-app/param"
)

type repository interface {
	RegisterUser(user entity.User) (entity.User, error)
	GetUserByPhoneNumber(phoneNumber string) (entity.User, error)
	GetUserByID(UserID uint) (entity.User, error)
	AddFavoriteLaptop(userID int,laptopID int) (error)
	RemoveFavoriteLaptop(LaptopID int, UserID int)error
	GetLaptops(UserID uint) ([]entity.Laptop, error)
	GetLaptopByID(LaptopID uint) (entity.Laptop, error)
	UpdateUser(updatedUser entity.User) (error)
	Search() ([]param.LaptopInfo, error)
}
type AuthGenerator interface {
	CreateAccessToken(user entity.User) (string, error)
	CreateRefreshToken(user entity.User) (string, error)
}

type Service struct {
	auth AuthGenerator
	repo repository
}

func New(authGenerator AuthGenerator, repo repository) *Service {
	return &Service{auth: authGenerator, repo: repo}
}
