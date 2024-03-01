// package repository

// import (
// 	"context"
// 	"task_teams/model"

// 	"gorm.io/gorm"
// )

//	type UsersRepository interface {
//		Register(ctx context.Context, tx *gorm.DB, user model.User)
//		FindById(ctx context.Context, tx *gorm.DB, Username string) (*model.User, error)
//	}
package repository

import (
	"task_teams/model/domain"

	"gorm.io/gorm"
)

type UserRepository interface {
	Register(db *gorm.DB, user domain.User) domain.User
	Login(db *gorm.DB, username string) (domain.User, error)
}
