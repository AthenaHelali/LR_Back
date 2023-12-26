package user

import (
	"game-app/entity"
	"game-app/param"
	"testing"
)
type repo struct{

}
func(r *repo)RegisterUser(user entity.User) (entity.User, error){
	panic("")

}
func(r *repo)GetUserByPhoneNumber(phoneNumber string) (entity.User, error){
	panic("")
}
func(r *repo)GetUserByID(UserID uint) (entity.User, error){
	panic("")
}
func(r *repo)RegisterUserAddFavoriteLaptop(userID int,laptopID int) (error){
	panic("")
}
func(r *repo)RegisterUserRemoveFavoriteLaptop(LaptopID int, UserID int)error{
	panic("")
}
func(r *repo)RegisterUserGetLaptops(UserID uint) ([]entity.Laptop, error){
	panic("")
}
func(r *repo)RegisterUserGetLaptopByID(LaptopID uint) (entity.Laptop, error){
	panic("")
}
func(r *repo)RegisterUserUpdateUser(updatedUser entity.User) (error){
	panic("")
}
func(r *repo)RegisterUserSearch(IDs[]int) ([]param.LaptopInfo, error){
	panic("")
}

type authGenerator struct {
}

func(a *authGenerator)	CreateAccessToken(user entity.User) (string, error){
	panic("")
}
func(a *authGenerator)CreateRefreshToken(user entity.User) (string, error){
	panic("")
}

type searchService struct{
}

func(a *searchService)TestSearchRequest(inf param.SearchInfo) (param.SearchResponse, error){

}

func TestUpdateUser(t *testing.T){
	svc := New(&authGenerator{},&repo{},&searchService{})
	

    got := Add(4, 6)
    want := 10

    if got != want {
        t.Errorf("got %q, wanted %q", got, want)
    }
}