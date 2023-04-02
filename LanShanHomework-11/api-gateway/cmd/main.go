package main

import (
	"LanShanHomework-11/api-gateway/internal/handler"
	"LanShanHomework-11/api-gateway/internal/service"
	"LanShanHomework-11/api-gateway/middleware"
	"LanShanHomework-11/config"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	config.InitConfig()
	grpcAddress := viper.GetString("server.grpcAddress")
	Userconn, err := grpc.Dial(grpcAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("连接服务失败")
	}
	userService := service.NewUserServiceClient(Userconn)

	r := gin.Default()
	r.Use(middleware.CORS(), middleware.InitMiddleware(userService))
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, "success")
	})
	//用户服务
	r.POST("/register", handler.Register)
	r.POST("/login", handler.Login)

	r.Run(":8000")
}
