package portfolio

import (
	"errors"
	"strings"
)

type Portfolio struct {
	id     string
	userID string
	name   string
}

func New(id, userID, name string) (*Portfolio, error) {
	if strings.TrimSpace(id) == "" ||
		strings.TrimSpace(userID) == "" ||
		strings.TrimSpace(name) == "" {
		return nil, errors.New("invalid portfolio")
	}
	return &Portfolio{
		id:     id,
		userID: userID,
		name:   name,
	}, nil
}

func (p *Portfolio) ID() string {
	return p.id
}

func (p *Portfolio) UserID() string {
	return p.userID
}

func (p *Portfolio) Name() string {
	return p.name
}

func (p *Portfolio) Rename(name string) error {
	if strings.TrimSpace(name) == "" {
		return errors.New("portfolio name cannot be empty")
	}
	p.name = name
	return nil
}
