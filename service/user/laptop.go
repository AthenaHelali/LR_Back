package user

import (
	"fmt"
	"game-app/entity"
	"game-app/param"
	"time"
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
			Price:       l.Price,
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

func (s Service) Search(req param.SearchRequest) ([]param.LaptopInfo, error) {
	resp, err := s.search.SearchRequest(req.Info)
	if err != nil {
		return nil, fmt.Errorf("unexpected error: %w", err)
	}
	result, err := s.repo.Search(resp.Result)
	if err != nil {
		return nil, fmt.Errorf("unexpected error: %w", err)
	}
	return result, nil
}

func (s Service) RemoveFavoriteLaptop(req param.RemoveFavoriteLaptopRequest) (param.RemoveFavoriteLaptopResponse, error) {
	err := s.repo.RemoveFavoriteLaptop(req.LaptopID, req.UserID)
	if err != nil {
		return param.RemoveFavoriteLaptopResponse{}, fmt.Errorf("unexpected error: %w", err)
	}

	return param.RemoveFavoriteLaptopResponse{Success: true}, nil
}

func (s Service) AddLaptop(req param.SellerLaptopRequest) (param.SellerLaptopResponse, error){
	laptopInfo := &param.LaptopInfo{
		CPU:        req.CPU,
		RAM:        req.RAM,
		SSD:        req.SSD,
		HDD:        req.HDD,
		Graphic:    req.Graphic,
		ScreenSize: req.ScreenSize,
		Company:    req.Company,
		Price:      req.Price,
		CreatedAt:  time.Now(),
		ImageURL:   req.ImageURL,
	}

	LaptopID, err := s.repo.AddLaptop(*laptopInfo, req.UserID)
	if err != nil {
		return param.SellerLaptopResponse{}, err
	}
	err = s.repo.AddSellerLaptop(LaptopID, req.UserID)
	if err != nil {
		return param.SellerLaptopResponse{}, err
	}

	return param.SellerLaptopResponse{ID: LaptopID, Success: true}, nil
}
func (s Service) GetSellerLaptops(req param.LaptopsRequest) (param.LaptopsResponse, error) {
	laptops, err := s.repo.GetSellerLaptops(uint(req.UserID))
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
			Price:       l.Price,
			CreatedAt:   l.CreatedAt,
			ImageURL:    l.ImageURL,
			RedirectURL: l.RedirectURL,
		})

	}

	return param.LaptopsResponse{Laptops: laptopsInfo}, nil
}

func (s Service) UpdateLaptop(req param.UpdateLaptopRequest) (param.UpdateLaptopResponse, error) {
	laptop := &entity.Laptop{
		ID:          req.ID,
		CPU:         req.CPU,
		RAM:         req.RAM,
		SSD:         req.SSD,
		HDD:         req.HDD,
		Graphic:     req.Graphic,
		ScreenSize:  req.ScreenSize,
		Company:     req.Company,
		Price:       req.Price,
		RedirectURL: req.RedirectURL,
	}
	err := s.repo.UpdateLaptop(*laptop)
	if err != nil {
		return param.UpdateLaptopResponse{}, fmt.Errorf("unexpected error: %w", err)
	}
	return param.UpdateLaptopResponse{ID: req.ID, Success: true}, err

}

func (s Service) RemoveSellerLaptop(req param.RemoveSellerLaptopRequest) (param.RemoveSellerLaptopResponse, error) {
	err := s.repo.RemoveSellerLaptop(req.LaptopID, req.UserID)
	if err != nil {
		return param.RemoveSellerLaptopResponse{}, fmt.Errorf("unexpected error: %w", err)
	}

	return param.RemoveSellerLaptopResponse{Success: true}, nil
}
