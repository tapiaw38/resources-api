package card

import (
	"context"

	"github.com/tapiaw38/resources-api/database"
	"github.com/tapiaw38/resources-api/models"
)

// GetCard returns a card from the database
func GetCards(ctx context.Context) ([]models.Card, error) {

	var cards []models.Card

	q := `
	SELECT id, width, height, color, created_at, updated_at
		FROM card
		ORDER BY id ASC;
	`

	rows, err := database.Data().QueryContext(ctx, q)
	if err != nil {
		return cards, err
	}

	defer rows.Close()

	for rows.Next() {
		var card models.Card

		err := rows.Scan(
			&card.ID,
			&card.Width,
			&card.Height,
			&card.Color,
			&card.CreatedAt,
			&card.UpdatedAt,
		)
		if err != nil {
			return cards, err
		}

		cards = append(cards, card)
	}

	return cards, nil
}
