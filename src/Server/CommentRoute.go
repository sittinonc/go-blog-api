package FiberServer

import (
	"strconv"
	"web-api/src/Controllers"
	"web-api/src/Types/Api/Request"

	"github.com/gofiber/fiber/v2"
)

type CommentRoute struct {
	CommentController Controllers.CommentController
}

func (f FiberServer) initCommentRoutes(base fiber.Router) {
	r := base.Group("/comment")
	r.Get("/:commentId", f.GetCommentByID)
	r.Post("/", f.CreateComment)
	r.Put("/", f.UpdateComment)
	r.Delete("/:commentId", f.DeleteComment)
}

func (f FiberServer) GetCommentByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("commentId"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid Comment ID"})
	}

	data := f.Controllers.CommentController.GetCommentByID(uint(id))
	if !data.Success {
		return c.Status(500).JSON(data)
	}

	return c.Status(200).JSON(data)
}

func (f FiberServer) CreateComment(c *fiber.Ctx) error {
	var request Request.CreateCommentRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	data := f.Controllers.CommentController.CreateComment(request)
	if !data.Success {
		return c.Status(500).JSON(data)
	}

	return c.Status(200).JSON(data)
}

func (f FiberServer) UpdateComment(c *fiber.Ctx) error {
	var request Request.UpdateCommentRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	data := f.Controllers.CommentController.UpdateComment(request)
	if !data.Success {
		return c.Status(500).JSON(data)
	}

	return c.Status(200).JSON(data)
}

func (f FiberServer) DeleteComment(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("commentId"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid Comment ID"})
	}

	data := f.Controllers.CommentController.DeleteComment(uint(id))
	if !data.Success {
		return c.Status(500).JSON(data)
	}

	return c.Status(200).JSON(data)
}
