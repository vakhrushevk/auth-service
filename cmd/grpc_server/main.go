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

// CreateUser - Создать пользователя
//
//nolint:revive
func (s server) CreateUser(ctx context.Context, request *user_v1.CreateUserRequest) (*user_v1.CreateUserResponse, error) {
	log.Println("Received Create request")
	log.Printf("Name: %s\n Email: %s\n Password: %s\n PasswordConfirm %s\n Role %s",
		request.Name, request.Email, request.Password, request.PasswordConfirm, request.Role.String())
	return &user_v1.CreateUserResponse{Id: 1}, nil
}

// GetUserById - Получить пользователя по идентификатору
//
//nolint:revive
func (s server) GetUserById(ctx context.Context, request *user_v1.GetUserByIdRequest) (*user_v1.GetUserByIdResponse, error) {
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

// UpdateUser - Обновить пользователя
//
//nolint:revive
func (s server) UpdateUser(ctx context.Context, request *user_v1.UpdateUserRequest) (*emptypb.Empty, error) {
	log.Println("Received Update request")
	log.Fatalf("id: %d\n name: %s\n email: %s\n", request.GetId(), request.GetName(), request.GetEmail())

	return nil, nil
}

// DeleteUserById - Удалить пользователя по идентификатору
//
//nolint:revive
func (s server) DeleteUserById(ctx context.Context, request *user_v1.DeleteUserRequest) (*emptypb.Empty, error) {
	log.Println("Received Delete request")
	log.Println("id: ", request.GetId())
	return nil, nil
}
