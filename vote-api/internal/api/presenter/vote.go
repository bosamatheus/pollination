package presenter

import (
	"time"

	"github.com/bosamatheus/pollination/vote-api/internal/entity"
)

// Vote presentation.
type Vote struct {
	PollID        int    `json:"pollId"`
	AlternativeID int    `json:"alternativeId"`
	Username      string `json:"username"`
}

// MapToEntity maps Vote presentation to domain.
func (p *Vote) MapToEntity() entity.Vote {
	return entity.Vote{
		PollID:        p.PollID,
		AlternativeID: p.AlternativeID,
		Username:      p.Username,
		CreatedAt:     time.Now(),
	}
}
