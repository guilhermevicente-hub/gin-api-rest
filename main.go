package main

import (
	"github.com/guilhermevicente-hub/api-go-gin/database"
	"github.com/guilhermevicente-hub/api-go-gin/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequests()
}
