package main

import (
	"fmt"
	"grpc-server-streaming-example/cache"
	pb "grpc-server-streaming-example/proto/user"
	"grpc-server-streaming-example/userService/service"
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {
	godotenv.Load()

	// create listener
	host := fmt.Sprintf("localhost:%s", os.Getenv("USER_PORT"))
	listener, err := net.Listen("tcp", host)

	if err != nil {
		panic("error building user service: " + err.Error())
	}

	// create gRPC server
	s := grpc.NewServer()
	redisClient, _ := cache.NewRedis()
	pb.RegisterUserServiceServer(s, service.NewService(redisClient))

	log.Printf("start user service in: %s", host)

	if err := s.Serve(listener); err != nil {
		panic("error building user service: " + err.Error())
	}

}
