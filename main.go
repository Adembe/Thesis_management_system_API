package main

import (
	"go-rest-api/database"
	"go-rest-api/server"
)

func main() {
	database.StartDatabase()
	server := server.NewServer()
	server.Run()
}
