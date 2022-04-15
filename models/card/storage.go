package card

import "context"

// CardStorage is the interface for the storage of card
type Storage interface {
	CreateCard(ctx context.Context, card *Card) (Card, error)
	GetCards(ctx context.Context) ([]Card, error)
	UpdateCard(ctx context.Context, id string, card Card) (Card, error)
}
