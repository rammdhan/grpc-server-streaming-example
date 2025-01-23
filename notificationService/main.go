package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "grpc-server-streaming-example/proto/user"
)

func main() {
	godotenv.Load()

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Please enter user email:")
	email, _ := reader.ReadString('\n')
	email = strings.Trim(email, "\n")
	fmt.Println()
	
	// dial to server
	host := fmt.Sprintf("localhost:%s", os.Getenv("USER_PORT"))
	conn, err := grpc.NewClient(host, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println("Error connecting to gRPC server: ", err.Error())
	}

	defer conn.Close()

	// create the stream
	client := pb.NewUserServiceClient(conn)

	stream, err := client.GetUserLastLogin(context.Background(), &pb.GetUserLastLoginRequest{
		Email: email,
	})
	if err != nil {
		panic(err) // dont use panic in your real project
	}

	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf(err.Error())
		}

		log.Printf("User %s was logged in at: %s , please contact support if thats not you", email, resp.Content)
	}
}
