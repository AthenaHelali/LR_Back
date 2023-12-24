package user

import (
	"fmt"
	"game-app/entity"
	"game-app/param"
)

func (s Service) AddFavoriteLaptop(req param.AddLaptopRequest) (param.AddLaptopResponse, error) {

	//create new laptop in storage
	laptop := &entity.Laptop{
		ID:          req.Laptop.ID,
		CPU:         req.Laptop.CPU,
		RAM:         req.Laptop.RAM,
		SSD:         req.Laptop.SSD,
		HDD:         req.Laptop.HDD,
		Graphic:     req.Laptop.Graphic,
		ScreenSize:  req.Laptop.ScreenSize,
		Company:     req.Laptop.Company,
		Price:       req.Laptop.Price,
		ImageURL:    req.Laptop.ImageURL,
		RedirectURL: req.Laptop.RedirectURL,
	}
	res, err := s.repo.AddFavoriteLaptop(*laptop, req.UserID)
	if err != nil {
		return param.AddLaptopResponse{Success: false}, fmt.Errorf("unexpected error: %w", err)
	}
	//return created laptop ID
	var resp param.AddLaptopResponse
	resp.ID = int(res.ID)
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
	laptops = append(laptops,
		param.LaptopInfo{
			ID:          1,
			CPU:         "corei7",
			RAM:         16,
			SSD:         512,
			HDD:         512,
			ScreenSize:  "13",
			Company:     "hp",
			Graphic:     8,
			ImageURL:    "sdfghjkl.jpg",
			RedirectURL: "goole.com",
		},
		param.LaptopInfo{
			ID:          2,
			CPU:         "corei9",
			RAM:         32,
			SSD:         512,
			HDD:         512,
			ScreenSize:  "14",
			Company:     "lenovo",
			Graphic:     2,
			ImageURL:    "sdfghjkl.jpg",
			RedirectURL: "goole.com",
		})

	return param.SearchResponse{Laptops: laptops}, nil

}

func (s Service) RemoveFavoriteLaptop(req param.LaptopRequest) (error) {

	err := s.repo.RemoveFavoriteLaptop(req.LaptopID)
	if err != nil {
		return fmt.Errorf("unexpected error: %w", err)
	}

	return nil
}

