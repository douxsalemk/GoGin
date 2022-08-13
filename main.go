package main

import (
	"rest-api-go-gin/database"
	"rest-api-go-gin/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequests()
}