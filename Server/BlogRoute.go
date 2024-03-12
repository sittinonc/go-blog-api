package FiberServer

import (
	"strconv"
	"web-api/Controllers"
	"web-api/Types/Api/Request"

	"github.com/gofiber/fiber/v2"
)

type BlogRoute struct {
	BlogController Controllers.BlogController
}

func (f FiberServer) initBlogRoutes(base fiber.Router) {
	r := base.Group("/blog")
	r.Get("/:blogId", f.GetBlogByID)
	r.Post("/", f.CreateBlog)
	r.Put("/", f.UpdateBlog)
	r.Delete("/:blogId", f.DeleteBlog)
}

func (f *FiberServer) GetBlogByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("blogId"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid Blog ID"})
	}

	data := f.Controllers.BlogController.GetBlogByID(uint(id))
	if !data.Success {
		return c.Status(500).JSON(data)
	}

	return c.Status(200).JSON(data)
}

func (f *FiberServer) CreateBlog(c *fiber.Ctx) error {
	var request Request.CreateBlogRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	data := f.Controllers.BlogController.CreateBlog(request)
	if !data.Success {
		return c.Status(500).JSON(data)
	}

	return c.Status(200).JSON(data)
}

func (f *FiberServer) UpdateBlog(c *fiber.Ctx) error {
	var request Request.UpdateBlogRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	data := f.Controllers.BlogController.UpdateBlog(request)
	if !data.Success {
		return c.Status(500).JSON(data)
	}

	return c.Status(200).JSON(data)
}

func (f *FiberServer) DeleteBlog(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("blogId"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid Blog ID"})
	}

	data := f.Controllers.BlogController.DeleteBlog(uint(id))
	if !data.Success {
		return c.Status(500).JSON(data)
	}

	return c.Status(200).JSON(data)
}
