// Package usecase implements application business logic. Each logic group in own file.
package usecase

import (
	"context"
	"vladmsnk/taskrec/internal/dto"

	"vladmsnk/taskrec/internal/entity"
)

type (
	// Selection -.
	Selection interface {
		PostActivity(context.Context, dto.ActivityDTO) error
		GetSelection(ctx context.Context) error
	}

	// SelectionRepo -.
	SelectionRepo interface {
		Store(context.Context, entity.Activity) error
		GetRandomActivities(ctx context.Context) ([]entity.Activity, error)
		StoreSelection(ctx context.Context, title string, activitiesIDs []entity.Activity) error
	}
)
