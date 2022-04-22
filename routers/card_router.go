package routers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tapiaw38/resources-api/models/card"
)

type CardRouter struct {
	Storage card.Storage
}

// CreateCardHandler is a handler for creating a card
func (cd *CardRouter) CreateCardHandler(w http.ResponseWriter, r *http.Request) {

	var c card.Card

	err := json.NewDecoder(r.Body).Decode(&c)

	if c.Color == "" {
		http.Error(w, "An error occurred, color is required", http.StatusBadRequest)
		return
	}

	if err != nil {
		http.Error(w, "An error occurred when trying to enter a card "+err.Error(), 400)
		return
	}

	defer r.Body.Close()

	ctx := r.Context()
	card, err := cd.Storage.CreateCard(ctx, &c)

	if err != nil {
		http.Error(w, "An error occurred when trying to enter a card in database "+err.Error(), 400)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(card)

}

// GetCardHandler is a handler for getting a card
func (cd *CardRouter) GetCardsHandler(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	cards, err := cd.Storage.GetCards(ctx)

	if err != nil {
		http.Error(w, "An error occurred when trying to get cards "+err.Error(), 400)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cards)

}

// UpdateCardHandler is a handler for updating a card
func (cd *CardRouter) UpdateCardHandler(w http.ResponseWriter, r *http.Request) {

	var c card.Card

	id := mux.Vars(r)["id"]

	if id == "" {
		http.Error(w, "An error occurred, id is required", http.StatusBadRequest)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&c)

	// color are required
	if c.Color == "" {
		http.Error(w, "An error occurred, color is required", http.StatusBadRequest)
		return
	}

	if err != nil {
		http.Error(w, "An error occurred when trying to enter a card "+err.Error(), 400)
		return
	}

	defer r.Body.Close()

	ctx := r.Context()
	card, err := cd.Storage.UpdateCard(ctx, id, c)

	if err != nil {
		http.Error(w, "An error occurred when trying to update a card in database "+err.Error(), 400)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(card)
}
