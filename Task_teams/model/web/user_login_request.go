package web

// type UserLoginRequest struct {
// 	Username string `json:"username"`
// 	Password string `json:"password"`
// }

type UserLoginRequest struct {
	Username string `validate:"required" gorm:"column:username" json:"username"`
	Password string `validate:"required" gorm:"column:password;type:varchar(255);not null" json:"password"`
}
