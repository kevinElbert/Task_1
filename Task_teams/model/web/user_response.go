package web

type UserResponse struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Id       int    `json:"id"`
	Nama     string `json:"nama"`
	Token    string `json:"token"`
}
