package types

import (
	"time"

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

type Post struct {
	ID        primitive.ObjectID   `json:"id" bson:"_id"`
	UserID    primitive.ObjectID   `json:"userId" bson:"userId"`
	CreatedAt time.Time            `json:"createdAt" bson:"createdAt"`
	Content   string               `json:"content" bson:"content"`
	Likes     []primitive.ObjectID `json:"likes" bson:"likes"`
	Comments  []Comment            `json:"comments" bson:"comments"`
}

func NewPost(userID primitive.ObjectID, content string) *Post {
	return &Post{
		ID:        primitive.NewObjectID(),
		UserID:    userID,
		CreatedAt: time.Now(),
		Content:   content,
		Likes:     make([]primitive.ObjectID, 0),
		Comments:  make([]Comment, 0),
	}
}

type Comment struct {
	ID      primitive.ObjectID   `json:"id" bson:"_id"`
	UserID  primitive.ObjectID   `json:"userId" bson:"userId"`
	PostID  primitive.ObjectID   `json:"postId" bson:"postId"`
	Likes   []primitive.ObjectID `json:"likes" bson:"likes"`
	Content string               `json:"content" bson:"content"`
}

func NewComment(userID, postID primitive.ObjectID, content string) *Comment {
	return &Comment{
		ID:      primitive.NewObjectID(),
		UserID:  userID,
		PostID:  postID,
		Likes:   make([]primitive.ObjectID, 0),
		Content: content,
	}
}

type UserResponse struct {
	ID        primitive.ObjectID   `json:"id" bson:"_id"`
	Name      string               `json:"name" bson:"name"`
	Email     string               `json:"email" bson:"email"`
	Gender    string               `json:"gender" bson:"gender"`
	Followers []primitive.ObjectID `json:"followers" bson:"followers"`
	Following []primitive.ObjectID `json:"following" bson:"following"`
}

func (u User) GetResponse() UserResponse {
	return UserResponse{
		ID:        u.ID,
		Name:      u.Name,
		Email:     u.Email,
		Gender:    u.Gender,
		Followers: u.Followers,
		Following: u.Following,
	}
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
