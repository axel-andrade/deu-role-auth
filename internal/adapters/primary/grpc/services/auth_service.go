package grpc_services

import (
	"context"

	pb "github.com/axel-andrade/deu-role-auth/internal/adapters/primary/grpc/pb"
	"github.com/axel-andrade/deu-role-auth/internal/core/usecases/auth/login"
	"github.com/axel-andrade/deu-role-auth/internal/core/usecases/auth/signup"
)

type AuthGrpcService struct {
	pb.UnimplementedAuthServiceServer
	loginUC  login.LoginUC
	signupUC signup.SignupUC
}

func NewAuthGrpcService(l login.LoginUC, s signup.SignupUC) *AuthGrpcService {
	return &AuthGrpcService{loginUC: l, signupUC: s}
}

func (s *AuthGrpcService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	input := login.LoginInputDTO{
		Email:    req.Email,
		Password: req.Password,
	}

	output, err := s.loginUC.Execute(input)
	if err != nil {
		return nil, err
	}
	return &pb.LoginResponse{AccessToken: output.AccessToken, RefressToken: output.RefreshToken}, nil
}

func (s *AuthGrpcService) Signup(ctx context.Context, req *pb.SignupRequest) (*pb.SignupResponse, error) {
	input := signup.SignupInputDTO{
		Email:    req.Email,
		Password: req.Password,
	}

	output, err := s.signupUC.Execute(input)
	if err != nil {
		return nil, err
	}
	return &pb.SignupResponse{UserId: output.User.ID}, nil
}
