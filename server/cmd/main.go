package main

import (
	"log"
	"server/db"
)

func main() {
	_, err := db.NewDatabase()
	if err != nil {
		log.Fatalf("Couldn't initialize database connection: %s", err)
	}
}
