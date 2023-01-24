package usecase

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
