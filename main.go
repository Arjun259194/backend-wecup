package main

import (
	"log"
	"os"

	"github.com/Arjun259194/wecup-go/api"
	"github.com/joho/godotenv"
)

func getEnv(key string) string {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error while loading enviroment variables - %v", err)
	}
	return os.Getenv(key)
}

func main() {
	var (
		PORT            string = getEnv("PORT")
		DBCONNECTIONURI string = getEnv("DB_CONNECTION_URI")
	)
	server := api.NewServer(PORT, DBCONNECTIONURI)
	log.Fatal(server.Start())
}
