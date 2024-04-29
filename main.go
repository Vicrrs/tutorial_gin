package main

import (
	"fmt"

	"github.com/Vicrrs/tutorial_gin/routes"
)

func main() {
	fmt.Println("Iniciando servidor Rest com Go")
	routes.HandleResquest()
}
