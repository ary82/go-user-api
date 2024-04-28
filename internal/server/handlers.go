package server

import (
	"encoding/base64"
	"os"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/argon2"
)

func (s *FiberServer) HelloWorldHandler(c *fiber.Ctx) error {
	resp := fiber.Map{
		"message": "Hello World!",
	}
	return c.JSON(resp)
}

func (s *FiberServer) InsertUserHandler(c *fiber.Ctx) error {
	user := new(NewUserReq)

	err := c.BodyParser(user)
	if err != nil || user.Name == nil || user.Password == nil || user.Email == nil {
		return c.Status(422).JSON(fiber.Map{
			"error": fiber.ErrUnprocessableEntity.Error(),
		})
	}

	salt := []byte(os.Getenv("SALT"))
	key := argon2.IDKey([]byte(*user.Password), salt, 1, 64*1024, 4, 32)
	encoded_key := base64.StdEncoding.EncodeToString(key)

	err = s.DB.CreateUser(*user.Email, *user.Name, encoded_key)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(user)
}
