package card

import (
	"encoding/json"
	"net/http"

	"github.com/tapiaw38/resources-api/models"

	card "github.com/tapiaw38/resources-api/database/card"
)

// CreateCardHandler is a handler for creating a card
func CreateCardHandler(w http.ResponseWriter, r *http.Request) {

	var c models.Card

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
	card, err := card.CreateCard(ctx, &c)

	if err != nil {
		http.Error(w, "An error occurred when trying to enter a card in database "+err.Error(), 400)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(card)

}
