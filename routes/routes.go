package routes

import (
	"log"
	"net/http"

	"github.com/Vicrrs/tutorial_gin/controllers"
)

func HandleResquest() {
	http.HandleFunc("/", controllers.Home)
	http.HandleFunc("/api/personalidades", controllers.TodasPersonalidade)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
