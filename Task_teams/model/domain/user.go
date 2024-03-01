package domain

type User struct {
	Username string `gorm:"primary_key;column:username"`
	Password string `gorm:"column:password"`
	RoleUser int    `gorm:"column:roleUser"`
	Role     Role   `gorm:"foreignKey:roleUser"`
}
