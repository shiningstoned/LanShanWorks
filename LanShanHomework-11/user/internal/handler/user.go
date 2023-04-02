package handler

import (
	"LanShanHomework-11/user/internal/repository"
	"LanShanHomework-11/user/internal/service"
	"context"
	"errors"
)

type UserService struct {
	service.UnimplementedUserServiceServer
}

func NewUserService() *UserService {
	return &UserService{}
}

func (u *UserService) Register(ctx context.Context, req *service.Request) (resp *service.RegisterResponse, err error) {
	resp = new(service.RegisterResponse)
	if exist := repository.IsUserExist(req); exist {
		resp.Msg = "用户已存在"
		return resp, errors.New("用户已存在")
	}
	err = repository.UserCreate(req)
	if err != nil {
		resp.Msg = "新建用户失败"
		return resp, err
	} else {
		resp.Msg = "新建用户成功"
		return resp, err
	}
}

func (u *UserService) Login(ctx context.Context, req *service.Request) (resp *service.LoginResponse, err error) {
	if exist := repository.IsUserExist(req); !exist {
		resp.Msg = "用户还未登陆"
		return resp, errors.New("用户不存在")
	}
	resp, err = repository.CheckPasswd(req)
	if err != nil {
		return resp, err
	}
	resp.UserId = repository.MatchId(req)
	return resp, nil
}
