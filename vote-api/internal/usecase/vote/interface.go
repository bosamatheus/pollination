package vote

import "github.com/bosamatheus/pollination/vote-api/internal/entity"

// UseCase interface.
type UseCase interface {
	// SendVote sends a Vote to be processed.
	Send(vote entity.Vote) error
}

// Publisher interface.
type Publisher interface {
	// Publish publishes a Vote to the queue.
	Publish(vote entity.Vote) error
}
