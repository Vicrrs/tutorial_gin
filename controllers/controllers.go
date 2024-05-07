package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Vicrrs/tutorial_gin/database"
	"github.com/Vicrrs/tutorial_gin/models"
	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home page")
}

func TodasPersonalidade(w http.ResponseWriter, r *http.Request) {
	var p []models.Personalidade
	database.DB.Find(&p)
	json.NewEncoder(w).Encode(p)
}

func RetornaUmaPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var personalidade models.Personalidade
	database.DB.First(&personalidade, id)
	json.NewEncoder(w).Encode(personalidade)
}

func CriaUmaNovaPersonalidade(w http.ResponseWriter, r *http.Request) {
	var novaPersonalidade models.Personalidade

	// USar json.newencoder para ler o rbody
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&novaPersonalidade)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Criando uma nova personalidade no db
	database.DB.Create(&novaPersonalidade)
	// Retorna a personalidae criada como resposta em json
	json.NewEncoder(w).Encode(novaPersonalidade)
}
