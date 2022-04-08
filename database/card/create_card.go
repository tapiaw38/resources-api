package card

import (
	"context"
	"time"

	"github.com/tapiaw38/resources-api/database"
	"github.com/tapiaw38/resources-api/models"
)

// CreateCard creates a new card in the database
func CreateCard(ctx context.Context, c *models.Card) (models.Card, error) {

	var card models.Card

	q := `
	INSERT INTO card (width, height, color, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, width, height, color, created_at, updated_at;
	`

	row := database.Data().QueryRowContext(
		ctx, q, c.Width, c.Height, c.Color, time.Now(), time.Now(),
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
