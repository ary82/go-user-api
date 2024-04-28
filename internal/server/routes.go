package server

import "github.com/gofiber/fiber/v2/middleware/logger"

func (s *FiberServer) RegisterRoutes() {
	// Logger middleware
	s.App.Use(logger.New())

	// Sends a hello world message
	s.App.Get("/", s.HelloWorldHandler)

	// Create new user
	s.App.Post("/user", s.InsertUserHandler)
}
