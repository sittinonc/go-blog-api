package FiberServer

import (
	"web-api/src/Controllers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type ServerConfig struct {
	Port            string
	DatabaseAddress string
}

type ServerControllers struct {
	BlogController Controllers.BlogController
	UserController Controllers.UserController
	TagController  Controllers.TagController
}

type FiberServer struct {
	Config      *ServerConfig
	Fiber       *fiber.App
	Controllers *ServerControllers
}

func NewFiberServer(config *ServerConfig, controllers *ServerControllers) *FiberServer {
	server := fiber.New(fiber.Config{
		CaseSensitive: false,
	})

	// for local development
	server.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "*",
		AllowMethods: "GET, POST, PUT, PATCH, DELETE, OPTIONS",
	}))

	f := &FiberServer{
		Config:      config,
		Fiber:       server,
		Controllers: controllers,
	}

	f.initBlogRoutes(server)
	f.initUserRoutes(server)
	f.initTagRoutes(server)

	return f
}
