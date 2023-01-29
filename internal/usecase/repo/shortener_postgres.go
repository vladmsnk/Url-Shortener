package repo

import (
	"context"
	"time"
	"vladmsnk/urlshort/internal/body"
	"vladmsnk/urlshort/pkg/postgres"
)

// Repo -.
type Repo struct {
	*postgres.Postgres
}

// New -.
func New(pg *postgres.Postgres) *Repo {
	return &Repo{pg}
}

// StoreURL -.
func (r *Repo) StoreURL(ctx context.Context, longURL, shortURL string) error {
	_, err := r.Pool.Exec(ctx, body.InsertCreatedURL, longURL, shortURL, time.Now())
	if err != nil {
		return err
	}
	return nil
}

// GetLongURL -.
func (r *Repo) GetLongURL(ctx context.Context, shortURL string) (string, error) {
	var longURL string

	err := r.Pool.QueryRow(ctx, body.SelectLongURLByShortURL, shortURL).Scan(&longURL)
	if err != nil {
		return "", err
	}

	return longURL, err
}

// GetShortURL -.
func (r *Repo) GetShortURL(ctx context.Context, longURL string) (string, error) {
	var short string

	err := r.Pool.QueryRow(ctx, body.SelectShortURLByLongURL, longURL).Scan(&short)
	if err != nil {
		return "", err
	}

	return short, err
}
