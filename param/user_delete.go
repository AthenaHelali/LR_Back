package param

type DeleteUserRequest struct{
	UserID int  `json:"id"`
}

type DeleteUserResponse struct{
	Success bool `json:"success"`
 
}