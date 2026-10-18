package memory

import (
	"context"
	"sync"

	"github.com/jermisha/walyas-api/internal/portfolio"
)

type Repository struct {
	mu         sync.RWMutex
	portfolios map[string]*portfolio.Portfolio
}

// Compile-time interface assertion.
var _ portfolio.Repository = (*Repository)(nil)

func New() *Repository {
	return &Repository{
		portfolios: make(map[string]*portfolio.Portfolio),
	}
}

func (r *Repository) Save(ctx context.Context, p *portfolio.Portfolio) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	r.portfolios[p.ID()] = p
	return nil
}

func (r *Repository) GetByIDForUser(
	ctx context.Context,
	portfolioID string,
	userID string,
) (*portfolio.Portfolio, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	r.mu.RLock()
	defer r.mu.RUnlock()

	p, exists := r.portfolios[portfolioID]

	if !exists || p.UserID() != userID {
		return nil, portfolio.ErrNotFound
	}

	return p, nil
}
