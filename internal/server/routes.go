package server

import (
	"github.com/ary82/go-user-api/internal/auth"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func (s *FiberServer) RegisterRoutes() {
	// Logger middleware
	s.App.Use(logger.New())

	// Sends a hello world message
	s.App.Get("/", s.helloWorldHandler)

	// Create new user
	s.App.Post("/user", s.insertUserHandler)

	// Login
	s.App.Post("/login", s.loginHandler)

	// Get login info
	s.App.Get("/login", auth.Middleware, s.getLoginInfoHandler)
}
