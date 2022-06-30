package presenter

import (
	"time"

	"github.com/bosamatheus/pollination/voting-api/internal/domain"
)

// Vote presentation.
type Vote struct {
	Option   VoteOption `json:"option"`
	Username string     `json:"username"`
	Email    string     `json:"email"`
}

// Vote Option presentation.
type VoteOption struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
}

// MapToDomain maps Vote presentation to domain.
func (p *Vote) MapToDomain() *domain.Vote {
	return &domain.Vote{
		Option: domain.VoteOption{
			ID:          p.Option.ID,
			Description: p.Option.Description,
		},
		Username:  p.Username,
		Email:     p.Email,
		CreatedAt: time.Now(),
	}
}
