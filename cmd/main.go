package main

import (
	"log"
	"os"

	"image-server/internal/handler"
)

func main() {
	port := "8000"
	if len(os.Args) > 1 {
		port = os.Args[1]
	}

	// Initialize and start the server
	server := handler.NewServer(port)
	print("Server running on http://localhost:", port, "\n")
	log.Fatal(server.Start())
}
