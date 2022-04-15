package card

import "time"

// card struct model
type Card struct {
	ID        int64     `json:"id"`
	Width     float64   `json:"width"`
	Height    float64   `json:"height"`
	Color     string    `json:"color"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
