package main

import (
	"log"
	"os"

	grpc_server "github.com/axel-andrade/deu-role-auth/internal/adapters/primary/grpc/server"
	mongo_database "github.com/axel-andrade/deu-role-auth/internal/adapters/secondary/database/mongo"
	redis_database "github.com/axel-andrade/deu-role-auth/internal/adapters/secondary/database/redis"
	"github.com/axel-andrade/deu-role-auth/internal/infra"
	"github.com/joho/godotenv"
)

/*
* The init function is called after all the variable declarations in the package have evaluated their initializers, and
* those are evaluated only after all the imported packages have been initialized.
 */
func init() {
	if os.Getenv("ENV") != "production" {
		err := godotenv.Load()
		if err != nil {
			log.Fatalf("Error loading .env file: %v", err)
		}
	}

	mongo_database.ConnectDB()
	redis_database.ConnectRedisDB()
}

func main() {
	d := infra.LoadDependencies()

	grpcPort := os.Getenv("GRPC_PORT")
	grpcAddress := ":50051" // Default gRPC server port

	if grpcPort != "" {
		grpcAddress = ":" + grpcPort
	}

	// Run gRPC server
	grpc_server.RunGRPCServer(grpcAddress, d)
}
