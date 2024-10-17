package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/vakhrushevk/auth-service/internal/config"
	"github.com/vakhrushevk/auth-service/internal/config/env"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/vakhrushevk/auth-service/pkg/user_v1"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config-path", ".env", "path to config file")
}

type server struct {
	user_v1.UnimplementedUserV1Server
}

func main() {
	ctx := context.Background()
	// CONFIG ++
	flag.Parse()
	err := config.Load(configPath)
	if err != nil {
		log.Fatalf("failed to load config %v", err)
	}
	grpcConfig, err := env.NewGRPCConfig()
	if err != nil {
		log.Fatalf("Failet to get grpc Config: %v", err)
	}
	pgConfig, err := env.NewPGConfig()
	if err != nil {
		log.Fatalf("Failet to get pg Config: %v", err)
	}
	_ = pgConfig
	// CONFIG --
	// POSTGRES INIT ++
	con, err := pgx.Connect(ctx, pgConfig.DSN())

	if err != nil {
		log.Fatalf("failed to connect to database%v", err)
	}
	defer func() {
		err = con.Close(ctx)
		if err != nil {
			log.Panicf("failed to close connection: %v", err)
		}
	}()

	// POSTGRES --
	lis, err := net.Listen("tcp", grpcConfig.Address()) // fmt.Sprintf("%d", grpcPort)
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

// GetUserByID - This method is designed to get the user by ID
func (s server) GetUserByID(_ context.Context, request *user_v1.GetUserByIDRequest) (*user_v1.GetUserByIDResponse, error) {
	if request.Id == 0 {
		return &user_v1.GetUserByIDResponse{}, errors.New("failed to request: empty id")
	}

	fmt.Println("Method Get | request: ", request.GetId())

	return &user_v1.GetUserByIDResponse{
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

// DeleteUserByID - This method allows delete User by id
func (s server) DeleteUserByID(_ context.Context, request *user_v1.DeleteUserByIDRequest) (*emptypb.Empty, error) {
	log.Println("Received Delete request")
	log.Println("id: ", request.GetId())
	return nil, nil
}
