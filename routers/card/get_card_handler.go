package card

import (
	"encoding/json"
	"net/http"

	card "github.com/tapiaw38/resources-api/database/card"
)

// GetCardHandler is a handler for getting a card
func GetCardsHandler(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	cards, err := card.GetCards(ctx)

	if err != nil {
		http.Error(w, "An error occurred when trying to get cards "+err.Error(), 400)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cards)

}
