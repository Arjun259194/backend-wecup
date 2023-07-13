package controllers

import "github.com/Arjun259194/wecup-go/database"

type Controller struct {
	Storage *database.Storage
}

func NewController(str *database.Storage) Controller {
	return Controller{
		Storage: str,
	}
}
