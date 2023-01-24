// Package usecase implements application business logic. Each logic group in own file.
package usecase

import "context"

type (
	ShortenerRepo interface {
		StoreURL(ctx context.Context, longURL, shortURL string) error
		GetLongURL(ctx context.Context, shortURL string) (string, error)
		GetShortURL(ctx context.Context, longURL string) (string, error)
	}
)
