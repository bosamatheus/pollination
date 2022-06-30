package publisher

import "github.com/bosamatheus/pollination/voting-api/internal/domain"

type VotingRabbitMQ struct {
}

func NewVotingRabbitMQ() *VotingRabbitMQ {
	return &VotingRabbitMQ{}
}

func (p *VotingRabbitMQ) Publish(vote *domain.Vote) error {
	return nil
}
