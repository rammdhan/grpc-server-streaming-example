package service

import (
	"context"
	"fmt"
	"grpc-server-streaming-example/cache"
	pb "grpc-server-streaming-example/proto/auth"
	"time"

	"google.golang.org/protobuf/types/known/emptypb"
)

type authService struct {
	redis cache.RedisInterface
	pb.UnimplementedAuthServiceServer
}

func NewService(redis cache.RedisInterface) pb.AuthServiceServer {
	return &authService{redis: redis}
}

func (u *authService) Login(ctx context.Context, request *pb.LoginRequest) (response *emptypb.Empty, err error) {
	err = u.redis.Publish(ctx, request.Email, time.Now().UnixMilli())
	if err != nil {
		fmt.Println("got error when publishing message:", err)
	}

	return
}
