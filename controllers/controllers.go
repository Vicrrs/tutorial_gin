package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Vicrrs/tutorial_gin/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home page")
}

func TodasPersonalidade(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.Personalidades)
}
