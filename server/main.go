package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type Recipe struct {
	ID           string   `json:"id"`
	Name         string   `json:"name"`
	Ingredients  []string `json:"ingredients"`
	Instructions string   `json:"instructions"`
	Cuisine      string   `json:"cuisine"`
	PrepTime     int      `json:"prepTime"`
	CookTime     int      `json:"cookTime"`
}

var recipes = []Recipe{}

func listRecipes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(recipes)
}

func addRecipe(w http.ResponseWriter, r *http.Request) {
	var recipe Recipe
	if err := json.NewDecoder(r.Body).Decode(&recipe); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	recipes = append(recipes, recipe)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(recipe)
}

func getRecipe(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, recipe := range recipes {
		if recipe.ID == params["id"] {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(recipe)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}

func updateRecipe(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var updated Recipe
	if err := json.NewDecoder(r.Body).Decode(&updated); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	for i, recipe := range recipes {
		if recipe.ID == params["id"] {
			recipes[i] = updated
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(updated)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}

func deleteRecipe(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for i, recipe := range recipes {
		if recipe.ID == params["id"] {
			recipes = append(recipes[:i], recipes[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/recipes", listRecipes).Methods("GET")
	r.HandleFunc("/recipes", addRecipe).Methods("POST")
	r.HandleFunc("/recipes/{id}", getRecipe).Methods("GET")
	r.HandleFunc("/recipes/{id}", updateRecipe).Methods("PUT")
	r.HandleFunc("/recipes/{id}", deleteRecipe).Methods("DELETE")

	http.ListenAndServe(":3000", r)
}
