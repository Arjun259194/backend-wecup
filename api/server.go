package api

import (
	"github.com/Arjun259194/wecup-go/controllers"
	"github.com/Arjun259194/wecup-go/database"
	"github.com/gofiber/fiber/v2"
)

var controller controllers.Controller

type Server struct {
	Addr      string
	DBConnUrl string
}

func NewServer(addr, dburl string) *Server {
	return &Server{
		Addr:      addr,
		DBConnUrl: dburl,
	}
}

func (s *Server) Start() error {
	storage := database.NewStorage(s.DBConnUrl)
	storage.Connect()
	defer storage.Close()

	controller = controllers.NewController(storage)

	server := fiber.New()

	setRoutes(server)

	return server.Listen(s.Addr)
}
