package main

import (
	"LanShanHomework-11/config"
	"LanShanHomework-11/user/internal/handler"
	"LanShanHomework-11/user/internal/repository"
	"LanShanHomework-11/user/internal/service"
	"fmt"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	config.InitConfig()
	repository.InitDB()

	grpcServer := grpc.NewServer()
	service.RegisterUserServiceServer(grpcServer, handler.NewUserService())

	grpcAddress := viper.GetString("server.grpcAddress")
	fmt.Println(grpcAddress)
	listener, err := net.Listen("tcp", grpcAddress)
	if err != nil {
		log.Fatal("监听服务失败")
	}

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("启动服务失败")
	}
}
