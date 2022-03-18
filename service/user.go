package service

import (
	"github.com/jinzhu/gorm"
	"todo_list/model"
	"todo_list/pkg/e"
	"todo_list/pkg/utils"
	"todo_list/serializer"
)

type UserService struct {
	Username string `form:"username" json:"username" binding:"required,min=3,max=15"`
	Password string `form:"password" json:"password" binding:"required,min=5,max=16"`
}

func (service *UserService) Register() serializer.Response {
	var user model.User
	var count int
	model.DB.Model(&model.User{}).Where("username=?", service.Username).First(&user).Count(&count)
	if count == 1 {
		return serializer.Response{
			Status: e.InvalidParams,
			Msg:    "当前用户名已存在",
		}
	}
	user.Username = service.Username
	//密码加密
	if err := user.SetPassword(service.Password); err != nil {
		return serializer.Response{
			Status: e.InvalidParams,
			Msg:    err.Error(),
		}
	}
	//创建用户
	if err := model.DB.Create(&user).Error; err != nil {
		return serializer.Response{
			Status: e.ERROR,
			Msg:    "数据库操作错误",
		}
	}
	return serializer.Response{
		Status: e.SUCCESS,
		Msg:    "用户注册成功",
	}
}

func (service *UserService) Login() serializer.Response {
	var user model.User
	if err := model.DB.Model(&model.User{}).Where("username=?", service.Username).First(&user).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return serializer.Response{
				Status: e.InvalidParams,
				Msg:    "用户不存在",
			}
		}
		return serializer.Response{
			Status: e.ERROR,
			Msg:    "数据库错误",
		}
	}
	if user.CheckPassword(service.Password) == false {
		return serializer.Response{
			Status: e.InvalidParams,
			Msg:    "密码错误",
		}
	}
	//发送一个token
	token, err := utils.GenerateToken(user.ID, user.Username)
	if err != nil {
		return serializer.Response{
			Status: e.ERROR,
			Msg:    "Token签发错误",
		}
	}
	return serializer.Response{
		Status: e.SUCCESS,
		Data: serializer.TokenData{
			User:  serializer.BuildUser(&user),
			Token: token,
		},
		Msg: "登陆成功",
	}
}
