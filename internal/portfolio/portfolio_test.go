package portfolio_test

import (
	"testing"

	"github.com/jermisha/walyas-api/internal/portfolio"
)

func TestNewPortfolio(t *testing.T) {
	p, err := portfolio.New(
		"P100",
		"U100",
		"Growth",
	)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if p.ID() != "P100" {
		t.Errorf("expected ID P100, got %s", p.ID())
	}

	if p.UserID() != "U100" {
		t.Errorf("expected user U100, got %s", p.UserID())
	}

	if p.Name() != "Growth" {
		t.Errorf("expected name Growth, got %s", p.Name())
	}

	// if p.Version() != 1 {
	// 	t.Errorf("expected version 1, got %d", p.Version())
	// }
}

func TestRenamePortfolio(t *testing.T) {
	p, err := portfolio.New("P100", "U100", "Growth")
	if err != nil {
		t.Fatal(err)
	}

	err = p.Rename("Retirement")
	if err != nil {
		t.Fatal(err)
	}

	if p.Name() != "Retirement" {
		t.Errorf("expected Retirement, got %s", p.Name())
	}
}

func TestRenameRejectsEmptyName(t *testing.T) {
	p, err := portfolio.New("P100", "U100", "Growth")
	if err != nil {
		t.Fatal(err)
	}

	err = p.Rename("")

	// if !errors.Is(err, portfolio.ErrInvalidName) {
	// 	t.Fatalf("expected ErrInvalidName, got %v", err)
	// }

	if p.Name() != "Growth" {
		t.Fatal("portfolio name changed after failed rename")
	}
}
