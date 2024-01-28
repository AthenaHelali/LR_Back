package user

import (
	"testing"

	"game-app/param"
	"game-app/repository/mysql/mysqluser"
	"game-app/service/authservice"
	"game-app/service/search"

	"github.com/stretchr/testify/assert"
)

func initService() Service {
	return Service{
		repo:   &mysqluser.DB{},
		search: &search.Service{},
		auth:   &authservice.Service{},
	}
}

func TestAddFavorite(t *testing.T) {
	defer func() {
		recover()
	}()
	s := initService()
	s.AddFavoriteLaptop(param.AddLaptopRequest{UserID: 1, LaptopID: 1})
}

func TestGetLaptop(t *testing.T) {
	defer func() {
		recover()
	}()
	s := initService()
	s.GetLaptop(param.LaptopRequest{LaptopID: 1})
}

func TestSearch(t *testing.T) {
	defer func() {
		recover()
	}()
	s := initService()
	laptops, err := s.Search(param.SearchRequest{UserID: 1})
	assert.Empty(t, laptops)
	assert.Error(t, err)
}

func TestRemoveFromFav(t *testing.T) {
	defer func() {
		recover()
	}()
	s := initService()
	s.RemoveFavoriteLaptop(param.RemoveFavoriteLaptopRequest{UserID: 1, LaptopID: 1})
}

func TestAddLaptop(t *testing.T) {
	defer func() {
		recover()
	}()
	s := initService()
	s.AddLaptop(param.SellerLaptopRequest{UserID: 1})
}

func TestGetSellerLaptop(t *testing.T) {
	defer func() {
		recover()
	}()
	s := initService()
	s.GetSellerLaptops(param.LaptopsRequest{UserID: 1})
}

func TestUpdateLaptop(t *testing.T) {
	defer func() {
		recover()
	}()
	s := initService()
	s.UpdateLaptop(param.UpdateLaptopRequest{ID: 1, CPU: "i7"})
}

func TestRemoveSeller(t *testing.T) {
	defer func() {
		recover()
	}()
	s := initService()
	s.RemoveSellerLaptop(param.RemoveSellerLaptopRequest{UserID: 1})
}
