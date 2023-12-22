package backofficeuserservice

import (
	"game-app/param"
)

func (s Service) ListLaptops() (param.ListLaptopsResponse, error) {
	// const op = "backofficeuserservice.listlaptops"

	// //TODO - implement me
	// l, err := s.repo.ListAllLaptops()
	// if err != nil {
	// 	return param.ListLaptopsResponse{}, richerror.New(op).WithError(err)
	// }
	
	return param.ListLaptopsResponse{Laptops: nil}, nil

}