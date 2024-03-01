package service

import (
	"errors"
	"fmt"
	"task_teams/helper"
	"task_teams/model/domain"
	"task_teams/model/web"
	"task_teams/repository"

	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserServiceImpl struct {
	UserRepo repository.UserRepository
	Validate *validator.Validate
	DB       *gorm.DB
}

func NewUserService(userRepo repository.UserRepository, DB *gorm.DB, validate *validator.Validate) UserService {
	return &UserServiceImpl{
		UserRepo: userRepo,
		Validate: validate,
		DB:       DB,
	}
}

func (us *UserServiceImpl) Register(request web.UserCreateRequest) web.UserResponse {
	err := us.Validate.Struct(request)
	helper.PanicIfError(err)

	tx := us.DB.Begin()
	defer tx.Rollback()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		helper.PanicIfError(err)
		// return web.UserResponse{}
	}

	// Membuat user baru dengan password yang di-hash
	user := domain.User{
		Username: request.Username,
		Password: string(hashedPassword),
		RoleUser: request.RoleUser,
	}

	// Menyimpan user ke dalam database
	userTemp := us.UserRepo.Register(tx, user) // Harus di check
	// if err != nil {
	// 	return web.UserResponse{}
	// }
	err = tx.Commit().Error
	helper.PanicIfError(err)
	// Mengembalikan respons sukses
	response := web.UserResponse{
		Username: userTemp.Username,
		Password: userTemp.Password,
		Id:       userTemp.RoleUser,
		Nama:     userTemp.Role.Name,
	}
	return response

	// user.Password = string(hashedPassword)
	// return us.UserRepo.Register(us.DB, user)
}

func (us *UserServiceImpl) Login(request web.UserLoginRequest) (web.UserResponse, error) {
	tx := us.DB.Begin()
	user, err := us.UserRepo.Login(tx, request.Username)
	if err != nil {
		fmt.Println(request.Username)
		return web.UserResponse{}, err
	}

	defer tx.Rollback()

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		return web.UserResponse{}, errors.New("invalid username or password")
	}

	err = tx.Commit().Error
	helper.PanicIfError(err)

	response := web.UserResponse{
		Username: user.Username,
		// Password: user.Password,
		Id:   user.Role.Id,
		Nama: user.Role.Name,
	}
	return response, nil
}
