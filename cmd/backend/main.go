package main

import (
	"log"

	"github.com/hypertrace/demo-events-app/pkg/backend"
)

func main() {
	router, err := backend.NewRouter()
	if err != nil {
		log.Fatalf("cannot create router: %v", err)
	}

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("cannot run backend server: %v", err)
	}
}
