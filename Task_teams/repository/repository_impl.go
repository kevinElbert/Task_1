package repository

import (
	"errors"
	"task_teams/model/domain"

	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (repository *UserRepositoryImpl) Register(DB *gorm.DB, user domain.User) domain.User {
	result := DB.Create(&user)

	// Memeriksa error
	if result.Error != nil {
		panic(result.Error)
	}

	return user
}

func (repository *UserRepositoryImpl) Login(DB *gorm.DB, userID string) (domain.User, error) {
	var user domain.User

	// Menggunakan GORM untuk mencari user berdasarkan Username
	// result := DB.Where("Username = ?", userID).First(&user)

	result := DB.Where("Username = ?", userID).Preload("Role").First(&user)

	// Memeriksa apakah terdapat error saat mencari
	if result.Error != nil {
		// Menangani kasus ketika user tidak ditemukan
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return user, errors.New("user is not found")
		}
		// Menangani error lainnya
		return user, result.Error
	}

	// Mengembalikan user jika ditemukan
	return user, nil
}
