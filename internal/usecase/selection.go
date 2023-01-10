package usecase

import (
	"context"
	"vladmsnk/taskrec/internal/dto"
)

// SelectionUseCase -.
type SelectionUseCase struct {
	repo SelectionRepo
}

// New -.
func New(r SelectionRepo) *SelectionUseCase {
	return &SelectionUseCase{
		repo: r,
	}
}

// PostActivity -.
func (uc *SelectionUseCase) PostActivity(ctx context.Context, request dto.ActivityDTO) error {

	err := uc.repo.Store(ctx, request.FromDto())
	if err != nil {
		return err
	}
	return nil
}

// GetSelection -.
func (uc *SelectionUseCase) GetSelection(ctx context.Context) (dto.GetSelectionResponse, error) {

	title := "title"
	activities, err := uc.repo.GetRandomActivities(ctx)

	if err != nil {
		return dto.GetSelectionResponse{}, err
	}
	selection := dto.GetSelectionResponse{}.ToDto(title, activities)

	err = uc.repo.StoreSelection(ctx, title, activities)
	if err != nil {
		return dto.GetSelectionResponse{}, err
	}
	return selection, nil
}
