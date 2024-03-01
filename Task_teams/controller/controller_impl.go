package controller

import (
	"net/http"
	"time"

	// "task_teams/controller"
	"task_teams/helper"
	"task_teams/model/web"
	"task_teams/service"

	"github.com/golang-jwt/jwt/v4"
	"github.com/julienschmidt/httprouter"
)

type UserControllerImpl struct {
	UserService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &UserControllerImpl{
		UserService: userService,
	}
}

type JwtClaims struct {
	UserID int `json:"user_id"`
	jwt.RegisteredClaims
}

func generateJWT(userResponse web.UserResponse) (string, error) {
	// Buat payload JWT
	ExpiresAts := jwt.NewNumericDate(time.Now().Add(time.Minute * 10))
	claims := JwtClaims{
		userResponse.Id,
		jwt.RegisteredClaims{
			ExpiresAt: ExpiresAts, // Token kedaluwarsa dalam 24 jam
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	// Buat token JWT dengan payload dan secret key
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Tandatangani token dengan secret key dan dapatkan string token
	tokenString, err := token.SignedString([]byte("your_secret_key_here"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (controller *UserControllerImpl) Register(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userCreateRequest := web.UserCreateRequest{}
	helper.ReadFromRequestBody(request, &userCreateRequest)

	userResponse := controller.UserService.Register(userCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   userResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *UserControllerImpl) Login(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	userLoginRequest := web.UserLoginRequest{}
	helper.ReadFromRequestBody(request, &userLoginRequest)

	userResponse, err := controller.UserService.Login(userLoginRequest)
	if err != nil {
		webResponse := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "Invalid username or password",
			Data:   userResponse,
		}
		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	// Generate JWT setelah login berhasil
	token, err := generateJWT(userResponse)
	if err != nil {
		http.Error(writer, "Failed to generate JWT", http.StatusInternalServerError)
		return
	}

	// Tambahkan token ke respons
	userResponse.Token = token

	Expires := jwt.NewNumericDate(time.Now().Add(time.Hour * 24))
	ExpiredTimes := Expires.UTC()
	http.SetCookie(writer, &http.Cookie{
		Name:    "Failed",
		Value:   userResponse.Token,
		Expires: ExpiredTimes,
	})

	// Login successful
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   userResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *UserControllerImpl) Logout(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	_, err := request.Cookie("Failed")
	helper.PanicIfError(err)

	http.SetCookie(writer, &http.Cookie{
		Name:    "Failed",
		Expires: time.Now(),
	})

	cookie, err := request.Cookie("Failed")
	helper.PanicIfError(err)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   cookie,
	}
	helper.WriteToResponseBody(writer, webResponse)
}
