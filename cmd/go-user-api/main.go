package main

import (
	"log"
	"os"

	"github.com/ary82/go-user-api/internal/database"
	"github.com/ary82/go-user-api/internal/server"

	"github.com/joho/godotenv"
)

func main() {
	// Load env vars
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	// DB Env Vars
	var (
		dbAddr    = os.Getenv("DB_ADDR")
		namespace = os.Getenv("DB_NAMESPACE")
	)

	// Server Env Vars
	port := os.Getenv("PORT")

	db, err := database.NewScyllaDB(dbAddr, namespace)
	if err != nil {
		log.Fatal(err)
	}

	server := server.NewServer(port, db)
	err = server.Init()
	if err != nil {
		log.Fatal(err)
	}
}
