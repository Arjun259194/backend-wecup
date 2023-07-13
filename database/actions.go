package database

import (
	"github.com/Arjun259194/wecup-go/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *Storage) AddUser(user types.User) (*mongo.InsertOneResult, error) {
	return s.UserModel.InsertOne(s.Ctx, user)
}

func (s *Storage) FindOneUser(filter bson.M) *mongo.SingleResult {
	return s.UserModel.FindOne(s.Ctx, filter)
}

func (s *Storage) FindUser(filter bson.M) (cur *mongo.Cursor, err error) {
	return s.UserModel.Find(s.Ctx, filter)
}
