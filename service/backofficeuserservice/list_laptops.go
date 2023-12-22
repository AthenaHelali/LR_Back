package backofficeuserservice

import (
	"game-app/param"
	"game-app/pkg/richerror"
)

func (s Service) ListLaptops() (param.ListLaptopsResponse, error) {
	const op = "backofficeuserservice.listlaptops"

	l, err := s.repo.ListAllLaptops()
	if err != nil {
		return param.ListLaptopsResponse{}, richerror.New(op).WithError(err)
	}
	
	return param.ListLaptopsResponse{Laptops: l}, nil

}