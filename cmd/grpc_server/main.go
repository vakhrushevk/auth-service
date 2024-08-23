package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/vakhrushevk/auth-service/pkg/user_v1"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const grpcPort = 50051

type server struct {
	user_v1.UnimplementedUserV1Server
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

// CreateUser - This method is intended to create a new User
func (s server) CreateUser(_ context.Context, request *user_v1.CreateUserRequest) (*user_v1.CreateUserResponse, error) {
	log.Println("Received Create request")
	log.Printf("Name: %s\n Email: %s\n Password: %s\n PasswordConfirm %s\n Role %s",
		request.Name, request.Email, request.Password, request.PasswordConfirm, request.Role.String())
	return &user_v1.CreateUserResponse{Id: 1}, nil
}

// GetUserById - This method is designed to get the user by ID
func (s server) GetUserById(_ context.Context, request *user_v1.GetUserByIdRequest) (*user_v1.GetUserByIdResponse, error) {
	if request.Id == 0 {
		return &user_v1.GetUserByIdResponse{}, errors.New("failed to request: empty id")
	}

	fmt.Println("Method Get | request: ", request.GetId())

	return &user_v1.GetUserByIdResponse{
		Id:        request.GetId(),
		Name:      "Konstantin",
		Email:     "Random@gmail.g",
		Role:      user_v1.Role_USER,
		UpdatedAt: timestamppb.New(time.Now()),
		CreatedAt: timestamppb.New(time.Unix(3200, 1211)),
	}, nil
}

// UpdateUser - This method allows you to update user fields
func (s server) UpdateUser(_ context.Context, request *user_v1.UpdateUserRequest) (*emptypb.Empty, error) {
	log.Println("Received Update request")
	log.Fatalf("id: %d\n name: %s\n email: %s\n", request.GetId(), request.GetName(), request.GetName())

	return nil, nil
}

// DeleteUserById - This method allows delete User by id
func (s server) DeleteUserById(_ context.Context, request *user_v1.DeleteUserRequest) (*emptypb.Empty, error) {
	log.Println("Received Delete request")
	log.Println("id: ", request.GetId())
	return nil, nil
}
