package entity

import (
	"time"
)

type Laptop struct {
	ID          uint
	CPU         string
	RAM         int64
	SSD         int64
	HDD         int64
	Graphic     int64
	ScreenSize  string
	Company     string
	Price       string
	CreatedAt   time.Time
	ImageURL    string
	RedirectURL string
}
