package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/vakhrushevk/auth-service/pkg/user_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const grpcPort = 50051

type server struct {
	user_v1.UnimplementedUserV1Server
}

//nolint:revive
func (s server) Create(ctx context.Context, request *user_v1.CreateRequest) (*user_v1.CreateResponse, error) {
	log.Println("Received Create request")
	log.Printf("Name: %s\n Email: %s\n Password: %s\n PasswordConfirm %s\n Role %s",
		request.Name, request.Email, request.Password, request.PasswordConfirm, request.Role.String())

	return &user_v1.CreateResponse{Id: 1}, nil
}

//nolint:revive
func (s server) Get(ctx context.Context, request *user_v1.GetRequest) (*user_v1.GetResponse, error) {
	if request.Id == 0 {
		return &user_v1.GetResponse{}, errors.New("failed to request: empty id")
	}

	fmt.Println("Method Get | request: ", request.GetId())

	return &user_v1.GetResponse{
		Id:        request.GetId(),
		Name:      "Konstantin",
		Email:     "Random@gmail.g",
		Role:      user_v1.Role_USER,
		UpdatedAt: timestamppb.New(time.Now()),
		CreatedAt: timestamppb.New(time.Unix(3200, 1211)),
	}, nil
}

//nolint:revive
func (s server) Update(ctx context.Context, request *user_v1.UpdateRequest) (*emptypb.Empty, error) {
	log.Println("Received Update request")
	log.Println("id: ", request.GetId())
	log.Println("name: ", request.GetName())
	log.Println("email: ", request.GetEmail())

	return nil, nil
}

//nolint:revive
func (s server) Delete(ctx context.Context, request *user_v1.DeleteRequest) (*emptypb.Empty, error) {
	log.Println("Received Delete request")
	log.Println("id: ", request.GetId())

	return nil, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	srv := grpc.NewServer()
	reflection.Register(srv)
	user_v1.RegisterUserV1Server(srv, &server{})

	log.Printf("server listening on %s", lis.Addr().String())

	if err := srv.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
