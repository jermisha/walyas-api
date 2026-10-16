package portfolio

import (
	"context"
	"errors"
)

var ErrNotFound = errors.New("portfolio not found")

type Repository interface {
	Save(ctx context.Context, p *Portfolio) error

	GetByIDForUser(
		ctx context.Context,
		portfolioID string,
		userID string,
	) (*Portfolio, error)
}
