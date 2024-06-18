package main

import (
	"go_api_project/database"
	"go_api_project/routes"
	"log"
)

func main() {
	database.GetDB()
	log.Println("Starting server...")
	routes.HandleRequest()
}
