package controllers

import (
	"github.com/Arjun259194/wecup-go/database"
	"github.com/go-playground/validator/v10"
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
