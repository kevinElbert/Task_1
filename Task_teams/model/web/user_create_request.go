package web

type UserCreateRequest struct {
	Username string `validate:"required" gorm:"column:username" json:"username"`
	Password string `validate:"required" gorm:"column:password;type:varchar(255);not null" json:"password"`
	RoleUser int    `gorm:"column:roleUser" json:"roleUser"`
}
