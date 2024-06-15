package grpc_server

import (
	"log"
	"net"

	pb "github.com/axel-andrade/deu-role-auth/internal/adapters/primary/grpc/pb"
	"github.com/axel-andrade/deu-role-auth/internal/infra"

	"google.golang.org/grpc"
)

func RunGRPCServer(address string, d *infra.Dependencies) {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterAuthServiceServer(grpcServer, d.AuthGrpcService)

	log.Printf("gRPC server listening on %s", address)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
