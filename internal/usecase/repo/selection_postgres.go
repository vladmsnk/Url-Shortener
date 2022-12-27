package repo

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"vladmsnk/taskrec/internal/entity"
	"vladmsnk/taskrec/pkg/postgres"
)

const _defaultEntityCap = 64

// SelectionRepo -.
type SelectionRepo struct {
	*postgres.Postgres
}

// New -.
func New(pg *postgres.Postgres) *SelectionRepo {
	return &SelectionRepo{pg}
}

// Store -.
func (r *SelectionRepo) Store(ctx context.Context, t entity.Activity) error {
	sql, args, err := r.Builder.
		Insert("activities").
		Columns("id, s_title, description, price, available_from, available_to, created_at").
		Values(uuid.New(), t.Title, t.Description, t.Price, t.AvailableFrom, t.AvailableTo, time.Now()).
		ToSql()
	if err != nil {
		return fmt.Errorf("Selectionnepo - Store - r.Builder: %w", err)
	}

	_, err = r.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("TranslationRepo - Store - r.Pool.Exec: %w", err)
	}

	return nil
}
