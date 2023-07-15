package database

import (
	"github.com/Arjun259194/wecup-go/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *Storage) FindUser(filter bson.M) (cur *mongo.Cursor, err error) {
	return s.UserModel.Find(s.Ctx, filter)
}

func (s *Storage) AddUser(user types.User) (*mongo.InsertOneResult, error) {
	return s.UserModel.InsertOne(s.Ctx, user)
}

func (s *Storage) FindOneUser(filter bson.M) (*types.User, error) {

	result := s.UserModel.FindOne(s.Ctx, filter)

	if err := result.Err(); err != nil {
		return nil, err
	}

	var foundUser types.User

	if err := result.Decode(&foundUser); err != nil {
		return nil, err
	}

	return &foundUser, nil
}

func (s *Storage) FindByIDAndUpdateUser(ID primitive.ObjectID, body types.UpdateRequest) error {
	filter := bson.M{
		"_id": ID,
	}

	update := bson.M{
		"$set": bson.M{
			"name":   body.Name,
			"email":  body.Email,
			"gender": body.Gender,
		},
	}

	result := s.UserModel.FindOneAndUpdate(s.Ctx, filter, update)

	if err := result.Err(); err != nil {
		return err
	}

	return nil
}

// This function will update follower and following status in client and requested user.
// if isFollowing is false then it will make clientID follow userID and if true then unfollow
func (s *Storage) FindByIDAndFollowOrUnfollow(userID, clientID primitive.ObjectID, isFollowing bool) error {

	var userUpdateQuery bson.M   // user to follow
	var clientUpdateQuery bson.M // requesting user

	if isFollowing == true {
		userUpdateQuery = bson.M{"$pull": bson.M{"followers": clientID}}
		clientUpdateQuery = bson.M{"$pull": bson.M{"following": userID}}
	} else {
		userUpdateQuery = bson.M{"$push": bson.M{"followers": clientID}}
		clientUpdateQuery = bson.M{"$push": bson.M{"following": userID}}
	}

	clientFilter := bson.M{
		"_id": clientID,
	}

	userFilter := bson.M{
		"_id": userID,
	}

	result := s.UserModel.FindOneAndUpdate(s.Ctx, clientFilter, clientUpdateQuery)
	if err := result.Err(); err != nil {
		return err
	}

	result = s.UserModel.FindOneAndUpdate(s.Ctx, userFilter, userUpdateQuery)
	if err := result.Err(); err != nil {
		return err
	}

  return nil
}
