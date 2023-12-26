package backofficeuserservice

import (
	"game-app/param"
	"game-app/pkg/richerror"
)

func (s Service)DeleteLaptop(DeleteLaptopReq param.DeleteLaptopRequest) (error){
	const op = "backofficeuserservice.DeleteLaptop"
	err := s.repo.DeleteLaptop(DeleteLaptopReq.LaptopID)
	if err != nil {
		return richerror.New(op).WithError(err)
	}

	return nil

}