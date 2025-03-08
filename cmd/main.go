package main

import (
	"log"

	"github.com/dhfai/go-rest-api/cmd/api"
	"github.com/dhfai/go-rest-api/db"
)

func main() {
	db, err := db.NewDBStorage()
	if err != nil {
		log.Fatal(err)
	}

	sqlDB, _ := db.DB.DB()
	defer sqlDB.Close()

	server := api.NewAPIServer(":8080", nil)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
