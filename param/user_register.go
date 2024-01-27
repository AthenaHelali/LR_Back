package param

// user role =1, seller role = 3
type RegisterRequest struct {
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
	Role int `json:"role"`
}

type RegisterResponse struct {
	User UserInfo `json:"user"`
}
