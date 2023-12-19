package entity
import(
	"time"
)

type Laptop struct{
	ID uint
	CPU  string
	RAM int64
	SSD int64
	HDD int64
	GraphicCard string
	ScreenSize string
	Company string
	CreatedAt time.Time
	ImageURL string
	RedirectURL string
}