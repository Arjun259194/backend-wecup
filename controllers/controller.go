package controllers

import (
	"github.com/Arjun259194/wecup-go/database"
	"github.com/Arjun259194/wecup-go/types"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Controller struct {
	Storage *database.Storage
}

var val *validator.Validate = validator.New()

func NewController(str *database.Storage) Controller {
	return Controller{
		Storage: str,
	}
}

func (ctrl *Controller) GetUserByID(ID primitive.ObjectID) (*types.User, error) {
	filter := bson.M{
		"_id": ID,
	}
	result, err := ctrl.Storage.FindOneUser(filter)

	return result, err
}
