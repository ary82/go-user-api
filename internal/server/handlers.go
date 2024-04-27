package server

import "github.com/gofiber/fiber/v2"

func (s *FiberServer) HelloWorldHandler(c *fiber.Ctx) error {
	resp := fiber.Map{
		"message": "Hello World!",
	}
	return c.JSON(resp)
}
