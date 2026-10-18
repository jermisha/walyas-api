package memory_test

import (
	"context"
	"errors"
	"testing"

	"github.com/jermisha/walyas-api/internal/portfolio"
	"github.com/jermisha/walyas-api/internal/portfolio/memory"
)

func TestPortfolioOwnership(t *testing.T) {
	ctx := context.Background()
	repo := memory.New()

	_, err := portfolio.New(
		"P200",
		"UserB",
		"Algorithmic Trading",
	)
	if err != nil {
		t.Fatal(err)
	}

	// User A must not access User B's portfolio.
	_, err = repo.GetByIDForUser(ctx, "P200", "UserA")

	if !errors.Is(err, portfolio.ErrNotFound) {
		t.Fatalf("expected ErrNotFound, got %v", err)
	}

	// The owner should be able to access it.
	got, err := repo.GetByIDForUser(ctx, "P200", "UserB")
	if err != nil {
		t.Fatal(err)
	}

	if got.ID() != "P200" {
		t.Fatalf("unexpected portfolio ID: %s", got.ID())
	}
}
