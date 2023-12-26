package param
type DeleteLaptopRequest struct{
	LaptopID uint64 `json:"laptop_id"`
}

type DeleteLaptopResponse struct{
	LaptopID uint64 `json:"id"`
}