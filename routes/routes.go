package routes

import (
	"log"
	"net/http"

	"github.com/Vicrrs/tutorial_gin/controllers"
)

func HandleResquest() {
	http.HandleFunc("/", controllers.Home)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
