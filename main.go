package main

import (
	"log"

	"github.com/Arjun259194/wecup-go/api"
	"github.com/Arjun259194/wecup-go/config"
)

func main() {
	server := api.NewServer(config.PORT, config.DB_CONNECTION_URI)
	log.Fatal(server.Start())
}
