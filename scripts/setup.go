package main

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ctx = context.TODO()
var client *mongo.Client

func main() {
	// Carrega variáveis de ambiente do arquivo .env
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Erro ao carregar variáveis de ambiente: %v", err)
	}

	// Conecta ao MongoDB
	clientOptions := options.Client().ApplyURI(os.Getenv("DB_URI"))
	client, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		log.Fatalf("Erro ao conectar ao MongoDB: %v", err)
	}

	defer client.Disconnect(ctx)

	// Obtém a coleção
	collection := client.Database(os.Getenv("DB_NAME")).Collection("users")

	// Criação do índice
	indexName, err := collection.Indexes().CreateOne(
		ctx,
		mongo.IndexModel{
			Keys:    bson.D{{Key: "email", Value: 1}},
			Options: options.Index().SetUnique(true),
		},
	)

	// If index already exists, ignore error
	if err != nil && err.Error() != "index with name: email_1 already exists with different options" {
		log.Fatalf("Erro ao criar índice: %v", err)
	}

	log.Printf("Índice criado com sucesso: %s", indexName)
}
