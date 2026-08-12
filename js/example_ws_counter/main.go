package main

import "log"

func main() {
	server := NewServer()
	log.Println("Server starting on :3000")

	if err := server.Run(":3000"); err != nil {
		log.Fatal(err)
	}
}
