package router

import (
	"context"
	"user-service/domain"
	"user-service/pb"
	"user-service/usecase"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type userGRPCServer struct {
	pb.UnimplementedUserServiceServer
	uu usecase.IUserUsecase
}

func NewUserGRPCServer(grpcServer *grpc.Server, uu usecase.IUserUsecase) pb.UserServiceServer {
	s := userGRPCServer{uu: uu}
	pb.RegisterUserServiceServer(grpcServer, &s)
	reflection.Register(grpcServer)
	return &s
}

func (s *userGRPCServer) Signup(ctx context.Context, req *pb.SignupRequest) (*pb.SignupResponse, error) {
	res := pb.SignupResponse{}
	signupRequest := domain.SignupRequest{Username: req.Username, Email: req.Email, Password: req.Password}
	signupResponse, err := s.uu.Signup(signupRequest)
	if err != nil {
		return nil, err
	}
	res.Id = int32(signupResponse.ID)
	res.Username = signupResponse.Username
	res.Email = signupResponse.Email

	return &res, nil
}

func (s *userGRPCServer) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	res := pb.LoginResponse{}
	loginRequest := domain.LoginRequest{Email: req.Email, Password: req.Password}
	loginResponse, err := s.uu.Login(loginRequest)
	if err != nil {
		return nil, err
	}
	res.UserId = int32(loginResponse.UserID)
	res.Token = loginResponse.Token

	return &res, nil
}
