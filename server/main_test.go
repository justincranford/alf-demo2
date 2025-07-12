package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestRecipeAPI(t *testing.T) {
	recipes = []Recipe{} // Reset global recipes

	recipe := Recipe{
		ID:           "1",
		Name:         "Pasta",
		Ingredients:  []string{"noodles", "sauce"},
		Instructions: "Boil noodles. Add sauce.",
		Cuisine:      "Italian",
		PrepTime:     10,
		CookTime:     15,
	}

	// Test POST /recipes
	body, _ := json.Marshal(recipe)
	req := httptest.NewRequest("POST", "/recipes", bytes.NewReader(body))
	w := httptest.NewRecorder()
	addRecipe(w, req)
	if w.Code != http.StatusCreated {
		t.Fatalf("expected 201, got %d", w.Code)
	}

	// Test GET /recipes
	req = httptest.NewRequest("GET", "/recipes", nil)
	w = httptest.NewRecorder()
	listRecipes(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}

	// Test GET /recipes/{id}
	muxCtx := mux.SetURLVars(req, map[string]string{"id": "1"})
	w = httptest.NewRecorder()
	getRecipe(w, muxCtx)
	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}

	// Test PUT /recipes/{id}
	recipe.Name = "Spaghetti"
	body, _ = json.Marshal(recipe)
	req = httptest.NewRequest("PUT", "/recipes/1", bytes.NewReader(body))
	muxCtx = mux.SetURLVars(req, map[string]string{"id": "1"})
	w = httptest.NewRecorder()
	updateRecipe(w, muxCtx)
	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}

	// Test DELETE /recipes/{id}
	req = httptest.NewRequest("DELETE", "/recipes/1", nil)
	muxCtx = mux.SetURLVars(req, map[string]string{"id": "1"})
	w = httptest.NewRecorder()
	deleteRecipe(w, muxCtx)
	if w.Code != http.StatusNoContent {
		t.Fatalf("expected 204, got %d", w.Code)
	}
}

func TestRecipeIDUniqueness(t *testing.T) {
	recipes = []Recipe{} // Reset global recipes

	recipe := Recipe{
		ID:           "tx001",
		Name:         "Texas-Style Beef Brisket",
		Ingredients:  []string{"brisket", "salt", "pepper"},
		Instructions: "Smoke it low and slow",
		Cuisine:      "Texas BBQ",
		PrepTime:     60,
		CookTime:     720,
	}

	// Add first recipe
	body, _ := json.Marshal(recipe)
	req := httptest.NewRequest("POST", "/recipes", bytes.NewReader(body))
	w := httptest.NewRecorder()
	addRecipe(w, req)
	if w.Code != http.StatusCreated {
		t.Fatalf("expected 201, got %d", w.Code)
	}

	// Try to add the same recipe again (same ID)
	req = httptest.NewRequest("POST", "/recipes", bytes.NewReader(body))
	w = httptest.NewRecorder()
	addRecipe(w, req)

	// Should get a conflict status
	if w.Code != http.StatusConflict {
		t.Fatalf("expected 409 conflict for duplicate ID, got %d", w.Code)
	}

	// Verify we still only have one recipe
	req = httptest.NewRequest("GET", "/recipes", nil)
	w = httptest.NewRecorder()
	listRecipes(w, req)

	var responseRecipes []Recipe
	if err := json.Unmarshal(w.Body.Bytes(), &responseRecipes); err != nil {
		t.Fatalf("failed to parse response: %v", err)
	}
	if len(responseRecipes) != 1 {
		t.Fatalf("expected 1 recipe, got %d", len(responseRecipes))
	}
}
