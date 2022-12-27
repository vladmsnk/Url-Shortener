// Package entity defines main entities for business logic (services), data base mapping and
// HTTP response objects if suitable. Each logic group entities in own file.
package entity

import (
	"github.com/jackc/pgtype"
	"time"
)

// Activity -.
type Activity struct {
	ID            pgtype.UUID `db:"id"`
	Title         string      `db:"s_title"`
	Description   string      `db:"description"`
	Price         float64     `db:"original"`
	AvailableFrom time.Time   `db:"available_from"`
	AvailableTo   time.Time   `db:"available_to"`
}
