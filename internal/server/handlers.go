package server

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"os"

	"github.com/ary82/go-user-api/internal/auth"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/argon2"
)

func (s *FiberServer) helloWorldHandler(c *fiber.Ctx) error {
	resp := fiber.Map{
		"message": "Hello World!",
	}
	return c.JSON(resp)
}

func (s *FiberServer) insertUserHandler(c *fiber.Ctx) error {
	user := new(NewUserReq)
	err := c.BodyParser(user)
	if err != nil || user.Name == nil || user.Password == nil || user.Email == nil {
		return c.Status(422).JSON(fiber.Map{
			"error": fiber.ErrUnprocessableEntity.Error(),
		})
	}

	// Hash the password
	salt := []byte(os.Getenv("SALT"))
	key := argon2.IDKey([]byte(*user.Password), salt, 1, 64*1024, 4, 32)
	encoded_key := base64.StdEncoding.EncodeToString(key)

	// Create user
	err = s.DB.CreateUser(*user.Email, *user.Name, encoded_key)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(user)
}

func (s *FiberServer) loginHandler(c *fiber.Ctx) error {
	loginReq := new(LoginReq)
	err := c.BodyParser(loginReq)
	if err != nil || loginReq.Email == nil || loginReq.Password == nil {
		return c.Status(422).JSON(fiber.Map{
			"error": fiber.ErrUnprocessableEntity.Error(),
		})
	}

	// Get User details
	user, err := s.DB.GetUser(*loginReq.Email)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Generate key from supplied password
	salt := []byte(os.Getenv("SALT"))
	key := argon2.IDKey([]byte(*loginReq.Password), salt, 1, 64*1024, 4, 32)

	// Decode actual password
	decoded, err := base64.StdEncoding.DecodeString(user.HashedPass)
	if err != nil {
		return fiber.ErrInternalServerError
	}

	// Compare
	if !bytes.Equal(key, decoded) {
		return c.Status(400).JSON(fiber.Map{
			"error": "wrong pass",
		})
	}
	ss, err := auth.GenerateJWT(user.Email)
	if err != nil {
		fmt.Println("HERE")
		return fiber.ErrInternalServerError
	}

	cookie := &fiber.Cookie{
		Name:     "jwt",
		Value:    ss,
		Path:     "/",
		MaxAge:   24 * 3600,
		HTTPOnly: true,
		Secure:   true,
		SameSite: fiber.CookieSameSiteLaxMode,
	}

	c.Cookie(cookie)

	return c.JSON(fiber.Map{
		"message": "logged in",
	})
}

func (s *FiberServer) getLoginInfoHandler(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"current_user": c.Locals("email").(string),
	})
}
