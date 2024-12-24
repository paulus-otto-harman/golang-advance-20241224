package service

import (
	pb "20241224/auth-service/proto"
	"context"
	"log"
)

type AuthService struct {
	pb.UnimplementedAuthServiceServer
}

type User struct {
	Username string
	Password string
}

func (a *AuthService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	login := User{
		Username: req.Username,
		Password: req.Password,
	}
	log.Println(login)
	return &pb.LoginResponse{Success: true, Token: "abcd"}, nil
}

func (a *AuthService) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	return nil, nil
}
