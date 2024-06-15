package main

import (
	"context"
	"log"

	pb "github.com/axel-andrade/deu-role-auth/internal/adapters/primary/grpc/pb"
	"google.golang.org/grpc"
)

const (
	serverAddr = "localhost:50051"
)

// AuthServiceClient holds the gRPC client and server address
type AuthServiceClient struct {
	conn   *grpc.ClientConn
	client pb.AuthServiceClient
}

// NewAuthServiceClient creates a new instance of AuthServiceClient
func NewAuthServiceClient(serverAddr string) (*AuthServiceClient, error) {
	// Estabelecer conexão com o servidor gRPC
	conn, err := grpc.Dial(serverAddr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return &AuthServiceClient{
		conn:   conn,
		client: pb.NewAuthServiceClient(conn),
	}, nil
}

// Close closes the gRPC connection
func (c *AuthServiceClient) Close() error {
	return c.conn.Close()
}

// Login performs the login operation
func (c *AuthServiceClient) Login(email, password string) (*pb.LoginResponse, error) {
	resp, err := c.client.Login(context.Background(), &pb.LoginRequest{
		Email:    email,
		Password: password,
	})
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// Signup performs the signup operation
func (c *AuthServiceClient) Signup(name, email, password string) (*pb.SignupResponse, error) {
	resp, err := c.client.Signup(context.Background(), &pb.SignupRequest{
		Name:     name,
		Email:    email,
		Password: password,
	})
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func main() {
	// Criar um cliente gRPC para AuthService
	authClient, err := NewAuthServiceClient(serverAddr)
	if err != nil {
		log.Fatalf("Falha ao criar cliente gRPC: %v", err)
	}
	defer authClient.Close()

	// Exemplo de utilização do método Login
	respLogin, err := authClient.Login("ajaxeljunio@gmail.com", "password123")
	if err != nil {
		log.Fatalf("Falha ao chamar Login: %v", err)
	}
	log.Printf("Token de acesso: %s", respLogin.AccessToken)
	log.Printf("Token de atualização: %s", respLogin.RefreshToken)

	// Exemplo de utilização do método Signup
	// respSignup, err := authClient.Signup("Axel", "ajaxeljunio@gmail.com", "password123")
	// if err != nil {
	// 	log.Fatalf("Falha ao chamar Signup: %v", err)
	// }
	// log.Printf("Novo usuário criado com ID: %s", respSignup.UserId)
}
