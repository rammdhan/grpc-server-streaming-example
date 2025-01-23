package service

import (
	"fmt"
	"grpc-server-streaming-example/cache"
	pb "grpc-server-streaming-example/proto/user"

	"google.golang.org/grpc"
)

type userService struct {
	redis cache.RedisInterface
	pb.UnimplementedUserServiceServer
}

func NewService(redis cache.RedisInterface) pb.UserServiceServer {
	return &userService{redis: redis}
}

func (u *userService) GetUserLastLogin(request *pb.GetUserLastLoginRequest, server grpc.ServerStreamingServer[pb.GetUserLastLoginResponse]) (err error) {
	subs := u.redis.Subscribe(server.Context(), request.Email)

	for {
		select {
		case <-server.Context().Done():
			return server.Context().Err()
		case message := <-subs.Channel():
			err := server.Send(&pb.GetUserLastLoginResponse{
				Content: message.Payload,
			})

			if err != nil {
				return fmt.Errorf(err.Error())
			}
		}
	}
}
