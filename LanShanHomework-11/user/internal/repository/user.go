package repository

import (
	"LanShanHomework-11/user/internal/service"
	"errors"
	"fmt"
)

type User struct {
	UserId   int32 `gorm:"primarykey"`
	Username string
	Password string
}

func IsUserExist(req *service.Request) bool {
	var count int64
	DB.Where("username=?", req.Username).Count(&count)
	if count == 0 {
		return false
	}
	return true
}

func UserCreate(req *service.Request) error {
	var user User
	user.Username = req.Username
	user.Password = req.Password
	err := DB.Create(&user).Error
	if err != nil {
		fmt.Println("新建用户失败")
		return err
	}
	return nil
}

func MatchId(req *service.Request) int32 {
	var user User
	DB.Where("username=?", req.Username).First(&user)
	return user.UserId
}

func CheckPasswd(req *service.Request) (resp *service.LoginResponse, err error) {
	var user User
	DB.Where("username=?", req.Username).First(&user)
	if user.Password == req.Password {
		resp.Msg = "登陆成功"
		return resp, nil
	}
	resp.Msg = "登陆失败，密码错误"
	return resp, errors.New("密码错误")
}
