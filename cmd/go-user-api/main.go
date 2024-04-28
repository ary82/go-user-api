package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/ary82/go-user-api/internal/database"
	"github.com/ary82/go-user-api/internal/server"

	_ "github.com/joho/godotenv/autoload"
)

// Env Vars
var (
	dbAddr    = os.Getenv("DB_ADDR")
	namespace = os.Getenv("DB_NAMESPACE")
	port      = os.Getenv("PORT")
)

func main() {
	// Make a new ScyllaDB Service
	db, err := database.NewScyllaStore(dbAddr, namespace)
	if err != nil {
		log.Fatal(err)
	}
	err = db.InitTables()
	if err != nil {
		log.Fatal(err)
	}

	// Make and Run Server
	server := server.New(db)
	server.Run(port)

	// Create channel to signify a signal being sent
	c := make(chan os.Signal, 1)
	// notify the channel on interrent/terminate
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	// Wait for thread until interrupt is received
	<-c

	fmt.Println("Gracefully shutting down...")
	_ = server.App.Shutdown()

	fmt.Println("Running cleanup tasks...")
	db.Close()

	fmt.Println("Shutdown Successful")
}
