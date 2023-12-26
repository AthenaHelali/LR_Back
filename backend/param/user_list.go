package param

import "game-app/entity"

type ListUsersResponse struct{
	Users []entity.User `json:"users"`
}