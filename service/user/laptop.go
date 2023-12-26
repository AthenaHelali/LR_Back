package user

import (
	"fmt"
	"game-app/param"
)

func (s Service) AddFavoriteLaptop(req param.AddLaptopRequest) (param.AddLaptopResponse, error) {

	err := s.repo.AddFavoriteLaptop(req.UserID, req.LaptopID)
	if err != nil {
		return param.AddLaptopResponse{Success: false}, fmt.Errorf("unexpected error: %w", err)
	}
	//return created laptop ID
	var resp param.AddLaptopResponse   
	resp.Success = true

	return resp, nil
}

func (s Service) GetLaptops(req param.LaptopsRequest) (param.LaptopsResponse, error) {
	laptops, err := s.repo.GetLaptops(uint(req.UserID))
	if err != nil {
		return param.LaptopsResponse{}, fmt.Errorf("unexpected error: %w", err)
	}
	var laptopsInfo []param.LaptopInfo
	for _, l := range laptops {
		laptopsInfo = append(laptopsInfo, param.LaptopInfo{
			ID:          l.ID,
			CPU:         l.CPU,
			RAM:         l.RAM,
			SSD:         l.SSD,
			HDD:         l.HDD,
			ScreenSize:  l.ScreenSize,
			Company:     l.Company,
			Graphic:     l.Graphic,
			Price: l.Price,
			CreatedAt:   l.CreatedAt,
			ImageURL:    l.ImageURL,
			RedirectURL: l.RedirectURL,
		})

	}

	return param.LaptopsResponse{Laptops: laptopsInfo}, nil
}

func (s Service) GetLaptop(req param.LaptopRequest) (param.LaptopResponse, error) {
	laptop, err := s.repo.GetLaptopByID(uint(req.LaptopID))
	if err != nil {
		return param.LaptopResponse{}, fmt.Errorf("unexpected error: %w", err)
	}

	laptopsInfo := &param.LaptopInfo{
		ID:          laptop.ID,
		CPU:         laptop.CPU,
		RAM:         laptop.RAM,
		SSD:         laptop.SSD,
		HDD:         laptop.HDD,
		ScreenSize:  laptop.ScreenSize,
		Company:     laptop.Company,
		Graphic:     laptop.Graphic,
		CreatedAt:   laptop.CreatedAt,
		ImageURL:    laptop.ImageURL,
		RedirectURL: laptop.RedirectURL,
	}

	return param.LaptopResponse{Laptop: *laptopsInfo}, nil
}

func (s Service) Search(req param.SearchRequest) (param.SearchResponse, error) {
	// TODO - get laptops from ml component
	var laptops []param.LaptopInfo
	laptops,_ = s.repo.Search()

	return param.SearchResponse{Laptops: laptops}, nil

}

func (s Service) RemoveFavoriteLaptop(req param.RemoveFavoriteLaptopRequest) (param.RemoveFavoriteLaptopResponse, error) {

	err := s.repo.RemoveFavoriteLaptop(req.LaptopID, req.UserID)
	if err != nil {
		return param.RemoveFavoriteLaptopResponse{},fmt.Errorf("unexpected error: %w", err)
	}

	return param.RemoveFavoriteLaptopResponse{Success: true}, nil
}

