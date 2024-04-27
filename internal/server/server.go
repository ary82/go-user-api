package server

import (
	"log"

	"github.com/ary82/go-user-api/internal/database"
	"github.com/gofiber/fiber/v2"
)

type FiberServer struct {
	App *fiber.App
	DB  database.Store
}

func New(db database.Store) *FiberServer {
	return &FiberServer{
		App: fiber.New(),
		DB:  db,
	}
}

func (s *FiberServer) Run(port string) {
	s.RegisterRoutes()

	// Start the Server in a goroutine
	go func() {
		err := s.App.Listen(port)
		if err != nil {
			log.Fatal(err)
		}
	}()
}
