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
func (uc *SelectionUseCase) PostActivity(ctx context.Context, request dto.PostActivityRequest) error {

	err := uc.repo.Store(ctx, request.FromDto())
	if err != nil {
		return err
	}
	return nil
}

// GetSelection -.
func (uc *SelectionUseCase) GetSelection(ctx context.Context) (dto.GetSelectionResponse, error) {
	var selection dto.GetSelectionResponse

}
