package database

import (
	"github.com/Arjun259194/wecup-go/types"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *Storage) AddUser(user types.User) (*mongo.InsertOneResult, error) {
	return s.UserModel.InsertOne(s.Ctx, user)
}
