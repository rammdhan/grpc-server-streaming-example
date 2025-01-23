package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"

	"grpc-server-streaming-example/authService/service"
	"grpc-server-streaming-example/cache"
	pb "grpc-server-streaming-example/proto/auth"
)

func main() {
	godotenv.Load()

	// create listener
	host := fmt.Sprintf("localhost:%s", os.Getenv("AUTH_PORT"))
	listener, err := net.Listen("tcp", host)

	if err != nil {
		panic("error building auth service: " + err.Error())
	}

	// create gRPC server
	s := grpc.NewServer()
	redisClient, _ := cache.NewRedis()
	pb.RegisterAuthServiceServer(s, service.NewService(redisClient))

	log.Printf("start auth service in: %s", host)

	if err := s.Serve(listener); err != nil {
		panic("error building auth service: " + err.Error())
	}
}
