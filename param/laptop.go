package param
import (
	"time"
)

type LaptopInfo struct{
	ID uint `json:"id"`
	CPU  string `json:"cpu"`
	RAM int64`json:"ram"`
	SSD int64 `json:"ssd"`
	HDD int64 `json:"hdd"`
	GraphicCard string `json:"graphic_card"`
	ScreenSize string `json:"screen_size"`
	Company string `json:"company"`
	CreatedAt time.Time `json:"created_at"`
	ImageURL string `json:"image_url"`
	RedirectURL string `json:"redirect_url"`
}

type AddLaptopRequest struct {
	UserID int `json:"user_id"`
	Laptop LaptopInfo `json:"laptop"`
}

type AddLaptopResponse struct {
	ID int `json:"id"`
	Success bool `json:"success"`
}

type LaptopsRequest struct{
	UserID int `json:"user_id"`
}

type LaptopsResponse struct{
	Laptops []LaptopInfo `json:"laptops"`
}

type SearchRequest struct{
	UserID int `json:"id"`
	Info LaptopInfo `json:"laptop"`
}

type SearchResponse struct{
	Laptops []LaptopInfo `json:"laptop"`
}