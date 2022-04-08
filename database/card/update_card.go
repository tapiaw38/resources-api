package card

import (
	"context"
	"time"

	"github.com/tapiaw38/resources-api/database"
	"github.com/tapiaw38/resources-api/models"
)

// UpdateCard updates a card in the database
func UpdateCard(ctx context.Context, id string, c models.Card) (models.Card, error) {

	var card models.Card

	q := `
	UPDATE card SET width = $1, height = $2, color = $3, updated_at = $4
		WHERE id = $5
		RETURNING id, width, height, color, created_at, updated_at;
	`

	row := database.Data().QueryRowContext(
		ctx, q, c.Width, c.Height, c.Color, time.Now(), id,
	)

	err := row.Scan(
		&card.ID,
		&card.Width,
		&card.Height,
		&card.Color,
		&card.CreatedAt,
		&card.UpdatedAt,
	)

	if err != nil {
		return card, err
	}

	return card, nil
}
