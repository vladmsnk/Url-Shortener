package entity

import (
	"github.com/jackc/pgtype"
)

// Activity -.
type Selection struct {
	ID     pgtype.UUID `db:"id"`
	UserId pgtype.UUID `db:"user_id"`
	Title  string      `db:"title"`
}
