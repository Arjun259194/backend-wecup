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

func (s *Storage) CreateNewUser(user types.User) error {
	_, err := s.UserModel.InsertOne(s.Ctx, user)
	if err != nil {
		return err
	}
	return nil
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

func (s *Storage) FindByIDAndUpdateUser(ID primitive.ObjectID, body types.UserUpdateRequest) error {
	filter := bson.M{"_id": ID}

	update := bson.M{
		"$set": bson.M{"name": body.Name, "email": body.Email, "gender": body.Gender},
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

	clientFilter := bson.M{"_id": clientID}

	userFilter := bson.M{"_id": userID}

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

func (s *Storage) CreateNewPost(p types.Post) error {
	_, err := s.PostModel.InsertOne(s.Ctx, p)
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) FindOnePostByID(ID primitive.ObjectID) (*types.Post, error) {
	filter := bson.M{"_id": ID}

	resutl := s.PostModel.FindOne(s.Ctx, filter)

	if err := resutl.Err(); err != nil {
		return nil, err
	}

	foundPost := new(types.Post)
	if err := resutl.Decode(foundPost); err != nil {
		return nil, err
	}

	return foundPost, nil
}

func (s *Storage) FindOnePostByIDAndUpdate(ID primitive.ObjectID, req *types.UpdatePostRequest) error {
	filter := bson.M{"_id": ID}

	update := bson.M{
		"$set": bson.M{"content": req.Content},
	}

	result := s.PostModel.FindOneAndUpdate(s.Ctx, filter, update)
	if err := result.Err(); err != nil {
		return err
	}

	return nil
}

func (s *Storage) FindOnePostByIDAndDelete(ID primitive.ObjectID) error {
	filter := bson.M{"_id": ID}

	result := s.PostModel.FindOneAndDelete(s.Ctx, filter)

	if err := result.Err(); err != nil {
		return err
	}

	return nil
}

func (s *Storage) LikeOrUnlikePost(postID, clientID primitive.ObjectID, isLiked bool) error {
	filter := bson.M{"_id": postID}

	var update bson.M

	if isLiked {
		update = bson.M{"$pull": bson.M{"likes": clientID}}
	} else {
		update = bson.M{"$push": bson.M{"likes": clientID}}
	}

	result := s.PostModel.FindOneAndUpdate(s.Ctx, filter, update)

	if err := result.Err(); err != nil {
		return err
	}

	return nil
}

func (s *Storage) CreateNewComment(postID primitive.ObjectID, comment *types.Comment) error {
	filter := bson.M{"_id": postID}

	update := bson.M{
		"$push": bson.M{"comments": comment},
	}

	result := s.PostModel.FindOneAndUpdate(s.Ctx, filter, update)
	if err := result.Err(); err != nil {
		return err
	}

	return nil
}
