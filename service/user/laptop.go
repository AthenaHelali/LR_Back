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
		GraphicCard: req.Laptop.GraphicCard,
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
			GraphicCard: l.GraphicCard,
			CreatedAt:   l.CreatedAt,
			ImageURL:    l.ImageURL,
			RedirectURL: l.RedirectURL,
		})

	}

	return param.LaptopsResponse{Laptops: laptopsInfo}, nil
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
			GraphicCard: "gforce",
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
			GraphicCard: "gforce",
			ImageURL:    "sdfghjkl.jpg",
			RedirectURL: "goole.com",
		})

		return param.SearchResponse{Laptops: laptops}, nil

}
