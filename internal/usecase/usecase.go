package usecase

import (
	"context"
	"github.com/jackc/pgx/v4"
)

// ShortenerUseCase -.
type ShortenerUseCase struct {
	repo ShortenerRepo
}

// New -.
func New(r ShortenerRepo) *ShortenerUseCase {
	return &ShortenerUseCase{
		repo: r,
	}
}

func (uc *ShortenerUseCase) CreateURL(ctx context.Context, longURL string) (string, error) {
	url, err := uc.repo.GetShortURL(ctx, longURL)
	if len(url) > 0 {
		return url, nil
	}

	var shortURL string
	for {

		shortURL = GenerateShortURL()
		_, err = uc.repo.GetLongURL(ctx, shortURL)
		if err == pgx.ErrNoRows {
			break
		}
	}
	err = uc.repo.StoreURL(ctx, longURL, shortURL)
	if err != nil {
		return "", err
	}
	return shortURL, nil
}

func (uc *ShortenerUseCase) GetURL(ctx context.Context, shortURL string) (string, error) {
	longURL, err := uc.repo.GetLongURL(ctx, shortURL)
	if err != nil {
		return "", err
	}
	return longURL, nil
}
