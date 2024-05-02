package main

import (
	"fmt"

	"github.com/Vicrrs/tutorial_gin/database"
	"github.com/Vicrrs/tutorial_gin/models"
	"github.com/Vicrrs/tutorial_gin/routes"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "Nome 1", Historia: "Historia 1"},
		{Id: 2, Nome: "Nome 2", Historia: "Historia 2"},
	}

	database.ConectaComBancoDeDados()
	fmt.Println("Iniciando servidor Rest com Go")
	routes.HandleResquest()
}
