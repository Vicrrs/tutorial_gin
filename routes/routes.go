package routes

import (
	"log"
	"net/http"

	"github.com/Vicrrs/tutorial_gin/controllers"
	"github.com/gorilla/mux"
)

func HandleResquest() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalidades", controllers.TodasPersonalidade)
	log.Fatal(http.ListenAndServe(":8000", r))
}
