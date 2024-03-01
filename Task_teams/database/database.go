package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDB() *gorm.DB {
	dialect := mysql.Open("root:@tcp(localhost:3306)/task1?charset=utf8mb4&parseTime=True&loc=Local")
	db, err := gorm.Open(dialect, &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

//  var db = NewDB()

// func TestOpenConnection(t *testing.T) {
// 	assert.NotNil(t, db)
// }
