package repo

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
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

func (r *SelectionRepo) GetRandomActivities(ctx context.Context) ([]entity.Activity, error) {
	tx, _ := r.Pool.BeginTx(ctx, pgx.TxOptions{})
	rows, err := tx.Query(ctx, SelectRandomActivities)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var activities []entity.Activity
	for rows.Next() {
		var activity entity.Activity
		err := rows.Scan(&activity.ID, &activity.Title, &activity.Description, &activity.Price, &activity.AvailableFrom,
			&activity.AvailableTo)
		if err != nil {
			tx.Rollback(ctx)
			return nil, err
		}
		activities = append(activities, activity)
	}
	tx.Commit(ctx)
	return activities, nil
}
