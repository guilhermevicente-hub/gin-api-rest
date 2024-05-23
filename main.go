package main

import (
	"github.com/guilhermevicente-hub/api-go-gin/database"
	"github.com/guilhermevicente-hub/api-go-gin/models"
	"github.com/guilhermevicente-hub/api-go-gin/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	models.Alunos = []models.Aluno{
		{Nome: "Guilherme Vicente", CPF: "000000000", RG: "4700000000"},
		{Nome: "Ariel Facciolli", CPF: "111111111", RG: "4800000000"},
	}
	routes.HandleRequests()
}
