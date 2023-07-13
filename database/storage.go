package database

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Storage struct {
	ConnectionString string
	Client           *mongo.Client
	UserModel        *mongo.Collection
	Ctx              context.Context
}

func NewStorage(connectionString string) *Storage {
	return &Storage{
		ConnectionString: connectionString,
		Ctx:              context.Background(),
	}
}

func (s *Storage) Connect() {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(s.ConnectionString).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(s.Ctx, opts)
	if err != nil {
		log.Fatalf("There was an error while connecting to database - %v", err)
	}

	if err = client.Ping(s.Ctx, nil); err != nil {
		log.Fatalf("Error while sending ping to database - %v", err)
	}

	fmt.Println("Ping! Connected to database")

	s.Client = client
	s.UserModel = client.Database("wecup").Collection("User")
}

func (s *Storage) Close() {
	if err := s.Client.Disconnect(s.Ctx); err != nil {
		log.Fatalf("Error while disconnecting database in storage.go - %v", err)
	}
}
