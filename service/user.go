package service

import (
	"errors"
	"gin_cli/global"
	"gin_cli/model"
	"gin_cli/service/request"
	"gin_cli/service/response"
	"gin_cli/utils"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type UserService struct {
}

func (s *UserService) Register(params request.RegisterParams) (userInter model.User, err error) {
	var user model.User
	if !errors.Is(global.DB.Where("username = ?", params.Username).First(&user).Error, gorm.ErrRecordNotFound) { // 判断用户名是否注册
		return userInter, errors.New("用户名已注册")
	}

	user.Username = params.Username
	user.UUID = uuid.NewV4()
	user.SetPssword(params.Password)
	if err = global.DB.Create(&user).Error; err != nil {
		return userInter, err
	}

	return user, nil
}

func (s *UserService) Login(params request.RegisterParams) (res response.LoginResponse, err error) {
	var user model.User
	if errors.Is(global.DB.Where("username = ?", params.Username).First(&user).Error, gorm.ErrRecordNotFound) { // 判断用户名是否注册
		return res, errors.New("用户名或密码错误")
	}
	if !user.CheckPassword(params.Password) {
		return res, errors.New("密码错误")
	}

	token, err := utils.GenerateToken(user.ID, user.Username, user.RoleID)
	if err != nil {
		return res, errors.New("token签发错误")
	}

	return response.LoginResponse{UserInfo: user, Token: token}, nil
}
