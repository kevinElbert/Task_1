package domain

type Role struct {
	Id   int    `gorm:"primay_key;column:id"`
	Name string `gorm:"column:name"`
}
