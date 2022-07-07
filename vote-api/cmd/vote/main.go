package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bosamatheus/pollination/vote-api/internal/api/handler"
	"github.com/bosamatheus/pollination/vote-api/internal/infrastructure/pubsub"
	"github.com/bosamatheus/pollination/vote-api/internal/usecase/vote"
	fiber "github.com/gofiber/fiber/v2"
	fiberCORS "github.com/gofiber/fiber/v2/middleware/cors"
	fiberLogger "github.com/gofiber/fiber/v2/middleware/logger"
	fiberRecover "github.com/gofiber/fiber/v2/middleware/recover"
	amqp "github.com/rabbitmq/amqp091-go"
	log "github.com/sirupsen/logrus"
)

func main() {
	// logger
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})

	// messaging
	url := fmt.Sprintf("amqps://zxvfqrgz:%s@jackal.rmq.cloudamqp.com/zxvfqrgz", "pwd")
	conn, err := amqp.Dial(url)
	failOnError(err, "failed to connect to RabbitMQ")
	defer conn.Close()
	ch, err := conn.Channel()
	failOnError(err, "failed to open RabbitMQ channel")
	defer ch.Close()
	queueName := "create_votes_q"
	queue, err := ch.QueueDeclare(queueName, true, false, false, false, nil)
	failOnError(err, "failed to declare RabbitMQ queue")

	// app
	app := fiber.New()

	// middlewares
	app.Use(fiberCORS.New())
	app.Use(fiberRecover.New())
	app.Use(fiberLogger.New(fiberLogger.Config{
		TimeFormat: time.RFC3339,
	}))

	// pubsub
	pub := pubsub.NewVotingRabbitMQ(ch, queue)

	// use cases
	service := vote.NewService(pub)

	// routers
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})
	app.Post("/api/v1/votes", handler.SendVote(service))

	// start server
	go func() {
		if err := app.Listen(":8080"); err != nil {
			log.Fatal(err)
		}
	}()

	// graceful shutdown
	quit := make(chan os.Signal, 1)                    // create channel to signify a signal being sent
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM) // when an interrupt or termination signal is sent, notify the channel
	<-quit                                             // blocks the main thread until an interrupt is received

	err = app.Shutdown()
	if err != nil {
		log.Fatal(err)
	}
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
