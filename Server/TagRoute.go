package FiberServer

import (
	"strconv"
	"web-api/Controllers"
	"web-api/Types/Api/Request"

	"github.com/gofiber/fiber/v2"
)

type TagRoute struct {
	TagController Controllers.TagController
}

func (f *FiberServer) initTagRoutes(base fiber.Router) {
	r := base.Group("/tag")
	r.Get("/", f.GetAllTags)
	r.Get("/", f.GetTagByIDs)
	r.Post("/", f.CreateTag)
	r.Put("/", f.UpdateTag)
	r.Delete("/:tagId", f.DeleteTag)
}

func (f *FiberServer) GetAllTags(c *fiber.Ctx) error {
	data := f.Controllers.TagController.GetAllTags()
	if !data.Success {
		return c.Status(500).JSON(data)
	}

	return c.Status(200).JSON(data)
}

func (f *FiberServer) GetTagByIDs(c *fiber.Ctx) error {
	var request Request.GetTagByIDsRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	data := f.Controllers.TagController.GetTagByIDs(request)
	if !data.Success {
		return c.Status(500).JSON(data)
	}

	return c.Status(200).JSON(data)
}

func (f *FiberServer) CreateTag(c *fiber.Ctx) error {
	var request Request.CreateTagRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	data := f.Controllers.TagController.CreateTag(request)
	if !data.Success {
		return c.Status(500).JSON(data)
	}

	return c.Status(200).JSON(data)
}

func (f *FiberServer) UpdateTag(c *fiber.Ctx) error {
	var request Request.UpdateTagRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	data := f.Controllers.TagController.UpdateTag(request)
	if !data.Success {
		return c.Status(500).JSON(data)
	}

	return c.Status(200).JSON(data)
}

func (f *FiberServer) DeleteTag(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("tagId"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid Tag ID"})
	}

	data := f.Controllers.TagController.DeleteTag(uint(id))
	if !data.Success {
		return c.Status(500).JSON(data)
	}

	return c.Status(200).JSON(data)
}
