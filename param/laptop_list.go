package param

import "game-app/entity"
type ListLaptopsResponse struct{
	Laptops []entity.Laptop `json:"laptops"`
}