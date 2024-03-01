package service

import (
	"task_teams/model/web"
)

type UserService interface {
	Register(request web.UserCreateRequest) web.UserResponse
	Login(request web.UserLoginRequest) (web.UserResponse, error)
}
