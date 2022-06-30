package voting

import "github.com/bosamatheus/pollination/voting-api/internal/domain"

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

func (s *Service) SendVote(vote *domain.Vote) error {
	return s.Pub.Publish(vote)
}
