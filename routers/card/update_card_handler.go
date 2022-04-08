package card

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	card "github.com/tapiaw38/resources-api/database/card"
	"github.com/tapiaw38/resources-api/models"
)

// UpdateCardHandler is a handler for updating a card
func UpdateCardHandler(w http.ResponseWriter, r *http.Request) {

	var c models.Card

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
	card, err := card.UpdateCard(ctx, id, c)

	if err != nil {
		http.Error(w, "An error occurred when trying to update a card in database "+err.Error(), 400)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(card)
}
