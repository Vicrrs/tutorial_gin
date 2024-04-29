package main

import (
	"fmt"

	"github.com/Vicrrs/tutorial_gin/models"
	"github.com/Vicrrs/tutorial_gin/routes"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Nome: "Nome 1", Historia: "Historia 1"},
		{Nome: "Nome w", Historia: "Historia 2"},
	}

	fmt.Println("Iniciando servidor Rest com Go")
	routes.HandleResquest()
}
