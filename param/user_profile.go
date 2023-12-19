package param

type ProfileRequest struct {
	UserID uint
}
type ProfileResponse struct {
	Info UserInfo `json:"Info"`

}
