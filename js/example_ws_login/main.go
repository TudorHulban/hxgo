package main

import "log"

func main() {
	server := NewServer()

	log.Println("Server starting on :3001")

	if errRun := server.Run(":3001"); errRun != nil {
		log.Fatal(errRun)
	}
}
