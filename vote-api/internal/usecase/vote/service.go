package vote

import "github.com/bosamatheus/pollination/vote-api/internal/entity"

// Service is the use case for Voting.
type Service struct {
	Pub Publisher
}

// NewService returns a new Voting use case.
func NewService(pub Publisher) *Service {
	return &Service{
		Pub: pub,
	}
}

func (s *Service) Send(vote entity.Vote) error {
	return s.Pub.Publish(vote)
}
