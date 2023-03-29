package main

import (
	"log"

	"github.com/danielwangai/todo-app/internal/transport"
)

func main() {
	err := transport.RunServer()
	if err != nil {
		log.Fatalf("Could not initialize server: %v", err)
	}
}
