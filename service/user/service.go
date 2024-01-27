package user

import (
	"game-app/entity"
	"game-app/param"
)

type repository interface {
	RegisterUser(user entity.User) (entity.User, error)
	GetUserByPhoneNumber(phoneNumber string) (entity.User, error)
	GetUserByID(UserID uint) (entity.User, error)
	AddFavoriteLaptop(userID int, laptopID int) error
	RemoveFavoriteLaptop(LaptopID int, UserID int) error
	GetLaptops(UserID uint) ([]entity.Laptop, error)
	GetLaptopByID(LaptopID uint) (entity.Laptop, error)
	UpdateUser(updatedUser entity.User) error
	Search(IDs []int) ([]param.LaptopInfo, error)
	AddLaptop(LaptopInfo param.LaptopInfo, UserID uint) (uint, error)
	AddSellerLaptop(LaptopID, UserID uint) error
	GetSellerLaptops(UserID uint) ([]entity.Laptop, error)
	UpdateLaptop(updatedLaptop entity.Laptop) (error)
	RemoveSellerLaptop(LaptopID int, SellerID int) error
}
type AuthGenerator interface {
	CreateAccessToken(user entity.User) (string, error)
	CreateRefreshToken(user entity.User) (string, error)
}

type SearchService interface {
	SearchRequest(inf param.SearchInfo) (param.SearchResponse, error)
}

type Service struct {
	auth   AuthGenerator
	repo   repository
	search SearchService
}

func New(authGenerator AuthGenerator, repo repository, search SearchService) *Service {
	return &Service{auth: authGenerator, repo: repo, search: search}
}
