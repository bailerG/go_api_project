package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go_api_project/database"
	"go_api_project/entities"
	"log"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintln(w, "Home Screen")
}

func GetPersonalities(w http.ResponseWriter, r *http.Request) {
	var p []entities.Personality
	database.DB.Table("personalities").Find(&p)

	err := json.NewEncoder(w).Encode(p)
	if err != nil {
		log.Panic(err.Error())
	}
}

func GetSinglePersonality(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]
	var p entities.Personality

	database.DB.First(&p, id)

	err := json.NewEncoder(w).Encode(p)
	if err != nil {
		log.Fatal(err.Error())
	}

}

func CreateNewPersonality(w http.ResponseWriter, r *http.Request) {
	var newPersonality entities.Personality
	err := json.NewDecoder(r.Body).Decode(&newPersonality)
	if err != nil {
		log.Panic(err.Error())
	}
	database.DB.Create(&newPersonality)
	e := json.NewEncoder(w).Encode(newPersonality)
	if err != nil {
		log.Panic(e.Error())
	}
}

func DeletePersonality(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var personality entities.Personality
	database.DB.Delete(&personality, id)
	err := json.NewEncoder(w).Encode(personality)
	if err != nil {
		log.Panic(err.Error())
	}
}

func EditPersonality(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var p entities.Personality
	database.DB.First(&p, id)
	_ = json.NewDecoder(r.Body).Decode(&p)
	database.DB.Save(&p)
	_ = json.NewEncoder(w).Encode(p)
}
