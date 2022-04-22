package storage

import (
	"context"
	"time"

	"github.com/tapiaw38/resources-api/models/card"
)

// CardStorage is the interface that wraps the methods to manage the cards
type CardStorage struct {
	Data *Data
}

// CreateCard creates a new card in the database
func (cd *CardStorage) CreateCard(ctx context.Context, c *card.Card) (card.Card, error) {

	var card card.Card

	q := `
	INSERT INTO card (width, height, color, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, width, height, color, created_at, updated_at;
	`

	row := cd.Data.DB.QueryRowContext(
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

// GetCard returns a card from the database
func (cd *CardStorage) GetCards(ctx context.Context) ([]card.Card, error) {

	var cards []card.Card

	q := `
	SELECT id, width, height, color, created_at, updated_at
		FROM card
		ORDER BY id ASC;
	`

	rows, err := cd.Data.DB.QueryContext(ctx, q)
	if err != nil {
		return cards, err
	}

	defer rows.Close()

	for rows.Next() {
		var card card.Card

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

// UpdateCard updates a card in the database
func (cd *CardStorage) UpdateCard(ctx context.Context, id string, c card.Card) (card.Card, error) {

	var card card.Card

	q := `
	UPDATE card SET width = $1, height = $2, color = $3, updated_at = $4
		WHERE id = $5
		RETURNING id, width, height, color, created_at, updated_at;
	`

	row := cd.Data.DB.QueryRowContext(
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
