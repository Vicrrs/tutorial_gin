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

func ApagaPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var personalidade models.Personalidade
	if err := database.DB.Delete(&personalidade, id).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent) // 204 No content
}

func EditaPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var personalidadeAtualizada models.Personalidade
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&personalidadeAtualizada); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Atualiza a personalidade com o ID especificado
	if err := database.DB.Model(&models.Personalidade{}).Where("id = ?", id).Updates(personalidadeAtualizada).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(personalidadeAtualizada)
}
