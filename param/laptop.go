package param

import (
	"time"
)

type LaptopInfo struct {
	ID          uint      `json:"id"`
	CPU         string    `json:"cpu"`
	RAM         int64     `json:"ram"`
	SSD         int64     `json:"ssd"`
	HDD         int64     `json:"hdd"`
	Graphic     int64     `json:"graphic"`
	ScreenSize  string    `json:"screen_size"`
	Company     string    `json:"company"`
	Price       string    `json:"price"`
	CreatedAt   time.Time `json:"created_at"`
	ImageURL    string    `json:"image_url"`
	RedirectURL string    `json:"redirect_url"`
}
type SearchInfo struct {
	CPU        string `json:"cpu"`
	RAM        int64  `json:"ram"`
	SSD        int64  `json:"ssd"`
	HDD        int64  `json:"hdd"`
	Graphic    int64  `json:"graphic"`
	ScreenSize string `json:"screen_size"`
	Company    string `json:"company"`
}

type UpdateLaptopRequest struct {
	ID          uint      `json:"id"`
	CPU         string    `json:"cpu"`
	RAM         int64     `json:"ram"`
	SSD         int64     `json:"ssd"`
	HDD         int64     `json:"hdd"`
	Graphic     int64     `json:"graphic"`
	ScreenSize  string    `json:"screen_size"`
	Company     string    `json:"company"`
	Price       string    `json:"price"`
	CreatedAt   time.Time `json:"created_at"`
	RedirectURL string    `json:"redirect_url"`
}

type UpdateLaptopResponse struct {
	ID      uint `json:"id"`
	Success bool `json:"success"`
}

type SellerLaptopRequest struct {
	UserID      uint   `json:"user_id"`
	CPU         string `json:"cpu"`
	RAM         int64  `json:"ram"`
	SSD         int64  `json:"ssd"`
	HDD         int64  `json:"hdd"`
	Graphic     int64  `json:"graphic"`
	ScreenSize  string `json:"screen_size"`
	Company     string `json:"company"`
	Price       string `json:"price"`
	ImageURL    string `json:"image_url"`
	RedirectURL string `json:"redirect_url"`
}

type SellerLaptopResponse struct {
	ID      uint `json:"id"`
	Success bool `json:"success"`
}
type AddLaptopRequest struct {
	UserID   int `json:"user_id"`
	LaptopID int `json:"laptop_id"`
}

type AddLaptopResponse struct {
	Success bool `json:"success"`
}

type LaptopsRequest struct {
	UserID int `json:"user_id"`
}

type LaptopsResponse struct {
	Laptops []LaptopInfo `json:"laptops"`
}

type LaptopRequest struct {
	LaptopID int `json:"laptop_id"`
}

type LaptopResponse struct {
	Laptop LaptopInfo `json:"laptop"`
}

type SearchRequest struct {
	UserID int        `json:"id"`
	Info   SearchInfo `json:"laptop"`
}

type SearchResponse struct {
	Success bool  `json:"success"`
	Result  []int `json:"result"`
}

type RemoveFavoriteLaptopRequest struct {
	UserID   int `json:"user_id"`
	LaptopID int `json:"laptop_id"`
}

type RemoveFavoriteLaptopResponse struct {
	Success bool `json:"success"`
}

type RemoveSellerLaptopRequest struct {
	UserID   int `json:"user_id"`
	LaptopID int `json:"laptop_id"`
}

type RemoveSellerLaptopResponse struct {
	Success bool `json:"success"`
}
