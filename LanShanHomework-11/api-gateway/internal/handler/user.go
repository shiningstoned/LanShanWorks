package handler

import (
	"LanShanHomework-11/api-gateway/internal/service"
	"LanShanHomework-11/api-gateway/pkg/utils"
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Register(c *gin.Context) {
	var userReq service.Request
	err := c.ShouldBind(&userReq)
	if err != nil {
		log.Fatal("错误请求")
	}
	userService := c.Keys["userService"].(service.UserServiceClient)
	userResp, err := userService.Register(context.Background(), &userReq)
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(http.StatusOK, userResp)
}

func Login(c *gin.Context) {
	var userReq service.Request
	err := c.Bind(&userReq)
	if err != nil {
		log.Fatal("错误请求")
	}
	userService := c.Keys["userService"].(service.UserServiceClient)
	userResp, err := userService.Login(context.Background(), &userReq)
	if err != nil {
		log.Fatal(err)
	}
	token, err := utils.GenerateToken(userResp.UserId)
	if err != nil {
		log.Fatal("解析token失败")
	}
	userResp.Token = token
	c.JSON(http.StatusOK, userResp)
}
