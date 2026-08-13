package main

import "log"

func main() {
	server := NewServer()

	log.Println("Server starting on :3000")

	if errRun := server.Run(":3000"); errRun != nil {
		log.Fatal(errRun)
	}
}
