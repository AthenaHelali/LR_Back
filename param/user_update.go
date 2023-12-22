package param

type UpdateUserRequest struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}

type UpdateUserResponse struct {
	Success bool `json:"success"`
}