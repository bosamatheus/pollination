package voting

import "github.com/bosamatheus/pollination/voting-api/internal/domain"

// UseCase interface.
type UseCase interface {
	// SendVote sends a Vote to be processed.
	SendVote(vote *domain.Vote) error
}

// Publisher interface.
type Publisher interface {
	// Publish publishes a Vote to the queue.
	Publish(vote *domain.Vote) error
}
