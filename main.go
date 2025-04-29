package main

import (
	"log"

	"github.com/darendsen/test-gin/internal/database"
	"github.com/darendsen/test-gin/internal/routes"
)

func main() {
	// Initialize database
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}

	// Setup and run router
	router := routes.SetupRouter(db)
	router.Run() // listen and serve on 0.0.0.0:8080
}
