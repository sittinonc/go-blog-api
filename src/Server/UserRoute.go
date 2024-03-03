package FiberServer

import (
	"strconv"
	"web-api/src/Controllers"
	"web-api/src/Types/Api/Request"

	"github.com/gofiber/fiber/v2"
)

type UserRoute struct {
	UserController Controllers.UserController
}

func (f FiberServer) initUserRoutes(base fiber.Router) {
	r := base.Group("/user")
	r.Get("/:userId", f.GetUserByID)
	r.Post("/", f.CreateUser)
	r.Put("/", f.UpdateUser)
}

func (f FiberServer) GetUserByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("userId"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid User ID"})
	}

	data := f.Controllers.UserController.GetUserByID(uint(id))
	if !data.Success {
		return c.Status(500).JSON(data)
	}

	return c.Status(200).JSON(data)
}

func (f FiberServer) CreateUser(c *fiber.Ctx) error {
	var request Request.CreateUserRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	data := f.Controllers.UserController.CreateUser(request)
	if !data.Success {
		return c.Status(500).JSON(data)
	}

	return c.Status(200).JSON(data)
}

func (f FiberServer) UpdateUser(c *fiber.Ctx) error {
	var request Request.UpdateUserRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	data := f.Controllers.UserController.UpdateUser(request)
	if !data.Success {
		return c.Status(500).JSON(data)
	}

	return c.Status(200).JSON(data)
}
