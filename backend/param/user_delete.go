package param

type DeleteUserRequest struct{
	UserID string  `json:"id"`
}

type DeleteUserResponse struct{
	Success bool `json:"success"`
 
}