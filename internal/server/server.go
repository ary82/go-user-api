package server

import (
	"github.com/ary82/go-user-api/internal/database"
	"github.com/gofiber/fiber/v2"
)

type Server struct {
	Addr string
	DB   database.Database
}

func NewServer(addr string, db database.Database) *Server {
	return &Server{
		Addr: addr,
		DB:   db,
	}
}

func (s *Server) Init() error {
	app := fiber.New()

	err := app.Listen(s.Addr)
	return err
}
