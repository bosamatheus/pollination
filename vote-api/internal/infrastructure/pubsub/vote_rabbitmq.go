package pubsub

import (
	"encoding/json"

	"github.com/bosamatheus/pollination/vote-api/internal/entity"
	amqp "github.com/rabbitmq/amqp091-go"
)

type VotingRabbitMQ struct {
	ch    *amqp.Channel
	queue amqp.Queue
}

func NewVotingRabbitMQ(ch *amqp.Channel, queue amqp.Queue) *VotingRabbitMQ {
	return &VotingRabbitMQ{
		ch:    ch,
		queue: queue,
	}
}

func (p *VotingRabbitMQ) Publish(vote entity.Vote) error {
	body, err := json.Marshal(vote)
	if err != nil {
		return err
	}

	return p.ch.Publish(
		"",
		p.queue.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
}
