package main

import (
	"fmt"
	"os"
	FiberServer "web-api/src/Server"

	"github.com/joho/godotenv"
)

func initEnvironment() FiberServer.ServerConfig {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file:", err)
		return FiberServer.ServerConfig{}
	}
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	dbname := os.Getenv("DB_NAME")
	password := os.Getenv("DB_PASSWORD")
	sslmode := os.Getenv("DB_SSLMODE")

	return FiberServer.ServerConfig{
		Port:            os.Getenv("SERVER_PORT"),
		DatabaseAddress: fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", host, port, user, dbname, password, sslmode),
	}
}

func main() {
	config := initEnvironment()
	controllers := FiberServer.ServerControllers{}
	server := FiberServer.NewFiberServer(&config, &controllers)

	server.Fiber.Listen(":" + config.Port)
}
