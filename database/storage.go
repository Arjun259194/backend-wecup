package database

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Storage struct {
	ConnectionString string
	Client           *mongo.Client
	Database         *mongo.Database
	UserModel        *mongo.Collection
	PostModel        *mongo.Collection
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
	s.Database = client.Database("wecup")
	s.UserModel = s.Database.Collection("User")
	s.PostModel = s.Database.Collection("Post")

	// Create unique index on email field
	indexModel := mongo.IndexModel{
		Keys:    bson.M{"email": 1},
		Options: options.Index().SetUnique(true),
	}

	_, err = s.UserModel.Indexes().CreateOne(s.Ctx, indexModel)
	if err != nil {
		log.Fatalf("Error while creating unique index on email field - %v", err)
	}
}

func (s *Storage) Close() {
	if err := s.Client.Disconnect(s.Ctx); err != nil {
		log.Fatalf("Error while disconnecting database in storage.go - %v", err)
	}
}
