package types

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID   `json:"id" bson:"_id"`
	Name      string               `json:"name" bson:"name"`
	Email     string               `json:"email" bson:"email"`
	Gender    string               `json:"gender" bson:"gender"`
	Password  string               `json:"password" bson:"password"`
	Followers []primitive.ObjectID `json:"followers" bson:"followers"`
	Following []primitive.ObjectID `json:"following" bson:"following"`
}

func NewUser(name, email, password, gender string) User {
	return User{
		Name:      name,
		Email:     email,
		Gender:    gender,
		Password:  password,
		ID:        primitive.NewObjectID(),
		Followers: make([]primitive.ObjectID, 0),
		Following: make([]primitive.ObjectID, 0),
	}
}