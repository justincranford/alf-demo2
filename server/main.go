package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// @title Food Recipes API
// @version 1.0
// @description API for managing food recipes
// @host localhost:3000
// @BasePath /

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

	// Check for duplicate ID
	for _, existingRecipe := range recipes {
		if existingRecipe.ID == recipe.ID {
			w.WriteHeader(http.StatusConflict)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{
				"error": fmt.Sprintf("Recipe with ID %s already exists", recipe.ID),
			})
			return
		}
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

	// Serve Swagger UI
	fs := http.FileServer(http.Dir("./swagger"))
	r.PathPrefix("/swagger/").Handler(http.StripPrefix("/swagger/", fs))

	// Serve OpenAPI spec
	r.HandleFunc("/openapi.yaml", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "../openapi.yaml")
	})

	log.Println("Server started on http://localhost:3000")
	log.Println("Swagger UI available at http://localhost:3000/swagger/")
	log.Fatal(http.ListenAndServe(":3000", r))
}
